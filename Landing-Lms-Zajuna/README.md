# 🎓 LMS ZAJUNA

<div align="center">

![Version](https://img.shields.io/badge/version-3.3.7-blue.svg)
![React](https://img.shields.io/badge/React-19.1.1-61DAFB?logo=react&logoColor=white)
![Vite](https://img.shields.io/badge/Vite-7.1.12-646CFF?logo=vite&logoColor=white)
![React Router](https://img.shields.io/badge/React_Router-6.8.0-CA4245?logo=react-router&logoColor=white)
![License](https://img.shields.io/badge/license-Private-red.svg)

**Sistema de Gestión de Aprendizaje (LMS) moderno, accesible y optimizado**

[Características](#-características-principales) • [Instalación](#-instalación) • [Uso](#-uso) • [Tecnologías](#️-stack-tecnológico) • [Estructura](#-estructura-del-proyecto)

</div>

---

## 📖 Descripción

**LMS ZAJUNA** es una aplicación web moderna desarrollada con React 19 que proporciona una plataforma completa de gestión de aprendizaje. El sistema ofrece acceso a múltiples programas educativos, incluyendo formación técnica y tecnológica, cursos de bilingüismo, programas especiales como Campesena, y recursos de soporte para estudiantes e instructores.

La aplicación está diseñada como una **Single Page Application (SPA)** con navegación fluida, optimizada para rendimiento y accesibilidad, cumpliendo con los estándares modernos de desarrollo web.

---

## ✨ Características Principales

### 🎯 Interfaz y Experiencia de Usuario
- **Diseño Moderno**: Interfaz intuitiva y profesional con React 19
- **Responsive Design**: Adaptación completa a dispositivos móviles, tablets y desktop
- **Navegación Fluida**: SPA con React Router para transiciones sin recargas
- **Contenido Multimedia**: Videos demostrativos y recursos educativos integrados

### 📚 Programas Educativos
- **Formación Titulada**: Programas técnicos y tecnológicos completos
- **Bilingüismo**: 13 niveles de cursos de inglés estructurados
- **Programa Campesena**: Formación especializada para comunidades rurales
- **Cursos Virtuales**: Catálogo extenso de cursos disponibles en línea
- **Centro de Soporte**: Manuales, guías y recursos de ayuda

### ♿ Accesibilidad
- **Controles de Zoom**: Aumento y disminución de tamaño de fuente
- **Modo Alto Contraste**: Mejora la legibilidad para usuarios con necesidades visuales
- **Navegación por Teclado**: Soporte completo para navegación sin mouse
- **Atributos ARIA**: Implementación correcta para lectores de pantalla
- **Textos Alternativos**: Todas las imágenes incluyen descripciones apropiadas

### 🚀 Rendimiento y Optimización
- **Build Optimizado**: Vite 7 con Hot Module Replacement (HMR) ultrarrápido
- **Assets Optimizados**: Imágenes en formato WebP para carga rápida
- **Lazy Loading**: Carga diferida de imágenes y recursos
- **Code Splitting**: División automática del código para mejor rendimiento
- **Tree Shaking**: Eliminación de código no utilizado en producción

### 🔐 Seguridad y Autenticación
- **Sistema de Login**: Autenticación segura con tokens JWT
- **Gestión de Sesiones**: Almacenamiento seguro de credenciales
- **Integración Backend**: Comunicación con API REST externa
- **Validación de Formularios**: Validación client-side y server-side

---

## 🚀 Instalación

### Requisitos Previos

Asegúrate de tener instalado en tu sistema:

- **Node.js**: versión 18.0.0 o superior
- **npm**: versión 9.0.0 o superior (incluido automáticamente con Node.js)
- **Git**: para clonar el repositorio

### Pasos de Instalación

1. **Clonar el repositorio**
   ```bash
   git clone https://github.com/tu-usuario/zajuna-lms-react.git
   cd zajuna-lms-react
   ```

2. **Instalar dependencias**
   ```bash
   npm install
   ```
   
   Este comando instalará todas las dependencias necesarias:
   - React 19.1.1
   - React Router DOM 6.8.0
   - Vite 7.1.12
   - Y sus dependencias relacionadas

3. **Configurar variables de entorno (opcional)**
   
   Crea un archivo `.env` en la raíz del proyecto si necesitas configurar URLs personalizadas:
   ```env
   VITE_API_BASE_URL=http://localhost:8080
   VITE_ZAJUNA_DASHBOARD_URL=http://localhost:3000
   ```

4. **Iniciar servidor de desarrollo**
   ```bash
   npm run dev
   ```

5. **Abrir en el navegador**
   
   La aplicación estará disponible en:
   ```
   http://localhost:3001
   ```

---

## 💻 Uso

### Scripts Disponibles

El proyecto incluye los siguientes comandos npm:

| Comando | Descripción | Puerto |
|---------|-------------|--------|
| `npm run dev` | Inicia el servidor de desarrollo con HMR | 3001 |
| `npm run build` | Genera una versión optimizada para producción | - |
| `npm run preview` | Previsualiza el build de producción localmente | 3000 |

### Desarrollo

```bash
# Iniciar servidor de desarrollo
npm run dev
```

El servidor de desarrollo se iniciará en `http://localhost:3001` con las siguientes características:
- **Hot Module Replacement (HMR)**: Los cambios se reflejan instantáneamente sin recargar la página
- **Fast Refresh**: Recarga rápida de componentes React manteniendo el estado
- **Source Maps**: Mapeo de código para debugging fácil

### Producción

```bash
# Generar build de producción
npm run build

# Previsualizar el build
npm run preview
```

El comando `build` generará una carpeta `dist/` con:
- Código JavaScript minificado y optimizado
- CSS procesado y minificado
- Assets estáticos optimizados
- Archivos HTML procesados

El build está optimizado para:
- ✅ Tamaño mínimo de bundle
- ✅ Carga rápida de recursos
- ✅ Compatibilidad con navegadores modernos
- ✅ SEO optimizado

---

## 🛠️ Stack Tecnológico

### Tecnologías Principales

#### Frontend Framework
- **[React](https://react.dev/)** `19.1.1`
  - Framework principal para construcción de UI
  - Componentes funcionales con Hooks
  - Virtual DOM para rendimiento optimizado

#### Enrutamiento
- **[React Router DOM](https://reactrouter.com/)** `^6.8.0`
  - Navegación declarativa entre páginas
  - Soporte para rutas anidadas y parámetros
  - Navegación programática

#### Build Tool
- **[Vite](https://vitejs.dev/)** `^7.1.12`
  - Build tool ultrarrápido
  - Hot Module Replacement (HMR)
  - Optimización automática de assets
  - Soporte nativo para ES Modules

#### Procesamiento de Imágenes
- **[Sharp](https://sharp.pixelplumbing.com/)** `^0.34.4`
  - Optimización y conversión de imágenes
  - Conversión a formato WebP
  - Compresión avanzada

### Características Técnicas

- **JavaScript ES6+**: Arrow functions, destructuring, async/await, módulos
- **CSS Modular**: Estilos organizados por componentes
- **Custom Hooks**: Lógica reutilizable encapsulada
- **Componentes Funcionales**: Arquitectura moderna de React
- **Responsive Design**: Mobile-first approach
- **Accesibilidad**: WCAG 2.1 compliant

---

## 📁 Estructura del Proyecto

```
zajuna-lms-react/
│
├── public/                          # Assets estáticos públicos
│   ├── favicon.ico                 # Favicon del sitio
│   ├── img/                        # Imágenes
│   │   ├── banners/               # Banners de cursos (WebP, ~50 archivos)
│   │   ├── icons/                 # Iconos (SVG/WebP/PNG, ~51 archivos)
│   │   ├── logos/                 # Logos institucionales (SVG/WebP, 27 archivos)
│   │   ├── posts/                 # Imágenes de posts (WebP, 7 archivos)
│   │   └── backgrounds/          # Fondos y texturas (SVG)
│   ├── video/                     # Videos MP4 demostrativos (7 archivos)
│   ├── pdfs/                      # Documentos PDF (44 archivos)
│   │   ├── titulada/
│   │   │   ├── manuales/
│   │   │   ├── tecgnologias/    # Tecnologías (13 PDFs)
│   │   │   └── tecnico/         # Programas técnicos (28 PDFs)
│   └── fonts/                     # Fuentes tipográficas
│       ├── WorkSans-Regular.ttf
│       ├── WorkSans-Bold.ttf
│       └── WorkSans-SemiBold.otf
│
├── src/                            # Código fuente
│   ├── components/                # Componentes React
│   │   ├── pages/                 # Páginas principales
│   │   │   ├── Index.jsx         # Página principal
│   │   │   ├── Bilinguismo.jsx   # Cursos de inglés
│   │   │   ├── Campesena.jsx     # Programa Campesena
│   │   │   ├── Titulada.jsx      # Formación titulada
│   │   │   └── Soporte.jsx       # Centro de soporte
│   │   ├── layouts/               # Componentes de layout
│   │   │   ├── Navbar.jsx        # Barra de navegación
│   │   │   ├── Menu.jsx          # Menú lateral
│   │   │   ├── Footer.jsx        # Pie de página
│   │   │   ├── Accesibilidad.jsx # Controles de accesibilidad
│   │   │   ├── Slider.jsx        # Carrusel de imágenes
│   │   │   └── LogosInstitucionales.jsx
│   │   ├── courses/               # Componentes de cursos
│   │   │   ├── CourseCard.jsx
│   │   │   ├── CourseCardContent.jsx
│   │   │   ├── EnglishLevel.jsx
│   │   │   └── VideoCard.jsx
│   │   ├── programs/              # Componentes de programas
│   │   │   ├── ProgramCard.jsx
│   │   │   └── CampesenaProgram.jsx
│   │   ├── sections/              # Secciones reutilizables
│   │   │   ├── LandingHero.jsx
│   │   │   ├── VideoCarousel.jsx
│   │   │   └── InformationPosts.jsx
│   │   └── forms/                 # Formularios
│   │       └── LoginForm.jsx
│   │
│   ├── data/                      # Datos estáticos
│   │   ├── courses.js            # Catálogo de cursos
│   │   ├── englishLevels.js      # Niveles de inglés
│   │   ├── programasCampesena.js
│   │   └── programasTitulada.js
│   │
│   ├── hooks/                     # Custom Hooks
│   │   ├── useAccesibilidad.js  # Hook de accesibilidad
│   │   └── useSlider.js          # Hook para sliders
│   │
│   ├── utils/                     # Utilidades
│   │   └── authClient.js         # Cliente de autenticación
│   │
│   ├── styles/                    # Estilos CSS modulares
│   │   ├── index.css            # Estilos principales
│   │   ├── themes.css           # Variables y temas
│   │   ├── bilinguismo.css
│   │   ├── campesena.css
│   │   ├── titulada.css
│   │   ├── soporte.css
│   │   └── layouts/             # Estilos de layouts
│   │       ├── accesibilidad.css
│   │       ├── footer.css
│   │       ├── logos-institucionales.css
│   │       ├── menu.css
│   │       ├── navbar.css
│   │       └── slider.css
│   │
│   ├── App.jsx                    # Componente raíz
│   └── main.jsx                   # Punto de entrada
│
├── scripts/                       # Scripts de utilidad
│   └── optimize-images.mjs       # Script de optimización de imágenes
│
├── index.html                     # HTML raíz
├── package.json                   # Dependencias y scripts
├── package-lock.json              # Lock file de dependencias
├── vite.config.js                 # Configuración de Vite
├── .gitignore                     # Archivos ignorados por Git
└── README.md                      # Este archivo
```

---

## 🎨 Páginas y Rutas

La aplicación incluye las siguientes rutas principales:

| Ruta | Componente | Descripción |
|------|------------|-------------|
| `/` | `Index` | Página principal con hero section, carousel de videos y posts informativos |
| `/bilinguismo` | `Bilinguismo` | Catálogo completo de cursos de inglés (13 niveles disponibles) |
| `/campesena` | `Campesena` | Programa especial Campesena con información detallada |
| `/titulada` | `Titulada` | Formación técnica y tecnológica con programas estructurados |
| `/soporte` | `Soporte` | Centro de ayuda con manuales, guías y recursos de soporte |

---

## 🔐 Autenticación y Backend

### Integración con Backend

El proyecto se conecta con un backend externo para la autenticación y gestión de datos:

- **URL Base**: Configurable mediante variable de entorno `VITE_API_BASE_URL` (por defecto: `http://localhost:8080`)
- **Endpoint de Login**: `/api/v1/auth/login`
- **Método**: POST con credenciales JSON
- **Autenticación**: Token JWT almacenado en localStorage

### Funcionalidades de Autenticación

El módulo `authClient.js` proporciona:

- ✅ Login de estudiantes (por número de documento)
- ✅ Login administrativo (por usuario)
- ✅ Almacenamiento seguro de tokens
- ✅ Gestión de sesiones de usuario
- ✅ Headers de autenticación para peticiones
- ✅ Logout completo

### Configuración

Para configurar las URLs del backend y dashboard, crea un archivo `.env` en la raíz del proyecto:

```env
# URL base del backend API
VITE_API_BASE_URL=http://localhost:8080

# URL del dashboard de ZAJUNA (opcional)
VITE_ZAJUNA_DASHBOARD_URL=http://localhost:3000
```

**Nota**: El archivo `.env` está incluido en `.gitignore` y no se subirá al repositorio. Asegúrate de crear un `.env.example` si necesitas documentar las variables requeridas.

---

## 📊 Compatibilidad de Navegadores

### Navegadores Soportados

El proyecto está optimizado para los siguientes navegadores:

| Navegador | Versión Mínima | Estado |
|-----------|----------------|--------|
| Chrome | Últimas 2 versiones | ✅ Soportado |
| Firefox | Últimas 2 versiones | ✅ Soportado |
| Safari | 14+ | ✅ Soportado |
| Edge | Últimas 2 versiones | ✅ Soportado |
| Opera | Últimas 2 versiones | ✅ Soportado |
| Navegadores móviles | Versiones modernas | ✅ Soportado |

**Cobertura estimada**: >95% de usuarios globales

### Características Requeridas

- ES6+ support
- CSS Grid y Flexbox
- LocalStorage API
- Fetch API
- CSS Custom Properties (Variables)

---

## 🎯 Características Técnicas Detalladas

### Accesibilidad (WCAG 2.1)

- ✅ **Controles de Zoom**: Incremento/disminución de tamaño de fuente (1.4.4)
- ✅ **Alto Contraste**: Modo de alto contraste para mejor legibilidad (1.4.6)
- ✅ **Navegación por Teclado**: Navegación completa sin mouse (2.1.1)
- ✅ **Atributos ARIA**: Implementación correcta de roles y propiedades (4.1.2)
- ✅ **Textos Alternativos**: Todas las imágenes tienen alt text apropiado (1.1.1)
- ✅ **Contraste de Colores**: Cumplimiento de ratios mínimos (1.4.3)

### Rendimiento

- ✅ **Lazy Loading**: Carga diferida de imágenes (`loading="lazy"`)
- ✅ **Code Splitting**: División automática del código por rutas
- ✅ **Tree Shaking**: Eliminación de código no utilizado
- ✅ **Minificación**: Código minificado en producción
- ✅ **Compresión**: Assets comprimidos (WebP para imágenes)
- ✅ **Caching**: Headers de caché apropiados

### SEO

- ✅ **Meta Tags**: Descripciones y títulos apropiados
- ✅ **URLs Semánticas**: Rutas descriptivas y amigables
- ✅ **Estructura HTML**: Jerarquía correcta de encabezados
- ✅ **Performance**: Carga rápida (Core Web Vitals)

---

## 🤝 Contribución

Las contribuciones son bienvenidas y apreciadas. Para contribuir al proyecto:

### Proceso de Contribución

1. **Fork el proyecto**
   ```bash
   git fork https://github.com/tu-usuario/zajuna-lms-react.git
   ```

2. **Crea una rama para tu feature**
   ```bash
   git checkout -b feature/nueva-funcionalidad
   ```

3. **Realiza tus cambios**
   - Sigue los estándares de código establecidos
   - Añade comentarios cuando sea necesario
   - Actualiza la documentación si es requerido

4. **Commit tus cambios**
   ```bash
   git commit -m 'feat: añade nueva funcionalidad X'
   ```

5. **Push a tu rama**
   ```bash
   git push origin feature/nueva-funcionalidad
   ```

6. **Abre un Pull Request**
   - Describe claramente los cambios realizados
   - Menciona cualquier breaking change
   - Incluye screenshots si aplica

### Estándares de Código

- **JavaScript**: ES6+ con camelCase para variables y funciones
- **React**: Componentes funcionales con PascalCase
- **CSS**: Metodología BEM-like con variables CSS
- **Commits**: Mensajes descriptivos siguiendo Conventional Commits
- **Comentarios**: Código autodocumentado con comentarios cuando sea necesario

### Convenciones de Nombres

- **Componentes**: PascalCase (`CourseCard.jsx`)
- **Hooks**: camelCase con prefijo `use` (`useAccesibilidad.js`)
- **Utilidades**: camelCase (`authClient.js`)
- **Archivos CSS**: kebab-case (`navbar.css`)

---

## 📝 Licencia

Este proyecto es **privado** y de uso interno. Todos los derechos reservados.

---


## 🐛 Solución de Problemas

### Problemas Comunes

#### El servidor no inicia
```bash
# Verifica que el puerto 3001 esté disponible
lsof -i :3001

# O cambia el puerto en vite.config.js
```

#### Error al instalar dependencias
```bash
# Limpia la caché de npm
npm cache clean --force

# Elimina node_modules y reinstala
rm -rf node_modules package-lock.json
npm install
```

#### Problemas de conexión con el backend
- Verifica que el backend esté corriendo en el puerto 8080
- Revisa la configuración de CORS en el backend
- Verifica las variables de entorno en `.env`
- Revisa la consola del navegador para mensajes de error detallados

---

## 👥 Equipo y Contacto

Desarrollado con ❤️ para el sistema educativo ZAJUNA.

Para soporte técnico o consultas sobre el proyecto, por favor contacta al equipo de desarrollo.

---

## 🔄 Changelog

### Versión 3.3.7
- Migración completa a React 19.1.1
- Actualización a Vite 7.1.12
- Mejoras en sistema de autenticación
- Optimización de assets (conversión a WebP)
- Implementación de características de accesibilidad
- Mejoras en diseño responsive

---

<div align="center">

**⭐ Si este proyecto te resulta útil, considera darle una estrella**

*Última actualización: 2025*

[⬆ Volver arriba](#-lms-zajuna)

</div>
