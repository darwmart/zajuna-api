#!/bin/bash

# Script de diagnóstico de permisos para Zajuna API
# Verifica los permisos de un usuario administrador

echo "============================================"
echo "DIAGNÓSTICO DE PERMISOS - ZAJUNA API"
echo "============================================"
echo ""

# Leer usuario
read -p "Ingrese el username o idnumber del usuario: " USERNAME

echo ""
echo "Consultando información del usuario..."
echo ""

# 1. Verificar que el usuario existe
mysql -uroot -p -D zajuna <<EOF
SELECT
    id AS user_id,
    username,
    firstname,
    lastname,
    email,
    idnumber,
    deleted,
    suspended,
    confirmed
FROM mdl_user
WHERE (username = '$USERNAME' OR idnumber = '$USERNAME')
LIMIT 1;
EOF

# Obtener el user_id
USER_ID=$(mysql -uroot -p -D zajuna -se "SELECT id FROM mdl_user WHERE (username = '$USERNAME' OR idnumber = '$USERNAME') LIMIT 1")

if [ -z "$USER_ID" ]; then
    echo "❌ Usuario no encontrado"
    exit 1
fi

echo ""
echo "✅ Usuario encontrado: ID=$USER_ID"
echo ""

# 2. Verificar si es Site Admin
echo "============================================"
echo "2. VERIFICANDO SI ES SITE ADMIN"
echo "============================================"
mysql -uroot -p -D zajuna <<EOF
SELECT
    ra.id AS assignment_id,
    r.id AS role_id,
    r.shortname AS role_shortname,
    r.name AS role_name,
    r.archetype,
    c.id AS context_id,
    c.contextlevel,
    c.path
FROM mdl_role_assignments ra
JOIN mdl_role r ON ra.roleid = r.id
JOIN mdl_context c ON ra.contextid = c.id
WHERE ra.userid = $USER_ID
  AND r.archetype = 'manager'
  AND c.contextlevel = 10;
EOF

# 3. Ver todas las asignaciones de roles
echo ""
echo "============================================"
echo "3. TODAS LAS ASIGNACIONES DE ROLES"
echo "============================================"
mysql -uroot -p -D zajuna <<EOF
SELECT
    ra.id AS assignment_id,
    r.id AS role_id,
    r.shortname AS role_shortname,
    r.name AS role_name,
    r.archetype,
    c.id AS context_id,
    c.contextlevel,
    c.instanceid,
    c.path
FROM mdl_role_assignments ra
JOIN mdl_role r ON ra.roleid = r.id
JOIN mdl_context c ON ra.contextid = c.id
WHERE ra.userid = $USER_ID
ORDER BY c.depth ASC;
EOF

# 4. Ver capabilities específicas
echo ""
echo "============================================"
echo "4. CAPABILITIES DEL USUARIO EN CONTEXTO SISTEMA"
echo "============================================"
echo "Consultando capabilities para roles asignados en contexto sistema..."
echo ""

mysql -uroot -p -D zajuna <<EOF
SELECT
    rc.capability,
    rc.permission,
    CASE rc.permission
        WHEN 1 THEN 'ALLOW'
        WHEN 0 THEN 'INHERIT'
        WHEN -1 THEN 'PREVENT'
        WHEN -1000 THEN 'PROHIBIT'
        ELSE 'UNKNOWN'
    END AS permission_name,
    r.shortname AS role_shortname,
    c.contextlevel
FROM mdl_role_capabilities rc
JOIN mdl_role r ON rc.roleid = r.id
JOIN mdl_context c ON rc.contextid = c.id
WHERE rc.roleid IN (
    SELECT ra.roleid
    FROM mdl_role_assignments ra
    WHERE ra.userid = $USER_ID
)
AND rc.capability IN (
    'moodle/category:manage',
    'moodle/category:viewcourselist',
    'moodle/user:update',
    'moodle/user:viewalldetails',
    'moodle/course:update',
    'moodle/course:delete'
)
AND c.contextlevel = 10
ORDER BY rc.capability, rc.permission DESC;
EOF

# 5. Ver contexto del sistema
echo ""
echo "============================================"
echo "5. CONTEXTO DEL SISTEMA"
echo "============================================"
mysql -uroot -p -D zajuna <<EOF
SELECT
    id,
    contextlevel,
    instanceid,
    path,
    depth,
    locked
FROM mdl_context
WHERE contextlevel = 10
LIMIT 1;
EOF

echo ""
echo "============================================"
echo "DIAGNÓSTICO COMPLETADO"
echo "============================================"
echo ""
echo "Si el usuario es administrador, debería tener:"
echo "  - Un role con archetype='manager' en contextlevel=10"
echo "  - Capabilities con permission=1 (ALLOW) para las operaciones"
echo ""
