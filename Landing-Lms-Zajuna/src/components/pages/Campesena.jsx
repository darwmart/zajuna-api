import React from 'react';
import '../../styles/campesena.css';
import Navbar from '../layouts/Navbar';
import Accesibilidad from '../layouts/Accesibilidad';
import Slider from '../layouts/Slider';
import Footer from '../layouts/Footer';
import CampesenaProgram from '../programs/CampesenaProgram';
import { getProgramasCampesena } from '../../data/programasCampesena';

// Imágenes del slider de Campesena
const CAMPESENA_SLIDER_IMAGES = ['/img/banners/campesena-radial.webp'];

function Campesena() {
	const programas = getProgramasCampesena();

	return (
		<>
			<Navbar />
			<Accesibilidad />
			<div className="campesena">
				<div className="campesena__slider">
					<Slider images={CAMPESENA_SLIDER_IMAGES} altText="Banner Campesena Radial" />
				</div>
				<div className="campesena__card">
					<div className="campesena__info">
						<img loading="lazy" src="/img/icons/campesena.svg" className="campesena__info-img" alt="Icono Campesena" />
						<p className="campesena__info-tittle">¿Que es Campesena Radial?</p>
						<div className="campesena__info-tittle2">
							<p className="campesena__info-text">
								Es una iniciativa transformadora para nuestros campesinos y campesinas, marcando un antes y un después en su desarrollo profesional y personal.
								<br />
								<br />
								Esta estrategia resalta la importancia de la radio como un poderoso medio de difusión cultural y educativo, capaz de fortalecer competencias, conocimientos y habilidades  esenciales en el campo, mejorando así la competitividad y productividad del sector rural.
								<br />
								<br />
								Para estimular el aprendizaje, la estrategia cuenta con diferentes materiales y recursos que buscan una participación activa de la comunidad campesina como:
							</p>
							<div className="campesena__info-video">
								<iframe 
									width="560" 
									height="315" 
									style={{borderRadius:'12px'}} 
									src="https://www.youtube.com/embed/gtYKOstz4E8?si=TAg1tGTibg_gehEg" 
									title="YouTube video player" 
									frameBorder="0" 
									allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" 
									referrerPolicy="strict-origin-when-cross-origin" 
									allowFullScreen
								/>
							</div>
						</div>
					</div>
				</div>

				<div className="campesena__informativa">
					<div className="campesena__informativa-cont">
						<div className="campesena__informativa-container">
							<div className="campesena__informativa-img">
								<img loading="lazy" className="campesena__informativa-icon" src="/img/icons/1.webp" alt="Icono 1" width="60px" />
							</div>
							<div className="campesena__informativa-post">
								<p className="campesena__informativa-text">
									<strong>Material de apoyo:</strong>
									<span style={{fontWeight:100}}>son las cartillas digital e impresa en el que se encuentra el contenido técnico para fortalecer las competencias de cada programa de formación.</span>
								</p>
							</div>
						</div>
						<div className="campesena__informativa-container">
							<div className="campesena__informativa-img">
								<img loading="lazy" className="campesena__informativa-icon" src="/img/icons/2.webp" alt="Icono 2" width="60px" />
							</div>
							<div className="campesena__informativa-post">
								<p className="campesena__informativa-text">
									<strong>Efectos de sonido y música ambiental:</strong>
									<span style={{fontWeight:100}}>se recrean ambientes rurales para crear una experiencia auditiva inmersiva y atractiva,
										manteniendo la atención y motivación de los participantes</span>
								</p>
							</div>
						</div>
						<div className="campesena__informativa-container">
							<div className="campesena__informativa-img">
								<img loading="lazy" className="campesena__informativa-icon" src="/img/icons/3.webp" alt="Icono 3" width="60px" />
							</div>
							<div className="campesena__informativa-post">
								<p className="campesena__informativa-text">
									<strong>Encuentros presenciales de interacción:</strong>
									<span style={{fontWeight:100}}>se fomentan espacios presenciales para que los campesinos intercambien ideas,
										compartan experiencias y se apoyen mutuamente en su proceso de aprendizaje.</span>
								</p>
							</div>
						</div>
						<div className="campesena__informativa-container">
							<div className="campesena__informativa-img">
								<img loading="lazy" className="campesena__informativa-icon" src="/img/icons/4.webp" alt="Icono 4" width="60px" />
							</div>
							<div className="campesena__informativa-post">
								<p className="campesena__informativa-text">
									<strong>Material de apoyo:</strong>
									<span style={{fontWeight:100}}>son las cartillas digital e impresa en el que se encuentra el
										contenido técnico para fortalecer las competencias de cada programa de formación.</span>
								</p>
							</div>
						</div>
						<div className="campesena__informativa-container">
							<div className="campesena__informativa-img">
								<img loading="lazy" className="campesena__informativa-icon" src="/img/icons/5.webp" alt="Icono 5" width="60px" />
							</div>
							<div className="campesena__informativa-post">
								<p className="campesena__informativa-text">
									<strong>Programas de radio:</strong>
									<span style={{fontWeight:100}}>una parrilla de programas radiales que se transmitirán a través de convenios con emisoras de todo el país, donde los aprendices podrán escuchar las experiencias y el contenido diseñado para apoyar el proceso formativo</span>
								</p>
							</div>
						</div>
						<div className="campesena__informativa-container">
							<div className="campesena__informativa-img">
								<img loading="lazy" className="campesena__informativa-icon" src="/img/icons/6.webp" alt="Icono 6" width="60px" />
							</div>
							<p className="campesena__informativa-text">
								<strong>Aplicación móvil:</strong>
								<span style={{fontWeight:100}}>una aplicación que contiene podcast, cartilla digital, glosario y
									actividad interactiva, permitiendo que el aprendiz consulte el material sin necesidad de tener acceso a internet.</span>
							</p>
						</div>
					</div>

					{/* Sección de Programas 2024 */}
					<div className="campesena__program">
						<div className="campesena__program-tittle">
							<p className="campesena__program-text">Programas 2024</p>
						</div>
						{programas.map((programa) => (
							<CampesenaProgram
								key={programa.id}
								title={programa.title}
								icon={programa.icon}
								courses={programa.courses}
							/>
						))}
					</div>
				</div>

				<div className="campesena__banner">
					<img loading="lazy" className="campesena__banner-img" src="/img/banners/banner-elementos-campesena.webp" alt="Banner elementos Campesena" />
				</div>

				<div className="campesena__buttons-cont">
					<div className="campesena__buttons">
						<img loading="lazy" className="campesena__buttons-img" src="/img/icons/Formación-tecnica.webp" alt="Formación técnica" width="50px" />
						<p className="campesena__buttons-text">Formación Técnica</p>
					</div>
					<div className="campesena__buttons">
						<img loading="lazy" className="campesena__buttons-img" src="/img/icons/Historia-escenarios.webp" alt="Historia en escenarios" width="50px" />
						<p className="campesena__buttons-text">Historia en escenarios
							<br />
							del campo colombiano</p>
					</div>
					<div className="campesena__buttons">
						<img loading="lazy" className="campesena__buttons-img" src="/img/icons/aprendizaje-sensorial.webp" alt="Aprendizaje sensorial" width="50px" />
						<p className="campesena__buttons-text">Aprendizaje
							sensorial</p>
					</div>
					<div className="campesena__buttons">
						<img loading="lazy" className="campesena__buttons-img" src="/img/icons/actividades-contexto.webp" alt="Actividades con contexto" width="50px" />
						<p className="campesena__buttons-text">Actividades con
							<br />
							contextos en región</p>
					</div>
				</div>

				<div className="campesena__cafetera-post">
					<img loading="lazy" className="campesena__cafetera-img" src="/img/banners/practicas-sostenibles.webp" alt="Prácticas sostenibles" />
					<div className="campesena__cafetera-buttons">
						<a href="" className="campesena__cafetera-programas">
							<img loading="lazy" className="campesena__cafetera-programas-img" src="/img/icons/programas.webp" alt="Programas" />
							<p>Programas radiales
								<br />
								4 programas de 30 minutos</p>
						</a>
						<a href="" className="campesena__cafetera-podcast">
							<img loading="lazy" className="campesena__cafetera-podcast-img" src="/img/icons/podcast.webp" alt="Podcast" />
							<p>podcast
								<br />
								18 audios</p>
						</a>
						<a href="" className="campesena__cafetera-descargar">
							<img loading="lazy" className="campesena__cafetera-descargar-img" src="/img/icons/Descargar.webp" alt="Descargar" />
							<p>Descargar
								<br />
								Cartilla</p>
						</a>
						<a href="" className="campesena__cafetera-aplicacion">
							<img loading="lazy" className="campesena__cafetera-aplicacion-img" src="/img/icons/Aplicacion.webp" alt="Aplicación" />
							<p>Descargar
								<br />
								Aplicacion</p>
						</a>
					</div>
				</div>
			</div>
			<Footer />
		</>
	);
}

export default Campesena;
