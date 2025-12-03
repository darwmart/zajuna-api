# Análisis del Sistema de Permisos de Moodle

## Resumen Ejecutivo

El sistema de permisos de Moodle se basa en el concepto de **capabilities** (capacidades) asignadas a **roles** (roles) en **contexts** (contextos) específicos. Este documento analiza la implementación del core de Moodle para replicarla en el proyecto Zajuna API.

## Conceptos Fundamentales

### 1. Capabilities (Capacidades)

- Son permisos atómicos con nombres como `moodle/course:view`, `moodle/category:manage`
- Se almacenan en la tabla `mdl_capabilities`
- Tienen niveles de riesgo (RISK_XSS, RISK_CONFIG, RISK_DATALOSS, etc.)
- Tienen tipos: 'read' o 'write'

### 2. Roles

- Son conjuntos de capabilities
- Tienen arquetipos (archetypes): 'manager', 'teacher', 'student', 'guest', etc.
- Se almacenan en `mdl_role`
- Cada rol puede tener capabilities con diferentes permisos en diferentes contextos

### 3. Contexts (Contextos)

Jerarquía de contextos (de más general a más específico):
- **CONTEXT_SYSTEM (10)**: Sistema completo
- **CONTEXT_USER (30)**: Usuario individual
- **CONTEXT_COURSECAT (40)**: Categoría de cursos
- **CONTEXT_COURSE (50)**: Curso
- **CONTEXT_MODULE (70)**: Módulo dentro de un curso
- **CONTEXT_BLOCK (80)**: Bloque

Cada contexto tiene un `path` que representa su jerarquía: `/1/3/45/` significa que el contexto 45 tiene como padres a 3 y 1.

### 4. Permissions (Permisos)

Valores posibles para una capability en un contexto:
- **CAP_INHERIT (0)**: Heredar del contexto padre
- **CAP_ALLOW (1)**: Permitir (sobrescribe PREVENT de contextos padre)
- **CAP_PREVENT (-1)**: Prevenir (sobrescribe ALLOW de contextos padre)
- **CAP_PROHIBIT (-1000)**: Prohibir (sobrescribe TODO, no se puede anular)

## Estructura de Datos

### Tablas principales:

```sql
-- Roles
mdl_role (id, name, shortname, description, sortorder, archetype)

-- Asignación de roles a usuarios en contextos
mdl_role_assignments (id, roleid, contextid, userid, timemodified, modifierid, component, itemid, sortorder)

-- Capabilities de cada rol en cada contexto
mdl_role_capabilities (id, contextid, roleid, capability, permission, timemodified, modifierid)

-- Contextos
mdl_context (id, contextlevel, instanceid, path, depth, locked)

-- Definición de capabilities
mdl_capabilities (id, name, captype, contextlevel, component, riskbitmask)
```

## Algoritmo de Verificación de Permisos

El core de Moodle usa la siguiente lógica en `has_capability()`:

### 1. Validaciones iniciales
```php
// Verificar si es instalación inicial
if (during_initial_install()) return true/false;

// Verificar que la capability existe
if (!$capinfo = get_capability_info($capability)) return false;

// Verificar que el usuario existe
if (!context_user::instance($userid, IGNORE_MISSING)) return false;

// Verificar contexto válido
if (empty($context->path) or $context->depth == 0) return false;
```

### 2. Verificar si es Site Admin
```php
if ($doanything && is_siteadmin($userid)) {
    // Los admins tienen todos los permisos (a menos que hayan hecho role switch)
    if (!has_switched_role_in_context($context)) {
        return true;
    }
}
```

### 3. Obtener accessdata del usuario
El `accessdata` es una estructura que contiene:
```php
$accessdata = [
    'ra' => [
        '/1/' => [5 => 5, 7 => 7],  // roleids por context path
        '/1/3/' => [8 => 8],
        '/1/3/45/' => [9 => 9]
    ],
    'rsw' => [],  // role switches
    'time' => 1234567890
];
```

### 4. Evaluar capability con has_capability_in_accessdata()

Pasos:
1. **Construir lista de paths** desde el contexto actual hasta el sistema:
   ```php
   // Si contexto es /1/3/45/, paths = ['/1/3/45/', '/1/3/', '/1/']
   ```

2. **Obtener todos los roles del usuario** en estos contextos:
   ```php
   // Recorrer paths y obtener roleids de $accessdata['ra'][path]
   ```

3. **Obtener definiciones de roles** (get_role_definitions):
   ```php
   // Cargar todas las capabilities de cada rol en todos los contextos
   $rdefs[$roleid][$path][$capability] = $permission;
   ```

4. **Evaluar permiso final**:
   ```php
   $allowed = false;
   foreach ($roles as $roleid) {
       foreach ($paths as $path) {  // De más específico a más general
           if (isset($rdefs[$roleid][$path][$capability])) {
               $perm = $rdefs[$roleid][$path][$capability];

               // PROHIBIT siempre gana
               if ($perm === CAP_PROHIBIT) return false;

               // PREVENT anula ALLOW previos
               if ($perm === CAP_PREVENT) {
                   $allowed = false;
               }

               // ALLOW solo cuenta si no hay PREVENT
               if ($perm === CAP_ALLOW && !hay_prevent_previo) {
                   $allowed = true;
               }
           }
       }
   }
   return $allowed;
   ```

## Diferencias con la Implementación Actual

### Problemas identificados:

1. **No se usa get_role_definitions**:
   - Actual: consulta directa por cada capability
   - Moodle: carga todas las capabilities de los roles y usa caché

2. **No se evalúa correctamente la jerarquía**:
   - Actual: itera contextos de profundo a superficial
   - Moodle: evalúa PROHIBIT primero en cualquier nivel, luego PREVENT/ALLOW

3. **Falta soporte para role switching**:
   - Moodle permite cambiar de rol temporalmente
   - Actual: no implementado

4. **No hay caché de accessdata**:
   - Moodle cachea el accessdata del usuario
   - Actual: consulta la BD en cada verificación

5. **No se validan capabilities peligrosas para guest**:
   - Moodle: guest nunca puede tener capabilities de 'write' o con riesgos
   - Actual: no valida esto

## Implementación Recomendada

### 1. Crear get_role_definitions en repository
```go
// GetRoleDefinitions obtiene todas las capabilities de una lista de roleIDs
func (r *PermissionRepository) GetRoleDefinitions(roleIDs []int) (map[int]map[string]map[string]int, error)
```

### 2. Implementar has_capability_in_accessdata
```go
// HasCapabilityInAccessData evalúa si un usuario tiene una capability usando accessdata
func (r *PermissionRepository) HasCapabilityInAccessData(
    capability string,
    context *models.Context,
    accessData *models.AccessData
) bool
```

### 3. Actualizar GetUserCapabilityInContext
```go
// Usar el algoritmo exacto de Moodle:
// 1. Verificar is_siteadmin
// 2. Construir accessdata
// 3. Llamar a HasCapabilityInAccessData
```

### 4. Agregar validaciones de seguridad
```go
// Validar que guest no tenga capabilities peligrosas
// Validar que el contexto tenga path y depth válidos
// Validar que la capability exista
```

## Beneficios de la Nueva Implementación

1. **Compatibilidad Total**: Funcionará exactamente como Moodle
2. **Performance**: Caché de role definitions reduce queries a BD
3. **Seguridad**: Validaciones adicionales para guest y capabilities peligrosas
4. **Mantenibilidad**: Código más claro que sigue el patrón de Moodle
5. **Escalabilidad**: Preparado para role switching y otras features avanzadas

## Referencias

- `/home/darwin/Escritorio/zajuna-nube/zajuna-nube/zajuna/lib/accesslib.php`
  - Líneas 432-582: `has_capability()`
  - Líneas 788-850: `has_capability_in_accessdata()`
  - Líneas 303-380: `get_role_definitions()`
  - Líneas 915-950: `get_user_roles_sitewide_accessdata()`

## Próximos Pasos

1. Implementar `GetRoleDefinitions()` en `permission_repository.go`
2. Implementar `HasCapabilityInAccessData()` en `permission_repository.go`
3. Refactorizar `GetUserCapabilityInContext()` para usar el nuevo algoritmo
4. Agregar struct `AccessData` en `models`
5. Agregar validaciones de seguridad
6. Implementar tests unitarios
7. Documentar el nuevo sistema
