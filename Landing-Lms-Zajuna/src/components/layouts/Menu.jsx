import React, { useState } from 'react';
import { Link, useLocation } from 'react-router-dom';
import '../../styles/layouts/menu.css';

// Usar rutas públicas /img

function Menu() {
  const [isMenuActive, setIsMenuActive] = useState(false);
  const location = useLocation();

  const toggleMenu = () => {
    setIsMenuActive(!isMenuActive);
  };

  const isActive = (path) => {
    if (path === '/soporte') {
      return location.pathname === path;
    }
    return false;
  };

  return (
    <div className={`list ${isMenuActive ? 'list--active' : ''}`}>
      <button 
        className="list__option list__option--menu" 
        id="list__menu-button"
        onClick={toggleMenu}
      >
        <img 
          loading="lazy" 
          className="list__option-button-icon" 
          src="/img/icons/hamburger-icon.svg" 
          alt="icono hamburguesa" 
        />
        <h4 className="list__option-info">Menú principal</h4>
        <div className="list_false-color">&nbsp;</div>
      </button>
      <a 
        className={`list__option ${location.pathname === '/' ? 'list__option--active' : ''}`}
        target="_blank" 
        href="https://zajuna.sena.edu.co/Repositorio/Titulada/institution/SENA/Tutoriales/Aprendiz/Manual_LMS_Aprendiz.pdf"
        rel="noopener noreferrer"
      >
        <img 
          loading="lazy" 
          className="list__option-button-icon" 
          src="/img/icons/manual-del-aprendiz-logo.svg" 
          alt="manual del aprendiz logo" 
        />
        <h4 className="list__option-info">Manual LMS Aprendiz</h4>
        <div className="list_false-color">&nbsp;</div>
      </a>
      <Link to="/soporte" className={`list__option ${isActive('/soporte') ? 'list__option--active' : ''}`}>
        <img 
          loading="lazy" 
          className="list__option-button-icon" 
          src="/img/icons/suport-online-icon.svg" 
          alt="soporte online icono" 
        />
        <h4 className="list__option-info">Soporte en línea</h4>
        <div className="list_false-color">&nbsp;</div>
      </Link>
      <a 
        href="https://betowa.sena.edu.co/oferta" 
        target="_blank" 
        className="list__option"
        rel="noopener noreferrer"
      >
        <img 
          loading="lazy" 
          className="list__option-button-icon" 
          src="/img/icons/known-course-icon.svg" 
          alt="nuestros programas icono" 
        />
        <h4 className="list__option-info">Nuestros Programas</h4>
        <div className="list_false-color">&nbsp;</div>
      </a>
      <a 
        href="https://biblioteca.sena.edu.co/" 
        target="_blank" 
        className="list__option"
        rel="noopener noreferrer"
      >
        <img 
          loading="lazy" 
          className="list__option-button-icon" 
          src="/img/icons/digital-librery-icon.svg" 
          alt="libresria digital icono" 
        />
        <h4 className="list__option-info">Biblioteca Digital</h4>
        <div className="list_false-color">&nbsp;</div>
      </a>
      <a 
        href="https://www.microsoft.com/es-co/microsoft-365/outlook/email-and-calendar-software-microsoft-outlook" 
        target="_blank" 
        className="list__option"
        rel="noopener noreferrer"
      >
        <img 
          loading="lazy" 
          className="list__option-button-icon" 
          src="/img/icons/email-mi-sena-icon.svg" 
          alt="correo icono" 
        />
        <h4 className="list__option-info">Correo @soy.sena</h4>
        <div className="list_false-color">&nbsp;</div>
      </a>
      <a 
        href="https://oferta.senasofiaplus.edu.co/sofia-oferta/certificaciones.html" 
        target="_blank" 
        className="list__option"
        rel="noopener noreferrer"
      >
        <img 
          loading="lazy" 
          className="list__option-button-icon" 
          src="/img/icons/online-certificate-icon.svg" 
          alt="certificado online icono" 
        />
        <h4 className="list__option-info">Certificados</h4>
        <div className="list_false-color">&nbsp;</div>
      </a>
      <a 
        href="http://sciudadanos.sena.edu.co/SolicitudIndex.aspx" 
        target="_blank" 
        className="list__option"
        rel="noopener noreferrer"
      >
        <img 
          loading="lazy" 
          className="list__option-button-icon" 
          src="/img/icons/pqrs-icon.svg" 
          alt="pqrs icono" 
        />
        <h4 className="list__option-info">PQRS SENA</h4>
        <div className="list_false-color">&nbsp;</div>
      </a>
    </div>
  );
}

export default Menu;
