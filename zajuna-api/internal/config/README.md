# Configuración del Proyecto

Este directorio contiene la configuración de la aplicación basada en variables de entorno.

## Variables de Entorno Requeridas

### Configuración de la Aplicación

| Variable | Descripción | Valores | Requerida |
|----------|-------------|---------|-----------|
| `APP_ENV` | Entorno de ejecución | `development`, `production` | Sí |
| `APP_PORT` | Puerto donde corre la aplicación | Número (ej: `8080`) | Sí |

### Configuración de Base de Datos

| Variable | Descripción | Ejemplo | Requerida |
|----------|-------------|---------|-----------|
| `DB_HOST` | Host del servidor PostgreSQL | `localhost`, `db` | Sí |
| `DB_PORT` | Puerto de PostgreSQL | `5432` | Sí |
| `DB_USER` | Usuario de la base de datos | `postgres` | Sí |
| `DB_PASSWORD` | Contraseña del usuario | - | Sí |
| `DB_NAME` | Nombre de la base de datos | `zajuna` | Sí |
| `SSL_MODE` | Modo SSL para conexión | `disable`, `require`, `verify-full` | Sí |

## Configuración por Entorno

### Development (Desarrollo)

```bash
APP_ENV=development
APP_PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=tu_password_desarrollo
DB_NAME=zajuna
SSL_MODE=disable
```

### Production (Producción)

```bash
APP_ENV=production
APP_PORT=8080
DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=tu_password_produccion_seguro
DB_NAME=zajuna
SSL_MODE=require
```

## Instrucciones de Configuración

### Primer Uso

1. **Copia el archivo de ejemplo correspondiente a tu entorno:**

   ```bash
   # Para desarrollo
   cp internal/config/.env.development.example internal/config/.env.development

   # Para producción
   cp internal/config/.env.production.example internal/config/.env.production
   ```

2. **Edita el archivo `.env` con tus credenciales reales:**

   ```bash
   # Abre el archivo con tu editor favorito
   nano internal/config/.env.development
   ```

3. **Configura la variable de entorno antes de ejecutar:**

   ```bash
   export APP_ENV=development
   ./server
   ```

### Notas de Seguridad

⚠️ **IMPORTANTE:**
- Los archivos `.env` contienen información sensible y **NO deben** ser incluidos en el control de versiones
- Mantén contraseñas seguras en producción
- Usa `SSL_MODE=require` o `verify-full` en producción
- Nunca compartas tus archivos `.env` reales

### Archivos en este Directorio

- ✅ `.env.example` - Plantilla general (sin datos sensibles)
- ✅ `.env.development.example` - Plantilla para desarrollo
- ✅ `.env.production.example` - Plantilla para producción
- ❌ `.env` - Tu archivo real (ignorado por git)
- ❌ `.env.development` - Tu archivo de desarrollo (ignorado por git)
- ❌ `.env.production` - Tu archivo de producción (ignorado por git)

## Solución de Problemas

### Error: "No se encontró .env.{environment}"

La aplicación buscará el archivo según `APP_ENV`. Si no lo encuentra, usará las variables del sistema.

**Solución:**
1. Asegúrate de que existe el archivo `.env.{environment}` en `internal/config/`
2. Verifica que `APP_ENV` esté configurado correctamente

### Error de Conexión a la Base de Datos

**Verifica:**
1. Que PostgreSQL esté corriendo
2. Que las credenciales sean correctas
3. Que el usuario tenga permisos en la base de datos
4. Que el puerto no esté bloqueado por firewall

### Variables de Entorno No Se Cargan

Si la aplicación no carga las variables del archivo `.env`:

1. Verifica que el archivo esté en la ubicación correcta: `internal/config/.env.{environment}`
2. Verifica que `APP_ENV` esté configurado
3. Reinicia la aplicación después de cambiar variables
