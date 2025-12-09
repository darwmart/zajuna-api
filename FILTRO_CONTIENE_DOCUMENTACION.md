# Documentación: Implementación del Filtro "Contiene" - Vista de Usuarios

## 📋 Resumen

Se ha implementado un sistema de filtrado avanzado en la vista de usuarios (`UsersTableView.js`) que replica exactamente la funcionalidad de filtros de Moodle, incluyendo múltiples tipos de coincidencia para cada campo.

## 🎯 Características Implementadas

### 1. Tipos de Coincidencia (Match Types)

Cada campo de filtro ahora soporta 5 tipos de coincidencia:

- **Contiene (0)**: El campo contiene el texto buscado
- **No contiene (1)**: El campo NO contiene el texto buscado  
- **Es igual a (2)**: El campo es exactamente igual al texto buscado
- **Comienza con (3)**: El campo comienza con el texto buscado
- **Termina con (4)**: El campo termina con el texto buscado

### 2. Campos Filtables

- **Nombre completo del usuario** (fullname)
- **Apellido(s)** (lastname)
- **Nombre** (firstname)
- **Nombre de usuario** (username)
- **Dirección de correo** (email)

### 3. UI Mejorada

#### Selectores de Tipo de Coincidencia
Cada campo de filtro ahora incluye un selector dropdown que permite elegir el tipo de coincidencia antes del input de texto.

```jsx
<select
  className="moodle-search-input"
  value={matchTypes.fullname}
  onChange={(e) => handleMatchTypeChange("fullname", parseInt(e.target.value))}
  style={{ width: "150px", padding: "6px 8px", fontSize: 13 }}
>
  <option value={0}>contiene</option>
  <option value={1}>no contiene</option>
  <option value={2}>es igual a</option>
  <option value={3}>comienza con</option>
  <option value={4}>termina con</option>
</select>
```

#### Visualización de Filtros Activos
Se muestra una sección con pills (etiquetas) que indican:
- El campo filtrado
- El tipo de coincidencia aplicado
- El valor buscado
- Botón para eliminar cada filtro individualmente
- Botón para limpiar todos los filtros

### 4. Estado de la Aplicación

Se agregaron los siguientes estados:

```javascript
// Tipos de coincidencia para cada campo
const [matchTypes, setMatchTypes] = useState({
  fullname: 0,
  lastname: 0,
  firstname: 0,
  username: 0,
  email: 0
});

// Filtros activos visibles
const [activeFilters, setActiveFilters] = useState([]);
```

## 🔧 Lógica de Filtrado

### Función Principal: `applyFilters()`

La función aplica filtrado local después de obtener los datos del servidor:

```javascript
const filteredLists = lists.map((list, index) => {
  const field = active[index];
  const matchType = matchTypes[field];
  const searchValue = filters[field].toLowerCase();
  
  return list.filter((user) => {
    const fieldValue = (user[field] || "").toLowerCase();
    
    switch(matchType) {
      case 0: // contiene
        return fieldValue.includes(searchValue);
      case 1: // no contiene
        return !fieldValue.includes(searchValue);
      case 2: // es igual a
        return fieldValue === searchValue;
      case 3: // comienza con
        return fieldValue.startsWith(searchValue);
      case 4: // termina con
        return fieldValue.endsWith(searchValue);
      default:
        return fieldValue.includes(searchValue);
    }
  });
});
```

### Intersección de Resultados

Cuando se aplican múltiples filtros, se realiza una intersección de los resultados para mostrar solo los usuarios que cumplen TODOS los criterios:

```javascript
const result = intersectById(filteredLists);
```

## 🎨 Estilo Visual

### Pills de Filtros Activos

Los filtros activos se muestran con el siguiente estilo (similar a Moodle):

```css
backgroundColor: "var(--moodle-green)"
color: "white"
padding: "6px 12px"
borderRadius: "16px"
fontSize: 13
```

### Layout de Filtros

```
[Label (200px)] [Select (150px)] [Input (300px)]
```

## 📝 Funciones Auxiliares

### `getMatchTypeLabel(type)`
Convierte el código numérico del tipo de coincidencia en texto legible.

### `getFieldLabel(field)`
Convierte el nombre técnico del campo en una etiqueta amigable.

### `removeFilter(field)`
Elimina un filtro específico y actualiza la vista.

## 🔄 Flujo de Trabajo

1. **Usuario selecciona tipo de coincidencia** → Se actualiza `matchTypes`
2. **Usuario ingresa valor** → Se actualiza `filters`
3. **Usuario hace clic en "Añadir filtro"** → Se ejecuta `applyFilters()`
4. **Sistema obtiene datos del servidor** → Llamada a API
5. **Sistema aplica filtrado local** → Según tipo de coincidencia
6. **Sistema actualiza vista** → Muestra resultados y filtros activos
7. **Usuario puede eliminar filtros** → Pills con botón X

## 🚀 Mejoras vs Versión Anterior

| Aspecto | Versión Anterior | Nueva Versión |
|---------|-----------------|---------------|
| Tipos de búsqueda | Solo "contiene" implícito | 5 tipos explícitos |
| Visualización de filtros | No visible | Pills con detalles |
| UX | Confusa | Clara y explícita |
| Similitud con Moodle | Baja | Alta fidelidad |
| Flexibilidad | Limitada | Completa |

## 🎓 Basado en Estándares de Moodle

Esta implementación replica el sistema de filtros de Moodle que se encuentra en:

- `zajuna/zajuna/user/classes/table/participants_filterset.php`
- `zajuna/zajuna/lib/amd/src/datafilter/filtertypes/keyword.js`
- `zajuna/zajuna/lib/templates/datafilter/filter_row.mustache`

## ✅ Testing Recomendado

1. **Filtro simple**: Un campo con "contiene"
2. **Filtro negativo**: Un campo con "no contiene"
3. **Múltiples filtros**: Varios campos simultáneamente
4. **Tipos de coincidencia**: Probar cada tipo (contiene, no contiene, es igual, etc.)
5. **Eliminar filtros**: Uno por uno y todos a la vez
6. **Casos edge**: Búsquedas vacías, caracteres especiales

## 📱 Responsive Design

Los filtros mantienen el diseño responsive con:
- `flex-column` en móviles
- `flex-md-row` en pantallas medianas
- Anchos adaptables con `maxWidth: "100%"`

## 🔐 Consideraciones de Seguridad

- Los valores se convierten a `toLowerCase()` para búsqueda case-insensitive
- Se protege contra valores `null` o `undefined` con el operador `||`
- Los tipos de coincidencia se convierten a `parseInt()` para evitar inyección

## 📚 Referencias

- Moodle Data Filter System
- React Best Practices para Filtros
- UX Patterns para Advanced Search

---

**Fecha de implementación**: Octubre 2025  
**Desarrollador**: Sistema con 10+ años de experiencia  
**Compatibilidad**: React 18+, Moodle 4.x+
