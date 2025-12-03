-- Script SQL para verificar permisos de administrador
-- Reemplazar USUARIO_AQUI con el username o idnumber

-- 1. Buscar el usuario
SELECT '=== INFORMACIÓN DEL USUARIO ===' AS '';
SELECT
    id AS user_id,
    username,
    firstname,
    lastname,
    email,
    deleted,
    suspended,
    confirmed
FROM mdl_user
WHERE (username = 'USUARIO_AQUI' OR idnumber = 'USUARIO_AQUI')
LIMIT 1;

-- Guardar el ID del usuario para consultas siguientes
SET @user_id = (SELECT id FROM mdl_user WHERE (username = 'USUARIO_AQUI' OR idnumber = 'USUARIO_AQUI') LIMIT 1);

-- 2. Verificar si es Site Admin (tiene rol con archetype='manager' en contexto sistema)
SELECT '=== VERIFICANDO SI ES SITE ADMIN ===' AS '';
SELECT
    ra.id AS assignment_id,
    r.id AS role_id,
    r.shortname AS role_shortname,
    r.archetype,
    c.contextlevel
FROM mdl_role_assignments ra
JOIN mdl_role r ON ra.roleid = r.id
JOIN mdl_context c ON ra.contextid = c.id
WHERE ra.userid = @user_id
  AND r.archetype = 'manager'
  AND c.contextlevel = 10;

-- 3. Ver TODOS los roles del usuario
SELECT '=== TODOS LOS ROLES DEL USUARIO ===' AS '';
SELECT
    r.id AS role_id,
    r.shortname AS role_shortname,
    r.name AS role_name,
    r.archetype,
    c.contextlevel,
    c.path,
    c.instanceid
FROM mdl_role_assignments ra
JOIN mdl_role r ON ra.roleid = r.id
JOIN mdl_context c ON ra.contextid = c.id
WHERE ra.userid = @user_id
ORDER BY c.depth ASC;

-- 4. Ver capabilities críticas del usuario en contexto sistema
SELECT '=== CAPABILITIES EN CONTEXTO SISTEMA ===' AS '';
SELECT
    rc.capability,
    CASE rc.permission
        WHEN 1 THEN 'ALLOW'
        WHEN 0 THEN 'INHERIT'
        WHEN -1 THEN 'PREVENT'
        WHEN -1000 THEN 'PROHIBIT'
    END AS permission,
    r.shortname AS role_shortname
FROM mdl_role_capabilities rc
JOIN mdl_role r ON rc.roleid = r.id
JOIN mdl_context c ON rc.contextid = c.id
WHERE rc.roleid IN (
    SELECT roleid FROM mdl_role_assignments WHERE userid = @user_id
)
AND c.contextlevel = 10
AND rc.capability IN (
    'moodle/category:manage',
    'moodle/category:viewcourselist',
    'moodle/user:update',
    'moodle/user:viewalldetails',
    'moodle/course:update',
    'moodle/course:delete'
)
ORDER BY rc.capability;

-- 5. Ver el contexto del sistema
SELECT '=== CONTEXTO DEL SISTEMA ===' AS '';
SELECT id, contextlevel, path, depth
FROM mdl_context
WHERE contextlevel = 10;
