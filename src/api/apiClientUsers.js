import axios from "axios";

const apiClientUsers = axios.create({
  baseURL: process.env.REACT_APP_USERS_API || "http://localhost:8080/api",
  headers: {
    "Content-Type": "application/json",
    "X-API-Key": process.env.REACT_APP_API_KEY || "14c878cb0d468b13d6ecc00bb6d332af",
  },
  withCredentials: true, // IMPORTANTE: Permite enviar cookies cross-site
});

// Interceptor para agregar el token JWT en cada petición
apiClientUsers.interceptors.request.use(
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

export default apiClientUsers;