# Testing del Sistema de Permisos

## 🧪 Guía de Pruebas

### 1. Obtener un Token de Sesión (SID) de Moodle

Primero necesitas obtener un Session ID (SID) válido de Moodle. Hay dos formas:

#### Opción A: Desde el Navegador (Desarrollo)
1. Inicia sesión en Moodle
2. Abre las herramientas de desarrollo (F12)
3. Ve a la pestaña "Application" o "Storage"
4. Busca las cookies
5. Encuentra la cookie `MoodleSession` - ese es tu SID

#### Opción B: Desde la Base de Datos (Desarrollo)
```sql
-- Obtener sesiones activas recientes
SELECT sid, userid, timemodified, state
FROM mdl_sessions
WHERE state = 0
ORDER BY timemodified DESC
LIMIT 10;
```

---

## 📡 Ejemplos de Peticiones HTTP

### 2. Probar Autenticación Básica

#### ✅ Con Token Válido
```bash
curl -X GET "http://localhost:8080/api/users?page=1&limit=25" \
  -H "Authorization: Bearer TU_SESSION_ID_AQUI"
```

**Respuesta esperada (si tienes permisos):**
```json
{
  "users": [...],
  "total": 50,
  "currentPage": 1,
  ...
}
```

**Respuesta esperada (si NO tienes permisos):**
```json
{
  "error": "Permisos insuficientes",
  "details": "required capability 'moodle/user:viewalldetails' not satisfied for user X in context Y"
}
```

#### ❌ Sin Token
```bash
curl -X GET "http://localhost:8080/api/users?page=1&limit=25"
```

**Respuesta esperada:**
```json
{
  "error": "Missing Authorization header"
}
```

#### ❌ Token Inválido
```bash
curl -X GET "http://localhost:8080/api/users?page=1&limit=25" \
  -H "Authorization: Bearer token_invalido_12345"
```

**Respuesta esperada:**
```json
{
  "error": "Invalid or expired session"
}
```

---

## 🔐 Pruebas de Permisos por Endpoint

### 3. GET /api/categories
**Requiere:** `moodle/category:viewcourselist`

```bash
curl -X GET "http://localhost:8080/api/categories" \
  -H "Authorization: Bearer TU_SID"
```

**Usuarios que pueden acceder:**
- Administradores
- Teachers (editingteacher)
- Managers
- Usuarios con el rol que tenga la capacidad `moodle/category:viewcourselist`

---

### 4. POST /api/categories
**Requiere:** `moodle/category:manage`

```bash
curl -X POST "http://localhost:8080/api/categories" \
  -H "Authorization: Bearer TU_SID" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Nueva Categoría",
    "description": "Descripción de prueba"
  }'
```

**Usuarios que pueden acceder:**
- Administradores
- Managers
- Usuarios con capacidad de gestionar categorías

---

### 5. DELETE /api/courses
**Requiere:** `moodle/course:delete`

```bash
curl -X DELETE "http://localhost:8080/api/courses?ids=123,456" \
  -H "Authorization: Bearer TU_SID"
```

**Usuarios que pueden acceder:**
- Administradores
- Teachers con permiso de eliminar cursos
- Managers

---

### 6. DELETE /api/users (Solo Admins)
**Requiere:** Ser administrador del sitio

```bash
curl -X DELETE "http://localhost:8080/api/users?ids=10,20" \
  -H "Authorization: Bearer TU_SID"
```

**Usuarios que pueden acceder:**
- **SOLO** administradores del sitio (usuarios con rol `manager` en contexto de sistema)

**Respuesta si no eres admin:**
```json
{
  "error": "Se requieren permisos de administrador"
}
```

---

### 7. GET /api/enrollments/course/:courseid
**Requiere:** `moodle/course:viewparticipants` en el curso específico

```bash
curl -X GET "http://localhost:8080/api/enrollments/course/5" \
  -H "Authorization: Bearer TU_SID"
```

**Usuarios que pueden acceder:**
- Administradores
- Teachers del curso específico (ID 5)
- Usuarios inscritos con rol que permita ver participantes

---

## 🧪 Casos de Prueba Específicos

### 8. Probar Herencia de Permisos

Los permisos se heredan jerárquicamente. Por ejemplo:

```
Sistema (contextlevel=10, path="/1/")
  └── Categoría (contextlevel=40, path="/1/3/")
      └── Curso (contextlevel=50, path="/1/3/45/")
```

Si un usuario tiene `moodle/course:view` con `CAP_ALLOW` en la categoría (contexto `/1/3/`), también puede ver todos los cursos dentro de esa categoría.

**Prueba:**
```sql
-- Asignar rol de teacher en una categoría
INSERT INTO mdl_role_assignments (roleid, contextid, userid, timemodified)
VALUES (3, <context_id_categoria>, <tu_userid>, EXTRACT(EPOCH FROM NOW())::bigint);
```

Luego probar acceso a cursos de esa categoría.

---

### 9. Probar PROHIBIT (No Sobrescribible)

El nivel `CAP_PROHIBIT (-1000)` no puede ser sobrescrito en contextos hijo.

**Ejemplo:**
```sql
-- Prohibir editar cursos en contexto de sistema (nadie puede, ni siquiera en cursos específicos)
INSERT INTO mdl_role_capabilities (contextid, roleid, capability, permission)
VALUES (<system_context_id>, <roleid>, 'moodle/course:update', -1000);
```

Incluso si tienes `CAP_ALLOW` en un curso específico, el `PROHIBIT` del sistema gana.

---

### 10. Probar PREVENT vs ALLOW

`CAP_PREVENT (-1)` sobrescribe `CAP_ALLOW` del padre.

**Ejemplo:**
```sql
-- En categoría: ALLOW editar cursos
INSERT INTO mdl_role_capabilities (contextid, roleid, capability, permission)
VALUES (<category_context>, 3, 'moodle/course:update', 1);

-- En curso específico: PREVENT editar
INSERT INTO mdl_role_capabilities (contextid, roleid, capability, permission)
VALUES (<course_context>, 3, 'moodle/course:update', -1);
```

El usuario podrá editar cursos de la categoría, **excepto** ese curso específico.

---

## 🛠️ Scripts de Testing

### Script Bash para Testing Automatizado

```bash
#!/bin/bash

# Configuración
API_URL="http://localhost:8080"
SID="TU_SESSION_ID_AQUI"

echo "🧪 Testing Sistema de Permisos Zajuna API"
echo "=========================================="

# Test 1: Sin autenticación
echo "Test 1: GET /api/users sin token"
curl -s -X GET "$API_URL/api/users?page=1&limit=5" | jq

# Test 2: Con autenticación
echo -e "\nTest 2: GET /api/users con token"
curl -s -X GET "$API_URL/api/users?page=1&limit=5" \
  -H "Authorization: Bearer $SID" | jq

# Test 3: Categorías
echo -e "\nTest 3: GET /api/categories"
curl -s -X GET "$API_URL/api/categories" \
  -H "Authorization: Bearer $SID" | jq

# Test 4: Endpoint de admin (debería fallar si no eres admin)
echo -e "\nTest 4: DELETE /api/users (requiere admin)"
curl -s -X DELETE "$API_URL/api/users?ids=999" \
  -H "Authorization: Bearer $SID" | jq

echo -e "\n✅ Tests completados"
```

Guarda como `test_permisos.sh` y ejecuta:
```bash
chmod +x test_permisos.sh
./test_permisos.sh
```

---

## 📊 Verificar Permisos de un Usuario

### Query SQL Útil

```sql
-- Ver todos los permisos de un usuario
SELECT
    u.id AS user_id,
    u.username,
    r.shortname AS role,
    rc.capability,
    rc.permission,
    c.contextlevel,
    c.path,
    CASE rc.permission
        WHEN 1 THEN 'ALLOW'
        WHEN -1 THEN 'PREVENT'
        WHEN -1000 THEN 'PROHIBIT'
        ELSE 'INHERIT'
    END AS permission_name
FROM mdl_user u
JOIN mdl_role_assignments ra ON u.id = ra.userid
JOIN mdl_role r ON ra.roleid = r.id
JOIN mdl_context c ON ra.contextid = c.id
LEFT JOIN mdl_role_capabilities rc ON r.id = rc.roleid AND c.id = rc.contextid
WHERE u.id = <TU_USER_ID>
ORDER BY c.path, rc.capability;
```

---

## 🐛 Debugging

### Activar Logs Detallados en Gin

```go
// En server.go, antes de router.Run()
gin.SetMode(gin.DebugMode)
```

### Ver Sesiones Activas

```sql
SELECT
    s.sid,
    s.userid,
    u.username,
    FROM_TIMESTAMP(s.timecreated) AS created,
    FROM_TIMESTAMP(s.timemodified) AS last_activity,
    s.state
FROM mdl_sessions s
JOIN mdl_user u ON s.userid = u.id
WHERE s.state = 0
ORDER BY s.timemodified DESC
LIMIT 20;
```

### Verificar Contextos

```sql
-- Ver estructura de contextos
SELECT
    id,
    contextlevel,
    CASE contextlevel
        WHEN 10 THEN 'SYSTEM'
        WHEN 30 THEN 'USER'
        WHEN 40 THEN 'CATEGORY'
        WHEN 50 THEN 'COURSE'
        WHEN 70 THEN 'MODULE'
        WHEN 80 THEN 'BLOCK'
    END AS level_name,
    instanceid,
    path,
    depth
FROM mdl_context
ORDER BY path;
```

---

## ✅ Checklist de Testing

- [ ] Autenticación funciona con token válido
- [ ] Autenticación rechaza token inválido
- [ ] Autenticación rechaza sesión expirada
- [ ] GET /api/users requiere permisos correctos
- [ ] GET /api/categories requiere permisos correctos
- [ ] DELETE /api/users solo permite admins
- [ ] POST /api/categories requiere gestionar categorías
- [ ] GET /api/enrollments/:courseid valida permisos por curso
- [ ] Administradores tienen acceso a todo
- [ ] Usuarios sin permisos reciben 403
- [ ] Herencia de permisos funciona correctamente
- [ ] PROHIBIT no puede ser sobrescrito
- [ ] PREVENT sobrescribe ALLOW del padre

---

## 📝 Notas Importantes

1. **Timeout de Sesión:** Por defecto 2 horas (7200 segundos). Puedes ajustarlo en `auth.go`.

2. **Formato del Token:** Siempre usar `Bearer <sid>` en el header `Authorization`.

3. **Permisos Jerárquicos:** Los permisos se calculan desde el contexto más específico al más general.

4. **Caché:** Moodle cachea permisos. Si modificas permisos en BD, puede tardar en reflejarse.

5. **Testing en Producción:** NUNCA uses session IDs reales de producción en desarrollo.

---

## 🚀 Próximos Pasos

1. Implementar tests unitarios para el repository
2. Implementar tests de integración para los middlewares
3. Agregar métricas y logging de permisos denegados
4. Implementar caché de permisos para optimizar performance
5. Crear herramienta CLI para gestionar roles desde zajuna-api

---

**Documentación Relacionada:**
- `SISTEMA_PERMISOS_MOODLE.md` - Guía completa del sistema
- `CAMBIOS_SISTEMA_PERMISOS.md` - Registro de cambios
