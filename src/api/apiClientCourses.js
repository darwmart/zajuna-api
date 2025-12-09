import axios from "axios";

const apiClientCourses = axios.create({
  baseURL: process.env.REACT_APP_COURSES_API || process.env.REACT_APP_API_URL || "http://localhost:8080/api",
  headers: {
    "Content-Type": "application/json",
    "X-API-Key": process.env.REACT_APP_API_KEY || "14c878cb0d468b13d6ecc00bb6d332af",
  },
  withCredentials: true, // IMPORTANTE: Permite enviar cookies cross-site
});

// Interceptor para agregar el token JWT en cada petición
apiClientCourses.interceptors.request.use(
  (config) => {
    // Obtener el token del localStorage (almacenado por la Landing al hacer login)
    const token = localStorage.getItem('zajuna_token');

    if (token) {
      // Agregar el token en el header Authorization
      config.headers.Authorization = `Bearer ${token}`;
    }

    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// Interceptor para logs y manejo de errores
apiClientCourses.interceptors.response.use(
  (response) => response,
  (error) => {
    console.error("API Error:", error.response?.data || error.message);
    return Promise.reject(error);
  }
);

// Función para obtener los cursos del usuario autenticado
// Endpoint: GET /api/courses/my-courses
// Retorna todos los cursos donde el usuario está matriculado (independiente del rol)
export const getMyCourses = async () => {
  try {
    const response = await apiClientCourses.get('/courses/my-courses');
    return response.data;
  } catch (error) {
    console.error('Error al obtener mis cursos:', error);
    throw error;
  }
};

// Función para obtener el contenido de un curso (secciones y módulos)
// Endpoint: GET /api/courses/:id/content
// Retorna todas las secciones con sus actividades/módulos
export const getCourseContent = async (courseId) => {
  try {
    const response = await apiClientCourses.get(`/courses/${courseId}/content`);
    return response.data;
  } catch (error) {
    console.error(`Error al obtener contenido del curso ${courseId}:`, error);
    throw error;
  }
};

export default apiClientCourses;
