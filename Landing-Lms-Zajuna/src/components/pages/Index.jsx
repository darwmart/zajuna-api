import React from 'react';
import '../../styles/index.css';
import Navbar from '../layouts/Navbar';
import Accesibilidad from '../layouts/Accesibilidad';
import Footer from '../layouts/Footer';
import LandingHero from '../sections/LandingHero';
import VideoCarousel from '../sections/VideoCarousel';
import InformationPosts from '../sections/InformationPosts';

/**
 * Index - Página principal de la aplicación (Landing Page)
 * 
 * Componente orquestador que organiza las diferentes secciones:
 * - Navbar y accesibilidad (layouts globales)
 * - LandingHero (menú, slider, login, redes sociales)
 * - VideoCarousel (carrusel de videos de cursos)
 * - InformationPosts (posts informativos)
 * - Footer (layout global)
 * 
 * La lógica compleja de cada sección está encapsulada en sus respectivos componentes,
 * mejorando la mantenibilidad y legibilidad del código.
 */
function Index() {
	return (
		<>
			<Navbar />
			<Accesibilidad />
			<LandingHero />
			<VideoCarousel />
			<InformationPosts />
			<Footer />
		</>
	);
}

export default Index;
