# DTOs (Data Transfer Objects)

Esta capa contiene todos los objetos de transferencia de datos utilizados en la API.

## Estructura

```
dto/
├── request/           # DTOs para peticiones HTTP
│   ├── user_request.go
│   ├── course_request.go
│   └── category_request.go
├── response/          # DTOs para respuestas HTTP
│   ├── common.go
│   ├── user_response.go
│   ├── course_response.go
│   └── category_response.go
└── mapper/            # Conversores entre modelos y DTOs
    ├── user_mapper.go
    ├── course_mapper.go
    └── category_mapper.go
```

## Request DTOs

Los DTOs de request incluyen:
- **Validación automática** con tags de Gin/validator
- **Valores por defecto** con método `SetDefaults()`
- **Métodos helper** para conversión de datos
- **Validaciones de negocio** con método `Validate()`

### Ejemplo: GetUsersRequest

```go
type GetUsersRequest struct {
    Firstname string `form:"firstname" binding:"omitempty,min=2,max=100"`
    Lastname  string `form:"lastname" binding:"omitempty,min=2,max=100"`
    Username  string `form:"username" binding:"omitempty,min=2,max=100"`
    Email     string `form:"email" binding:"omitempty,email"`
    Page      int    `form:"page" binding:"omitempty,min=1"`
    Limit     int    `form:"limit" binding:"omitempty,min=1,max=100"`
}
```

**Uso en handler:**

```go
func (h *UserHandler) GetUsers(c *gin.Context) {
    var req request.GetUsersRequest

    // Bind query parameters con validación automática
    if err := c.ShouldBindQuery(&req); err != nil {
        c.JSON(400, response.NewErrorResponse("INVALID_INPUT", "Parámetros inválidos", err.Error()))
        return
    }

    // Establecer valores por defecto
    req.SetDefaults()

    // Convertir a filtros para el repository
    filters := req.ToFilterMap()

    // Llamar al servicio
    users, total, err := h.service.GetUsers(filters, req.Page, req.Limit)
    // ...
}
```

## Response DTOs

Los DTOs de response proporcionan:
- **Consistencia** en el formato de respuestas
- **Funciones helper** para crear respuestas
- **Paginación estandarizada**
- **Manejo de errores uniforme**

### Respuestas Comunes

#### ErrorResponse
```go
response.NewErrorResponse("NOT_FOUND", "Usuario no encontrado", nil)
```

**Salida JSON:**
```json
{
  "code": "NOT_FOUND",
  "message": "Usuario no encontrado",
  "timestamp": "2024-10-23T10:30:00Z"
}
```

#### SuccessResponse
```go
response.NewSuccessResponse("Usuario actualizado correctamente", userData)
```

**Salida JSON:**
```json
{
  "message": "Usuario actualizado correctamente",
  "data": { ... }
}
```

#### PaginatedResponse
```go
response.NewPaginatedResponse(users, 1, 15, 100)
```

**Salida JSON:**
```json
{
  "data": [ ... ],
  "pagination": {
    "page": 1,
    "limit": 15,
    "total": 100,
    "total_pages": 7,
    "has_next": true,
    "has_previous": false
  }
}
```

## Mappers

Los mappers convierten entre modelos de dominio y DTOs de respuesta.

### Ejemplo: UserMapper

```go
// Convertir un usuario
userResponse := mapper.UserToResponse(&user)

// Convertir lista de usuarios
usersResponse := mapper.UsersToResponse(users)
```

**Uso en handler:**
```go
func (h *UserHandler) GetUsers(c *gin.Context) {
    // ... obtener usuarios del servicio ...

    // Convertir modelos a DTOs
    usersResponse := mapper.UsersToResponse(users)

    // Crear respuesta paginada
    response := response.NewPaginatedResponse(usersResponse, req.Page, req.Limit, total)

    c.JSON(200, response)
}
```

## Validación

### Validación Automática (Gin)

Los DTOs usan tags de validación de Gin:

```go
type UpdateUserRequest struct {
    ID        uint   `json:"id" binding:"required,min=1"`
    FirstName string `json:"firstname" binding:"required,min=2,max=100"`
    Email     string `json:"email" binding:"required,email"`
}
```

### Validación Personalizada

Para validaciones más complejas, implementa el método `Validate()`:

```go
func (r *DeleteCoursesRequest) Validate() error {
    seen := make(map[int]bool)
    for _, id := range r.CourseIDs {
        if id == 1 {
            return &ValidationError{
                Field:   "courseids",
                Message: "No se puede eliminar el curso site (ID=1)",
            }
        }
        if seen[id] {
            return &ValidationError{
                Field:   "courseids",
                Message: "IDs duplicados detectados",
            }
        }
        seen[id] = true
    }
    return nil
}
```

**Uso en handler:**
```go
var req request.DeleteCoursesRequest
if err := c.ShouldBindJSON(&req); err != nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
}

// Validación adicional
if err := req.Validate(); err != nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
}
```

## Tags de Validación Disponibles

| Tag | Descripción | Ejemplo |
|-----|-------------|---------|
| `required` | Campo obligatorio | `binding:"required"` |
| `min` | Valor/longitud mínima | `binding:"min=2"` |
| `max` | Valor/longitud máxima | `binding:"max=100"` |
| `len` | Longitud exacta | `binding:"len=2"` |
| `email` | Email válido | `binding:"email"` |
| `oneof` | Uno de los valores | `binding:"oneof=0 1"` |
| `omitempty` | Opcional | `binding:"omitempty,min=2"` |
| `dive` | Validar array elements | `binding:"required,dive,min=1"` |

## Validaciones Personalizadas

El validador en `internal/validator/validator.go` incluye:

- `moodle_username` - Username válido de Moodle (alfanuméricos, -, _, ., @)
- `moodle_shortname` - Shortname válido (sin espacios)

**Uso:**
```go
type CreateUserRequest struct {
    Username string `json:"username" binding:"required,moodle_username"`
}
```

## Mejores Prácticas

1. **Siempre usa DTOs en handlers** - Nunca expongas modelos de dominio directamente
2. **Valida en el DTO primero** - Usa tags de binding para validación básica
3. **Validaciones de negocio en Validate()** - Para lógica más compleja
4. **Usa mappers** - No conviertas manualmente en los handlers
5. **Mantén DTOs simples** - Solo datos, sin lógica de negocio
6. **Documenta validaciones** - Comenta reglas de negocio complejas

## Ejemplo Completo

```go
// Handler
func (h *UserHandler) UpdateUser(c *gin.Context) {
    var req request.UpdateUserRequest

    // 1. Bind y validación automática
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, response.NewErrorResponse(
            "INVALID_INPUT",
            "Datos de entrada inválidos",
            err.Error(),
        ))
        return
    }

    // 2. Validación de negocio adicional (si existe)
    if err := req.Validate(); err != nil {
        c.JSON(400, response.NewErrorResponse(
            "VALIDATION_ERROR",
            err.Error(),
            nil,
        ))
        return
    }

    // 3. Llamar al servicio
    updated, err := h.service.UpdateUser(req.ID, req)
    if err != nil {
        c.JSON(500, response.NewErrorResponse(
            "UPDATE_FAILED",
            "Error al actualizar usuario",
            err.Error(),
        ))
        return
    }

    // 4. Responder con DTO
    c.JSON(200, response.NewSuccessResponse(
        "Usuario actualizado correctamente",
        updated,
    ))
}
```
