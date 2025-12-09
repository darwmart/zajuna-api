# 🎨 Vista de Usuarios - Tipo Zajuna SENA / Moodle

## ✅ Implementación Completa

Tu aplicación ahora cuenta con una **vista de usuarios profesional** que replica fielmente el flujo de plataformas LMS como **Zajuna SENA** y **Moodle**.

---

## 📦 Componentes Creados/Mejorados

### 1. **UsersTableView.js** (Principal) ⭐
- Componente principal con diseño tipo Moodle
- Tabla completa con todas las funcionalidades
- Sistema de paginación avanzado
- Gestión de usuarios individual y masiva

### 2. **MoodleBreadcrumb.js** (Nuevo)
- Navegación tipo migas de pan
- Iconos de home y flechas
- Estilo consistente con Moodle

### 3. **UserStatsCards.js** (Nuevo)
- Tarjetas de estadísticas
- 4 métricas principales:
  - Total usuarios
  - Usuarios activos
  - Usuarios inactivos  
  - Acceso reciente (7 días)
- Diseño moderno con iconos y colores

### 4. **UserEditModal.js** (Existente)
- Modal de Material-UI
- Formulario de edición
- Integrado con el servicio de API

---

## 🎯 Funcionalidades Implementadas

### ✨ Características Principales

#### 🔍 **Búsqueda y Filtrado**
- [x] Campo de búsqueda con debounce (400ms)
- [x] Búsqueda por nombre, apellido o email
- [x] Filtros expandibles/contraíbles
- [x] Botón para limpiar filtros
- [x] Contador de resultados en tiempo real

#### 📊 **Visualización de Datos**
- [x] Tabla con filas alternas (mejor legibilidad)
- [x] Tarjetas de estadísticas en la parte superior
- [x] Breadcrumb de navegación
- [x] Marcador lateral verde decorativo
- [x] Tooltips informativos en botones

#### ⚙️ **Acciones sobre Usuarios**
- [x] **Ver**: Vista rápida del perfil
- [x] **Editar**: Modal para modificar datos
- [x] **Eliminar individual**: Con confirmación
- [x] **Eliminar masivo**: Selección múltiple
- [x] **Exportar CSV**: Descarga de datos

#### 📄 **Paginación**
- [x] Paginación superior e inferior
- [x] Botones anterior/siguiente
- [x] Máximo 5 páginas visibles
- [x] Puntos suspensivos cuando hay más páginas
- [x] Contador de registros (ej: "Mostrando 1-20 de 150")

#### ✅ **Selección Múltiple**
- [x] Checkbox en cada fila
- [x] Checkbox "Seleccionar todos"
- [x] Contador de seleccionados
- [x] Resaltado de filas seleccionadas

---

## 🎨 Diseño Visual

### Colores (Paleta Moodle)
```css
Verde oscuro header: #1B5E20
Verde principal:     #2E7D32
Verde claro:         #E6EFDC
Fondo gris:          #F5F7F6
Azul breadcrumb:     #0D47A1
Rojo eliminar:       #D32F2F
```

### Elementos Visuales
- ✅ Estilos CSS personalizados en `moodle-theme.css`
- ✅ Iconos de React Icons (FaTrash, FaEye, FaCog, etc.)
- ✅ Spinner de carga animado
- ✅ Transiciones suaves en botones y hover
- ✅ Diseño responsive para móviles

---

## 🚀 Cómo Usar

### Opción 1: Vista tipo Moodle (Recomendada)
```javascript
import UsersTableView from './components/UsersTableView';

function App() {
  return <UsersTableView />;
}
```

### Opción 2: Vista con Material-UI (Alternativa)
```javascript
import UsersList from './components/UsersList';

function App() {
  return <UsersList />;
}
```

### Ya está integrado en tu App.js
Accede mediante la pestaña **"Vista tipo Moodle"** en tu aplicación.

---

## 📱 Capturas de Funcionalidades

### **Búsqueda y Filtrado**
```
┌─────────────────────────────────────────┐
│ 🏠 Inicio > Admin > Usuarios > Lista    │
├─────────────────────────────────────────┤
│ [📊 Total: 150] [✓ Activos: 120] [...]  │
├─────────────────────────────────────────┤
│ 150 Usuarios    [🔍 Filtros] [⬇ Export] │
│                                          │
│ Filtros de búsqueda:                     │
│ Nombre: [________________] [Limpiar]     │
│ 💡 Buscar por nombre, email...           │
└─────────────────────────────────────────┘
```

### **Tabla de Usuarios**
```
┌──────────────────────────────────────────────────┐
│ ☑ | Nombre      | Email          | Acciones      │
├──────────────────────────────────────────────────┤
│ ☑ | Juan Pérez  | juan@test.com  | 👁 ⚙ 🗑       │
│ □ | Ana García  | ana@test.com   | 👁 ⚙ 🗑       │
│ ☑ | Luis Torres | luis@test.com  | 👁 ⚙ 🗑       │
└──────────────────────────────────────────────────┘
│ Mostrando 1-20 de 150    [« 1 2 3 4 5 ... »]   │
└──────────────────────────────────────────────────┘
```

---

## 🔧 Servicios API Utilizados

```javascript
// services/usersService.js

getUsers(query, page, perPage, deleted)
  → Obtener lista paginada de usuarios

deleteUsers(userIds)
  → Eliminar uno o varios usuarios

updateUser(userData)
  → Actualizar información de usuario

getUsersByField(field, values)
  → Buscar usuarios por campo específico
```

---

## 📊 Flujo de Trabajo Típico

### Escenario 1: Buscar y Editar Usuario
1. Usuario escribe "Juan" en el buscador
2. Sistema espera 400ms → Busca automáticamente
3. Muestra resultados filtrados
4. Click en ⚙ (editar)
5. Modal se abre con datos precargados
6. Modificar campos → "Guardar"
7. Lista se actualiza automáticamente

### Escenario 2: Eliminación Masiva
1. Marcar checkbox de 5 usuarios
2. Aparece botón "Eliminar (5)"
3. Click → Confirmación
4. Usuarios eliminados
5. Checkboxes se limpian
6. Lista se recarga

### Escenario 3: Exportar Datos
1. Aplicar filtros (opcional)
2. Click en "Exportar"
3. Se genera archivo CSV
4. Descarga automática: `usuarios_2025-10-17.csv`

---

## 🎯 Mejoras Adicionales Sugeridas

### Corto Plazo
- [ ] Filtro por país (dropdown)
- [ ] Filtro por fecha de registro
- [ ] Ordenamiento de columnas (click en header)
- [ ] Vista de detalles en modal (no solo edición)

### Mediano Plazo
- [ ] Asignación masiva de roles
- [ ] Exportar a Excel (XLSX)
- [ ] Importar usuarios desde CSV
- [ ] Historial de cambios

### Largo Plazo
- [ ] Dashboard con gráficas
- [ ] Reportes personalizados
- [ ] Notificaciones por email
- [ ] Auditoría de acciones

---

## 🐛 Troubleshooting

### Problema: No se muestran usuarios
**Solución**: Verifica que el backend esté corriendo y responda en la URL configurada.

### Problema: La búsqueda no funciona
**Solución**: Revisa que el hook `useDebounce` esté funcionando correctamente.

### Problema: Los estilos no se aplican
**Solución**: Asegúrate de que `moodle-theme.css` esté importado en el componente.

### Problema: Error al eliminar usuarios
**Solución**: Verifica que el usuario tenga permisos de administrador en el backend.

---

## 📚 Documentación Adicional

- **Documentación completa**: `USERS_VIEW_DOCUMENTATION.md`
- **Código fuente**: `/src/components/UsersTableView.js`
- **Estilos**: `/src/styles/moodle-theme.css`

---

## 🎓 Conceptos Aplicados

- ✅ **React Hooks**: useState, useEffect, useCallback
- ✅ **Custom Hooks**: useDebounce
- ✅ **Async/Await**: Peticiones asíncronas
- ✅ **CSS Variables**: Tema personalizable
- ✅ **Responsive Design**: Media queries
- ✅ **UX Best Practices**: Loading states, confirmaciones
- ✅ **Accessibility**: ARIA labels, roles

---

## ✨ Resultado Final

Has conseguido replicar exitosamente el flujo frontend de **Zajuna SENA / Moodle** con:

- 🎨 Diseño visual idéntico
- ⚡ Funcionalidades completas
- 📱 Responsive design
- 🔒 Confirmaciones de seguridad
- 📊 Estadísticas en tiempo real
- 🎯 UX optimizada

**¡Tu aplicación está lista para producción!** 🚀

---

**Versión**: 2.0 Final  
**Fecha**: Octubre 2025  
**Stack**: React 19 + React Icons + CSS Custom
