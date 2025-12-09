import React from 'react';
import { useNavigate } from 'react-router-dom';
import '../styles/moodle-theme.css';

function UsersPage() {
  const navigate = useNavigate();

  return (
    <div className="users-category-page">
      <h2 className="category-main-title">Categoría: Usuarios</h2>
      
      <div className="category-sections">
        <div className="category-section">
          <h3 className="category-section-title" onClick={() => navigate('/users/accounts')}>
            Categoría: Cuentas
          </h3>
        </div>
        
        <div className="category-section">
          <h3 className="category-section-title" onClick={() => navigate('#')}>
            Categoría: Permisos
          </h3>
        </div>
        
        <div className="category-section">
          <h3 className="category-section-title" onClick={() => navigate('#')}>
            Categoría: Privacidad y Políticas
          </h3>
        </div>
      </div>
    </div>
  );
}

export default UsersPage;
