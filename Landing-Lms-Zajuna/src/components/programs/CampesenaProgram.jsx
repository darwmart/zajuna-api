import React, { useState } from 'react';

/**
 * Componente genérico para mostrar programas de Campesena con información desplegable
 * 
 * @param {Object} props
 * @param {string} props.title - Título del programa
 * @param {string} props.icon - Ruta del icono del programa
 * @param {Array<string>} props.courses - Lista de cursos del programa
 */
function CampesenaProgram({ title, icon, courses }) {
	const [isExpanded, setIsExpanded] = useState(false);

	const handleToggle = () => {
		setIsExpanded(!isExpanded);
	};

	return (
		<div className="campesena__program-container">
			<div className="campesena__program-post">
				<img 
					className="campesena__program-img" 
					src={icon} 
					alt={`${title} icono`} 
				/>
				<div 
					className="campesena__program-post2"
					onClick={handleToggle}
					role="button"
					tabIndex={0}
					onKeyPress={(e) => {
						if (e.key === 'Enter' || e.key === ' ') {
							e.preventDefault();
							handleToggle();
						}
					}}
					aria-expanded={isExpanded}
				>
					<p className="campesena__program-texto2">{title}</p>
					<img 
						className="campesena__program-icon" 
						src="/img/icons/flechita.webp" 
						alt="flecha indicadora" 
					/>
				</div>
			</div>
			<div className={`campesena__program-info ${isExpanded ? 'active' : ''}`}>
				<div className="campesena__program-info-link">
					{courses.map((course, index) => (
						<a 
							key={index}
							className="campesena__program-info-text" 
							href=""
							onClick={(e) => e.preventDefault()}
						>
							{course}
						</a>
					))}
				</div>
			</div>
		</div>
	);
}

export default CampesenaProgram;

