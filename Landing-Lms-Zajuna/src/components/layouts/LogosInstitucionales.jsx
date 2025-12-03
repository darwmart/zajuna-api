import React, { useState } from 'react';
import '../../styles/layouts/logos-institucionales.css';

const LOGOS_DATA = [
	{
		superior: {
			src: '/img/logos/fondo-emplender-logo.svg',
			alt: 'fondo emprender log',
			href: 'https://www.sena.edu.co/es-co/trabajo/Paginas/fondo-emprender.aspx'
		},
		inferior: {
			src: '/img/logos/fondo-emplender-logo1.svg',
			alt: 'fondo emprender log',
			href: 'https://www.sena.edu.co/es-co/trabajo/Paginas/fondo-emprender.aspx'
		}
	},
	{
		superior: {
			src: '/img/logos/campesena-logo.svg',
			alt: 'campesena logo',
			href: 'https://sena.edu.co/es-co/campesena/Paginas/index.aspx'
		},
		inferior: {
			src: '/img/logos/campesena-logo1.svg',
			alt: 'campesena logo',
			href: 'https://sena.edu.co/es-co/campesena/Paginas/index.aspx'
		}
	},
	{
		superior: {
			src: '/img/logos/tecnoparque-logo.svg',
			alt: 'tecno parque',
			href: 'https://redtecnoparque.com/'
		},
		inferior: {
			src: '/img/logos/tecnoparque-logo1.svg',
			alt: 'tecno parque',
			href: 'https://redtecnoparque.com/'
		}
	},
	{
		superior: {
			src: '/img/logos/agencia-publica-logo.svg',
			alt: 'agencia publica de empleo logo',
			href: 'https://www.serviciodeempleo.gov.co/'
		},
		inferior: {
			src: '/img/logos/agencia-publica-logo1.svg',
			alt: 'agencia publica de empleo logo',
			href: 'https://www.serviciodeempleo.gov.co/'
		}
	}
];

function LogosInstitucionales() {
	const [hoveredIndex, setHoveredIndex] = useState(null);

	return (
		<div className="logos-institucionales">
			{LOGOS_DATA.map((item, index) => (
				<div key={index} className="logos-institucionales__item">
					<a
						target="_blank"
						href={item.superior.href}
						rel="noopener noreferrer"
						className={`logos-institucionales__icono ${hoveredIndex === index ? 'logos-institucionales__icono--hovered' : ''}`}
						onMouseEnter={() => setHoveredIndex(index)}
						onMouseLeave={() => setHoveredIndex(null)}
					>
						<img
							loading="lazy"
							src={item.superior.src}
							alt={item.superior.alt}
							className="logos-institucionales__icono-img"
						/>
					</a>
					<a
						target="_blank"
						href={item.inferior.href}
						rel="noopener noreferrer"
						className={`logos-institucionales__logo ${hoveredIndex === index ? 'logos-institucionales__logo--hovered' : ''}`}
						onMouseEnter={() => setHoveredIndex(index)}
						onMouseLeave={() => setHoveredIndex(null)}
					>
						<img
							loading="lazy"
							src={item.inferior.src}
							alt={item.inferior.alt}
							className="logos-institucionales__logo-img"
						/>
					</a>
				</div>
			))}
		</div>
	);
}

export default LogosInstitucionales;

