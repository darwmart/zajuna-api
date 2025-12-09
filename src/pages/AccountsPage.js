import React from 'react';
import { useNavigate } from 'react-router-dom';
import '../styles/moodle-theme.css';

function AccountsPage() {
  const navigate = useNavigate();

  return (
    <div className="users-category-page">
      <h2 className="category-main-title">Categoría: Cuentas</h2>
      
      <div className="category-sections">
        <div className="category-section">
          <h3 className="category-section-title" onClick={() => navigate('/')}>
            Examinar lista de usuarios
          </h3>
        </div>
        
        <div className="category-section">
          <h3 className="category-section-title" onClick={() => navigate('#')}>
            Acciones de usuario masivas
          </h3>
        </div>
        
        <div className="category-section">
          <h3 className="category-section-title" onClick={() => navigate('/users/new')}>
            Añade un nuevo usuario
          </h3>
        </div>
        
        <div className="category-section">
          <h3 className="category-section-title" onClick={() => navigate('#')}>
            Gestión de usuarios
          </h3>
        </div>
      </div>
    </div>
  );
}

export default AccountsPage;
