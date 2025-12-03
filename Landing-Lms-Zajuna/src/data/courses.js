/**
 * DATOS DE CURSOS NUEVOS
 * Cursos virtuales disponibles en la plataforma
 */

export const COURSES_DATA = [
  {
    id: 'water-sampling',
    img: '/img/banners/agua.webp',
    title: 'MUESTRAS DE AGUA PARA CONSUMO HUMANO',
    href: 'https://betowa.sena.edu.co/oferta/muestreo-de-agua-para-consumo-humano?enrollCourse=3336960&programId=226753',
    videoSrc: '/video/muestreo-agua.mp4',
  },
  {
    id: 'derby-pattern',
    img: '/img/banners/patronaje.webp',
    title: 'PATRONAJE DE CALZADO TIPO DERBY',
    href: 'https://betowa.sena.edu.co/oferta/patronaje-de-calzado-tipo-derby?enrollCourse=3336949&programId=226388',
    videoSrc: '/video/patronaje-calzado.mp4',
  },
  {
    id: 'financial-analysis',
    img: '/img/banners/financiero.webp',
    title: 'ANÁLISIS FINANCIERO EMPRESARIAL',
    href: 'https://betowa.sena.edu.co/oferta/analisis-financiero-empresarial?enrollCourse=3334530&programId=226673',
    videoSrc: '/video/analisis-financiero.mp4',
  },
  {
    id: 'environmental-measurement',
    img: '/img/banners/ambientales.webp',
    title: 'MEDICIÓN DE VARIABLES AMBIENTALES EN SISTEMAS AGRICOLAS',
    href: 'https://betowa.sena.edu.co/oferta/medicion-de-variables-ambientales-en-sistemas-agricolas?enrollCourse=3334706&programId=226996',
    videoSrc: '/video/variables-ambientales.mp4',
  },
  {
    id: 'ai-data-transformation',
    img: '/img/banners/modelos.webp',
    title: 'TRANSFORMACIÓN DE DATOS EN MODELOS DE INTELIGENCIA ARTIFICIAL',
    href: 'https://betowa.sena.edu.co/oferta/transformacion-de-datos-en-modelos-de-inteligencia-artificial?enrollCourse=3337258&programId=227158',
    videoSrc: '/video/ia-modelos.mp4',
  },
  {
    id: 'ai-data-integration',
    img: '/img/banners/datos.webp',
    title: 'APLICACIÓN DE LA INTELIGENCIA ARTIFICIAL EN LA INTEGRACION DE DATOS',
    href: 'https://betowa.sena.edu.co/oferta/aplicacion-de-la-inteligencia-artificial-en-la-integracion-de-datos?enrollCourse=3337262&programId=226510',
    videoSrc: '/video/ia-integracion-datos.mp4',
  },
];

/**
 * Obtener todos los cursos
 */
export function getAllCourses() {
  return COURSES_DATA;
}

/**
 * Obtener un curso por su ID
 */
export function getCourseById(id) {
  return COURSES_DATA.find(course => course.id === id);
}

/**
 * Obtener cursos con video
 */
export function getCoursesWithVideo() {
  return COURSES_DATA.filter(course => course.videoSrc);
}

