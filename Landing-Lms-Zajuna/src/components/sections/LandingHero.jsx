import React from 'react';
import Menu from '../layouts/Menu';
import Slider from '../layouts/Slider';
import LogosInstitucionales from '../layouts/LogosInstitucionales';
import LoginForm from '../forms/LoginForm';

// Imágenes del slider principal
const SLIDER_IMAGES = [
	'/img/banners/banner7-banner.webp',
	'/img/banners/banner8-banner.webp',
	'/img/banners/banner9-banner.webp',
	'/img/banners/banner10-banner.webp',
	'/img/banners/banner11-banner.webp'
];

// Datos de redes sociales
const SOCIAL_LINKS = [
	{
		id: 'facebook',
		href: 'https://www.facebook.com/SENA/',
		icon: '/img/icons/facebook-icon.svg',
		alt: 'Facebook SENA'
	},
	{
		id: 'instagram',
		href: 'https://www.instagram.com/senacomunica/',
		icon: '/img/icons/instagram-icon.svg',
		alt: 'Instagram SENA'
	},
	{
		id: 'youtube',
		href: 'https://www.youtube.com/channel/UCt5y885UFplu2okY39TBwCg',
		icon: '/img/icons/youtube-icon.svg',
		alt: 'YouTube SENA'
	},
	{
		id: 'twitter',
		href: 'https://twitter.com/SENAComunica',
		icon: '/img/icons/x-icon.svg',
		alt: 'Twitter SENA'
	}
];

/**
 * LandingHero - Sección principal de la landing page
 * Incluye:
 * - Menú de navegación
 * - Slider de banners informativos
 * - Logos institucionales
 * - Formulario de login
 * - Enlaces a redes sociales
 */
function LandingHero() {
	return (
		<div className="landing">
			{/* Menú de navegación */}
			<div className="landing__item--0">
				<Menu />
			</div>
			
			{/* Slider de banners */}
			<div className="landing__item--1">
				<Slider images={SLIDER_IMAGES} altText="Banner informativo" />
			</div>
			
			{/* Logos institucionales */}
			<div className="landing__item--2">
				<LogosInstitucionales />
			</div>
			
			{/* Login y redes sociales */}
			<div className="landing__item--4">
				<LoginForm />
				<div className="landing__2">
					<div className="landing__redes">
						{SOCIAL_LINKS.map((social) => (
							<a 
								key={social.id}
								target="_blank" 
								href={social.href} 
								rel="noopener noreferrer"
								aria-label={social.alt}
							>
								<img
									loading="lazy" 
									src={social.icon} 
									alt={social.alt} 
									className="landing__redes-img" 
								/>
							</a>
						))}
					</div>
				</div>
			</div>
		</div>
	);
}

export default React.memo(LandingHero);

