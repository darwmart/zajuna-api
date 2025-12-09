# 📖 Ejemplos de Uso - Filtro "Contiene" y Match Types

## 🎯 Escenarios Prácticos

### Ejemplo 1: Buscar usuarios cuyo apellido CONTIENE "García"

**Configuración:**
- Campo: Apellido(s)
- Tipo: `contiene`
- Valor: `García`

**Resultado esperado:**
- García Pérez
- María García
- García-López
- Juan García Martínez

---

### Ejemplo 2: Buscar usuarios cuyo email NO CONTIENE "gmail"

**Configuración:**
- Campo: Dirección de correo
- Tipo: `no contiene`
- Valor: `gmail`

**Resultado esperado:**
- usuario@yahoo.com ✓
- admin@empresa.com ✓
- test@hotmail.com ✓
- usuario@gmail.com ✗ (excluido)

---

### Ejemplo 3: Buscar usuario con username EXACTO

**Configuración:**
- Campo: Nombre de usuario
- Tipo: `es igual a`
- Valor: `admin`

**Resultado esperado:**
- admin ✓
- admin123 ✗
- superadmin ✗

---

### Ejemplo 4: Buscar usuarios cuyo nombre COMIENZA CON "Juan"

**Configuración:**
- Campo: Nombre
- Tipo: `comienza con`
- Valor: `Juan`

**Resultado esperado:**
- Juan ✓
- Juan Carlos ✓
- Juana ✓
- María Juan ✗

---

### Ejemplo 5: Buscar usuarios cuyo apellido TERMINA CON "ez"

**Configuración:**
- Campo: Apellido(s)
- Tipo: `termina con`
- Valor: `ez`

**Resultado esperado:**
- Pérez ✓
- Rodríguez ✓
- González ✓
- López ✓
- García ✗

---

## 🔗 Ejemplos de Filtros Combinados

### Escenario A: Instructores con email institucional

**Filtros aplicados:**
1. Campo: Nombre de usuario | Tipo: `comienza con` | Valor: `prof`
2. Campo: Dirección de correo | Tipo: `termina con` | Valor: `@universidad.edu`

**Resultado:** Solo profesores con email institucional

---

### Escenario B: Excluir usuarios de prueba

**Filtros aplicados:**
1. Campo: Nombre de usuario | Tipo: `no contiene` | Valor: `test`
2. Campo: Nombre de usuario | Tipo: `no contiene` | Valor: `demo`

**Resultado:** Usuarios reales sin cuentas de prueba

---

### Escenario C: Buscar familia específica

**Filtros aplicados:**
1. Campo: Apellido(s) | Tipo: `contiene` | Valor: `López`
2. Campo: Dirección de correo | Tipo: `contiene` | Valor: `familialopez`

**Resultado:** Miembros de la familia López con email familiar

---

## 💡 Casos de Uso por Tipo de Coincidencia

### 📍 CONTIENE (Opción más flexible)

**Cuándo usar:**
- Búsquedas generales
- No estás seguro de la estructura exacta
- Quieres incluir variaciones

**Ejemplos:**
```
"María" encontrará:
✓ María
✓ María José
✓ Ana María
✓ mariajose@email.com
```

---

### 🚫 NO CONTIENE (Filtro de exclusión)

**Cuándo usar:**
- Eliminar cuentas de prueba
- Excluir dominios de email
- Filtrar usuarios temporales

**Ejemplos:**
```
Excluir "test" de username encontrará:
✓ juan_rodriguez
✓ admin_sistema
✗ test_usuario
✗ usuario_test
```

---

### 🎯 ES IGUAL A (Búsqueda exacta)

**Cuándo usar:**
- Buscar un usuario específico
- Validar username único
- Coincidencia precisa

**Ejemplos:**
```
"admin" encontrará:
✓ admin
✗ admin123
✗ superadmin
✗ Admin (si es case-insensitive)
```

---

### ▶️ COMIENZA CON (Prefijo)

**Cuándo usar:**
- Usuarios con prefijo de rol (prof_, admin_)
- Nomenclatura organizacional
- Códigos de identificación

**Ejemplos:**
```
"prof_" encontrará:
✓ prof_matematicas
✓ prof_historia
✗ ayudante_prof
✗ supervisor_prof
```

---

### ◀️ TERMINA CON (Sufijo)

**Cuándo usar:**
- Filtrar por dominio de email
- Extensiones de username
- Patrones de nomenclatura

**Ejemplos:**
```
"@empresa.com" encontrará:
✓ usuario@empresa.com
✓ admin@empresa.com
✗ usuario@empresa.com.mx
✗ usuario@otraempresa.com
```

---

## 🎓 Casos de Uso Académicos (Específicos de Moodle/Zajuna)

### Caso 1: Encontrar todos los estudiantes de un programa

**Filtros:**
- Username | `comienza con` | `est_`
- Email | `termina con` | `@estudiantes.edu`

---

### Caso 2: Identificar cuentas suspendidas temporalmente

**Filtros:**
- Username | `contiene` | `_suspended`
- Username | `no contiene` | `_deleted`

---

### Caso 3: Auditoria de administradores

**Filtros:**
- Username | `comienza con` | `admin`
- Email | `no contiene` | `@gmail.com`

---

### Caso 4: Usuarios de un curso específico

**Filtros:**
- Nombre completo | `contiene` | `curso2024`
- Email | `termina con` | `@plataforma.edu`

---

## 🛠️ Tips de Búsqueda Avanzada

### Tip 1: Búsqueda Incremental
Comienza con filtros amplios y ve refinando:
1. `Apellido contiene "García"` → 500 resultados
2. + `Nombre comienza con "Juan"` → 50 resultados  
3. + `Email termina con "@empresa.com"` → 5 resultados

### Tip 2: Uso de "NO CONTIENE"
Muy útil para limpiar resultados:
```
Buscar profesores activos:
✓ Username comienza con "prof_"
✓ Username NO contiene "inactive"
✓ Username NO contiene "test"
```

### Tip 3: Combinación de Patrones
```
Buscar coordinadores de área:
✓ Username comienza con "coord_"
✓ Username termina con "_2024"
```

---

## ⚠️ Errores Comunes a Evitar

### ❌ Error 1: Usar "es igual a" cuando necesitas "contiene"
```
Búsqueda: Apellido "es igual a" "García"
Resultado: Solo "García" exacto
Perdido: "García Pérez", "García-López"

✅ Correcto: Apellido "contiene" "García"
```

### ❌ Error 2: Olvidar espacios en "comienza con" o "termina con"
```
❌ Email "termina con" " @gmail.com" (espacio extra)
✅ Email "termina con" "@gmail.com"
```

### ❌ Error 3: Usar múltiples "NO CONTIENE" sin sentido
```
❌ Username "no contiene" "a" (demasiado restrictivo)
✅ Username "no contiene" "test" (específico y útil)
```

---

## 📊 Comparación de Rendimiento

| Tipo de Filtro | Velocidad | Precisión | Casos de Uso |
|----------------|-----------|-----------|--------------|
| Contiene | ⚡⚡⚡ | 🎯🎯 | General |
| No contiene | ⚡⚡ | 🎯🎯🎯 | Exclusión |
| Es igual a | ⚡⚡⚡⚡ | 🎯🎯🎯🎯 | Exacto |
| Comienza con | ⚡⚡⚡ | 🎯🎯🎯 | Prefijos |
| Termina con | ⚡⚡⚡ | 🎯🎯🎯 | Sufijos |

---

## 🎬 Ejemplo Completo Paso a Paso

### Objetivo: Encontrar profesores activos del departamento de ciencias

**Paso 1:** Abrir "Mostrar más..." para ver todos los filtros

**Paso 2:** Configurar primer filtro
- Campo: Nombre de usuario
- Tipo: `comienza con`
- Valor: `prof_ciencias_`

**Paso 3:** Hacer clic en "Añadir filtro"

**Paso 4:** Configurar segundo filtro
- Campo: Dirección de correo
- Tipo: `termina con`
- Valor: `@universidad.edu`

**Paso 5:** Hacer clic en "Añadir filtro"

**Paso 6:** Configurar filtro de exclusión
- Campo: Nombre de usuario
- Tipo: `no contiene`
- Valor: `_inactive`

**Paso 7:** Hacer clic en "Añadir filtro"

**Resultado:** Lista de profesores activos de ciencias con email institucional

**Visualización de filtros activos:**
```
[Nombre de usuario] comienza con "prof_ciencias_" [×]
[Dirección de correo] termina con "@universidad.edu" [×]
[Nombre de usuario] no contiene "_inactive" [×]
[Limpiar todos los filtros]
```

---

## 🔍 Testing Checklist

- [ ] Filtro "contiene" funciona correctamente
- [ ] Filtro "no contiene" excluye correctamente
- [ ] Filtro "es igual a" es case-insensitive
- [ ] Filtro "comienza con" reconoce prefijos
- [ ] Filtro "termina con" reconoce sufijos
- [ ] Múltiples filtros se intersectan correctamente
- [ ] Pills de filtros activos se muestran
- [ ] Botón X elimina filtros individuales
- [ ] "Limpiar todos" resetea todo
- [ ] Búsqueda vacía no rompe la aplicación

---

**¡Feliz filtrado! 🎉**
