/**
 * Datos de los niveles del programa English Does Work del SENA
 * 13 niveles desde A1 hasta B2+ según el Marco Común Europeo (MCERL)
 */

export const ENGLISH_LEVELS_DATA = [
	{
		id: 1,
		level: 'Level 1',
		image: '/img/banners/ingles1-banner.webp',
		alt: 'ingles 1 banner',
		description: 'Afianzamiento de herramientas básicas para la comunicación en inglés.',
		href: 'https://oferta.senasofiaplus.edu.co/sofia-oferta/detalle-oferta.html?fm=0&fc=nUItZqt0tTk'
	},
	{
		id: 2,
		level: 'Level 2',
		image: '/img/banners/ingles2-banner.webp',
		alt: 'ingles 2 banner',
		description: 'Comunicación en contextos personales y laborales en inglés',
		href: 'https://oferta.senasofiaplus.edu.co/sofia-oferta/detalle-oferta.html?fm=0&fc=Fh5muxh85T4'
	},
	{
		id: 3,
		level: 'Level 3',
		image: '/img/banners/ingles3-banner.webp',
		alt: 'ingles 3 banner',
		description: 'Comunicación en contextos personales y laborales en inglés.',
		href: 'https://oferta.senasofiaplus.edu.co/sofia-oferta/detalle-oferta.html?fm=0&fc=3si-fIPR_0s'
	},
	{
		id: 4,
		level: 'Level 4',
		image: '/img/banners/ingles4-banner.webp',
		alt: 'ingles 4 banner',
		description: 'Consolidación y comprensión de diferentes textos orales y escritos en inglés.',
		href: 'https://oferta.senasofiaplus.edu.co/sofia-oferta/detalle-oferta.html?fm=0&fc=4y3yLXIHxjg'
	},
	{
		id: 5,
		level: 'Level 5',
		image: '/img/banners/ingles5-banner.webp',
		alt: 'ingles 5 banner',
		description: 'Interacción en diferentes contextos expresando gustos y preferencias en inglés.',
		href: 'https://oferta.senasofiaplus.edu.co/sofia-oferta/detalle-oferta.html?fm=0&fc=tO1MQ4wo--A'
	},
	{
		id: 6,
		level: 'Level 6',
		image: '/img/banners/ingles6-banner.webp',
		alt: 'ingles 6 banner',
		description: 'Afianzamiento de herramientas para la comunicación en inglés.',
		href: 'https://oferta.senasofiaplus.edu.co/sofia-oferta/detalle-oferta.html?fm=0&fc=v6A_n61Hv3E'
	},
	{
		id: 7,
		level: 'Level 7',
		image: '/img/banners/ingles7-banner.webp',
		alt: 'ingles 7 banner',
		description: 'Consolidación de herramientas para la comunicación efectiva en diferentes contextos.',
		href: 'https://oferta.senasofiaplus.edu.co/sofia-oferta/detalle-oferta.html?fm=0&fc=ejdeKOvpQl0'
	},
	{
		id: 8,
		level: 'Level 8',
		image: '/img/banners/ingles8-banner.webp',
		alt: 'ingles 8 banner',
		description: 'Construir textos orales y escritos de acuerdo con las características e intencionalidad del contexto.',
		href: 'https://oferta.senasofiaplus.edu.co/soerta-oferta/detalle-oferta.html?fm=0&fc=TemxAxoUAKM'
	},
	{
		id: 9,
		level: 'Level 9',
		image: '/img/banners/ingles9-banner.webp',
		alt: 'ingles 9 banner',
		description: 'Opinar de hechos ocurridos o planeados en inglés con base en textos narrativos.',
		href: 'https://oferta.senasofiaplus.edu.co/sofia-oferta/detalle-oferta.html?fm=0&fc=B_VEzMlZ9ZM'
	},
	{
		id: 10,
		level: 'Level 10',
		image: '/img/banners/ingles10-banner.webp',
		alt: 'ingles 10 banner',
		description: 'Construir textos orales y escritos en lengua inglesa acerca de sucesos futuros.',
		href: 'https://oferta.senasofiaplus.edu.co/sofia-oferta/detalle-oferta.html?fm=0&fc=OyMqRqMdNfY'
	},
	{
		id: 11,
		level: 'Level 11',
		image: '/img/banners/ingles11-banner.webp',
		alt: 'ingles 11 banner',
		description: 'Elaborar textos argumentativos en inglés con coherencia y cohesión según la intencionalidad comunicativa.',
		href: 'https://oferta.senasofiaplus.edu.co/sofia-oferta/detalle-oferta.html?fm=0&fc=kjW_ASKUywY'
	},
	{
		id: 12,
		level: 'Level 12',
		image: '/img/banners/ingles12-banner.webp',
		alt: 'ingles 12 banner',
		description: 'Justificar opiniones orales y escritas según el contexto social o laboral en inglés.',
		href: 'https://oferta.senasofiaplus.edu.co/sofia-oferta/detalle-oferta.html?fm=0&fc=TK2PpqUqL5w'
	},
	{
		id: 13,
		level: 'Level 13',
		image: '/img/banners/ingles13-banner.webp',
		alt: 'ingles 13 banner',
		description: 'Interactuar en actos comunicativos con independencia y fluidez a partir de contextos sociales actuales.',
		href: 'https://oferta.senasofiaplus.edu.co/sofia-oferta/detalle-oferta.html?fm=0&fc=4Y1CE0G5Rdk'
	}
];

/**
 * Obtener todos los niveles de inglés
 * @returns {Array} Array de objetos con datos de niveles
 */
export function getEnglishLevels() {
	return ENGLISH_LEVELS_DATA;
}

/**
 * Obtener un nivel específico por ID
 * @param {number} levelId - ID del nivel (1-13)
 * @returns {Object|null} Objeto con datos del nivel o null si no existe
 */
export function getEnglishLevelById(levelId) {
	return ENGLISH_LEVELS_DATA.find(level => level.id === levelId) || null;
}

