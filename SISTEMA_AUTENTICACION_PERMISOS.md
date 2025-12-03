# Sistema de Autenticación y Permisos - Zajuna

## Descripción General

Este documento describe el sistema completo de autenticación y validación de permisos implementado en el proyecto Zajuna. El sistema replica la arquitectura de permisos de **Moodle**, validando que solo usuarios con roles administrativos puedan acceder al Dashboard.

---

## Flujo de Autenticación Completo

```
┌─────────────────────────────────────────────────────────────────────┐
│                   FLUJO DE AUTENTICACIÓN CON PERMISOS               │
└─────────────────────────────────────────────────────────────────────┘

1. Usuario ingresa credenciales en Landing Page
   └─ POST http://localhost:5173/login
      ├─ username: "admin" | "123456789" (documento)
      └─ password: "******"

2. Landing envía credenciales a API
   └─ POST http://localhost:8080/api/login
      └─ Body: { "username": "admin", "password": "******" }

3. API valida credenciales (zajuna-api/internal/handlers/auth_handler.go)
   ├─ Busca usuario en mdl_user por username O idnumber
   ├─ Verifica que usuario NO esté: deleted, suspended, sin confirmar
   ├─ Compara password hasheado con bcrypt
   └─ Si credenciales son válidas → continúa

4. API verifica permisos del usuario (Sistema de Permisos de Moodle)
   ├─ IsSiteAdmin(userID) → ¿Es administrador del sitio?
   ├─ HasCapability("moodle/category:manage") → ¿Puede gestionar categorías?
   ├─ HasCapability("moodle/course:update") → ¿Puede gestionar cursos?
   └─ HasCapability("moodle/user:update") → ¿Puede gestionar usuarios?

5. API calcula canAccessDashboard
   └─ canAccessDashboard = isAdmin OR canManageCategories OR
                           canManageCourses OR canManageUsers

6. API responde con datos del usuario y permisos
   └─ Response:
      {
        "success": true,
        "token": "ZAJUNA_admin_20250101120000",
        "user": { "id": 2, "username": "admin", ... },
        "isAdmin": true,
        "canAccessDashboard": true
      }

7. Landing procesa respuesta
   ├─ Almacena token en localStorage
   ├─ Almacena user en localStorage
   └─ Evalúa canAccessDashboard:
      ├─ SI canAccessDashboard = true
      │  └─ Redirige a Dashboard: http://localhost:3000
      └─ SI canAccessDashboard = false
         └─ Muestra mensaje: "No tienes permisos administrativos"
```

---

## Arquitectura del Sistema

### Componentes Implementados

```
zajuna-api/
├── internal/
│   ├── handlers/
│   │   └── auth_handler.go              # Handler de login con validación de permisos
│   ├── middleware/
│   │   ├── auth.go                      # Middleware de autenticación JWT/Session
│   │   └── permission.go                # Middlewares de permisos (Moodle)
│   ├── models/
│   │   ├── user.go                      # Modelo de usuario
│   │   ├── context.go                   # Contextos de Moodle
│   │   ├── moodle_role.go               # Roles de Moodle
│   │   ├── role_assignment.go           # Asignaciones de roles
│   │   ├── role_capability.go           # Capacidades de roles
│   │   └── capability.go                # Definición de capacidades
│   ├── repository/
│   │   ├── user_repository.go           # Acceso a datos de usuarios
│   │   └── permission_repository.go     # Acceso a datos de permisos
│   ├── services/
│   │   └── permission_service.go        # Lógica de negocio de permisos
│   └── routes/
│       └── routes.go                    # Registro de rutas (incluye /api/login)

Landing-Lms-Zajuna/
├── src/
│   ├── components/
│   │   └── forms/
│   │       └── LoginForm.jsx            # Formulario de login con validación
│   └── utils/
│       └── authClient.js                # Cliente de autenticación API
```

---

## Sistema de Permisos de Moodle

### Niveles de Contexto

El sistema de permisos de Moodle usa contextos jerárquicos:

| Nivel | Constante        | Descripción         |
| ----- | ---------------- | ------------------- |
| 10    | ContextSystem    | Sistema completo    |
| 30    | ContextUser      | Usuario individual  |
| 40    | ContextCoursecat | Categoría de cursos |
| 50    | ContextCourse    | Curso específico    |
| 70    | ContextModule    | Módulo/actividad    |
| 80    | ContextBlock     | Bloque              |

### Capacidades Verificadas para Dashboard

El usuario puede acceder al Dashboard si tiene **AL MENOS UNA** de estas capacidades en el **contexto del sistema**:

1. **Es Site Admin** (`IsSiteAdmin`)

   - Usuario con rol `archetype='manager'` en contexto del sistema
   - Tiene TODOS los permisos automáticamente

2. **Puede gestionar categorías** (`moodle/category:manage`)

   - Crear, editar, eliminar categorías de cursos

3. **Puede gestionar cursos** (`moodle/course:update`)

   - Editar cursos existentes

4. **Puede gestionar usuarios** (`moodle/user:update`)
   - Actualizar datos de usuarios

### Lógica de Validación

```go
// En zajuna-api/internal/handlers/auth_handler.go:104-130

isAdmin, _ := h.permService.IsSiteAdmin(int(user.ID))

systemCtx, _ := h.permService.GetSystemContext()

canManageCategories, _ := h.permService.HasCapability(
    int(user.ID),
    "moodle/category:manage",
    systemCtx.ID
)

canManageCourses, _ := h.permService.HasCapability(
    int(user.ID),
    "moodle/course:update",
    systemCtx.ID
)

canManageUsers, _ := h.permService.HasCapability(
    int(user.ID),
    "moodle/user:update",
    systemCtx.ID
)

canAccessDashboard := isAdmin ||
                      canManageCategories ||
                      canManageCourses ||
                      canManageUsers
```

---

## Endpoint de Login

### POST /api/login

Endpoint público (NO requiere autenticación previa).

#### Request

```json
{
  "username": "admin", // O número de documento (idnumber)
  "password": "MyPassword123"
}
```

#### Response (Éxito)

```json
{
  "success": true,
  "token": "ZAJUNA_admin_20250101120000",
  "user": {
    "id": 2,
    "username": "admin",
    "firstname": "Juan",
    "lastname": "Pérez",
    "email": "admin@zajuna.com",
    "idnumber": ""
  },
  "isAdmin": true,
  "canAccessDashboard": true
}
```

#### Response (Usuario sin permisos)

```json
{
  "success": true,
  "token": "ZAJUNA_estudiante_20250101120000",
  "user": {
    "id": 123,
    "username": "123456789",
    "firstname": "María",
    "lastname": "González",
    "email": "maria@estudiante.com",
    "idnumber": "123456789"
  },
  "isAdmin": false,
  "canAccessDashboard": false
}
```

#### Response (Credenciales inválidas)

```json
{
  "success": false,
  "error": "Credenciales inválidas"
}
```

---

## Validación en Landing Page

### LoginForm.jsx (líneas 57-98)

El componente `LoginForm` valida la respuesta del backend:

```javascript
const processLogin = async (credentials) => {
  const response = await loginRequest(credentials);

  // Verificar si el usuario tiene permisos para acceder al dashboard
  const canAccessDashboard = response.canAccessDashboard || false;
  const isAdmin = response.isAdmin || false;

  // Almacenar token y usuario
  if (response.token) {
    storeToken(response.token);
  }
  if (response.user) {
    storeUser(response.user);
  }

  // Solo redirigir al dashboard si el usuario tiene permisos
  if (canAccessDashboard) {
    const zajunaUrl =
      import.meta.env.VITE_ZAJUNA_DASHBOARD_URL || "http://localhost:3000";

    setTimeout(() => {
      console.log(
        "Usuario con permisos administrativos. Redirigiendo al dashboard:",
        zajunaUrl
      );
      window.location.href = zajunaUrl;
    }, 1500);
  } else {
    // Usuario sin permisos administrativos - mostrar mensaje
    setTimeout(() => {
      setSuccess(false);
      setError(
        "Login exitoso, pero no tienes permisos para acceder al panel administrativo. Accede a tus cursos desde Moodle."
      );
    }, 1500);
  }
};
```

---

## Middleware de Autenticación

### AuthMiddleware (zajuna-api/internal/middleware/auth.go)

Valida la sesión y setea el `userID` en el contexto de Gin:

```go
func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
  return func(c *gin.Context) {
    // Obtener token desde Header Authorization
    authHeader := c.GetHeader("Authorization")

    // Formato esperado: "Bearer <token>"
    parts := strings.Split(authHeader, " ")
    sid := parts[1]

    // Validar sesión en la BD de Moodle
    var session struct {
      UserID       int   `gorm:"column:userid"`
      TimeModified int64 `gorm:"column:timemodified"`
      State        int   `gorm:"column:state"`
    }

    err := db.Table("mdl_sessions").
      Where("sid = ? AND state = 0", sid).
      First(&session).Error

    // CRÍTICO: Setear userID en el contexto
    c.Set("userID", session.UserID)

    c.Next()
  }
}
```

---

## Middleware de Permisos

### RequireAdmin (zajuna-api/internal/middleware/permission.go:127-159)

Requiere que el usuario sea Site Admin:

```go
func RequireAdmin(permService *services.PermissionService) gin.HandlerFunc {
  return func(c *gin.Context) {
    userID, _ := c.Get("userID")

    isAdmin, err := permService.IsSiteAdmin(userID.(int))
    if err != nil || !isAdmin {
      c.JSON(http.StatusForbidden, gin.H{
        "error": "Se requieren permisos de administrador"
      })
      c.Abort()
      return
    }

    c.Next()
  }
}
```

### RequireCapability

Requiere una capacidad específica en el contexto del sistema:

```go
func RequireCapability(permService *services.PermissionService, capability string) gin.HandlerFunc {
  return func(c *gin.Context) {
    userID, _ := c.Get("userID")

    systemCtx, _ := permService.GetSystemContext()

    err := permService.RequireCapability(userID.(int), capability, systemCtx.ID)
    if err != nil {
      c.JSON(http.StatusForbidden, gin.H{
        "error":   "Permisos insuficientes",
        "details": err.Error(),
      })
      c.Abort()
      return
    }

    c.Next()
  }
}
```

---

## Configuración de Rutas

### routes.go (zajuna-api/internal/routes/routes.go)

```go
func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
  // Inicializar servicios
  permRepo := repository.NewPermissionRepository(db)
  permService := services.NewPermissionService(permRepo)
  authHandler := handlers.NewAuthHandler(userRepo, permService)

  // --- Rutas públicas (sin autenticación) ---
  router.POST("/api/login", authHandler.Login)

  // --- Rutas protegidas (requieren autenticación) ---
  api := router.Group("/api")
  api.Use(middleware.AuthMiddleware(db))
  api.Use(middleware.PermissionMiddleware(permService))

  // Rutas con permisos específicos
  api.POST("/categories",
    middleware.CanManageCategories(permService),
    categoryHandler.CreateCategories)

  api.DELETE("/users",
    middleware.RequireAdmin(permService),
    userHandler.DeleteUsers)
}
```

---

## Casos de Uso

### Caso 1: Administrador ingresa al sistema

1. Admin ingresa en Landing: http://localhost:5173
2. Formulario: "Ingreso Administrativos Zajuna"
   - Username: admin
   - Password: Admin123!
3. Landing envía a API: `POST /api/login`
4. API valida:
   - Credenciales correctas
   - Usuario activo (not deleted, not suspended, confirmed)
   - Es Site Admin (`IsSiteAdmin = true`)
5. API responde:
   ```json
   {
     "success": true,
     "token": "ZAJUNA_admin_20250101120000",
     "isAdmin": true,
     "canAccessDashboard": true
   }
   ```
6. Landing redirige a Dashboard: http://localhost:3000

### Caso 2: Manager de categorías ingresa

1. Manager ingresa en Landing
2. Formulario: "Ingreso Administrativos Zajuna"
   - Username: manager_cat
   - Password: Manager123!
3. API valida:
   - Credenciales correctas
   - No es Site Admin
   - Tiene capacidad `moodle/category:manage`
4. API responde:
   ```json
   {
     "success": true,
     "isAdmin": false,
     "canAccessDashboard": true
   }
   ```
5. Landing redirige a Dashboard

### Caso 3: Estudiante intenta ingresar

1. Estudiante ingresa en Landing
2. Formulario: "Ingreso cursos Zajuna"
   - Tipo documento: CC
   - Documento: 123456789
   - Password: Student123!
3. API valida:
   - Credenciales correctas
   - No es Site Admin
   - NO tiene capacidades administrativas
4. API responde:
   ```json
   {
     "success": true,
     "isAdmin": false,
     "canAccessDashboard": false
   }
   ```
5. Landing NO redirige - muestra mensaje:
   > "Login exitoso, pero no tienes permisos para acceder al panel administrativo. Accede a tus cursos desde Moodle."

---

## Tablas de Base de Datos (Moodle)

### mdl_user

Almacena información de usuarios.

| Campo     | Tipo    | Descripción                   |
| --------- | ------- | ----------------------------- |
| id        | INT     | ID del usuario                |
| username  | VARCHAR | Nombre de usuario             |
| password  | VARCHAR | Password hasheado (bcrypt)    |
| idnumber  | VARCHAR | Número de documento           |
| email     | VARCHAR | Email                         |
| deleted   | TINYINT | 0=activo, 1=eliminado         |
| suspended | TINYINT | 0=activo, 1=suspendido        |
| confirmed | TINYINT | 0=no confirmado, 1=confirmado |

### mdl_role

Define los roles disponibles en el sistema.

| Campo     | Tipo    | Descripción                           |
| --------- | ------- | ------------------------------------- |
| id        | INT     | ID del rol                            |
| name      | VARCHAR | Nombre del rol                        |
| shortname | VARCHAR | Nombre corto                          |
| archetype | VARCHAR | Tipo: manager, teacher, student, etc. |

### mdl_role_assignments

Asigna roles a usuarios en contextos específicos.

| Campo     | Tipo | Descripción         |
| --------- | ---- | ------------------- |
| id        | INT  | ID de la asignación |
| roleid    | INT  | ID del rol          |
| contextid | INT  | ID del contexto     |
| userid    | INT  | ID del usuario      |

### mdl_role_capabilities

Define las capacidades que tiene cada rol en cada contexto.

| Campo      | Tipo    | Descripción                         |
| ---------- | ------- | ----------------------------------- |
| id         | INT     | ID                                  |
| contextid  | INT     | ID del contexto                     |
| roleid     | INT     | ID del rol                          |
| capability | VARCHAR | Nombre de la capacidad              |
| permission | INT     | 1=allow, -1=prevent, -1000=prohibit |

### mdl_context

Define los contextos jerárquicos del sistema.

| Campo        | Tipo    | Descripción                                 |
| ------------ | ------- | ------------------------------------------- |
| id           | INT     | ID del contexto                             |
| contextlevel | INT     | Nivel: 10=system, 50=course, etc.           |
| instanceid   | INT     | ID de la instancia (courseID, userID, etc.) |
| path         | VARCHAR | Ruta jerárquica: /1/3/45/                   |
| depth        | INT     | Profundidad en la jerarquía                 |

---

## Seguridad

### Hashing de Contraseñas

- Moodle usa **bcrypt** para hashear contraseñas
- La API usa `golang.org/x/crypto/bcrypt` para validar
- Nunca se almacenan contraseñas en texto plano

```go
// Comparar password con bcrypt
err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(req.Password))
```

### Validación de Usuario Activo

Solo se permite login si el usuario cumple:

- `deleted = 0` (no eliminado)
- `suspended = 0` (no suspendido)
- `confirmed = 1` (email confirmado)

```go
err := h.userRepo.DB.Where("username = ? OR idnumber = ?", req.Username, req.Username).
  Where("deleted = 0").
  Where("suspended = 0").
  Where("confirmed = 1").
  First(&user).Error
```

### Tokens de Sesión

- Token simple por ahora: `ZAJUNA_{username}_{timestamp}`
- **TODO**: Implementar JWT con firma y expiración
- Los tokens se almacenan en `localStorage` en el frontend

---

## Testing

### Test Manual del Login

```bash
# 1. Iniciar el sistema
./start-dev.sh

# 2. Test con curl - Administrador
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "password": "Admin123!"
  }'

# Response esperado:
# {
#   "success": true,
#   "token": "ZAJUNA_admin_...",
#   "isAdmin": true,
#   "canAccessDashboard": true
# }

# 3. Test con curl - Estudiante
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "123456789",
    "password": "Student123!"
  }'

# Response esperado:
# {
#   "success": true,
#   "token": "ZAJUNA_123456789_...",
#   "isAdmin": false,
#   "canAccessDashboard": false
# }
```

---

## Troubleshooting

### Error: "Credenciales inválidas"

**Causas posibles:**

1. Username o password incorrectos
2. Usuario eliminado (`deleted = 1`)
3. Usuario suspendido (`suspended = 1`)
4. Usuario no confirmado (`confirmed = 0`)

**Solución:**

```sql
-- Verificar estado del usuario en Moodle
SELECT id, username, idnumber, deleted, suspended, confirmed
FROM mdl_user
WHERE username = 'admin';
```

### Error: "Login exitoso, pero no tienes permisos..."

**Causa:** Usuario no tiene roles administrativos en Moodle

**Solución:**

```sql
-- Verificar roles del usuario
SELECT ra.userid, r.shortname, r.archetype, c.contextlevel
FROM mdl_role_assignments ra
JOIN mdl_role r ON ra.roleid = r.id
JOIN mdl_context c ON ra.contextid = c.id
WHERE ra.userid = 123;  -- ID del usuario

-- Asignar rol de manager si es necesario
-- (Esto debe hacerse desde Moodle UI, no directamente en BD)
```

### Error: "Error al verificar permisos de administrador"

**Causa:** Problema con el sistema de permisos o BD

**Solución:**

1. Verificar que la BD de Moodle esté accesible
2. Verificar que las tablas `mdl_role_assignments`, `mdl_role`, `mdl_context` existan
3. Revisar logs del backend

---

## Próximos Pasos

### Mejoras Recomendadas

1. **Implementar JWT real**

   - Usar librería `github.com/golang-jwt/jwt/v5`
   - Agregar firma con secret key
   - Implementar expiración de tokens
   - Agregar refresh tokens

2. **Agregar rate limiting**

   - Proteger endpoint de login contra ataques de fuerza bruta
   - Limitar intentos de login por IP

3. **Implementar logout**

   - Invalidar tokens
   - Limpiar localStorage en frontend

4. **Mejorar mensajes de error**

   - No revelar si un username existe o no
   - Usar mensajes genéricos para mejorar seguridad

5. **Agregar auditoría**

   - Registrar intentos de login exitosos y fallidos
   - Almacenar en tabla `mdl_logstore_standard_log`

6. **Proteger rutas del Dashboard**
   - Implementar guard de rutas en React
   - Verificar token en cada petición
   - Redirigir a login si sesión expirada

---

## Referencias

- [Documentación de permisos de Moodle](https://docs.moodle.org/en/Capabilities)
- [Roles y permisos en Moodle](https://docs.moodle.org/en/Roles_and_permissions)
- [Contextos en Moodle](https://docs.moodle.org/dev/Context)
- [Bcrypt en Go](https://pkg.go.dev/golang.org/x/crypto/bcrypt)

---

**Fecha:** 2025-12-01
**Versión:** 1.0
**Autor:** Sistema de Autenticación Zajuna
