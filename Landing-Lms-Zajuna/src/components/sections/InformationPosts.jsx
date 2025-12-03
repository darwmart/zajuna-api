import React from 'react';

// Datos de los posts informativos
const INFORMATION_POSTS = [
	{
		id: 'cursos-virtuales',
		href: 'https://ejecucionformacion.sena.edu.co/cursos-cortos',
		image: '/img/posts/cursos-virtuales-post.webp',
		alt: 'cursos virtuales posts'
	},
	{
		id: 'radio',
		href: 'https://www.sena.edu.co/es-co/Noticias/Paginas/frecuencia-sena.aspx',
		image: '/img/posts/radio-post.webp',
		alt: 'radio posts'
	},
	{
		id: 'soporte',
		href: 'https://sava.sena.edu.co/',
		image: '/img/posts/soporte-post.webp',
		alt: 'soporte posts'
	},
	{
		id: 'instructor',
		href: 'https://caplms.sena.edu.co/',
		image: '/img/posts/instructor-post.webp',
		alt: 'instructor posts'
	},
	{
		id: 'virtual-lms',
		href: '',
		image: '/img/posts/virtual-lms-post.webp',
		alt: 'virtual lms'
	},
	{
		id: 'aprendiz',
		href: 'https://ejecucionformacion.sena.edu.co/comunidad-aprendices',
		image: '/img/posts/aprendiz-post.webp',
		alt: 'aprendiz posts'
	}
];

/**
 * InformationPosts - Sección de posts informativos
 * Muestra un grid de enlaces a recursos y servicios del SENA
 */
function InformationPosts() {
	return (
		<>
			{/* Título de la sección */}
			<div className="text-informativo">
				<div className="text-informativo__container">
					<img 
						loading="lazy" 
						src="/img/icons/informacion-representativo-icon.svg" 
						alt="Icono información" 
						className="text-informativo__icon" 
					/>
					<h3 className="text-informativo__title">Informe Representativo</h3>
				</div>
			</div>
			
			{/* Grid de posts */}
			<div className="post-information">
				{INFORMATION_POSTS.map((post) => (
					<a 
						key={post.id}
						className="post-information__url" 
						target="_blank"
						href={post.href} 
						rel="noopener noreferrer"
						aria-label={post.alt}
					>
						<img 
							loading="lazy"
							src={post.image} 
							alt={post.alt}
							className="post-information__img" 
						/>
					</a>
				))}
			</div>
		</>
	);
}

export default React.memo(InformationPosts);

