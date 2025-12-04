import React, { useState } from 'react';
import '../../styles/index.css';
import { loginRequest, storeToken, storeUser } from '../../utils/authClient';

// Constantes para mensajes de error
const ERROR_MESSAGES = {
	EMPTY_FIELDS: 'Por favor completa todos los campos',
	NO_TOKEN: 'Login exitoso pero no se pudo obtener el token',
	FETCH_ERROR: 'No se pudo conectar con el servidor. Verifica que el backend esté corriendo en el puerto 8080.',
	ENDPOINT_404: 'El endpoint de autenticación no está disponible. Verifica que el backend tenga /api/login implementado',
	INVALID_CREDENTIALS: 'Credenciales inválidas. Por favor verifica tu número de documento y contraseña.',
	CORS_ERROR: 'Error de CORS. Verifica la configuración del backend.',
	DEFAULT: 'Error al conectar con el servidor. Verifica tu conexión a internet y que el backend esté corriendo.'
};

const REDIRECT_DELAY = 500;

function LoginForm() {
	const [activeForm, setActiveForm] = useState('cursos');
	const [loading, setLoading] = useState(false);
	const [error, setError] = useState(null);
	const [success, setSuccess] = useState(false);

	const handleFormToggle = (formType) => {
		setActiveForm(formType);
		setError(null); // Limpiar error al cambiar de formulario
	};

	/**
	 * Función genérica para manejar errores de login
	 * @param {Error} err - Error capturado
	 * @returns {string} Mensaje de error descriptivo
	 */
	const getErrorMessage = (err) => {
		const errorMsg = err.message;
		
		if (errorMsg.includes('Failed to fetch') || errorMsg.includes('No se pudo conectar')) {
			return ERROR_MESSAGES.FETCH_ERROR;
		} else if (errorMsg.includes('404') || errorMsg.includes('no encontrado') || errorMsg.includes('Endpoint de login')) {
			return ERROR_MESSAGES.ENDPOINT_404;
		} else if (errorMsg.includes('401') || errorMsg.includes('Credenciales')) {
			return ERROR_MESSAGES.INVALID_CREDENTIALS;
		} else if (errorMsg.includes('CORS')) {
			return ERROR_MESSAGES.CORS_ERROR;
		} else if (!errorMsg || errorMsg === 'Login failed') {
			return ERROR_MESSAGES.DEFAULT;
		}
		
		return errorMsg;
	};

	/**
	 * Función genérica para procesar el login
	 * @param {Object} credentials - Credenciales de usuario
	 * @returns {Promise<void>}
	 */
	const processLogin = async (credentials) => {
        const response = await loginRequest(credentials);

        // Aceptar login exitoso si hay token O si el backend usó cookie HttpOnly
        if (!response.success || (!response.token && !response.cookieBased)) {
            throw new Error(ERROR_MESSAGES.NO_TOKEN);
        }

        // Si hay token, ya fue almacenado por loginRequest; si no, es cookie-based -> no almacenar token
        if (response.token) {
            storeToken(response.token);
        }

        // Solo almacenar user si viene en la respuesta
        if (response.user) {
            storeUser(response.user);
        }

        // Verificar si el usuario tiene permisos para acceder al dashboard
        console.log('📊 Response completa:', response);
        console.log('🔑 canAccessDashboard:', response.canAccessDashboard);
        console.log('👤 isAdmin:', response.isAdmin);

        const canAccessDashboard = response.canAccessDashboard || false;
        const isAdmin = response.isAdmin || false;

        console.log('✅ canAccessDashboard final:', canAccessDashboard);
        console.log('✅ isAdmin final:', isAdmin);

        // Mostrar mensaje de éxito brevemente
        setSuccess(true);
        setError(null);

        // Redirigir según los permisos del usuario
        const zajunaUrl = import.meta.env.VITE_ZAJUNA_DASHBOARD_URL || 'http://localhost:3000';

        if (canAccessDashboard) {
            // Usuario con permisos administrativos - redirigir al dashboard de admin
            setTimeout(() => {
                console.log('Usuario con permisos administrativos. Redirigiendo al dashboard:', zajunaUrl);

                // SEGURIDAD: NO pasar datos sensibles en la URL
                // El token ya está guardado en localStorage por storeToken()
                // El Dashboard lo leerá de ahí usando getToken() en apiClientUsers.js

                // Solo redirigir sin parámetros en la URL
                window.location.href = zajunaUrl;
            }, 1500); // 1.5 segundos para mostrar el mensaje de éxito
        } else {
            // Usuario estudiante - redirigir al StudentDashboard
            setTimeout(() => {
                console.log('Usuario estudiante. Redirigiendo al dashboard de estudiante:', zajunaUrl + '/student-dashboard');
                window.location.href = zajunaUrl + '/student-dashboard';
            }, 1500);
        }
    };

	/**
	 * Función genérica para manejar submit de formularios de login
	 * @param {Event} e - Evento de submit
	 * @param {Function} getCredentials - Función que extrae las credenciales del FormData
	 * @param {Function} validateFields - Función que valida los campos del formulario
	 */
	const handleLoginSubmit = async (e, getCredentials, validateFields) => {
		e.preventDefault();
		setLoading(true);
		setError(null);

		const formData = new FormData(e.target);

		// Validar campos
		const validationError = validateFields(formData);
		if (validationError) {
			setError(validationError);
			setLoading(false);
			return;
		}

		try {
			const credentials = getCredentials(formData);
			await processLogin(credentials);
		} catch (err) {
			setError(getErrorMessage(err));
		} finally {
			setLoading(false);
		}
	};

	// Handler para formulario de cursos (usa idnumber/documento + password)
	const handleCursosSubmit = (e) => {
		const validateFields = (formData) => {
			const typeDocument = formData.get('typeDocument');
			const document = formData.get('document');
			const password = formData.get('password');
			
			if (!typeDocument || !document || !password) {
				return ERROR_MESSAGES.EMPTY_FIELDS;
			}
			return null;
		};

		const getCredentials = (formData) => {
			const document = formData.get('document');
			const password = formData.get('password');
			return { idnumber: document.trim(), password };
		};

		return handleLoginSubmit(e, getCredentials, validateFields);
	};

	// Handler para formulario administrativo (usa username + password)
	const handleAdministrativoSubmit = (e) => {
		const validateFields = (formData) => {
			const username = formData.get('username');
			const password = formData.get('password');
			
			if (!username || !password) {
				return ERROR_MESSAGES.EMPTY_FIELDS;
			}
			return null;
		};

		const getCredentials = (formData) => {
			const username = formData.get('username');
			const password = formData.get('password');
			return { username: username.trim(), password };
		};

		return handleLoginSubmit(e, getCredentials, validateFields);
	};

	return (
		<div className="login">
			<div className="login__ingresos">
				<button
					className="login__ingreso-cursos"
					onClick={() => handleFormToggle('cursos')}
					id="login__ingreso-cursos"
				>
					<p className={`login__ingreso-p ${activeForm === 'cursos' ? 'login__ingreso-p--activo' : ''}`}>
						Ingreso <br /> cursos Zajuna
					</p>
				</button>
				<div className="login__ingreso-administrativo" id="login__ingreso-administrativo">
					<div className="login__linea"></div>
					<button
						className="login__ingreso-administrativo-btn"
						onClick={() => handleFormToggle('administrativo')}
					>
						<p className={`login__ingreso-p ${activeForm === 'administrativo' ? 'login__ingreso-p--activo' : ''}`}>
							Ingreso <br /> Administrativos Zajuna
						</p>
					</button>
				</div>
			</div>
			<div className={`login__form__container ${error ? 'login__form__container--has-error' : ''} ${success ? 'login__form__container--has-success' : ''}`}>
				{error && (
					<div className="login__error">
						{error}
					</div>
				)}
				{success && (
					<div className="login__success" style={{
						backgroundColor: '#d4edda',
						color: '#155724',
						padding: '12px 16px',
						borderRadius: '4px',
						marginBottom: '16px',
						border: '1px solid #c3e6cb'
					}}>
						✅ Login exitoso. Redirigiendo al dashboard...
					</div>
				)}
			<form
				id="login__form-cursos"
				autoComplete="off"
				className={activeForm === 'cursos' ? 'login__form--cursos' : 'login__form--inactivo login__form--cursos'}
				onSubmit={handleCursosSubmit}
			>
					<label className="login__label" htmlFor="typeDocument">Tipo de documento</label>
					<select className="login__select" id="typeDocument" name="typeDocument" required>
						<option value="">Seleccione su Documento</option>
						<option value="CC">Cédula de Ciudadanía</option>
						<option value="TI">Tarjeta de Identidad</option>
						<option value="CE">Cédula de Extranjería</option>
						<option value="PEP">PEP</option>
						<option value="PPT">Permiso por Protección Temporal</option>
					</select>
					<label className="login__label" htmlFor="document">Número de Documento</label>
					<input
						className="login__input"
						id="document"
						name="document"
						type="number"
						placeholder="Ingresa el Documento"
						pattern="[0-9]+"
						title="Número documento"
						required
					/>
					<label className="login__label" htmlFor="password">Contraseña</label>
					<input
						className="login__input"
						id="password"
						name="password"
						type="password"
						placeholder="Ingresa la Contraseña"
						required
						autoComplete="off"
					/>
					<a
						className="login__link"
						target="_blank"
						href="https://betowa.sena.edu.co/restablecer-contrasena"
						rel="noopener noreferrer"
					>
						Olvidé mi Contraseña
					</a>
					<a
						className="login__link"
						target="_blank"
						href="https://portal.senasofiaplus.edu.co/index.php/seguridad/usuario-bloqueado-o-inactivo"
						rel="noopener noreferrer"
					>
						Mi usuario está bloqueado o inactivo
					</a>
					<button 
						type="submit" 
						name="form_login_user" 
						className="login__boton"
						disabled={loading}
					>
						{loading ? 'Iniciando sesión...' : 'Iniciar sesión'}
					</button>
				</form>
				<form
					id="login__form-administrativo"
					autoComplete="off"
					className={activeForm === 'administrativo' ? '' : 'login__form--inactivo'}
					onSubmit={handleAdministrativoSubmit}
				>
					<div className="login__form__administrativo">
						<label className="login__label" htmlFor="inputName">Usuario</label>
						<input
							className="login__input"
							type="text"
							id="inputName"
							name="username"
							placeholder="Ingrese el usuario"
							required
						/>
						<label className="login__label" htmlFor="inputPassword">Contraseña</label>
						<input
							className="login__input"
							type="password"
							id="inputPassword"
							name="password"
							placeholder="Ingrese la contraseña"
							autoComplete="off"
							required
						/>
					</div>
					<button 
						type="submit" 
						className="login__boton" 
						id="submit"
						disabled={loading}
					>
						{loading ? 'Iniciando sesión...' : 'Iniciar sesión'}
					</button>
				</form>
			</div>
		</div>
	);
}

export default LoginForm;
