import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import '../../styles/layouts/navbar.css';

// Importar imágenes
// Usar rutas públicas para imágenes

function Navbar() {
  const [isMenuOpen, setIsMenuOpen] = useState(false);

  const toggleMenu = () => {
    setIsMenuOpen(!isMenuOpen);
  };

  return (
    <>
      <header className="gov" id="inicio">
        <div className="gov__container">
          <a href="https://www.gov.co/" target="_blank" rel="noopener noreferrer">
            <img loading="lazy" src="/img/logos/gov-logo.svg" alt="Logo de la pagina gov.co" className="gov__img" />
          </a>
        </div>
      </header>
      <div className="navbar-brand">
        <img loading="lazy" src="/img/logos/zajuna-logo.svg" alt="Logo de zajuna" className="navbar-brand__logo navbar__cpv--logo" />
        <img loading="lazy" src="/img/logos/sena-logo.svg" alt="Logo del servicio nacional de aprendizaje (SENA)" className="navbar-brand__logo" />
      </div>
      <nav className="navbar">
        <div className="navbar__container">
          <Link to="/" className="navbar_logo-container">
            {/* Logo removido - no se necesita imagen aquí */}
          </Link>
          <button 
            className="navbar__menu-toggle" 
            aria-label="Toggle menu" 
            id="navbar__menu-toggle"
            onClick={toggleMenu}
          >
            <img loading="lazy" className="navbar__icon" src="/img/icons/hamburger-icon.svg" alt=" icono hamburguesa" />
          </button>
          <ul className={`navbar__links ${isMenuOpen ? 'navbar__links--active' : ''}`}>
            <li><Link to="/" className="navbar__link">Inicio</Link></li>
            <hr className="navbar__separadores" />
            <li><a target="_blank" href="https://ejecucionformacion.sena.edu.co/cursos-cortos" className="navbar__link" rel="noopener noreferrer">Cursos cortos</a></li>
            <hr className="navbar__separadores" />
            <li><Link to="/bilinguismo" className="navbar__link">Bilingüismo</Link></li>
            <hr className="navbar__separadores" />
            <li><Link to="/titulada" className="navbar__link">Titulada</Link></li>
            <hr className="navbar__separadores" />
            <li><a href="https://ejecucionformacion.sena.edu.co/comunidad-aprendices" className="navbar__link" target="_blank" rel="noopener noreferrer">Comunidad aprendices</a></li>
            <hr className="navbar__separadores" />
            <li><a href="https://ejecucionformacion.sena.edu.co/comunidad-instructores" className="navbar__link" target="_blank" rel="noopener noreferrer">Comunidad instructores</a></li>
            <hr className="navbar__separadores" />
            <li><Link to="/campesena" className="navbar__link">Comunidad campesena</Link></li>
          </ul>
        </div>
      </nav>
    </>
  );
}

export default Navbar;
