import React from 'react';
import '../../styles/zajuna-theme.css';

/**
 * CourseCard - Tarjeta de curso para vista de estudiante
 * Replica el diseño de las tarjetas de curso del theme Zajuna-Nube
 */
function CourseCard({ course, onClick }) {
  // Imagen por defecto si no hay imagen del curso
  const courseImage = course.courseimage || course.imageurl || 'https://via.placeholder.com/300x150/39a900/ffffff?text=Curso';

  // Calcular progreso (si existe)
  const progress = course.progress || 0;

  // Manejar click en la tarjeta
  const handleClick = () => {
    if (onClick) {
      onClick(course);
    }
  };

  return (
    <div className="zajuna-course-card" onClick={handleClick}>
      {/* Imagen del curso */}
      <img
        src={courseImage}
        alt={course.fullname || course.shortname}
        className="zajuna-course-card__image"
        onError={(e) => {
          e.target.src = 'https://via.placeholder.com/300x150/39a900/ffffff?text=Curso';
        }}
      />

      {/* Cuerpo de la tarjeta */}
      <div className="zajuna-course-card__body">
        {/* Nombre corto (opcional) */}
        {course.shortname && course.shortname !== course.fullname && (
          <div className="zajuna-course-card__category">
            {course.shortname}
          </div>
        )}

        {/* Nombre completo del curso */}
        <h3 className="zajuna-course-card__title">
          {course.fullname || course.shortname || 'Curso sin nombre'}
        </h3>

        {/* Categoría del curso */}
        {course.categoryname && (
          <div className="zajuna-course-card__category">
            {course.categoryname}
          </div>
        )}

        {/* Badge si está oculto */}
        {course.visible === 0 && (
          <div className="zajuna-mb-1">
            <span style={{
              backgroundColor: '#3498db',
              color: 'white',
              padding: '0.25rem 0.5rem',
              borderRadius: '4px',
              fontSize: '0.75rem',
              fontWeight: 'bold'
            }}>
              Oculto para estudiantes
            </span>
          </div>
        )}

        {/* Barra de progreso */}
        {progress > 0 && (
          <div className="zajuna-course-card__progress">
            <div className="zajuna-course-card__progress-bar">
              <div
                className="zajuna-course-card__progress-fill"
                style={{ width: `${progress}%` }}
              />
            </div>
            <div style={{
              fontSize: '0.75rem',
              color: 'var(--zajuna-secondary-text)',
              marginTop: '0.25rem'
            }}>
              {progress}% completado
            </div>
          </div>
        )}
      </div>
    </div>
  );
}

export default CourseCard;
