import { useEffect, useState } from 'react';
import { useLocation } from 'react-router-dom';
import CoursesView from '../components/CoursesView';
import '../styles/moodle-theme.css';

export default function CoursesManagePage() {
  const location = useLocation();
  const [successMessage, setSuccessMessage] = useState('');

  useEffect(() => {
    // Si viene de crear categoría, mostrar mensaje de éxito
    if (location.state?.successMessage) {
      setSuccessMessage(location.state.successMessage);
      // Limpiar el mensaje después de 5 segundos
      setTimeout(() => {
        setSuccessMessage('');
      }, 5000);
    }
  }, [location.state]);

  return (
    <div className="users-category-page">
      {/* Mensaje de éxito */}
      {successMessage && (
        <div style={{
          padding: '12px 16px',
          marginBottom: '16px',
          backgroundColor: '#d4edda',
          color: '#155724',
          border: '1px solid #c3e6cb',
          borderRadius: '4px',
          display: 'flex',
          alignItems: 'center',
          gap: '8px'
        }}>
          <svg width="20" height="20" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
            <circle cx="12" cy="12" r="10" fill="#28a745"/>
            <path d="M9 12l2 2 4-4" stroke="white" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"/>
          </svg>
          <span>{successMessage}</span>
        </div>
      )}

      <CoursesView initialCategories={location.state?.updatedCategories} />
    </div>
  );
}

