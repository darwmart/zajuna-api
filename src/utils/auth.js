/**
 * Utilidades de autenticación para el Dashboard
 */

/**
 * Realiza logout completo
 * - Elimina el token del localStorage
 * - Elimina los datos del usuario
 * - Redirige a la Landing Page
 */
export const logout = () => {
  // Eliminar token y datos de usuario del localStorage
  localStorage.removeItem('zajuna_token');
  localStorage.removeItem('zajuna_user');
  localStorage.removeItem('zajuna_refresh_token');
  localStorage.removeItem('zajuna_token_expires');

  // Redirigir a la Landing Page
  const landingUrl = process.env.REACT_APP_LANDING_URL || 'http://localhost:5173';
  window.location.href = landingUrl;
};

/**
 * Verifica si el usuario está autenticado
 * @returns {boolean}
 */
export const isAuthenticated = () => {
  const token = localStorage.getItem('zajuna_token');
  return !!token;
};

/**
 * Obtiene el token almacenado
 * @returns {string|null}
 */
export const getToken = () => {
  return localStorage.getItem('zajuna_token');
};

/**
 * Obtiene los datos del usuario almacenados
 * @returns {Object|null}
 */
export const getUser = () => {
  const userStr = localStorage.getItem('zajuna_user');
  return userStr ? JSON.parse(userStr) : null;
};
