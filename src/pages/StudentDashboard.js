import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { getMyCourses } from '../api/apiClientCourses';
import { getUser } from '../utils/auth';
import CourseCard from '../components/Student/CourseCard';
import '../styles/zajuna-theme.css';

/**
 * StudentDashboard - Vista principal para estudiantes
 * Replica el diseño del theme Zajuna-Nube de Moodle
 */
function StudentDashboard() {
  const navigate = useNavigate();
  const [courses, setCourses] = useState([]);
  const [loading, setLoading] = useState(true);
  const [user, setUser] = useState(null);

  useEffect(() => {
    // Cargar información del usuario
    const userData = getUser();
    if (!userData) {
      // Si no hay usuario, redirigir al login
      navigate('/');
      return;
    }
    setUser(userData);

    // Cargar cursos del estudiante
    loadCourses();
  }, [navigate]);

  const loadCourses = async () => {
    try {
      setLoading(true);
      const response = await getMyCourses();

      if (response && response.courses) {
        setCourses(response.courses);
      }
    } catch (error) {
      console.error('Error al cargar cursos:', error);
    } finally {
      setLoading(false);
    }
  };

  const handleCourseClick = (course) => {
    // Navegar al curso (por ahora solo mostramos en consola)
    console.log('Navegar al curso:', course);
    // TODO: Implementar navegación al curso
    // navigate(`/course/${course.id}`);
  };

  const handleLogout = () => {
    localStorage.clear();
    navigate('/');
  };

  // Obtener iniciales del usuario para el avatar
  const getUserInitials = () => {
    if (!user) return '?';

    const names = (user.firstname || '').split(' ');
    const lastnames = (user.lastname || '').split(' ');

    const firstInitial = names[0] ? names[0][0] : '';
    const lastInitial = lastnames[0] ? lastnames[0][0] : '';

    return (firstInitial + lastInitial).toUpperCase() || '?';
  };

  return (
    <div className="student-dashboard" style={{ minHeight: '100vh', backgroundColor: 'var(--zajuna-primary-background)' }}>
      {/* Header con logos */}
      <div className="zajuna-header">
        <img
          src="http://localhost/frontcmd/img/logos/zajuna-logo.svg"
          alt="Logo de Zajuna"
          className="zajuna-header__logo-zajuna"
          onError={(e) => {
            e.target.style.display = 'none';
          }}
        />
        <img
          src="https://zajuna.sena.edu.co/img/logos/sena-logo.svg"
          alt="Logo del SENA"
          className="zajuna-header__logo-sena"
          onError={(e) => {
            e.target.src = 'https://via.placeholder.com/150x60/04324d/ffffff?text=SENA';
          }}
        />
      </div>

      {/* Navbar azul oscuro */}
      <nav className="zajuna-navbar">
        <div className="zajuna-navbar__menu">
          <a href="#inicio" className="zajuna-navbar__link">Inicio</a>
          <a href="#cursos" className="zajuna-navbar__link">Mis Cursos</a>
        </div>

        <div className="zajuna-navbar__user">
          {user && (
            <>
              <span className="zajuna-navbar__user-name">
                {user.firstname} {user.lastname}
              </span>
              <div className="zajuna-navbar__user-initials">
                {getUserInitials()}
              </div>
              <button
                onClick={handleLogout}
                className="zajuna-btn zajuna-btn--secondary"
                style={{ marginLeft: '0.5rem' }}
              >
                Salir
              </button>
            </>
          )}
        </div>
      </nav>

      {/* Mensaje de bienvenida */}
      <div className="zajuna-welcome">
        <h2 className="zajuna-welcome__message">
          Hola, {user?.firstname || 'Estudiante'}
          <span role="img" aria-label="wave" style={{ fontSize: '2rem' }}>👋</span>
        </h2>
      </div>

      {/* Contenido principal con 3 columnas */}
      <div className="zajuna-page-content">
        {/* Sidebar izquierdo - Navegación simple para estudiantes */}
        <aside className="zajuna-sidebar zajuna-sidebar--left">
          <div className="zajuna-block">
            <h3 className="zajuna-block__title">Navegación</h3>
            <div className="zajuna-block__content">
              <ul className="zajuna-nav-tree">
                <li className="zajuna-nav-tree__item zajuna-nav-tree__item--active">
                  <a href="#inicio" className="zajuna-nav-tree__link">Área personal</a>
                </li>
                <li className="zajuna-nav-tree__item">
                  <a href="#cursos" className="zajuna-nav-tree__link">Mis cursos</a>
                </li>
              </ul>
            </div>
          </div>
        </aside>

        {/* Contenido central - Lista de cursos */}
        <main className="zajuna-main-content" id="cursos">
          <div className="zajuna-block">
            <h3 className="zajuna-block__title">Mis Cursos</h3>
            <div className="zajuna-block__content">
              {loading ? (
                <div style={{ textAlign: 'center', padding: '2rem' }}>
                  <p>Cargando cursos...</p>
                </div>
              ) : courses.length > 0 ? (
                <div className="zajuna-courses-grid">
                  {courses.map((course) => (
                    <CourseCard
                      key={course.id}
                      course={course}
                      onClick={handleCourseClick}
                    />
                  ))}
                </div>
              ) : (
                <div style={{ textAlign: 'center', padding: '2rem' }}>
                  <p>No estás matriculado en ningún curso todavía.</p>
                </div>
              )}
            </div>
          </div>
        </main>

        {/* Sidebar derecho - Calendario y actividades */}
        <aside className="zajuna-sidebar zajuna-sidebar--right">
          {/* Bloque de calendario */}
          <div className="zajuna-block">
            <h3 className="zajuna-block__title">Calendario</h3>
            <div className="zajuna-block__content">
              <p style={{ fontSize: '0.875rem', color: 'var(--zajuna-secondary-text)' }}>
                No hay eventos próximos
              </p>
            </div>
          </div>

          {/* Bloque de actividades próximas */}
          <div className="zajuna-block">
            <h3 className="zajuna-block__title">Próximas Actividades</h3>
            <div className="zajuna-block__content">
              <p style={{ fontSize: '0.875rem', color: 'var(--zajuna-secondary-text)' }}>
                No hay actividades pendientes
              </p>
            </div>
          </div>
        </aside>
      </div>

      {/* Footer */}
      <footer className="zajuna-footer">
        <div className="zajuna-footer__logo">
          <img
            src="https://zajuna.sena.edu.co/img/logos/zajuna-logo.svg"
            alt="Logo Zajuna"
            onError={(e) => {
              e.target.style.display = 'none';
            }}
          />
        </div>
        <div className="zajuna-footer__content">
          <p style={{ margin: 0, fontSize: '0.875rem' }}>
            © {new Date().getFullYear()} Zajuna - SENA. Todos los derechos reservados.
          </p>
          {user && (
            <p style={{ margin: 0, fontSize: '0.75rem', opacity: 0.8 }}>
              Conectado como: {user.username}
            </p>
          )}
        </div>
      </footer>
    </div>
  );
}

export default StudentDashboard;
