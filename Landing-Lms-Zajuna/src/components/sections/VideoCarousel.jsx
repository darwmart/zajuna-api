import React, { useState, useRef, useEffect, useMemo, useCallback } from 'react';
import CourseCardContent from '../courses/CourseCardContent';
import { getAllCourses } from '../../data/courses';

// Constantes de configuración de videos y tiempos
const VIDEO_TRANSITION_DELAY = 100;
const FADE_OUT_DURATION = 250;
const FADE_IN_DURATION = 50;
const VISIBLE_POSITIONS_COUNT = 3;

/**
 * VideoCarousel - Componente complejo que maneja el carrusel de videos de cursos
 * Incluye:
 * - Video principal de banner
 * - Carrusel de 3 videos que rotan automáticamente
 * - Lógica de hover, play/pause y transiciones
 */
function VideoCarousel() {
	// Estados para control de videos
	const [hoveredCard, setHoveredCard] = useState(null);
	const [activeOverlayIndex, setActiveOverlayIndex] = useState(null);
	const [hoverOverlayIndex, setHoverOverlayIndex] = useState(null);
	const [positionIndices, setPositionIndices] = useState([0, 1, 2]);
	const [playingPosition, setPlayingPosition] = useState(0);
	const [isTransitioning, setIsTransitioning] = useState(false);
	
	// Referencias
	const videoRef = useRef(null);
	const videoRefs = useRef([]);
	const hoveredCardRef = useRef(null);

	const courses = getAllCourses();
	
	// Crear array de videos asociados con sus cursos (memoizado)
	const videosWithCourses = useMemo(() => {
		return courses.map((course) => ({
			videoSrc: course.videoSrc || '/video/demo-cursos-virtuales.mp4', // Fallback a video demo
			course: course
		}));
	}, [courses]);

	// Función para obtener el siguiente índice en el ciclo de una posición (memoizada)
	const getNextIndexForPosition = useCallback((currentIndex, position) => {
		const baseIndex = position;
		const totalCards = videosWithCourses.length;
		if (currentIndex === baseIndex) {
			return baseIndex + VISIBLE_POSITIONS_COUNT < totalCards ? baseIndex + VISIBLE_POSITIONS_COUNT : baseIndex;
		} else {
			return baseIndex;
		}
	}, [videosWithCourses.length]);

	// Función para cambiar la tarjeta en una posición específica (memoizada)
	const changeCardAtPosition = useCallback((position) => {
		setPositionIndices(prev => {
			const newIndices = [...prev];
			const currentIndex = newIndices[position];
			const nextIndex = getNextIndexForPosition(currentIndex, position);
			newIndices[position] = nextIndex;
			
			// Si esta es la posición que está reproduciendo y no hay hover activo
			if (position === playingPosition && hoveredCardRef.current === null) {
				// Pausar todos los videos
				videoRefs.current.forEach((v) => {
					if (v) {
						v.pause();
					}
				});
				
				setTimeout(() => {
					const video = videoRefs.current[nextIndex];
					if (video) {
						video.currentTime = 0;
						video.play().catch(error => {
							console.log('Error al reproducir el video:', error);
						});
						setActiveOverlayIndex(nextIndex);
					}
				}, VIDEO_TRANSITION_DELAY);
			}
			
			return newIndices;
		});
	}, [getNextIndexForPosition, playingPosition]);

	// Efecto para el video principal (prueba7)
	useEffect(() => {
		const video = videoRef.current;
		if (video) {
			video.loop = false;
			video.play().catch(error => {
				console.log('Error al reproducir el video:', error);
			});
		}
	}, []);

	// Efecto para iniciar el primer video al montar
	useEffect(() => {
		const timer = setTimeout(() => {
			const firstIndex = positionIndices[0];
			const firstVideo = videoRefs.current[firstIndex];
			if (firstVideo) {
				firstVideo.play().catch(error => {
					console.log('Error al reproducir el primer video:', error);
				});
				setActiveOverlayIndex(firstIndex);
				setPlayingPosition(0);
			}
		}, VIDEO_TRANSITION_DELAY);
		
		return () => clearTimeout(timer);
	// eslint-disable-next-line react-hooks/exhaustive-deps
	}, []);

	// Efecto para manejar la reproducción de videos según la posición activa
	useEffect(() => {
		// Si hay un hover activo o está en transición, no hacer nada
		if (hoveredCardRef.current !== null || isTransitioning) {
			return;
		}

		// Pausar todos los videos
		videoRefs.current.forEach((video) => {
			if (video) {
				video.pause();
				video.currentTime = 0;
			}
		});

		// Reproducir el video de la posición activa
		const activeIndex = positionIndices[playingPosition];
		const activeVideo = videoRefs.current[activeIndex];
		if (activeVideo) {
			activeVideo.play().catch(error => {
				// Silenciar errores de reproducción que son normales cuando el video se pausa antes de empezar
				// Solo loggear si no es un AbortError (que es esperado cuando se interrumpe la reproducción)
				if (error.name !== 'AbortError') {
					console.warn('Error al reproducir el video:', error);
				}
			});
			setActiveOverlayIndex(activeIndex);
		}
	}, [playingPosition, positionIndices.join(','), isTransitioning]);

	// Handler cuando un video termina (memoizado)
	const handleVideoEnd = useCallback((index) => {
		// Verificar si el video que terminó es el que está reproduciendo en la posición activa
		const activeIndex = positionIndices[playingPosition];
		if (index === activeIndex) {
			// Activar transición (fade out)
			setIsTransitioning(true);
			
			// Esperar a que termine el fade out antes de cambiar la tarjeta
			setTimeout(() => {
				// Cambiar a la siguiente tarjeta en el ciclo de esta posición (sin activar el video)
				setPositionIndices(prev => {
					const newIndices = [...prev];
					const currentIndex = newIndices[playingPosition];
					const nextIndex = getNextIndexForPosition(currentIndex, playingPosition);
					newIndices[playingPosition] = nextIndex;
					return newIndices;
				});
				
				// Fade in y activar la siguiente posición
				setTimeout(() => {
					setIsTransitioning(false);
					// Avanzar a la siguiente posición (0 -> 1 -> 2 -> 0)
					setPlayingPosition((prev) => (prev + 1) % VISIBLE_POSITIONS_COUNT);
				}, FADE_IN_DURATION);
			}, FADE_OUT_DURATION);
		}
	}, [positionIndices, playingPosition, getNextIndexForPosition]);

	// Handler cuando se hace hover sobre un video (memoizado)
	const handleVideoHover = useCallback((index) => {
		setHoveredCard(index);
		hoveredCardRef.current = index;
		
		// Pausar todos los videos
		videoRefs.current.forEach((v) => {
			if (v) {
				v.pause();
			}
		});
		
		// Reproducir el video sobre el que se hace hover
		const video = videoRefs.current[index];
		if (video) {
			video.play().catch(error => {
				console.log('Error al reproducir el video:', error);
			});
		}
		setActiveOverlayIndex(index);
		setHoverOverlayIndex(index);
	}, []);

	// Handler cuando se quita el hover de un video (memoizado)
	const handleVideoLeave = useCallback(() => {
		setHoveredCard(null);
		hoveredCardRef.current = null;
		setHoverOverlayIndex(null);
		
		// Pausar todos los videos
		videoRefs.current.forEach((v) => {
			if (v) {
				v.pause();
			}
		});
		
		// Reproducir el video de la posición activa
		const activeIndex = positionIndices[playingPosition];
		const activeVideo = videoRefs.current[activeIndex];
		if (activeVideo) {
			activeVideo.play().catch(error => {
				console.log('Error al reanudar el video:', error);
			});
			setActiveOverlayIndex(activeIndex);
		}
	}, [positionIndices, playingPosition]);

	return (
		<div className="cursos-virtuales">
			<div className="cursos-virtuales__header">
				<div className="cursos-virtuales__header-container">
					<img loading="lazy" src="/img/icons/REDES-1.svg" alt="icono cursos virtuales" className="cursos-virtuales__header-icon" />
					<h3>Nuevos cursos cortos 100% virtuales</h3>
				</div>
			</div>
			<div className="cursos-virtuales__content-wrapper">
				<div className="cursos-virtuales__container">
					{/* Video principal del banner */}
					<div className="cursos-virtuales__graphic-container">
						<video 
							ref={videoRef}
							autoPlay 
							muted 
							playsInline
							className="cursos-virtuales__banner-img"
						>
							<source src="/video/demo-cursos-virtuales.mp4" type="video/mp4" />
						</video>
					</div>
					
					{/* Carrusel de videos de cursos */}
					<div className="cursos-virtuales__carousel-container">
						<div className="cursos-virtuales__carousel-wrapper">
							<div className="cursos-virtuales__carousel-track">
								{positionIndices.map((cardIndex, position) => {
									const item = videosWithCourses[cardIndex];
									if (!item) return null;
									
									const isTransitioningThis = isTransitioning && position === playingPosition;
									
									return (
										<div 
											key={`${item.course.id}-${position}-${cardIndex}`}
											className={`cursos-virtuales__video-container cursos-virtuales__card ${hoveredCard === cardIndex ? 'cursos-virtuales__card--hovered' : ''} ${isTransitioningThis ? 'cursos-virtuales__card--transitioning' : ''}`}
											onMouseEnter={() => handleVideoHover(cardIndex)}
											onMouseLeave={handleVideoLeave}
										>
											<button
												className="cursos-virtuales__floating-btn"
												onClick={() => changeCardAtPosition(position)}
												aria-label={`Cambiar tarjeta ${position + 1}`}
											>
												<svg width="18" height="18" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
													<rect x="3" y="5" width="16" height="18" rx="1.5" stroke="currentColor" strokeWidth="1.5" fill="none"/>
													<line x1="6" y1="9" x2="14" y2="9" stroke="currentColor" strokeWidth="1" strokeLinecap="round"/>
													<line x1="6" y1="12" x2="12" y2="12" stroke="currentColor" strokeWidth="1" strokeLinecap="round"/>
													<line x1="6" y1="15" x2="10" y2="15" stroke="currentColor" strokeWidth="1" strokeLinecap="round"/>
													<rect x="5" y="3" width="16" height="18" rx="1.5" stroke="currentColor" strokeWidth="1.5" fill="none" opacity="0.6"/>
													<rect x="7" y="1" width="16" height="18" rx="1.5" stroke="currentColor" strokeWidth="1.5" fill="none" opacity="0.3"/>
												</svg>
											</button>
											<video 
												ref={(el) => (videoRefs.current[cardIndex] = el)}
												className="cursos-virtuales__carousel-video cursos-virtuales__video"
												muted 
												playsInline
												onEnded={() => handleVideoEnd(cardIndex)}
											>
												<source src={item.videoSrc} type="video/mp4" />
											</video>
											<CourseCardContent 
												course={item.course} 
												isHovered={
													activeOverlayIndex === cardIndex ||
													hoverOverlayIndex === cardIndex
												} 
											/>
										</div>
									);
								})}
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
	);
}

export default React.memo(VideoCarousel);

