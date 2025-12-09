# Zajuna Frontend

Frontend web application for the Zajuna learning management system (LMS), built with a Moodle-inspired user interface. This project provides a modern, responsive interface for managing users, courses, and administrative tasks.

## 🚀 Características Principales

- **Navegación Tipo Moodle**: Sidebar jerárquico con navegación intuitiva y colapsable
- **Gestión de Usuarios**: CRUD completo con filtros avanzados y búsqueda
- **Sistema de Filtros**: Controles dinámicos con operadores personalizables
- **Páginas de Categorías**: Vistas organizadas para Usuarios, Cuentas y Cursos
- **Routing Hash-Based**: Compatible con la estructura de Moodle
- **Formularios Optimizados**: Validación, estados de focus y feedback visual
- **Tabla Compacta**: Visualización eficiente con paginación y hover states
- **Responsive Design**: Adaptable a diferentes tamaños de pantalla

## 🛠️ Tecnologías Utilizadas

- **React** 18.x
- **React Router DOM** v6 (HashRouter)
- **Axios** para llamadas a API REST
- **React Icons** para iconografía
- **CSS Custom** (moodle-theme.css) para estilos precisos
- **Create React App** como base

## 📋 Requisitos Previos

- Node.js >= 14.x
- npm >= 6.x
- Acceso al backend/API de Zajuna (Moodle)

## ⚙️ Instalación

1. Clonar el repositorio:
```bash
git clone https://github.com/Jr01sena/Zajuna.git
cd Zajuna
```

2. Instalar dependencias:
```bash
npm install
```

3. Configurar variables de entorno:
```bash
cp .env.example .env
```

Editar `.env` con tus credenciales:
```properties
REACT_APP_API=http://localhost/zajuna/zajuna/webservice/rest/server.php
REACT_APP_API_KEY=tu_api_key_aqui
```

4. Iniciar el servidor de desarrollo:
```bash
npm start
```

La aplicación se abrirá en [http://localhost:3000](http://localhost:3000)

## 📦 Comandos Disponibles

### `npm start`
Ejecuta la app en modo desarrollo.\
Los cambios se reflejan automáticamente sin recargar la página.

### `npm test`
Ejecuta el test runner en modo interactivo.

### `npm run build`
Genera la versión optimizada para producción en la carpeta `build/`.\
Los archivos están minificados y listos para deploy.

### `npm run eject`
⚠️ **Operación irreversible**. Expone la configuración interna de Webpack/Babel.

## 📁 Estructura del Proyecto

```
zajuna-frontend/
├── public/                    # Archivos estáticos
├── src/
│   ├── api/                  # Clientes de API (Axios)
│   │   ├── apiClientUsers.js
│   │   └── apiClientCourses.js
│   ├── assets/               # Logos y recursos
│   │   └── logos/
│   ├── components/           # Componentes reutilizables
│   │   ├── Sidebar/         # Navegación lateral
│   │   ├── UsersTableView.js
│   │   ├── UserForm.js
│   │   └── ...
│   ├── hooks/               # Custom hooks
│   │   └── useDebounce.js
│   ├── layout/              # Componentes de layout
│   │   ├── ZajunaHeader.js
│   │   ├── ZajunaSidebar.js
│   │   └── ZajunaLayout.js
│   ├── pages/               # Páginas principales
│   │   ├── UsersPage.js
│   │   ├── AccountsPage.js
│   │   ├── CoursesPage.js
│   │   ├── UserFormPage.js
│   │   └── UserEditPage.js
│   ├── services/            # Lógica de negocio/API
│   │   ├── usersService.js
│   │   └── coursesService.js
│   ├── styles/              # Estilos globales
│   │   └── moodle-theme.css
│   ├── App.js               # Componente raíz y routing
│   └── index.js             # Entry point
├── .env.example             # Plantilla de variables de entorno
├── .gitignore
├── package.json
└── README.md
```

## 🎨 Sistema de Estilos

El proyecto utiliza un tema custom (`moodle-theme.css`) con variables CSS:

```css
--moodle-green: #39A900        /* Verde Zajuna/Moodle */
--moodle-header-blue: #053449  /* Azul oscuro para iconos */
--moodle-grey-bg: #F5F7F6      /* Fondo de página */
--moodle-border: #D6D6D6       /* Bordes */
```

### Métricas de Componentes

**Filtros:**
- Select: 34px alto, 116px ancho, border-radius 10px
- Input: 34px alto, 200px ancho, border-radius 6px
- Gap: 4px entre controles

**Sidebar:**
- Width: 340px
- Padding por nivel: 20px (nivel 1), 32px (nivel 2), 44px (nivel 3)

**Tabla:**
- Filas body: padding vertical 3px, line-height 14px
- Columnas 2+: padding-left 44px

## 🔗 Rutas Disponibles

| Ruta | Componente | Descripción |
|------|-----------|-------------|
| `/` | UsersTableView | Lista principal de usuarios |
| `/users` | UsersPage | Categoría: Usuarios |
| `/users/accounts` | AccountsPage | Categoría: Cuentas |
| `/users/new` | UserFormPage | Formulario nuevo usuario |
| `/users/:id/edit` | UserEditPage | Editar usuario existente |
| `/courses` | CoursesPage | Categoría: Cursos |

## 🔧 Configuración de API

El proyecto consume la API REST de Moodle/Zajuna. Endpoints principales:

- **GET** `/webservice/rest/server.php?wsfunction=core_user_get_users`
- **POST** `/webservice/rest/server.php?wsfunction=core_user_create_users`
- **POST** `/webservice/rest/server.php?wsfunction=core_user_update_users`
- **POST** `/webservice/rest/server.php?wsfunction=core_user_delete_users`

Parámetros requeridos:
- `wstoken`: Token de autenticación
- `moodlewsrestformat`: json

## 📝 Notas de Desarrollo

### Decisiones Técnicas

1. **HashRouter vs BrowserRouter**: Se usa hash routing para compatibilidad con la estructura de directorios de Moodle.
2. **CSS Custom vs Library**: CSS puro permite control exacto del estilo Moodle sin sobrecarga.
3. **Client-side filtering**: `getUserById` filtra en cliente por limitaciones de la API (workaround temporal).
4. **Iconografía externa**: Iconos de ayuda/required fuera de inputs para mejor UX.

### Pendientes de Optimización

- [ ] Resolver ESLint warnings (unused vars, anchor validity)
- [ ] Optimizar `getUserById` con endpoint dedicado backend
- [ ] Implementar lazy loading en rutas
- [ ] Añadir tests unitarios para componentes clave

## 🤝 Contribución

1. Fork el proyecto
2. Crea una rama para tu feature (`git checkout -b feature/nueva-caracteristica`)
3. Commit tus cambios (`git commit -m 'feat: añadir nueva característica'`)
4. Push a la rama (`git push origin feature/nueva-caracteristica`)
5. Abre un Pull Request

## 📄 Licencia

Este proyecto es privado y pertenece a Zajuna/SENA.

## 👥 Autores

- Equipo de Desarrollo Zajuna

## 📞 Soporte

Para reportar bugs o solicitar features, abre un issue en el repositorio.

---

**Última actualización:** Octubre 2025  
**Estado:** ✅ Producción Ready
