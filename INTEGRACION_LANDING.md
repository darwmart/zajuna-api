# Integración Landing - Dashboard - API Zajuna

## Resumen

Este documento describe la integración completa de los tres proyectos del sistema Zajuna:

- **Landing Page**: Página de entrada y autenticación de usuarios
- **Dashboard Zajuna**: Panel administrativo para gestión de usuarios y cursos
- **API Backend**: Servicio backend en Go/Gin que maneja toda la lógica de negocio

## Arquitectura

```
┌─────────────────────────────────────────────────────────────┐
│                    FLUJO DE AUTENTICACIÓN                   │
└─────────────────────────────────────────────────────────────┘

1. Usuario accede a Landing Page (http://localhost:5173)
   │
   ├─ Formulario "Ingreso Cursos"
   │  └─ Usa: tipo documento + número documento + contraseña
   │     └─ Envía: { idnumber: "123456", password: "xxx" }
   │
   └─ Formulario "Ingreso Administrativos"
      └─ Usa: username + contraseña
         └─ Envía: { username: "admin", password: "xxx" }

2. Landing envía credenciales a API Backend
   │
   └─ POST http://localhost:8080/api/login
      └─ Respuesta: { token: "jwt...", user: {...} }

3. Landing recibe respuesta exitosa
   │
   └─ Almacena token en localStorage
   └─ Redirige al Dashboard: http://localhost:3000
```

## Configuración de Puertos

| Servicio     | Puerto | URL                   |
| ------------ | ------ | --------------------- |
| Backend API  | 8080   | http://localhost:8080 |
| Landing Page | 5173   | http://localhost:5173 |
| Dashboard    | 3000   | http://localhost:3000 |

## Estructura de Proyectos

```
Proyecto_Zajuna/
├── zajuna-api/              # Backend Go/Gin (Puerto 8080)
│   ├── cmd/server/main.go
│   ├── internal/
│   │   ├── config/.env.development
│   │   ├── models/
│   │   ├── routes/
│   │   └── services/
│   └── go.mod
│
├── Landing-Lms-Zajuna/      # Landing Page React + Vite (Puerto 5173)
│   ├── src/
│   │   ├── components/
│   │   │   └── forms/LoginForm.jsx
│   │   ├── utils/authClient.js
│   │   └── App.jsx
│   ├── .env
│   ├── vite.config.js
│   └── package.json
│
└── Zajuna/                  # Dashboard React (Puerto 3000)
    ├── src/
    │   ├── pages/
    │   ├── components/
    │   └── App.js
    └── package.json
```

## Archivos de Configuración

### Landing-Lms-Zajuna/.env

```env
# URL base del backend API
VITE_API_BASE_URL=http://localhost:8080

# URL del dashboard de Zajuna (para redirección después del login)
VITE_ZAJUNA_DASHBOARD_URL=http://localhost:3000
```

### Landing-Lms-Zajuna/vite.config.js

```javascript
import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

export default defineConfig({
  plugins: [react()],
  server: {
    port: 5173, // Puerto de la Landing
  },
  preview: {
    port: 5173,
  },
});
```

### zajuna-api/internal/config/.env.development

```env
# Configuración del backend
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=tu_password
DB_NAME=zajuna_db
```

## Funcionamiento del Login

### 1. LoginForm.jsx (Landing-Lms-Zajuna/src/components/forms/LoginForm.jsx)

El componente `LoginForm` maneja dos tipos de autenticación:

#### a) Login de Cursos (Estudiantes)

```javascript
// Formulario para estudiantes
const getCredentials = (formData) => {
  const document = formData.get("document");
  const password = formData.get("password");
  return { idnumber: document.trim(), password };
};
```

#### b) Login Administrativo

```javascript
// Formulario para administradores
const getCredentials = (formData) => {
  const username = formData.get("username");
  const password = formData.get("password");
  return { username: username.trim(), password };
};
```

### 2. authClient.js (Landing-Lms-Zajuna/src/utils/authClient.js)

Maneja la comunicación con el backend:

```javascript
export async function loginRequest({ username, idnumber, password }) {
  const body = {
    username: username || idnumber,
    password: password,
  };

  const url = `${API_BASE_URL}/api/login`;

  const response = await fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Accept: "application/json",
    },
    body: JSON.stringify(body),
    credentials: "include",
  });

  // Procesar respuesta y almacenar token
  return await processLoginResponse(response);
}
```

### 3. Redirección al Dashboard

Después de un login exitoso:

```javascript
// En LoginForm.jsx:75-86
const processLogin = async (credentials) => {
  const response = await loginRequest(credentials);

  if (response.token) {
    storeToken(response.token);
  }

  if (response.user) {
    storeUser(response.user);
  }

  setSuccess(true);

  // Redirigir al dashboard de Zajuna
  const zajunaUrl =
    import.meta.env.VITE_ZAJUNA_DASHBOARD_URL || "http://localhost:3000";

  setTimeout(() => {
    window.location.href = zajunaUrl;
  }, 1500);
};
```

## Script de Inicio

El script `start-dev.sh` inicia los tres servicios en el orden correcto:

### Orden de Inicio:

1. **Backend API** (puerto 8080) - Espera 5 segundos
2. **Landing Page** (puerto 5173) - Espera 3 segundos
3. **Dashboard** (puerto 3000)

### Uso:

```bash
# Dar permisos de ejecución
chmod +x start-dev.sh

# Ejecutar el script
./start-dev.sh
```

### Salida esperada:

```
Iniciando Zajuna - API, Landing y Dashboard
========================================

🔍 Verificando dependencias...
✓ Todas las dependencias están instaladas

🔍 Verificando PostgreSQL...
✓ PostgreSQL está corriendo

📡 Paso 1/3: Iniciando Backend API (Go/Gin)...
✓ Backend API iniciado en http://localhost:8080 (PID: 12345)

⏳ Esperando a que la API esté lista...

🌐 Paso 2/3: Iniciando Landing Page (React + Vite)...
✓ Landing Page iniciada en http://localhost:5173 (PID: 12346)

🎛️  Paso 3/3: Iniciando Dashboard Zajuna (React)...
✓ Dashboard iniciado en http://localhost:3000 (PID: 12347)

========================================
✅ Todos los servidores están corriendo

📍 URLs:
   Landing Page: http://localhost:5173
   Dashboard:    http://localhost:3000
   Backend API:  http://localhost:8080

📋 Logs:
   [API]       = Backend API (Go/Gin)
   [LANDING]   = Landing Page (React + Vite)
   [DASHBOARD] = Dashboard Admin (React)

🔐 Flujo de autenticación:
   1. Los usuarios ingresan por la Landing: http://localhost:5173
   2. Al hacer login, se autentican contra la API: http://localhost:8080
   3. Si es administrador, son redirigidos al Dashboard: http://localhost:3000

⚠️  Orden de inicio:
   1. Backend API (puerto 8080)
   2. Landing Page (puerto 5173)
   3. Dashboard (puerto 3000)

💡 Presiona Ctrl+C para detener todos los servidores
========================================
```

## Flujo Completo de Usuario

### Caso 1: Estudiante

1. Usuario accede a `http://localhost:5173`
2. Selecciona "Ingreso cursos Zajuna"
3. Ingresa:
   - Tipo de documento: CC
   - Número de documento: 123456789
   - Contraseña: **\*\*\*\***
4. Click en "Iniciar sesión"
5. Landing envía a API: `{ idnumber: "123456789", password: "..." }`
6. API valida credenciales
7. API responde con token JWT
8. Landing almacena token
9. Landing redirige a Dashboard: `http://localhost:3000`

### Caso 2: Administrador

1. Usuario accede a `http://localhost:5173`
2. Selecciona "Ingreso Administrativos Zajuna"
3. Ingresa:
   - Usuario: admin
   - Contraseña: **\*\*\*\***
4. Click en "Iniciar sesión"
5. Landing envía a API: `{ username: "admin", password: "..." }`
6. API valida credenciales
7. API responde con token JWT
8. Landing almacena token
9. Landing redirige a Dashboard: `http://localhost:3000`

## Almacenamiento de Sesión

El sistema almacena la sesión en `localStorage`:

```javascript
// Token JWT
localStorage.setItem("zajuna_token", token);

// Datos del usuario
localStorage.setItem("zajuna_user", JSON.stringify(user));
```

### Acceso al token en peticiones posteriores:

```javascript
// En el Dashboard u otras páginas
const token = localStorage.getItem("zajuna_token");
const user = JSON.parse(localStorage.getItem("zajuna_user"));

// Incluir en headers de peticiones
const headers = {
  Authorization: `Bearer ${token}`,
  "Content-Type": "application/json",
};
```

## Endpoints de API

### POST /api/login

**Request:**

```json
{
  "username": "123456789", // O "admin" para administrativos
  "password": "contraseña"
}
```

**Response (éxito):**

```json
{
  "success": true,
  "token": "eyJhbGciOiJIUzI1NiIs...",
  "user": {
    "id": 1,
    "username": "123456789",
    "email": "user@example.com",
    "role": "admin"
  }
}
```

**Response (error):**

```json
{
  "success": false,
  "error": "Credenciales inválidas"
}
```

## Troubleshooting

### Error: "No se pudo conectar con el servidor"

**Causa:** La API no está corriendo o no responde en el puerto 8080

**Solución:**

```bash
# Verificar que la API esté corriendo
curl http://localhost:8080/api/health

# Si no responde, iniciar la API
cd zajuna-api
go run cmd/server/main.go
```

### Error: "CORS Error"

**Causa:** El backend no permite peticiones desde el origen de la Landing

**Solución:** Verificar la configuración CORS en el backend (zajuna-api)

### Error: "Credenciales inválidas"

**Causa:** Usuario o contraseña incorrectos

**Solución:** Verificar las credenciales en la base de datos

### Puerto en uso

**Error:** "Port 5173 is already in use"

**Solución:**

```bash
# Encontrar el proceso usando el puerto
lsof -i :5173

# Matar el proceso
kill -9 <PID>
```

## Variables de Entorno Importantes

### Landing (.env)

- `VITE_API_BASE_URL`: URL base del backend API
- `VITE_ZAJUNA_DASHBOARD_URL`: URL del dashboard para redirección

### Backend (.env.development)

- `PORT`: Puerto del servidor backend
- `DB_HOST`: Host de PostgreSQL
- `DB_PORT`: Puerto de PostgreSQL
- `DB_NAME`: Nombre de la base de datos

## Próximos Pasos

1. Implementar refresh token para renovar sesiones
2. Agregar validación de roles en el Dashboard
3. Implementar logout en el Dashboard
4. Agregar protección de rutas privadas en el Dashboard
5. Implementar manejo de sesiones expiradas

## Soporte

Para más información o reportar problemas:

- Revisar logs en la consola de cada servicio
- Verificar la documentación del API en `zajuna-api/README.md`
- Contactar al equipo de desarrollo
