/**
 * Datos de programas de formación Campesena Radial 2024
 */

export const PROGRAMAS_CAMPESENA_2024 = [
	{
		id: 'camp-01',
		title: 'Desarrollo Comunitario',
		icon: '/img/icons/Desarrollo-comunitario.webp',
		courses: [
			'Atención integral en salud al recién nacido',
			'Manejo integrado de la desnutrición aguda en menor de 5 años',
			'Generación de ideas innovadoras con Design Thinking'
		]
	},
	{
		id: 'camp-02',
		title: 'Técnicas de Cultivos',
		icon: '/img/icons/Tecnicas-cultivos.webp',
		courses: [
			'Manejo de sustratos y fertilización en agricultura urbana',
			'Prácticas y aplicaciones de agricultura ecológica',
			'Aplicación de investigación de recursos naturales en comunidades étnicas',
			'Suelos en la agricultura',
			'Manejo de cosecha y poscosecha de frutas y hortalizas'
		]
	},
	{
		id: 'camp-03',
		title: 'Sistema Agroecológico',
		icon: '/img/icons/sistema-agroeco.webp',
		courses: [
			'Implementación de procesos agroecológicos para la transición de sistemas alimentarios',
			'Agroecología y desarrollo rural',
			'Agricultura ecológica: fertilización, suelos y cultivos'
		]
	},
	{
		id: 'camp-04',
		title: 'Ambiental',
		icon: '/img/icons/Ambiental.webp',
		courses: [
			'Planificación de acciones socio ambientales en comunidades étnicas',
			'Medición de variables ambientales en agroecosistemas',
			'Medición de la huella hídrica'
		]
	},
	{
		id: 'camp-05',
		title: 'Buenas Prácticas',
		icon: '/img/icons/Buenas-practicas.webp',
		courses: [
			'Buenas prácticas agrícolas',
			'Buenas prácticas agrícolas para mora y naranja',
			'Implementación de buenas prácticas ganaderas en la producción porcina'
		]
	},
	{
		id: 'camp-06',
		title: 'Desarrollo Sostenible',
		icon: '/img/icons/Desarrollo-sostenibles.webp',
		courses: [
			'Manejo de los residuos sólidos en la producción avícola',
			'Aplicación de conceptos de economía circular en contextos productivos',
			'Planeación de prácticas sostenibles para la finca cafetera',
			'Ganaderia sustentable.',
			'Emprendimiento de unidades productivas'
		]
	},
	{
		id: 'camp-07',
		title: 'Técnicas procesamiento',
		icon: '/img/icons/tecnicas.webp',
		courses: [
			'Procesamiento de frutas y verduras',
			'Aceites esenciales: extracción, usos y aplicaciones',
			'Protección y conservación de alimentos',
			'Agroindustria del plátano'
		]
	}
];

/**
 * Obtener todos los programas de Campesena 2024
 */
export function getProgramasCampesena() {
	return PROGRAMAS_CAMPESENA_2024;
}

