import React from 'react';

function CourseCardContent({ course, isHovered }) {
  return (
    <>
      {/* Link de inscripción */}
      <a
        className="cursos-virtuales__card-link"
        href={course.href}
        target="_blank"
        rel="noopener noreferrer"
        aria-label={`Inscribirse en ${course.title}`}
      />

      {/* Overlay con efecto de hover */}
      <div 
        className={`cursos-virtuales__card-overlay ${isHovered ? 'cursos-virtuales__card-overlay--visible' : ''}`}
      >
        <div className="cursos-virtuales__card-content">
          {/* Badge 100% virtual */}
          <div 
            className={`cursos-virtuales__badge ${isHovered ? 'cursos-virtuales__badge--visible' : ''}`}
          >
            <span className="cursos-virtuales__badge-text">
              <img 
                src="/img/icons/Globe.png" 
                alt="Virtual" 
                className="cursos-virtuales__badge-icon"
              />
              100% virtual
            </span>
          </div>

          {/* Título del curso */}
          <h3 
            className={`cursos-virtuales__card-title ${isHovered ? 'cursos-virtuales__card-title--visible' : ''}`}
          >
            {course.title}
          </h3>

          {/* Botón de inscripción */}
          <div 
            className={`cursos-virtuales__card-button-container ${isHovered ? 'cursos-virtuales__card-button-container--visible' : ''}`}
          >
            <a 
              href={course.href} 
              target="_blank"
              rel="noopener noreferrer"
              className="cursos-virtuales__card-button"
            >
              Inscribirme
            </a>
          </div>
        </div>
      </div>
    </>
  );
}

// Memoizar para evitar re-renders innecesarios
export default React.memo(CourseCardContent);

