# 📚 Documentación - Vista de Usuarios Tipo Moodle/Zajuna SENA

## 🎯 Descripción General

Esta implementación replica fielmente el flujo frontend de la gestión de usuarios de plataformas LMS como **Moodle** y **Zajuna SENA**, proporcionando una interfaz intuitiva y profesional para administradores.

## ✨ Características Implementadas

### 1. **Diseño Visual Tipo Moodle**
- ✅ Esquema de colores verde característico de Moodle
- ✅ Breadcrumb (migas de pan) para navegación
- ✅ Marcador lateral verde decorativo
- ✅ Tabla con filas alternas para mejor legibilidad
- ✅ Estilos responsive para dispositivos móviles

### 2. **Funcionalidades de Búsqueda y Filtrado**
- ✅ Búsqueda con debounce (400ms) para optimizar peticiones
- ✅ Filtro por nombre completo, email o username
- ✅ Área de filtros expandible/contraíble
- ✅ Contador de resultados en tiempo real
- ✅ Botón para limpiar todos los filtros

### 3. **Gestión de Usuarios**
- ✅ **Visualización**: Vista detallada de perfil de usuario
- ✅ **Edición**: Modal para editar información del usuario
- ✅ **Eliminación**: Individual y masiva con confirmación
- ✅ **Selección múltiple**: Checkboxes para operaciones en lote
- ✅ **Exportación**: Descarga de datos a formato CSV

### 4. **Paginación Avanzada**
- ✅ Paginación superior e inferior
- ✅ Navegación con botones anterior/siguiente
- ✅ Páginas limitadas a 5 visibles con puntos suspensivos
- ✅ Indicador de rango de registros mostrados
- ✅ Contador total de usuarios

### 5. **Experiencia de Usuario**
- ✅ Spinner de carga animado
- ✅ Mensajes informativos cuando no hay resultados
- ✅ Tooltips en botones de acción
- ✅ Iconos intuitivos (React Icons)
- ✅ Transiciones y animaciones suaves
- ✅ Indicador visual de filas seleccionadas

## 📁 Estructura de Archivos

```
src/
├── components/
│   ├── UsersTableView.js          # Componente principal
│   ├── UserEditModal.js            # Modal de edición
│   ├── MoodleBreadcrumb.js         # Breadcrumb de navegación
│   ├── UsersList.js                # Vista alternativa (Material-UI)
│   └── ...
├── services/
│   └── usersService.js             # Servicios API
├── hooks/
│   └── useDebounce.js              # Hook personalizado
└── styles/
    └── moodle-theme.css            # Estilos tipo Moodle
```

## 🎨 Paleta de Colores

```css
--moodle-green-strong: #1B5E20  /* Verde oscuro header */
--moodle-green: #2E7D32          /* Verde principal botones */
--moodle-green-light: #E6EFDC    /* Verde claro tabla header */
--moodle-grey-bg: #F5F7F6        /* Fondo de página */
--moodle-text: #333333           /* Texto principal */
--moodle-accent: #0D47A1         /* Azul breadcrumb */
--moodle-border: #D6D6D6         /* Bordes */
```

## 🔧 Uso del Componente

### Importación Básica

```javascript
import UsersTableView from './components/UsersTableView';

function App() {
  return <UsersTableView />;
}
```

### Funciones Principales

#### 1. **Búsqueda de Usuarios**
```javascript
const fetchUsers = useCallback(async () => {
  setLoading(true);
  const data = await getUsers(debouncedQuery, page, perPage);
  setUsers(data.items || []);
  setTotal(data.total || 0);
  setLoading(false);
}, [debouncedQuery, page, perPage]);
```

#### 2. **Eliminación Individual**
```javascript
const handleDelete = async (user) => {
  if (!window.confirm(`¿Eliminar a ${user.firstname} ${user.lastname}?`)) return;
  await deleteUsers([user.id]);
  fetchUsers();
};
```

#### 3. **Eliminación Masiva**
```javascript
const handleBulkDelete = async () => {
  if (selectedUsers.length === 0) return;
  if (!window.confirm(`¿Eliminar ${selectedUsers.length} usuario(s)?`)) return;
  await deleteUsers(selectedUsers);
  setSelectedUsers([]);
  fetchUsers();
};
```

#### 4. **Exportación a CSV**
```javascript
const exportToCSV = () => {
  const headers = ["ID", "Nombre", "Apellido", "Email", "Ciudad", "País"];
  const csvContent = [
    headers.join(","),
    ...users.map(u => [u.id, u.firstname, u.lastname, u.email, u.city, u.country].join(","))
  ].join("\n");
  
  const blob = new Blob([csvContent], { type: "text/csv" });
  const url = window.URL.createObjectURL(blob);
  const a = document.createElement("a");
  a.href = url;
  a.download = `usuarios_${new Date().toISOString().split('T')[0]}.csv`;
  a.click();
};
```

## 🚀 Mejoras Implementadas

### Comparación con la versión anterior:

| Característica | Antes | Ahora |
|---------------|-------|-------|
| Selección múltiple | ❌ | ✅ |
| Exportar CSV | ❌ | ✅ |
| Filtros colapsables | ❌ | ✅ |
| Breadcrumb | ❌ | ✅ |
| Paginación inteligente | Básica | Avanzada |
| Loading spinner | Texto | Animado |
| Botones de acción | Básicos | Con tooltips y estilos |
| Contador de registros | ❌ | ✅ |
| Responsive design | Limitado | Completo |

## 📊 API Endpoints Utilizados

```javascript
// Obtener usuarios con paginación
GET /api/core_user_get_users
params: { q, page, per_page, deleted }

// Eliminar usuarios
DELETE /api/core_user_delete_users
body: { userids: [id1, id2, ...] }

// Actualizar usuario
PUT /api/core_user_update_users
body: { users: [userData] }
```

## 🎯 Casos de Uso

### 1. **Administrador busca usuario por nombre**
1. Usuario escribe en el campo de búsqueda
2. Sistema espera 400ms (debounce)
3. Se ejecuta la búsqueda automáticamente
4. Resultados se muestran con contador

### 2. **Eliminación masiva de usuarios**
1. Administrador selecciona checkboxes de usuarios
2. Aparece botón "Eliminar (N)"
3. Click en el botón → Confirmación
4. Usuarios se eliminan en lote
5. Lista se actualiza automáticamente

### 3. **Exportación de datos**
1. Usuario filtra datos según necesidad
2. Click en botón "Exportar"
3. Se genera archivo CSV con los datos visibles
4. Descarga automática del archivo

## 🔐 Consideraciones de Seguridad

- ✅ Confirmación antes de eliminar usuarios
- ✅ Validación de permisos en el backend
- ✅ Sanitización de datos de entrada
- ✅ Manejo de errores con try-catch

## 📱 Responsive Design

El diseño se adapta a diferentes tamaños de pantalla:

- **Desktop** (>768px): Tabla completa con todos los elementos
- **Tablet/Mobile** (<768px): 
  - Tabla con scroll horizontal
  - Elementos reorganizados verticalmente
  - Botones adaptados

## 🎓 Próximas Mejoras Sugeridas

1. **Filtros avanzados adicionales**:
   - Filtro por país
   - Filtro por fecha de último acceso
   - Filtro por roles

2. **Ordenamiento de columnas**:
   - Click en headers para ordenar
   - Indicadores visuales de orden

3. **Acciones masivas adicionales**:
   - Asignar roles en lote
   - Enviar emails masivos
   - Cambiar estado (activo/inactivo)

4. **Visualización de perfil mejorada**:
   - Modal más detallado
   - Historial de actividad
   - Cursos inscritos

5. **Persistencia de filtros**:
   - Guardar filtros en localStorage
   - Recordar última página visitada

## 🤝 Contribución

Para agregar nuevas funcionalidades:

1. Actualiza `UsersTableView.js` con la nueva función
2. Agrega estilos necesarios en `moodle-theme.css`
3. Actualiza este documento con la nueva característica

## 📞 Soporte

Para dudas o problemas:
- Revisa la consola del navegador para errores
- Verifica que los endpoints API estén funcionando
- Comprueba que los servicios estén correctamente configurados

---

**Versión**: 2.0  
**Última actualización**: Octubre 2025  
**Autor**: Zajuna Frontend Team
