import React, { useState } from 'react';

/**
 * Componente genérico para mostrar tarjetas de programas académicos
 * con información desplegable (técnicos y tecnólogos)
 * 
 * @param {Object} props
 * @param {string} props.title - Título del programa
 * @param {string} props.sniesCode - Código SNIES (opcional)
 * @param {string} props.pdfUrl - URL del PDF del programa
 * @param {string} props.videoUrl - URL del video de YouTube
 * @param {string} props.offerUrl - URL de la oferta (opcional, comentado por defecto)
 */
function ProgramCard({ title, sniesCode, pdfUrl, videoUrl, offerUrl }) {
	const [isExpanded, setIsExpanded] = useState(false);

	const handleToggle = () => {
		setIsExpanded(!isExpanded);
	};

	return (
		<div className="programas__card">
			<div 
				className={`programas__desplegable ${isExpanded ? 'programas__desplegable--activo' : ''}`}
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
				<p>{title}</p>
				<img 
					loading="lazy" 
					className="programas__desplegable-flecha" 
					src={isExpanded ? '/img/icons/flecha-arriba.svg' : '/img/icons/flecha-icon.svg'}
					alt={isExpanded ? 'flecha arriba icono' : 'flecha abajo icono'}
				/>
			</div>

			<div className={`programas__informacion ${isExpanded ? '' : 'programas__informacion--oculto'}`}>
				{sniesCode && (
					<div className="programas__informacion-codigo">
						<p>Código SNIES: {sniesCode}</p>
					</div>
				)}

				<div className="programas__informacion__ver">
					{/* Link al PDF */}
					{pdfUrl && (
						<a 
							className="programas__informacion__ver-doc" 
							href={pdfUrl} 
							target="_blank" 
							rel="noopener noreferrer"
							aria-label={`Ver PDF de ${title}`}
						>
							<p>Ver PDF</p>
							<img 
								loading="lazy" 
								className="programas__informacion__ver-icon" 
								src="/img/icons/libro-icon.svg" 
								alt="libro icono" 
							/>
						</a>
					)}

					{/* Link al video */}
					{videoUrl && (
						<a 
							className="programas__informacion__ver-doc" 
							href={videoUrl} 
							target="_blank" 
							rel="noopener noreferrer"
							aria-label={`Ver video de ${title}`}
						>
							<p>Ver video</p>
							<img 
								loading="lazy" 
								className="programas__informacion__ver-icon" 
								src="/img/icons/computador-icon.svg" 
								alt="computador icono" 
							/>
						</a>
					)}

					{/* Link a la oferta (comentado por defecto, puede descomentarse si es necesario) */}
					{/* {offerUrl && (
						<a 
							className="programas__informacion__ver-doc" 
							href={offerUrl} 
							target="_blank" 
							rel="noopener noreferrer"
							aria-label={`Ver oferta de ${title}`}
						>
							<p>Ver oferta</p>
							<img 
								loading="lazy" 
								className="programas__informacion__ver-icon" 
								src="/img/icons/mundo.svg" 
								alt="mundo icono" 
							/>
						</a>
					)} */}
				</div>
			</div>
		</div>
	);
}

export default ProgramCard;

