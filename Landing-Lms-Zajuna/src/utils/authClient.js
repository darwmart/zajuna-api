/**
 * Cliente de autenticación para comunicación con el backend
 * Maneja el login, almacenamiento del token y headers de autenticación
 */

// ============================================================================
// CONSTANTES
// ============================================================================

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080';
const LOGIN_ENDPOINT = '/api/login';
const AUTH_COOKIE_NAME = 'Authorization';

// Keys de localStorage
const STORAGE_KEYS = {
	TOKEN: 'zajuna_token',
	USER: 'zajuna_user'
};

// Mensajes de error
const ERROR_MESSAGES = {
	INVALID_CREDENTIALS: 'Credenciales inválidas',
	ENDPOINT_NOT_FOUND: 'Endpoint de login no encontrado (404). Verifica que el backend esté corriendo en http://localhost:8080',
	METHOD_NOT_ALLOWED: 'Método no permitido. El backend puede no aceptar POST en /api/v1/auth/login',
	SERVER_ERROR: 'Error del servidor. Por favor intenta más tarde',
	INVALID_RESPONSE: 'El backend retornó una respuesta inválida. Verifica el formato de respuesta.',
	DEFAULT: 'Error al iniciar sesión'
};

// ============================================================================
// FUNCIONES HELPER PRIVADAS
// ============================================================================

/**
 * Extrae el token de las cookies del navegador
 * @returns {string|null} Token o null si no se encuentra
 */
function extractTokenFromCookie() {
	const cookies = document.cookie.split(';');
	
	for (let cookie of cookies) {
		const [name, value] = cookie.trim().split('=');
		if (name === AUTH_COOKIE_NAME) {
			return value;
		}
	}
	
	return null;
}

/**
 * Construye un mensaje de error apropiado según el código de estado HTTP
 * @param {Response} response - Objeto Response de fetch
 * @param {Object} errorData - Datos de error del backend (opcional)
 * @returns {string} Mensaje de error descriptivo
 */
function buildErrorMessage(response, errorData = null) {
	if (errorData && (errorData.error || errorData.message)) {
		return errorData.error || errorData.message;
	}
	
	const { status, statusText } = response;
	
	switch (status) {
		case 401:
			return ERROR_MESSAGES.INVALID_CREDENTIALS;
		case 404:
			return ERROR_MESSAGES.ENDPOINT_NOT_FOUND;
		case 405:
			return ERROR_MESSAGES.METHOD_NOT_ALLOWED;
		default:
			if (status >= 500) {
				return ERROR_MESSAGES.SERVER_ERROR;
			}
			return `Error del servidor (${status}): ${statusText}`;
	}
}

/**
 * Procesa una respuesta exitosa del login
 * @param {Response} response - Objeto Response de fetch
 * @returns {Promise<Object>} Objeto con success, token y user
 */
async function processLoginResponse(response) {
    try {
        // Intentar parsear JSON si el backend devuelve un body (caso original)
        let data = null;
        try {
            data = await response.json();
            console.log('Respuesta del backend (JSON):', data);
        } catch (e) {
            // No siempre habrá JSON, puede que el backend solo haya enviado Set-Cookie
            console.log('No se pudo parsear JSON del body (posible cookie-only response).', e);
        }
        
        // Si el backend sigue devolviendo token en el body, mantener comportamiento previo
        if (data && data.token) {
            const token = data.token;
            // Guardar token y refreshToken si están disponibles
            storeToken(token);
            if (data.refreshToken) {
                localStorage.setItem('zajuna_refresh_token', data.refreshToken);
            }
            
            if (data.expiresIn) {
                const expirationTime = Date.now() + (data.expiresIn * 1000);
                localStorage.setItem('zajuna_token_expires', expirationTime.toString());
            }
            
            return {
                success: true,
                token: token,
                refreshToken: data.refreshToken || null,
                user: data.user || null,
                expiresIn: data.expiresIn || null,
                isAdmin: data.isAdmin || false,
                canAccessDashboard: data.canAccessDashboard || false,
                cookieBased: false
            };
        }
        
        // Si no hay token en el body pero response.ok, se asume que el backend envió una cookie HttpOnly
        if (response.ok) {
            console.log('Respuesta OK sin token en body: asumiendo autenticación vía cookie HttpOnly.');
            // Nota: no podemos leer cookies HttpOnly desde JS. Para verificar sesión, el frontend
            // deberá hacer llamadas al backend (con credentials: "include") a endpoints que validen la sesión.
            return {
                success: true,
                token: null,
                refreshToken: null,
                user: data && data.user ? data.user : null,
                expiresIn: null,
                isAdmin: data && data.isAdmin ? data.isAdmin : false,
                canAccessDashboard: data && data.canAccessDashboard ? data.canAccessDashboard : false,
                cookieBased: true
            };
        }
        
        // Si llegamos aquí algo falló al parsear y el status no es OK
        console.warn('No se obtuvo token ni cookie; respuesta inválida.');
        throw new Error(ERROR_MESSAGES.INVALID_RESPONSE);
    } catch (parseError) {
        console.error('Error al procesar respuesta de login:', parseError);
        
        // Si el status es 200 pero no hay JSON ni cookie accesible, informar
        if (response.status === 200) {
            throw new Error(ERROR_MESSAGES.INVALID_RESPONSE);
        }
        
        throw parseError;
    }
}

/**
 * Maneja errores de red y genera un mensaje descriptivo
 * @param {Error} error - Error capturado
 * @returns {Error} Error con mensaje mejorado
 */
function handleNetworkError(error) {
	// Detectar errores de red (Failed to fetch)
	if (error instanceof TypeError && (error.message === 'Failed to fetch' || error.message.includes('fetch'))) {
		const frontendOrigin = window.location.origin;
		const errorMessage = `No se pudo conectar con el backend en ${API_BASE_URL}. 
		
Posibles causas:
1. El backend no está corriendo (verifica con: curl ${API_BASE_URL}/api/v1)
2. El endpoint ${LOGIN_ENDPOINT} no está implementado en el backend
3. Problema de CORS (el backend necesita permitir el origen del frontend: ${frontendOrigin})
4. Error de conexión de red
5. El backend puede estar corriendo en otro puerto

Solución: 
- Verifica que el backend esté corriendo: curl ${API_BASE_URL}/api/v1
- Verifica CORS en el backend permitiendo el origen: ${frontendOrigin}
- Revisa la consola del navegador para más detalles del error`;
		return new Error(errorMessage);
	}
	
	// Re-lanzar otros errores sin modificar
	return error;
}

// ============================================================================
// FUNCIÓN PRINCIPAL DE LOGIN
// ============================================================================

/**
 * Realiza una petición de login al backend
 * @param {Object} credentials - Credenciales de login
 * @param {string} credentials.username - Usuario (opcional, para login administrativo)
 * @param {string} credentials.idnumber - Número de documento (opcional, para login de cursos)
 * @param {string} credentials.password - Contraseña
 * @returns {Promise<Object>} Objeto con token y datos del usuario
 * @throws {Error} Si las credenciales son inválidas o hay un error en el servidor
 */
export async function loginRequest({ username, idnumber, password }) {
	// Preparar credenciales para el backend
	// El backend espera: { "username": "...", "password": "..." }
	// Para login de cursos, usamos idnumber como username
	// Para login administrativo, usamos username directamente
	const body = {
		username: username || idnumber, // El backend busca por username (puede ser número de documento)
		password: password
	};

	const url = `${API_BASE_URL}${LOGIN_ENDPOINT}`;
	
	// Log para debug
	console.log('Intentando login en:', url);
	console.log('Body enviado:', { username: body.username, password: '***' });
	
	try {
		// Realizar petición al backend
		console.log('🔵 Iniciando petición de login...');
		console.log('📍 URL:', url);
		console.log('📦 Body:', { username: body.username, password: '***' });
		
		const response = await fetch(url, {
            method: 'POST',
            headers: { 
                'Content-Type': 'application/json',
                'Accept': 'application/json'
            },
            body: JSON.stringify(body),
            // IMPORTANTE: para que el navegador acepte la cookie HttpOnly que envía el backend
            // hay que permitir credenciales en la petición.
            // Además el backend debe responder con Access-Control-Allow-Credentials: true
            // y no usar Access-Control-Allow-Origin: * (debe ser el origen exacto).
            credentials: 'include'
            // No usamos credentials: 'include' antes porque:
            // 1. No estábamos usando cookies HttpOnly (usábamos localStorage)
            // 2. El backend usaba Access-Control-Allow-Origin: * que no es compatible con credentials
            // 3. Los tokens se manejaban mediante headers Authorization
        });
        
		console.log('📥 Respuesta recibida:', {
			status: response.status,
			statusText: response.statusText,
			ok: response.ok,
			headers: Object.fromEntries(response.headers.entries())
		});

		// Manejar respuesta con error
		if (!response.ok) {
			let errorData = null;
			try {
				const text = await response.text();
				console.log('📄 Respuesta de error (texto):', text);
				errorData = JSON.parse(text);
				console.log('📄 Respuesta de error (JSON):', errorData);
			} catch (parseError) {
				console.warn('⚠️ No se pudo parsear respuesta de error como JSON:', parseError);
			}
			
			const errorMessage = buildErrorMessage(response, errorData);
			console.error('❌ Error en login:', errorMessage);
			throw new Error(errorMessage);
		}

		// Procesar respuesta exitosa
		console.log('✅ Login exitoso, procesando respuesta...');
		return await processLoginResponse(response);
	} catch (error) {
		// Log detallado del error para debug
		console.error('❌ Error completo en loginRequest:', {
			message: error.message,
			name: error.name,
			stack: error.stack,
			error: error
		});
		
		// Manejar errores de red
		const handledError = handleNetworkError(error);
		console.error('❌ Error manejado:', handledError.message);
		throw handledError;
	}
}

// ============================================================================
// FUNCIONES DE GESTIÓN DE TOKEN
// ============================================================================

/**
 * Almacena el token JWT en localStorage
 * @param {string} token - Token JWT recibido del backend
 */
export function storeToken(token) {
	if (token) {
		localStorage.setItem(STORAGE_KEYS.TOKEN, token);
	}
}

/**
 * Obtiene el token JWT almacenado
 * @returns {string|null} Token JWT o null si no existe
 */
export function getToken() {
	return localStorage.getItem(STORAGE_KEYS.TOKEN);
}

/**
 * Elimina el token JWT del localStorage (logout)
 */
export function removeToken() {
	localStorage.removeItem(STORAGE_KEYS.TOKEN);
}

/**
 * Verifica si hay un token almacenado (usuario autenticado)
 * @returns {boolean} true si hay token, false en caso contrario
 */
export function isAuthenticated() {
	return !!getToken();
}

/**
 * Obtiene los headers de autenticación para incluir en peticiones
 * @returns {Object} Objeto con header Authorization si existe token
 * Nota: Si el token está en una cookie HttpOnly, el navegador lo enviará automáticamente
 */
export function getAuthHeaders() {
	const token = getToken();
	if (token) {
		return { Authorization: `Bearer ${token}` };
	}
	// Si no hay token en localStorage, verificar si hay cookie
	// Nota: Las cookies HttpOnly no son accesibles desde JS, pero se enviarán automáticamente
	return {};
}

// ============================================================================
// FUNCIONES DE GESTIÓN DE USUARIO
// ============================================================================

/**
 * Almacena los datos del usuario en localStorage
 * @param {Object} user - Datos del usuario
 */
export function storeUser(user) {
	if (user) {
		localStorage.setItem(STORAGE_KEYS.USER, JSON.stringify(user));
	}
}

/**
 * Obtiene los datos del usuario almacenados
 * @returns {Object|null} Datos del usuario o null si no existen
 */
export function getUser() {
	const userStr = localStorage.getItem(STORAGE_KEYS.USER);
	return userStr ? JSON.parse(userStr) : null;
}

/**
 * Elimina los datos del usuario del localStorage
 */
export function removeUser() {
	localStorage.removeItem(STORAGE_KEYS.USER);
}

// ============================================================================
// FUNCIÓN DE LOGOUT
// ============================================================================

/**
 * Realiza logout completo
 * 1. Llama al backend para invalidar la sesión
 * 2. Elimina token y usuario de localStorage
 * 3. Redirige al login
 */
export async function logout() {
	try {
		// Llamar al backend para invalidar la sesión
		const token = getToken();
		if (token) {
			await fetch(`${API_BASE_URL}/api/logout`, {
				method: 'POST',
				headers: {
					'Authorization': `Bearer ${token}`,
					'Content-Type': 'application/json'
				},
				credentials: 'include'
			});
		}
	} catch (error) {
		console.error('Error al hacer logout en el backend:', error);
		// Continuar con el logout local aunque falle el backend
	} finally {
		// Siempre limpiar localStorage
		removeToken();
		removeUser();
		localStorage.removeItem('zajuna_refresh_token');
		localStorage.removeItem('zajuna_token_expires');
	}
}

