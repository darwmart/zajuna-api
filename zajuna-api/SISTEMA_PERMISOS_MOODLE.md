# Sistema de Permisos de Moodle en Zajuna API

## 📋 Descripción General

Este documento describe la implementación del sistema de permisos de **Moodle** en **zajuna-api** (Go + Gin + GORM). El sistema replica la arquitectura completa de permisos de Moodle, incluyendo:

- **Contextos jerárquicos** (System, User, Course, Module, etc.)
- **Roles y capacidades**
- **Herencia de permisos**
- **Niveles de permiso** (ALLOW, PREVENT, PROHIBIT)

---

## 🏗️ Arquitectura

### Componentes Implementados

```
zajuna-api/
├── internal/
│   ├── models/
│   │   └── permission.go          # Modelos de datos (Context, Role, RoleAssignment, etc.)
│   ├── repository/
│   │   └── permission_repository.go  # Acceso a datos de permisos
│   ├── services/
│   │   └── permission_service.go     # Lógica de negocio de permisos
│   ├── middleware/
│   │   └── permission.go          # Middlewares de autorización para Gin
│   └── routes/
│       └── routes.go              # Definición de rutas con permisos
```

---

## 📊 Modelos de Datos

### 1. Niveles de Contexto

```go
const (
    ContextSystem    = 10  // Contexto del sistema completo
    ContextUser      = 30  // Contexto de usuario individual
    ContextCoursecat = 40  // Contexto de categoría de cursos
    ContextCourse    = 50  // Contexto de curso
    ContextModule    = 70  // Contexto de módulo/actividad
    ContextBlock     = 80  // Contexto de bloque
)
```

### 2. Niveles de Permiso

```go
const (
    CapInherit  = 0     // Heredar del contexto padre
    CapAllow    = 1     // Permitir explícitamente
    CapPrevent  = -1    // Prevenir (sobrescribe ALLOW del padre)
    CapProhibit = -1000 // Prohibir (no puede ser sobrescrito)
)
```

### 3. Estructuras Principales

#### Context
```go
type Context struct {
    ID           int64  `gorm:"column:id;primaryKey"`
    ContextLevel int    `gorm:"column:contextlevel"`
    InstanceID   int64  `gorm:"column:instanceid"`
    Path         string `gorm:"column:path"`       // Ej: "/1/3/45/"
    Depth        int    `gorm:"column:depth"`
    Locked       int    `gorm:"column:locked"`
}
```

#### Role
```go
type Role struct {
    ID          int64  `gorm:"column:id;primaryKey"`
    Name        string `gorm:"column:name"`
    ShortName   string `gorm:"column:shortname"`
    Archetype   string `gorm:"column:archetype"`  // manager, teacher, student, etc.
}
```

#### RoleAssignment
```go
type RoleAssignment struct {
    ID           int64  `gorm:"column:id;primaryKey"`
    RoleID       int64  `gorm:"column:roleid"`
    ContextID    int64  `gorm:"column:contextid"`
    UserID       int64  `gorm:"column:userid"`
}
```

#### RoleCapability
```go
type RoleCapability struct {
    ID         int64  `gorm:"column:id;primaryKey"`
    ContextID  int64  `gorm:"column:contextid"`
    RoleID     int64  `gorm:"column:roleid"`
    Capability string `gorm:"column:capability"`  // Ej: "moodle/course:update"
    Permission int    `gorm:"column:permission"`  // CapAllow, CapPrevent, etc.
}
```

---

## 🔧 Funciones Principales del Servicio

### HasCapability
Verifica si un usuario tiene una capacidad específica en un contexto.

```go
func (s *PermissionService) HasCapability(userID int64, capability string, contextID int64) (bool, error)
```

**Ejemplo:**
```go
hasPermission, err := permService.HasCapability(userID, "moodle/course:update", contextID)
if hasPermission {
    // Usuario puede actualizar el curso
}
```

### RequireCapability
Lanza un error si el usuario NO tiene la capacidad.

```go
func (s *PermissionService) RequireCapability(userID int64, capability string, contextID int64) error
```

**Ejemplo:**
```go
err := permService.RequireCapability(userID, "moodle/course:delete", contextID)
if err != nil {
    // Usuario NO tiene permiso
    return err
}
// Usuario tiene permiso, continuar
```

### HasCapabilityByCourse
Verifica una capacidad específicamente en el contexto de un curso.

```go
func (s *PermissionService) HasCapabilityByCourse(userID int64, capability string, courseID int64) (bool, error)
```

**Ejemplo:**
```go
canView, err := permService.HasCapabilityByCourse(userID, "moodle/course:view", courseID)
```

### IsSiteAdmin
Verifica si un usuario es administrador del sitio.

```go
func (s *PermissionService) IsSiteAdmin(userID int64) (bool, error)
```

**Ejemplo:**
```go
isAdmin, err := permService.IsSiteAdmin(userID)
if isAdmin {
    // Usuario es administrador
}
```

### Funciones de Conveniencia

```go
// Cursos
CanViewCourse(userID, courseID int64) (bool, error)
CanEditCourse(userID, courseID int64) (bool, error)
CanDeleteCourse(userID, courseID int64) (bool, error)
CanCreateCourse(userID int64) (bool, error)
CanBackupCourse(userID, courseID int64) (bool, error)

// Categorías
CanManageCategories(userID int64) (bool, error)

// Usuarios
CanEnrolUsers(userID, courseID int64) (bool, error)
CanViewGrades(userID, courseID int64) (bool, error)
```

---

## 🛡️ Middlewares para Gin

### 1. RequireCapability
Requiere una capacidad específica en el contexto del sistema.

```go
api.POST("/categories",
    middleware.RequireCapability(permService, "moodle/category:manage"),
    categoryHandler.CreateCategories)
```

### 2. RequireCapabilityInCourse
Requiere una capacidad en un curso específico. El courseID puede venir de:
- Query param: `?courseId=123`
- Path param: `/courses/:courseid`
- Path param: `/courses/:id`

```go
api.GET("/enrollments/course/:courseid",
    middleware.RequireCapabilityInCourse(permService, "moodle/course:viewparticipants"),
    userHandler.GetEnrolledUsers)
```

### 3. RequireAdmin
Requiere que el usuario sea administrador del sitio.

```go
api.DELETE("/users",
    middleware.RequireAdmin(permService),
    userHandler.DeleteUsers)
```

### 4. Middlewares de Conveniencia

```go
// Verificar permisos de curso
middleware.CanViewCourse(permService)
middleware.CanEditCourse(permService)
middleware.CanDeleteCourse(permService)

// Verificar permisos de categoría
middleware.CanManageCategories(permService)

// Verificar permisos de inscripción
middleware.CanEnrolUsers(permService)
```

**Ejemplo de uso:**
```go
api.PUT("/courses",
    middleware.CanEditCourse(permService),
    courseHandler.UpdateCourses)
```

---

## 🚀 Ejemplos de Uso Completos

### Ejemplo 1: Proteger una ruta de curso

```go
// En routes/routes.go
api.DELETE("/courses/:id",
    middleware.RequireCapabilityInCourse(permService, "moodle/course:delete"),
    courseHandler.DeleteCourse)
```

**Flujo:**
1. Usuario hace DELETE `/api/courses/123`
2. Middleware obtiene `userID` del contexto (debe estar seteado por middleware de autenticación)
3. Middleware obtiene `courseID` del path param (`:id`)
4. Middleware verifica: `permService.RequireCapabilityByCourse(userID, "moodle/course:delete", 123)`
5. Si no tiene permiso → `403 Forbidden`
6. Si tiene permiso → ejecuta `courseHandler.DeleteCourse`

### Ejemplo 2: Verificar permisos en un handler

```go
// En un handler personalizado
func (h *CourseHandler) UpdateCourse(c *gin.Context) {
    // Obtener userID del contexto
    userID, _ := middleware.GetUserID(c)

    // Obtener courseID del request
    var req UpdateCourseRequest
    c.BindJSON(&req)

    // Obtener servicio de permisos
    permService, _ := middleware.GetPermissionService(c)

    // Verificar permiso manualmente
    canUpdate, err := permService.CanEditCourse(userID, req.CourseID)
    if err != nil || !canUpdate {
        c.JSON(403, gin.H{"error": "No tienes permiso para actualizar este curso"})
        return
    }

    // Continuar con la lógica...
}
```

### Ejemplo 3: Verificar múltiples capacidades

```go
// Verificar si el usuario puede editar O eliminar cursos
capabilities := []string{"moodle/course:update", "moodle/course:delete"}

hasAny, err := permService.HasAnyCapability(userID, capabilities, contextID)
if hasAny {
    // Usuario puede hacer al menos una de las dos acciones
}

// Verificar si el usuario puede AMBAS cosas
hasAll, err := permService.HasAllCapabilities(userID, capabilities, contextID)
if hasAll {
    // Usuario puede hacer ambas acciones
}
```

### Ejemplo 4: Asignar y Desasignar Roles

```go
// Asignar rol de "teacher" (rol ID 3) a un usuario en un curso
courseCtx, _ := permService.GetContextByLevelAndInstance(models.ContextCourse, courseID)
err := permService.AssignRole(3, userID, courseCtx.ID, "", 0)

// Desasignar rol
err = permService.UnassignRole(3, userID, courseCtx.ID)
```

---

## 🔑 Capacidades Comunes de Moodle

### Cursos
- `moodle/course:view` - Ver curso
- `moodle/course:update` - Actualizar curso
- `moodle/course:delete` - Eliminar curso
- `moodle/course:create` - Crear curso
- `moodle/course:viewparticipants` - Ver participantes
- `moodle/backup:backupcourse` - Hacer backup de curso

### Categorías
- `moodle/category:manage` - Gestionar categorías
- `moodle/category:viewcourselist` - Ver lista de cursos en categoría

### Usuarios
- `moodle/user:viewalldetails` - Ver detalles de todos los usuarios
- `moodle/user:update` - Actualizar usuarios
- `moodle/user:delete` - Eliminar usuarios
- `enrol/manual:enrol` - Inscribir usuarios manualmente

### Calificaciones
- `moodle/grade:viewall` - Ver todas las calificaciones
- `moodle/grade:edit` - Editar calificaciones

---

## ⚙️ Configuración e Integración

### 1. Prerequisito: Middleware de Autenticación

**IMPORTANTE:** Antes de usar los middlewares de permisos, debes tener un middleware que establezca el `userID` en el contexto de Gin.

**Ejemplo de middleware de autenticación:**

```go
// internal/middleware/auth.go
func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Obtener token/session del header
        token := c.GetHeader("Authorization")

        // Validar token y obtener userID
        userID, err := validateToken(token, db)
        if err != nil {
            c.JSON(401, gin.H{"error": "No autenticado"})
            c.Abort()
            return
        }

        // CRÍTICO: Setear userID en el contexto
        c.Set("userID", userID)
        c.Next()
    }
}
```

### 2. Integrar en el Router Principal

```go
// internal/server/server.go
func New() *Server {
    // ... configuración previa ...

    router := gin.Default()
    router.Use(middleware.EnableCORS())

    // Aplicar middleware de autenticación PRIMERO
    router.Use(middleware.AuthMiddleware(db))

    // Registrar rutas (que incluyen middlewares de permisos)
    routes.RegisterRoutes(router, db)

    return &Server{router: router, db: db, cfg: cfg}
}
```

### 3. Estructura Completa de Middlewares en Rutas

```go
// internal/routes/routes.go
func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
    // Inicializar servicios
    permRepo := repository.NewPermissionRepository(db)
    permService := services.NewPermissionService(permRepo)

    // Grupo API con middleware de permisos
    api := router.Group("/api")
    api.Use(middleware.PermissionMiddleware(permService))

    // Definir rutas con middlewares específicos
    api.POST("/courses",
        middleware.RequireCapability(permService, "moodle/course:create"),
        courseHandler.CreateCourse)
}
```

---

## 🧪 Testing

### Test de Permisos

```go
func TestHasCapability(t *testing.T) {
    // Setup
    db := setupTestDB()
    permRepo := repository.NewPermissionRepository(db)
    permService := services.NewPermissionService(permRepo)

    userID := int64(2)  // Usuario de prueba
    courseID := int64(5) // Curso de prueba

    // Test
    canView, err := permService.HasCapabilityByCourse(userID, "moodle/course:view", courseID)

    assert.NoError(t, err)
    assert.True(t, canView)
}
```

---

## 📝 Notas Importantes

### 1. Jerarquía de Contextos

Los contextos son jerárquicos:
```
Sistema (/)
  └── Categoría (/1/)
      └── Curso (/1/3/)
          └── Módulo (/1/3/45/)
```

Un permiso en un contexto padre puede afectar a los hijos, dependiendo de la configuración.

### 2. Prioridad de Permisos

La lógica de Moodle aplica los permisos en este orden de prioridad:

1. **PROHIBIT (-1000)** → Siempre gana, no puede ser sobrescrito
2. **PREVENT (-1)** → Sobrescribe ALLOW de contextos padre
3. **ALLOW (1)** → Permite la acción
4. **INHERIT (0)** → Hereda del contexto padre

### 3. Administradores

Los usuarios con rol `archetype='manager'` en el contexto del sistema son considerados administradores y tienen **TODOS** los permisos automáticamente.

### 4. Performance

El sistema implementa consultas optimizadas con GORM usando:
- JOINs para obtener datos relacionados
- Filtrado por path para jerarquías de contextos
- Caching (por implementar si es necesario)

---

## 🐛 Troubleshooting

### Error: "Usuario no autenticado"

**Causa:** El middleware de permisos no encuentra `userID` en el contexto.

**Solución:** Asegúrate de que el middleware de autenticación esté ejecutándose ANTES y que setee correctamente:
```go
c.Set("userID", userID)
```

### Error: "No se proporcionó courseId"

**Causa:** El middleware `RequireCapabilityInCourse` no encuentra el courseID.

**Solución:** Asegúrate de que el courseID esté presente en:
- Query param: `?courseId=123`
- Path param: `:courseid` o `:id`

### Error: "Permisos insuficientes"

**Causa:** El usuario no tiene la capacidad requerida.

**Solución:** Verifica en la BD de Moodle:
1. Que el usuario tenga un rol asignado (`mdl_role_assignments`)
2. Que ese rol tenga la capacidad (`mdl_role_capabilities`)
3. Que el contexto sea correcto

**Query de diagnóstico:**
```sql
SELECT ra.userid, r.shortname, rc.capability, rc.permission, c.path
FROM mdl_role_assignments ra
JOIN mdl_role r ON ra.roleid = r.id
JOIN mdl_role_capabilities rc ON ra.roleid = rc.roleid
JOIN mdl_context c ON ra.contextid = c.id
WHERE ra.userid = 2;  -- ID del usuario
```

---

## 📚 Recursos Adicionales

- [Documentación oficial de permisos de Moodle](https://docs.moodle.org/en/Capabilities)
- [Roles y permisos en Moodle](https://docs.moodle.org/en/Roles_and_permissions)
- [Contextos en Moodle](https://docs.moodle.org/dev/Context)

---

## ✅ Checklist de Implementación

- [x] Modelos de datos (Context, Role, RoleAssignment, RoleCapability)
- [x] Repository con lógica de consultas a BD
- [x] Service con funciones de negocio (HasCapability, RequireCapability, etc.)
- [x] Middlewares para Gin (RequireCapability, RequireAdmin, etc.)
- [x] Integración en routes.go
- [ ] Middleware de autenticación (prerequisito)
- [ ] Testing unitario
- [ ] Testing de integración
- [ ] Documentación de capacidades específicas del proyecto
- [ ] Optimización con caché (opcional)

---

**Fecha:** 2025-01-28
**Versión:** 1.0
**Autor:** Sistema de Permisos Zajuna API
