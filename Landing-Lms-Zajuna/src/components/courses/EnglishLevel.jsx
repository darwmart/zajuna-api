import React from 'react';

/**
 * EnglishLevel - Componente para mostrar un nivel de inglés del programa English Does Work
 * 
 * @param {Object} props
 * @param {string} props.image - Ruta de la imagen del nivel
 * @param {string} props.alt - Texto alternativo para la imagen
 * @param {string} props.description - Descripción del nivel
 * @param {string} props.href - URL para inscribirse al nivel
 */
function EnglishLevel({ image, alt, description, href }) {
	return (
		<div className="bilinguismo__ingles-niveles">
			<a 
				target="_blank" 
				href={href} 
				rel="noopener noreferrer"
				aria-label={`Inscribirse en ${alt}`}
			>
				<img 
					loading="lazy" 
					src={image} 
					alt={alt} 
					className="bilinguismo__ingles-imgs" 
				/>
			</a>
			<p className="bilinguismo__ingles-text">
				{description}
			</p>
		</div>
	);
}

// Memoizar para evitar re-renders innecesarios
export default React.memo(EnglishLevel);

