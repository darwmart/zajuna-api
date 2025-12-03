import React from 'react';
import '../../styles/layouts/slider.css';
import useSlider from '../../hooks/useSlider';

/**
 * Componente Slider genérico y reutilizable
 * @param {Object} props
 * @param {Array<string>} props.images - Array de rutas de imágenes a mostrar en el slider
 * @param {string} props.altText - Texto alternativo para las imágenes (opcional)
 */
function Slider({ images = [], altText = 'Slider image' }) {
	useSlider();

	// Si no hay imágenes, no renderizar nada
	if (!images || images.length === 0) {
		return null;
	}

	return (
		<div className="slider">
			<div className="slider__container" id="slider__container">
				<div className="slider__items" id="slider__items">
					{images.map((imageSrc, index) => (
						<img 
							key={`slider-img-${index}`}
							src={imageSrc} 
							alt={`${altText} ${index + 1}`} 
							className="slider__img" 
						/>
					))}
				</div>
			</div>
			<div className="slider__indicators" id="slider__indicators"></div>
		</div>
	);
}

export default Slider;
