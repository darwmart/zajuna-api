import React, { useState, useEffect } from "react";
import SidebarSection from "../components/Sidebar/SidebarSection";
import SidebarNavItem from "../components/Sidebar/SidebarNavItem";
import SidebarSearch from "../components/Sidebar/SidebarSearch";
import { logout } from "../utils/auth";
import { getMyCourses, getCourseContent } from "../api/apiClientCourses";
import "../styles/moodle-theme.css";
import { useLocation } from 'react-router-dom';
import SidebarCalendar from '../components/Sidebar/SidebarCalendar';

function ZajunaSidebar() {
  const [expandedItems, setExpandedItems] = useState({
    areaPersonal: true,
    misCursos: true,
    adminSite: true, // Administración del sitio
    adminUsers: false,
    adminAccounts: false, // Cuentas
    adminCourses: false,
  });

  const [myCourses, setMyCourses] = useState([]);
  const [loadingCourses, setLoadingCourses] = useState(true);
  const [expandedCourses, setExpandedCourses] = useState({}); // Para controlar qué cursos están expandidos
  const [expandedParticipants, setExpandedParticipants] = useState({}); // Para controlar Participantes expandido
  const [courseSections, setCourseSections] = useState({}); // Para almacenar las secciones de cada curso
  const [loadingSections, setLoadingSections] = useState({}); // Para controlar qué cursos están cargando secciones

  const toggleItem = (key) => {
    setExpandedItems((prev) => ({
      ...prev,
      [key]: !prev[key],
    }));
  };

  // show the calendar card only when viewing enrolled users (route: /courses/:id/enrolled)
  const location = useLocation();
  const showCalendar = location && location.pathname && /\/courses\/[\w-]+\/enrolled/.test(location.pathname);

  const toggleCourse = async (courseId) => {
    const isCurrentlyExpanded = expandedCourses[courseId];

    // Toggle el estado de expansión
    setExpandedCourses((prev) => ({
      ...prev,
      [courseId]: !prev[courseId],
    }));

    // Si está colapsando, no hacer nada más
    if (isCurrentlyExpanded) {
      return;
    }

    // Si ya tenemos las secciones cargadas, no volver a cargarlas
    if (courseSections[courseId]) {
      return;
    }

    // Cargar las secciones del curso desde el API
    try {
      setLoadingSections((prev) => ({ ...prev, [courseId]: true }));
      const sections = await getCourseContent(courseId);
      setCourseSections((prev) => ({
        ...prev,
        [courseId]: sections,
      }));
    } catch (error) {
      console.error(`Error al cargar secciones del curso ${courseId}:`, error);
      // Si hay error, mantener array vacío para evitar repetir la carga
      setCourseSections((prev) => ({
        ...prev,
        [courseId]: [],
      }));
    } finally {
      setLoadingSections((prev) => ({ ...prev, [courseId]: false }));
    }
  };

  const toggleParticipants = (courseId) => {
    setExpandedParticipants((prev) => ({
      ...prev,
      [courseId]: !prev[courseId],
    }));
  };

  const handleLogout = (e) => {
    e.preventDefault();
    // Confirmar antes de hacer logout
    if (window.confirm('¿Estás seguro de que deseas cerrar sesión?')) {
      logout();
    }
  };

  // Cargar cursos del usuario al montar el componente
  useEffect(() => {
    const fetchMyCourses = async () => {
      try {
        setLoadingCourses(true);
        const response = await getMyCourses();

        // El backend retorna { courses: [...] }
        if (response && response.courses) {
          setMyCourses(response.courses);
        }
      } catch (error) {
        console.error('Error al cargar mis cursos:', error);
        // En caso de error, mantener array vacío
        setMyCourses([]);
      } finally {
        setLoadingCourses(false);
      }
    };

    fetchMyCourses();
  }, []); // Solo ejecutar al montar el componente

  return (
    <aside className="zajuna-sidebar">
      {/* Sección NAVEGACIÓN */}
      <SidebarSection title="NAVEGACIÓN" defaultOpen={true} className="nav-section">
        <SidebarNavItem 
          text="Área personal" 
          hasChildren={true}
          isExpanded={expandedItems.areaPersonal}
          onToggle={() => toggleItem('areaPersonal')}
        />
        {expandedItems.areaPersonal && (
          <>
            <SidebarNavItem
              icon="home"
              text="Página principal del sitio"
              href="#"
              onClick={handleLogout}
              level={1}
            />
            <SidebarNavItem 
              text="Páginas del sitio" 
              href="/pages" 
              level={1}
              showBullet={true}
            />
            <SidebarNavItem
              text="Mis cursos"
              hasChildren={true}
              isExpanded={expandedItems.misCursos}
              onToggle={() => toggleItem('misCursos')}
              level={1}
            />
            {expandedItems.misCursos && (
              <>
                {loadingCourses ? (
                  <SidebarNavItem
                    text="Cargando cursos..."
                    level={2}
                    showBullet={true}
                  />
                ) : myCourses.length > 0 ? (
                  myCourses.map((course) => (
                    <React.Fragment key={course.id}>
                      <SidebarNavItem
                        text={course.fullname || course.shortname}
                        hasChildren={true}
                        isExpanded={expandedCourses[course.id]}
                        onToggle={() => toggleCourse(course.id)}
                        level={2}
                      />
                      {expandedCourses[course.id] && (
                        <>
                          {/* Participantes (expandible) */}
                          <SidebarNavItem
                            text="Participantes"
                            hasChildren={true}
                            isExpanded={expandedParticipants[course.id]}
                            onToggle={() => toggleParticipants(course.id)}
                            level={3}
                          />
                          {expandedParticipants[course.id] && (
                            <>
                              <SidebarNavItem
                                icon="user"
                                text="Usuarios inscritos"
                                href={`/courses/${course.id}/participants`}
                                level={4}
                              />
                              <SidebarNavItem
                                icon="group"
                                text="Grupos"
                                href={`/courses/${course.id}/groups`}
                                level={4}
                              />
                            </>
                          )}

                          {/* Calificaciones (no expandible) */}
                          <SidebarNavItem
                            icon="table"
                            text="Calificaciones"
                            href={`/courses/${course.id}/grades`}
                            level={3}
                            showBullet={false}
                          />

                          {/* Secciones del curso - Cargadas dinámicamente desde el API */}
                          {loadingSections[course.id] ? (
                            <SidebarNavItem
                              text="Cargando secciones..."
                              level={3}
                              showBullet={true}
                            />
                          ) : courseSections[course.id] && courseSections[course.id].length > 0 ? (
                            courseSections[course.id].map((section) => (
                              <SidebarNavItem
                                key={section.id}
                                text={section.name || `Tema ${section.section}`}
                                href={`/courses/${course.id}/section/${section.section}`}
                                level={3}
                                showBullet={true}
                              />
                            ))
                          ) : (
                            <SidebarNavItem
                              text="No hay secciones disponibles"
                              level={3}
                              showBullet={true}
                            />
                          )}
                        </>
                      )}
                    </React.Fragment>
                  ))
                ) : (
                  <SidebarNavItem
                    text="No hay cursos matriculados"
                    level={2}
                    showBullet={true}
                  />
                )}
              </>
            )}
          </>
        )}
      </SidebarSection>

      {/* Sección ADMINISTRACIÓN */}
      <SidebarSection title="ADMINISTRACIÓN" defaultOpen={true} className="admin-section">
        {/* Administración del sitio (colapsable) */}
        <SidebarNavItem
          text="Administración del sitio"
          hasChildren={true}
          isExpanded={expandedItems.adminSite}
          onToggle={() => toggleItem('adminSite')}
          showBullet={true}
        />
        {expandedItems.adminSite && (
          <>
            {/* Usuarios (colapsable) */}
            <SidebarNavItem
              text="Usuarios"
              hasChildren={true}
              isExpanded={expandedItems.adminUsers}
              onToggle={() => toggleItem('adminUsers')}
              level={1}
              href="/users"
            />
            {expandedItems.adminUsers && (
              <>
                <SidebarNavItem 
                  text="Cuentas" 
                  hasChildren={true}
                  isExpanded={expandedItems.adminAccounts}
                  onToggle={() => toggleItem('adminAccounts')}
                  level={2}
                  href="/users/accounts"
                />
                {expandedItems.adminAccounts && (
                  <>
                    <SidebarNavItem icon="cog" text="Examinar lista de usuarios" href="/" level={3} />
                    <SidebarNavItem icon="cog" text="Acciones de usuario masivas" href="#" level={3} />
                    <SidebarNavItem icon="cog" text="Añade un nuevo usuario" href="/users/new" level={3} />
                    <SidebarNavItem icon="cog" text="Gestión de usuarios" href="#" level={3} />
                  </>
                )}
                <SidebarNavItem text="Permisos" href="#" level={2} showBullet={true} />
                <SidebarNavItem text="Privacidad y Políticas" href="#" level={2} showBullet={true} />
              </>
            )}

            {/* Cursos (colapsable) */}
            <SidebarNavItem
              text="Cursos"
              hasChildren={true}
              isExpanded={expandedItems.adminCourses}
              onToggle={() => toggleItem('adminCourses')}
              level={1}
              href="/courses"
            />
            {expandedItems.adminCourses && (
              <>
                <SidebarNavItem icon="cog" text="Administrar cursos y categorías" href="/courses/manage" level={2} />
                <SidebarNavItem icon="cog" text="Añadir una categoría" href="/courses/categories/new" level={2} />
                <SidebarNavItem icon="cog" text="Añade un curso nuevo" href="#" level={2} />
                <SidebarNavItem icon="cog" text="Restaurar curso" href="#" level={2} />
                <SidebarNavItem icon="cog" text="Descargar contenido del curso" href="#" level={2} />
                <SidebarNavItem icon="cog" text="Solicitud de curso" href="#" level={2} />
                <SidebarNavItem icon="cog" text="Requerimientos pendientes" href="#" level={2} />
              </>
            )}
          </>
        )}

        <SidebarSearch placeholder="Ajustes de búsqueda" />
      </SidebarSection>

      {/* Calendario funcional (componente) */}
      {showCalendar && <SidebarCalendar />}

      {/* Próximos eventos: aparece solo en la vista de usuarios matriculados */}
      {showCalendar && (
        <SidebarSection title="PRÓXIMOS EVENTOS" defaultOpen={true} className="upcoming-section">
          <div className="upcoming-events-box">
            <div className="upcoming-events-empty">No hay eventos próximos</div>
            <div className="upcoming-events-footer"><a href="/calendar">Ir al calendario...</a></div>
          </div>
        </SidebarSection>
      )}
    </aside>
  );
}

export default ZajunaSidebar;
