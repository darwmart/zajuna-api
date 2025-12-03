import React, { useState } from 'react';
import '../../styles/soporte.css';
import Navbar from '../layouts/Navbar';
import Accesibilidad from '../layouts/Accesibilidad';
import Footer from '../layouts/Footer';

// Constantes para los tipos de manuales
const MANUAL_TYPES = {
	INSTRUCTOR: 'instructor',
	APRENDIZ: 'aprendiz',
	SAVA: 'sava'
};

function Soporte() {
	// Estado para controlar qué manual está activo
	const [activeManual, setActiveManual] = useState(MANUAL_TYPES.INSTRUCTOR);

	return (
		<>
			<Navbar />
			<Accesibilidad />
			<main className="soporte">
				<div className="soporte__container">
					<img src="/img/banners/banner-soporte.webp" alt="" className="soporte__img" />
				</div>

				<div className="soporte__informativo">
					<div className="soporte__informativo-item">
						<a className="soporte__informativo-container" href="https://gestorvirtual.sena.edu.co/agenti_lite_senaliz/"
							target="_blank" rel="noopener noreferrer">
							<img loading="lazy" src="/img/logos/agente-virtual-logo.webp" alt=""
								className="soporte__informativo-img" />
							<p className="soporte__informativo-text">Recibe ayuda de nuestro agente automático</p>
						</a>
					</div>

					<div className="soporte__informativo-item">
						<a className="soporte__informativo-container" href="https://sciudadanos.sena.edu.co/SolicitudIndex.aspx"
							target="_blank" rel="noopener noreferrer">
							<img loading="lazy" src="/img/logos/portal-pqrs-logo.webp" alt="" className="soporte__informativo-img" />
							<p className="soporte__informativo-text">Portal PQRS del SENA</p>
						</a>
					</div>

					<div className="soporte__informativo-item">
						<a className="soporte__informativo-container" href="https://sava.sena.edu.co/gfvd/" target="_blank" rel="noopener noreferrer">
							<img loading="lazy" src="/img/logos/sava-logo.webp" alt="" className="soporte__informativo-img" />
							<p className="soporte__informativo-text">Soporte Ambiente Virtual de Aprendizaje al
								instructor</p>
						</a>
					</div>
				</div>

				<div className="soporte__manual">
					<div className="soporte__manual-container">
						<h5 className="soporte__manual-text">Manuales de usuario</h5>
					</div>

					<div className="capsulas__informativas__margenes">
						<div className="capsulas__informativas__container">
							<div className="capsulas__informativas__navegacion">
								<div 
									className={`capsulas__informativas__manual__instructor ${activeManual === MANUAL_TYPES.INSTRUCTOR ? 'capsulas__informativas__manual__instructor--activo' : ''}`}
									onClick={() => setActiveManual(MANUAL_TYPES.INSTRUCTOR)}
								>
									Manual Instructor
								</div>
								<div 
									className={`capsulas__informativas__manual__aprendiz ${activeManual === MANUAL_TYPES.APRENDIZ ? 'capsulas__informativas__manual__aprendiz--activo' : ''}`}
									onClick={() => setActiveManual(MANUAL_TYPES.APRENDIZ)}
								>
									Manual Aprendiz
								</div>
								<div 
									className={`capsulas__informativas__manual__sava ${activeManual === MANUAL_TYPES.SAVA ? 'capsulas__informativas__manual__sava--activo' : ''}`}
									onClick={() => setActiveManual(MANUAL_TYPES.SAVA)}
								>
									Manual Sava
								</div>
							</div>

							<div className="capsulas__informativas__videos">
								{/* Manual Instructor */}
								<div className={`instructor__pdf ${activeManual !== MANUAL_TYPES.INSTRUCTOR ? 'manuales__container--oculto' : ''}`}>
									<div className="manuales__container">
										<div className="manuales__pdf">
											<div className="manual__ins__cont">
												<a target="_blank"
													href="/pdfs/titulada/manuales/MANUAL ZAJUNA INSTRUCTOR_compressed.pdf"
													className="capsulas__informativas__manuales__link" rel="noopener noreferrer">Ver PDF</a>
												<img src="/img/banners/manual-zajuna-instructor-banner.webp" alt="Manual Zajuna Instructor"
													className="capsulas__informativas__manuales__img" />
											</div>
										</div>
									</div>
								</div>
								
								{/* Manual Aprendiz */}
								<div className={`aprendiz__pdf ${activeManual !== MANUAL_TYPES.APRENDIZ ? 'manuales__container--oculto' : ''}`}>
									<div className="manuales__container">
										<div className="manuales__pdf">
											<div className="manual__ins__cont">
												<a target="_blank"
													href="/pdfs/titulada/manuales/MANUAL ZAJUNA - APRENDIZ_compressed.pdf"
													className="capsulas__informativas__manuales__link" rel="noopener noreferrer">Ver PDF</a>
												<img src="/img/banners/manual-zajuna-aprendiz-banner.webp" alt="Manual Zajuna Aprendiz"
													className="capsulas__informativas__manuales__img" />
											</div>
										</div>
									</div>
								</div>
								
								{/* Manual SAVA */}
								<div className={`sava__pdf ${activeManual !== MANUAL_TYPES.SAVA ? 'manuales__container--oculto' : ''}`}>
									<div className="manuales__container">
										<div className="manuales__pdf">
											<div className="manual__ins__cont">
												<a target="_blank" href="/pdfs/titulada/manuales/MANUAL SAVA.pdf"
													className="capsulas__informativas__manuales__link" rel="noopener noreferrer">Ver PDF</a>
												<img src="/img/banners/Portada-manual-sava.webp" alt="Manual SAVA"
													className="capsulas__informativas__manuales__img" />
											</div>
										</div>
									</div>
								</div>
							</div>
						</div>
					</div>
				</div>

				<div className="soporte__container-2">
					<div className="soporte__cursos">
						<div className="soporte__cursos-container">
							<h5 className="soporte__manual-text">¿Qué es SAVA?</h5>
						</div>
						<br />
						<iframe width="70%" style={{aspectRatio: '20 / 9', borderRadius:'16px', margin:'auto'}}
							src="https://www.youtube.com/embed/vLu8PDIKnag?si=hsjIq5NdjKkgsFpG" title="YouTube video player"
							frameBorder="0"
							allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share"
							allowFullScreen></iframe>
						<br />
						<p className="soporte__cursos-text">Estimado Instructor, el sistema de Soporte Ambiente Virtual
							de Aprendizaje, es un sistema de información desarrollado por el Equipo de Soporte Técnico LMS al
							interior del Grupo de Ejecución de la Formación Virtual SENA, para la atención de
							incidencias/requerimientos presentados en el ambiente virtual.</p>
					</div>
				</div>
			</main>
			<Footer />
		</>
	);
}

export default Soporte;

