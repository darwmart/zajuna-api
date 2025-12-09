import React from 'react';
import { useNavigate } from 'react-router-dom';
import '../styles/moodle-theme.css';

function CoursesPage() {
  const navigate = useNavigate();

  return (
    <div className="users-category-page">
      <h2 className="category-main-title">Categoría: Cursos</h2>
      
      <div className="category-sections">
        <div className="category-section">
          <h3 className="category-section-title" onClick={() => navigate('/courses/manage')}>
            Administrar cursos y categorías
          </h3>
        </div>
        
        <div className="category-section">
          <h3
            className="category-section-title"
            onClick={() => navigate('/courses/categories/new')}
          >
            Añadir una categoría
          </h3>
        </div>
        
        <div className="category-section">
          <h3 className="category-section-title" onClick={() => navigate('#')}>
            Añade un curso nuevo
          </h3>
        </div>
        
        <div className="category-section">
          <h3 className="category-section-title" onClick={() => navigate('#')}>
            Restaurar curso
          </h3>
        </div>
        
        <div className="category-section">
          <h3 className="category-section-title" onClick={() => navigate('#')}>
            Descargar contenido del curso
          </h3>
        </div>
        
        <div className="category-section">
          <h3 className="category-section-title" onClick={() => navigate('#')}>
            Solicitud de curso
          </h3>
        </div>
        
        <div className="category-section">
          <h3 className="category-section-title" onClick={() => navigate('#')}>
            Requerimientos pendientes
          </h3>
        </div>
      </div>
    </div>
  );
}

export default CoursesPage;
