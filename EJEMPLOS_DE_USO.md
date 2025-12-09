# 🎯 Ejemplos de Uso - UsersTableView

## Casos de Uso Prácticos

### 1️⃣ Integración Básica

```javascript
// App.js
import React from 'react';
import UsersTableView from './components/UsersTableView';

function App() {
  return (
    <div>
      <UsersTableView />
    </div>
  );
}

export default App;
```

---

### 2️⃣ Personalización de Paginación

Si quieres cambiar el número de usuarios por página:

```javascript
// En UsersTableView.js, cambia:
const [perPage] = useState(20);

// Por ejemplo a 50:
const [perPage] = useState(50);
```

---

### 3️⃣ Agregar Nuevos Filtros

Para agregar un filtro por país:

```javascript
// 1. Agregar estado
const [selectedCountry, setSelectedCountry] = useState("");

// 2. En la sección de filtros, agregar:
<div className="filter-row">
  <label>País</label>
  <select 
    value={selectedCountry}
    onChange={(e) => { 
      setSelectedCountry(e.target.value); 
      setPage(1); 
    }}
  >
    <option value="">Todos</option>
    <option value="Colombia">Colombia</option>
    <option value="México">México</option>
    <option value="Argentina">Argentina</option>
  </select>
</div>

// 3. Modificar fetchUsers para incluir el filtro
const data = await getUsers(debouncedQuery, page, perPage, selectedCountry);
```

---

### 4️⃣ Personalizar Colores del Tema

Edita `/src/styles/moodle-theme.css`:

```css
:root {
  /* Cambiar el verde por azul, por ejemplo */
  --moodle-green-strong: #0D47A1; /* Azul oscuro */
  --moodle-green: #1976D2;        /* Azul medio */
  --moodle-green-light: #E3F2FD;  /* Azul claro */
}
```

---

### 5️⃣ Agregar Columna Personalizada

```javascript
// En la tabla, agregar un <th>:
<thead>
  <tr>
    <th>...</th>
    <th>Nueva Columna</th>
  </tr>
</thead>

// Y el correspondiente <td>:
<tbody>
  {users.map((u) => (
    <tr key={u.id}>
      <td>...</td>
      <td>{u.nuevoValor || "N/A"}</td>
    </tr>
  ))}
</tbody>
```

---

### 6️⃣ Exportar con Más Campos

Modificar la función `exportToCSV`:

```javascript
const exportToCSV = () => {
  const headers = [
    "ID", 
    "Nombre", 
    "Apellido", 
    "Email", 
    "Ciudad", 
    "País", 
    "Teléfono",        // Nuevo
    "Fecha Registro",   // Nuevo
    "Último acceso"
  ];
  
  const csvContent = [
    headers.join(","),
    ...users.map(u => [
      u.id,
      u.firstname,
      u.lastname,
      u.email,
      u.city || "",
      u.country || "",
      u.phone || "",              // Nuevo
      u.timecreated || "",        // Nuevo
      u.lastlogin || "Nunca"
    ].join(","))
  ].join("\n");

  // ... resto del código
};
```

---

### 7️⃣ Agregar Notificaciones Toast

Instalar librería:
```bash
npm install react-toastify
```

Usar en el componente:
```javascript
import { toast, ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

// En handleDelete:
const handleDelete = async (u) => {
  if (!window.confirm(`¿Eliminar a ${u.firstname} ${u.lastname}?`)) return;
  try {
    await deleteUsers([u.id]);
    toast.success(`Usuario ${u.firstname} eliminado correctamente`);
    fetchUsers();
  } catch (e) {
    toast.error("Error al eliminar usuario");
  }
};

// En el return, agregar:
return (
  <div>
    <ToastContainer position="top-right" />
    {/* resto del código */}
  </div>
);
```

---

### 8️⃣ Modo Oscuro

Agregar en `moodle-theme.css`:

```css
/* Dark Mode */
[data-theme="dark"] {
  --moodle-green-strong: #66BB6A;
  --moodle-green: #81C784;
  --moodle-green-light: #2C2C2C;
  --moodle-grey-bg: #1a1a1a;
  --moodle-text: #E0E0E0;
  --moodle-border: #444;
}

[data-theme="dark"] .moodle-container {
  background: #242424;
}

[data-theme="dark"] .table-moodle tbody tr:nth-child(even) {
  background-color: #2a2a2a !important;
}
```

Toggle en el componente:
```javascript
const [darkMode, setDarkMode] = useState(false);

useEffect(() => {
  document.documentElement.setAttribute(
    'data-theme', 
    darkMode ? 'dark' : 'light'
  );
}, [darkMode]);

// Agregar botón:
<button onClick={() => setDarkMode(!darkMode)}>
  {darkMode ? "🌞 Modo Claro" : "🌙 Modo Oscuro"}
</button>
```

---

### 9️⃣ Búsqueda Avanzada con Múltiples Campos

```javascript
const [filters, setFilters] = useState({
  name: "",
  email: "",
  country: "",
  city: ""
});

const handleFilterChange = (field, value) => {
  setFilters({ ...filters, [field]: value });
  setPage(1);
};

// En el área de filtros:
<div className="filter-row">
  <label>Nombre</label>
  <input 
    value={filters.name}
    onChange={(e) => handleFilterChange('name', e.target.value)}
  />
</div>

<div className="filter-row">
  <label>Email</label>
  <input 
    value={filters.email}
    onChange={(e) => handleFilterChange('email', e.target.value)}
  />
</div>

// Modificar fetchUsers para enviar todos los filtros
```

---

### 🔟 Guardar Preferencias del Usuario

```javascript
// Guardar en localStorage
useEffect(() => {
  localStorage.setItem('usersTablePreferences', JSON.stringify({
    perPage,
    showFilters,
    page
  }));
}, [perPage, showFilters, page]);

// Cargar al inicio
useEffect(() => {
  const saved = localStorage.getItem('usersTablePreferences');
  if (saved) {
    const prefs = JSON.parse(saved);
    setShowFilters(prefs.showFilters);
    setPage(prefs.page);
  }
}, []);
```

---

## 🎨 Personalización de Estilos

### Cambiar el ancho de la tabla:
```css
.moodle-container {
  width: 95%; /* En vez de 92% */
  max-width: 1400px; /* Añadir límite máximo */
}
```

### Hacer la tabla más compacta:
```css
.table-moodle th,
.table-moodle td {
  padding: 6px 8px; /* En vez de 10px 12px */
  font-size: 12px;  /* En vez de 13px */
}
```

### Agregar bordes a la tabla:
```css
.table-moodle {
  border: 1px solid var(--moodle-border);
}

.table-moodle th,
.table-moodle td {
  border: 1px solid #e6e6e6;
}
```

---

## 🚀 Optimizaciones de Rendimiento

### 1. Memoizar funciones pesadas:
```javascript
import { useMemo } from 'react';

const processedUsers = useMemo(() => {
  return users.map(u => ({
    ...u,
    fullName: `${u.firstname} ${u.lastname}`,
    displayDate: new Date(u.lastlogin).toLocaleDateString()
  }));
}, [users]);
```

### 2. Virtualización para muchas filas:
```bash
npm install react-window
```

```javascript
import { FixedSizeList } from 'react-window';

// Reemplazar map por virtualización
<FixedSizeList
  height={600}
  itemCount={users.length}
  itemSize={50}
  width="100%"
>
  {({ index, style }) => (
    <div style={style}>
      {/* Renderizar fila */}
    </div>
  )}
</FixedSizeList>
```

---

## 📱 Adaptación para Móviles

### Tabla responsive con cards en móvil:

```css
@media (max-width: 768px) {
  .table-moodle thead {
    display: none;
  }
  
  .table-moodle tr {
    display: block;
    margin-bottom: 16px;
    border: 1px solid #ddd;
    border-radius: 8px;
  }
  
  .table-moodle td {
    display: block;
    text-align: right;
    padding: 8px;
    border-bottom: 1px solid #eee;
  }
  
  .table-moodle td::before {
    content: attr(data-label);
    float: left;
    font-weight: bold;
  }
}
```

```javascript
// En el JSX, agregar data-label:
<td data-label="Nombre">
  {`${u.firstname} ${u.lastname}`}
</td>
```

---

## 🔐 Control de Acceso por Roles

```javascript
const [userRole, setUserRole] = useState("admin"); // o "viewer"

// Condicionar botones según el rol
{userRole === "admin" && (
  <button onClick={handleDelete}>
    <FaTrash />
  </button>
)}

{userRole === "viewer" && (
  <p style={{ color: "#999" }}>
    No tienes permisos para eliminar usuarios
  </p>
)}
```

---

## 📊 Integración con Analytics

```javascript
// Al exportar CSV:
const exportToCSV = () => {
  // ... código de exportación
  
  // Tracking
  if (window.gtag) {
    window.gtag('event', 'export_users', {
      'event_category': 'engagement',
      'event_label': 'CSV Export',
      'value': users.length
    });
  }
};

// Al eliminar usuarios:
const handleDelete = async (u) => {
  // ... código de eliminación
  
  if (window.gtag) {
    window.gtag('event', 'delete_user', {
      'event_category': 'admin_action',
      'user_id': u.id
    });
  }
};
```

---

## 💡 Tips y Trucos

### 1. **Atajos de Teclado**
```javascript
useEffect(() => {
  const handleKeyPress = (e) => {
    if (e.ctrlKey && e.key === 'f') {
      e.preventDefault();
      document.querySelector('input[placeholder*="Buscar"]').focus();
    }
  };
  
  window.addEventListener('keydown', handleKeyPress);
  return () => window.removeEventListener('keydown', handleKeyPress);
}, []);
```

### 2. **Auto-refresh cada 30 segundos**
```javascript
useEffect(() => {
  const interval = setInterval(() => {
    fetchUsers();
  }, 30000); // 30 segundos
  
  return () => clearInterval(interval);
}, [fetchUsers]);
```

### 3. **Indicador de usuarios online**
```javascript
<td>
  {isUserOnline(u) && (
    <span style={{ 
      width: 8, 
      height: 8, 
      borderRadius: '50%', 
      background: '#4CAF50',
      display: 'inline-block',
      marginRight: 6
    }} />
  )}
  {u.firstname} {u.lastname}
</td>
```

---

¡Estos ejemplos te ayudarán a extender y personalizar tu vista de usuarios según tus necesidades! 🚀
