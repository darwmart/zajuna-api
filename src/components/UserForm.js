import React, { useMemo, useState } from "react";
import { FaQuestionCircle, FaExclamationCircle, FaEye, FaEyeSlash } from "react-icons/fa";
import { createUser, updateUser } from "../services/usersService";
import "../styles/moodle-theme.css";

export default function UserForm({ user = null, onSuccess, onCancel }) {
  const isEdit = Boolean(user);
  const [collapsed, setCollapsed] = useState(false);
  const [showPassword, setShowPassword] = useState(false);
  const [formData, setFormData] = useState(
    user ? {
      id: user.id,
      username: user.username || "",
      auth: user.auth || "manual",
      suspended: Boolean(user.suspended),
      createpassword: false,
      forcepasswordchange: Boolean(user.forcepasswordchange),
      password: "",
      firstname: user.firstname || "",
      lastname: user.lastname || "",
      email: user.email || "",
      maildisplay: user.maildisplay ?? 2,
      city: user.city || "",
      country: user.country || "CO",
    } : {
      username: "",
      auth: "manual",
      suspended: false,
      createpassword: false,
      forcepasswordchange: false,
      password: "",
      firstname: "",
      lastname: "",
      email: "",
      maildisplay: 2,
      city: "",
      country: "CO",
    }
  );
  const [error, setError] = useState(null);
  const [success, setSuccess] = useState(null);

  const requiredInvalid = useMemo(() => {
    const errs = {};
    if (!formData.username?.trim()) errs.username = true;
    if (!formData.firstname?.trim()) errs.firstname = true;
    if (!formData.lastname?.trim()) errs.lastname = true;
    if (!formData.email?.trim()) errs.email = true;
    if (!isEdit && !formData.createpassword && !formData.password?.trim()) errs.password = true;
    return errs;
  }, [formData, isEdit]);

  const handleChange = (e) => {
    const { name, value, type, checked } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: type === "checkbox" ? checked : value,
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError(null);
    setSuccess(null);

    const hasErrors = Object.keys(requiredInvalid).length > 0;
    if (hasErrors) {
      setError("Por favor complete los campos requeridos.");
      return;
    }

    try {
      const payload = { ...formData };
      // Moodle acepta 0/1; enviamos booleanos o enteros según corresponda
      payload.suspended = formData.suspended ? 1 : 0;
      payload.createpassword = formData.createpassword ? 1 : 0;
      payload.forcepasswordchange = formData.forcepasswordchange ? 1 : 0;
      if (formData.createpassword) {
        delete payload.password;
      }

      if (isEdit) {
        await updateUser(payload);
        setSuccess("Usuario actualizado correctamente");
      } else {
        await createUser(payload);
        setSuccess("Usuario creado correctamente");
        setFormData({
          username: "",
          auth: "manual",
          suspended: false,
          createpassword: false,
          forcepasswordchange: false,
          password: "",
          firstname: "",
          lastname: "",
          email: "",
          maildisplay: 2,
          city: "",
          country: "CO",
        });
      }
      if (onSuccess) onSuccess();
    } catch (err) {
      setError("Error al guardar el usuario");
    }
  };

  return (
    <div className="moodle-userform user-form">
      {isEdit && (
        <div className="mform-user-title">
          <h2>{formData.firstname} {formData.lastname}</h2>
        </div>
      )}
      <div className="mform-topbar">
        <div className="mform-topbar-right">
          <button type="button" className="link-btn" onClick={() => setCollapsed((v) => !v)}>
            {collapsed ? "Expandir todo" : "Colapsar todo"}
          </button>
        </div>
        <div className="mform-section-row">
          <button
            type="button"
            className="mform-section-toggle"
            onClick={() => setCollapsed((v) => !v)}
            aria-expanded={!collapsed}
          >
            <span className="moodle-chevron-circle">
              <svg className="chevron" width="30" height="30" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg"
                style={{ transform: collapsed ? 'rotate(-90deg)' : 'rotate(0deg)', transition: 'transform 0.2s' }}>
                <path d="M6 8L10 12L14 8" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round"/>
              </svg>
            </span>
            <span className="section-title">General</span>
          </button>
        </div>
      </div>

      {!collapsed && (
        <form onSubmit={handleSubmit} className="mform-body" autoComplete="off">
          {/* Nombre de usuario */}
          <div className="mform-row">
            <label className="mform-label">
              <span>Nombre de usuario</span>
            </label>
            <div className="mform-field">
              <div className="field-with-icons">
                <span className="field-icons">
                  <FaExclamationCircle color="#c62828" title="Campo obligatorio" />
                </span>
                <input
                  type="text"
                  name="username"
                  value={formData.username}
                  onChange={handleChange}
                  className={`mform-input ${requiredInvalid.username ? "invalid" : ""}`}
                  disabled={isEdit}
                />
              </div>
            </div>
          </div>

          {/* Método de identificación (auth) */}
          <div className="mform-row">
            <label className="mform-label">
              <span>Escoger un método de identificación:</span>
            </label>
            <div className="mform-field">
              <div className="field-with-icons">
                <span className="field-icons">
                  <FaQuestionCircle color="#0D47A1" title="Seleccione el método de autenticación del usuario." />
                </span>
                <select name="auth" value={formData.auth} onChange={handleChange} className="mform-select">
                  <option value="manual">Cuentas manuales</option>
                  <option value="oauth2">OAuth 2</option>
                  <option value="ldap">LDAP</option>
                </select>
              </div>
              <div className="mform-check">
                <label className="checkbox">
                  <div className="field-with-icons">
                    <span className="field-icons" aria-hidden="true"></span>
                    <input type="checkbox" name="suspended" checked={!!formData.suspended} onChange={handleChange} />
                    <span>Cuenta suspendida</span>
                  </div>
                  
                </label>
              </div>
              <div className="mform-check">
                <label className="checkbox">
                  <div className="field-with-icons">
                    <span className="field-icons" aria-hidden="true"></span>
                    <input type="checkbox" name="createpassword" checked={!!formData.createpassword} onChange={handleChange} />
                    <span>Generar contraseña y notificar al usuario</span>
                  </div>
                  
                </label>
              </div>
              <div className="mform-hint field-with-icons">
                <span className="field-icons" aria-hidden="true"></span>
                <span>
                  La contraseña debería tener al menos 8 carácter(es), al menos 1 dígito(s), al menos 1 minúscula(s), al menos 1 mayúscula(s), al menos 1 carácter(es) especial(es) como *, -, o #
                </span>
              </div>
            </div>
          </div>

          {/* Nueva contraseña */}
          {!isEdit && (
            <div className="mform-row">
              <label className="mform-label">
                <span>Nueva contraseña</span>
              </label>
              <div className="mform-field password-field">
                <div className="password-input-row">
                  <div className="field-with-icons">
                    <span className="field-icons">
                      {(!formData.createpassword) ? (
                        <FaExclamationCircle color="#c62828" title="Campo obligatorio" />
                      ) : (
                        <FaQuestionCircle color="#0D47A1" title="Establezca una contraseña para el usuario." />
                      )}
                    </span>
                    <input
                      type={showPassword ? "text" : "password"}
                      name="password"
                      value={formData.password}
                      onChange={handleChange}
                      className={`mform-input ${requiredInvalid.password ? "invalid" : ""}`}
                      placeholder="Haz click para insertar texto"
                      disabled={!!formData.createpassword}
                    />
                  </div>
                  <button type="button" className="icon-btn eye" onClick={() => setShowPassword((v) => !v)} disabled={!!formData.createpassword}>
                    {showPassword ? <FaEyeSlash /> : <FaEye />}
                  </button>
                </div>
                <div className="mform-check">
                  <label className="checkbox">
                    <div className="field-with-icons">
                      <span className="field-icons" aria-hidden="true"></span>
                      <input type="checkbox" name="forcepasswordchange" checked={!!formData.forcepasswordchange} onChange={handleChange} />
                      <span>Forzar cambio de contraseña</span>
                    </div>
                    
                  </label>
                </div>
              </div>
            </div>
          )}

          {/* Nombre */}
          <div className="mform-row">
            <label className="mform-label">
              <span>Nombre</span>
            </label>
            <div className="mform-field">
              <div className="field-with-icons">
                <span className="field-icons">
                  <FaExclamationCircle color="#c62828" title="Campo obligatorio" />
                </span>
                <input
                  type="text"
                  name="firstname"
                  value={formData.firstname}
                  onChange={handleChange}
                  className={`mform-input ${requiredInvalid.firstname ? "invalid" : ""}`}
                />
              </div>
            </div>
          </div>

          {/* Apellido(s) */}
          <div className="mform-row">
            <label className="mform-label">
              <span>Apellido(s)</span>
            </label>
            <div className="mform-field">
              <div className="field-with-icons">
                <span className="field-icons">
                  <FaExclamationCircle color="#c62828" title="Campo obligatorio" />
                </span>
                <input
                  type="text"
                  name="lastname"
                  value={formData.lastname}
                  onChange={handleChange}
                  className={`mform-input ${requiredInvalid.lastname ? "invalid" : ""}`}
                />
              </div>
            </div>
          </div>

          {/* Dirección de correo */}
          <div className="mform-row">
            <label className="mform-label">
              <span>Dirección de correo</span>
            </label>
            <div className="mform-field">
              <div className="field-with-icons">
                <span className="field-icons">
                  <FaExclamationCircle color="#c62828" title="Campo obligatorio" />
                </span>
                <input
                  type="email"
                  name="email"
                  value={formData.email}
                  onChange={handleChange}
                  className={`mform-input ${requiredInvalid.email ? "invalid" : ""}`}
                />
              </div>
            </div>
          </div>

          {/* Visibilidad del correo electrónico */}
          <div className="mform-row">
            <label className="mform-label">
              <span>Visibilidad del correo electrónico</span>
            </label>
            <div className="mform-field">
              <div className="field-with-icons">
                <span className="field-icons">
                  <FaQuestionCircle color="#0D47A1" title="Controla quién puede ver la dirección de correo." />
                </span>
                <select name="maildisplay" value={formData.maildisplay} onChange={handleChange} className="mform-select">
                  <option value={0}>Ocultar mi dirección a todos</option>
                  <option value={2}>Visible para los participantes en el curso</option>
                  <option value={1}>Permitir que todos vean mi dirección</option>
                </select>
              </div>
            </div>
          </div>

          {/* Ciudad/Pueblo */}
          <div className="mform-row">
            <label className="mform-label"><span>Ciudad</span></label>
            <div className="mform-field">
              <div className="field-with-icons">
                <span className="field-icons" aria-hidden="true"></span>
                <input type="text" name="city" value={formData.city} onChange={handleChange} className="mform-input" />
              </div>
            </div>
          </div>

          {/* País */}
          <div className="mform-row">
            <label className="mform-label"><span>Seleccione su país</span></label>
            <div className="mform-field">
              <div className="field-with-icons">
                <span className="field-icons" aria-hidden="true"></span>
                <select name="country" value={formData.country} onChange={handleChange} className="mform-select">
                  <option value="CO">Colombia</option>
                  <option value="US">Estados Unidos</option>
                  <option value="MX">México</option>
                  <option value="ES">España</option>
                  <option value="AR">Argentina</option>
                </select>
              </div>
            </div>
          </div>

          {/* Alerts */}
          {error && (
            <div className="mform-alert error">{error}</div>
          )}
          {success && (
            <div className="mform-alert success">{success}</div>
          )}
          {/* Actions - Always visible */}
          <div className="mform-row">
            <div className="mform-label" />
            <div className="mform-field">
              <div className="mform-actions">
                <button type="submit" className="btn-primary">{isEdit ? "Actualizar usuario" : "Crear usuario"}</button>
                <button type="button" className="btn-secondary" onClick={onCancel}>Cancelar</button>
              </div>
            </div>
          </div>
        </form>
      )}
    </div>
  );
}
