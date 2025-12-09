import React from 'react';

function MoreUserFilters(props) {
  const {
    visible,
    days, months, years,
    lastNameOp, setLastNameOp, lastNameRef, lastNameVal,
    firstNameOp, setFirstNameOp, firstNameRef, firstNameVal,
    usernameOp, setUsernameOp, usernameRef, usernameVal,
    emailOp, setEmailOp, emailRef, emailVal,
    cityOp, setCityOp, cityRef, cityVal,
    countryOp, countryVal, setCountryOp, setCountryVal,
    confirmedOp, setConfirmedOp, suspendedOp, setSuspendedOp,
    profileFieldName, setProfileFieldName, profileFieldOp, setProfileFieldOp, profileFieldRef, profileFieldVal,
    roleFieldOne, setRoleFieldOne, roleFieldTwo, setRoleFieldTwo, roleFieldRef, roleFieldVal, roleCategories,
    enrolledOp, setEnrolledOp, systemRoleOp, setSystemRoleOp,
    cohortOp, setCohortOp, cohortRef, cohortVal,
    firstAfterEnabled, setFirstAfterEnabled, firstAfterDay, setFirstAfterDay, firstAfterMonth, setFirstAfterMonth, firstAfterYear, setFirstAfterYear,
    firstBeforeEnabled, setFirstBeforeEnabled, firstBeforeDay, setFirstBeforeDay, firstBeforeMonth, setFirstBeforeMonth, firstBeforeYear, setFirstBeforeYear,
    lastAfterEnabled, setLastAfterEnabled, lastAfterDay, setLastAfterDay, lastAfterMonth, setLastAfterMonth, lastAfterYear, setLastAfterYear,
    lastBeforeEnabled, setLastBeforeEnabled, lastBeforeDay, setLastBeforeDay, lastBeforeMonth, setLastBeforeMonth, lastBeforeYear, setLastBeforeYear, lastNever, setLastNever,
    createdAfterEnabled, setCreatedAfterEnabled, createdAfterDay, setCreatedAfterDay, createdAfterMonth, setCreatedAfterMonth, createdAfterYear, setCreatedAfterYear,
    createdBeforeEnabled, setCreatedBeforeEnabled, createdBeforeDay, setCreatedBeforeDay, createdBeforeMonth, setCreatedBeforeMonth, createdBeforeYear, setCreatedBeforeYear,
    modifiedAfterEnabled, setModifiedAfterEnabled, modifiedAfterDay, setModifiedAfterDay, modifiedAfterMonth, setModifiedAfterMonth, modifiedAfterYear, setModifiedAfterYear,
    modifiedBeforeEnabled, setModifiedBeforeEnabled, modifiedBeforeDay, setModifiedBeforeDay, modifiedBeforeMonth, setModifiedBeforeMonth, modifiedBeforeYear, setModifiedBeforeYear, modifiedNever, setModifiedNever,
    identificationOp, setIdentificationOp, setIdentificationTouched,
    mnetProviderOp, setMnetProviderOp, setMnetProviderTouched,
    idNumberOp, setIdNumberOp, idNumberRef, idNumberVal,
    institutionOp, setInstitutionOp, institutionRef, institutionVal,
    departmentOp, setDepartmentOp, departmentRef, departmentVal,
    lastIpOp, setLastIpOp, lastIpRef, lastIpVal,
    afterCalendarIconRef, beforeCalendarIconRef, lastAfterCalendarIconRef, lastBeforeCalendarIconRef, createdAfterCalendarIconRef, createdBeforeCalendarIconRef, modifiedAfterCalendarIconRef, modifiedBeforeCalendarIconRef,
    openCalendarFor
  } = props;

  return (
    <div style={{ display: visible ? 'block' : 'none' }} aria-hidden={!visible}>
      {/* The content below was extracted from UsersTableView.js to reduce mount cost on open */}
      <div>
        <div className="filter-row" style={{ marginTop: 0 }}>
          <label>Apellido(s)</label>
          <div className="select-wrap select-wrapper form-inline bulk-caret">
            <select className="moodle-select" value={lastNameOp} onChange={(e) => setLastNameOp(e.target.value)}>
              <option value="contains">contiene</option>
              <option value="notcontains">no contiene</option>
              <option value="equals">es igual a</option>
              <option value="starts">comienza con</option>
              <option value="ends">termina en</option>
              <option value="empty">está vacío</option>
            </select>
          </div>
          <input
            ref={lastNameRef}
            className="moodle-search-input"
            type="text"
            defaultValue={lastNameVal}
            placeholder=""
          />
        </div>

        <div className="filter-row">
          <label>Nombre</label>
          <div className="select-wrap select-wrapper form-inline bulk-caret">
            <select className="moodle-select" value={firstNameOp} onChange={(e) => setFirstNameOp(e.target.value)}>
              <option value="contains">contiene</option>
              <option value="notcontains">no contiene</option>
              <option value="equals">es igual a</option>
              <option value="starts">comienza con</option>
              <option value="ends">termina en</option>
              <option value="empty">está vacío</option>
            </select>
          </div>
          <input
            ref={firstNameRef}
            className="moodle-search-input"
            type="text"
            defaultValue={firstNameVal}
            placeholder=""
          />
        </div>

        <div className="filter-row">
          <label>Nombre de usuario</label>
          <div className="select-wrap select-wrapper form-inline bulk-caret">
            <select className="moodle-select" value={usernameOp} onChange={(e) => setUsernameOp(e.target.value)}>
              <option value="contains">contiene</option>
              <option value="notcontains">no contiene</option>
              <option value="equals">es igual a</option>
              <option value="starts">comienza con</option>
              <option value="ends">termina en</option>
              <option value="empty">está vacío</option>
            </select>
          </div>
          <input
            ref={usernameRef}
            className="moodle-search-input"
            type="text"
            defaultValue={usernameVal}
            placeholder=""
          />
        </div>

        <div className="filter-row">
          <label>Dirección de correo</label>
          <div className="select-wrap select-wrapper form-inline bulk-caret">
            <select className="moodle-select" value={emailOp} onChange={(e) => setEmailOp(e.target.value)}>
              <option value="contains">contiene</option>
              <option value="notcontains">no contiene</option>
              <option value="equals">es igual a</option>
              <option value="starts">comienza con</option>
              <option value="ends">termina en</option>
              <option value="empty">está vacío</option>
            </select>
          </div>
          <input
            ref={emailRef}
            className="moodle-search-input"
            type="text"
            defaultValue={emailVal}
            placeholder=""
          />
        </div>

        <div className="filter-row">
          <label>Ciudad</label>
          <div className="select-wrap select-wrapper form-inline bulk-caret">
            <select className="moodle-select" value={cityOp} onChange={(e) => setCityOp(e.target.value)}>
              <option value="contains">contiene</option>
              <option value="notcontains">no contiene</option>
              <option value="equals">es igual a</option>
              <option value="starts">comienza con</option>
              <option value="ends">termina en</option>
              <option value="empty">está vacío</option>
            </select>
          </div>
          <input
            ref={cityRef}
            className="moodle-search-input"
            type="text"
            defaultValue={cityVal}
            placeholder=""
          />
        </div>

        <div className="filter-row">
          <label>País</label>
          <div className="select-wrap select-wrapper form-inline bulk-caret">
            <select className="moodle-select operator-select" value={countryOp} onChange={(e) => setCountryOp(e.target.value)}>
              <option value="any">cualquier valor</option>
              <option value="equals">es igual a</option>
              <option value="notequals">no es igual a</option>
            </select>
          </div>
          <div className="select-wrap select-wrapper form-inline bulk-caret">
            <select
              className="select custom-select ml-2 country-select"
              value={countryVal}
              onChange={(e) => setCountryVal(e.target.value)}
              disabled={countryOp === 'any'}
              aria-disabled={countryOp === 'any'}
              style={{ display: 'inline-block', verticalAlign: 'middle', maxWidth: 480, marginLeft: 0 }}
            >
              <option value="Colombia">Colombia</option>
            </select>
          </div>
        </div>

        <div className="filter-row">
          <label>Confirmado</label>
          <div className="select-wrap select-wrapper form-inline bulk-caret">
            <select
              className="moodle-select operator-select"
              value={confirmedOp}
              onChange={(e) => setConfirmedOp(e.target.value)}
            >
              <option value="any">cualquier valor</option>
              <option value="false">No</option>
              <option value="true">Sí</option>
            </select>
          </div>
        </div>

        <div className="filter-row">
          <label>Cuenta suspendida</label>
          <div className="select-wrap select-wrapper form-inline bulk-caret">
            <select
              className="moodle-select operator-select"
              value={suspendedOp}
              onChange={(e) => setSuspendedOp(e.target.value)}
            >
              <option value="any">cualquier valor</option>
              <option value="false">No</option>
              <option value="true">Sí</option>
            </select>
          </div>
        </div>

        <div className="filter-row">
          <label>Campos de perfil del usuario</label>
          <div className="select-wrap select-wrapper form-inline bulk-caret">
            <select className="moodle-select operator-select profile-field-name-select" value={profileFieldName} onChange={(e) => setProfileFieldName(e.target.value)}>
              <option value="any">cualquier campo</option>
              <option value="about">Acerca de mí</option>
              <option value="experience">Experiencia Profesional</option>
              <option value="profession">Profesión</option>
            </select>
          </div>
          <div className="select-wrap select-wrapper form-inline bulk-caret">
            <select className="moodle-select operator-select profile-field-op-select" value={profileFieldOp} onChange={(e) => setProfileFieldOp(e.target.value)}>
              <option value="contains">contiene</option>
              <option value="notcontains">no contiene</option>
              <option value="equals">es igual a</option>
              <option value="starts">comienza con</option>
              <option value="ends">termina en</option>
              <option value="empty">está vacío</option>
              <option value="notdefined">no está definido</option>
              <option value="defined">está definido</option>
            </select>
          </div>
          <input
            ref={profileFieldRef}
            className="moodle-search-input"
            type="text"
            defaultValue={profileFieldVal}
            placeholder=""
            style={{ display: 'inline-block', verticalAlign: 'middle', maxWidth: 480, marginLeft: 0 }}
          />
        </div>

        <div className="filter-row">
          <label>Rol de curso</label>
          <div className="select-wrap select-wrapper form-inline bulk-caret">
            <select className="moodle-select operator-select role-field-one-select" value={roleFieldOne} onChange={(e) => setRoleFieldOne(e.target.value)}>
              <option value="any">cualquier rol</option>
              <option value="learner">Aprendiz</option>
              <option value="instructor_no_edit">Instructor sin permiso de edición</option>
              <option value="instructor">Instructor</option>
              <option value="manager">Gestor</option>
            </select>
          </div>
          <div className="select-wrap select-wrapper form-inline bulk-caret">
            <select className="moodle-select operator-select role-field-two-select" value={roleFieldTwo} onChange={(e) => setRoleFieldTwo(e.target.value)}>
              <option value="any">cualquier categoría</option>
              {roleCategories && roleCategories.map((c) => (
                <option key={c.id ?? c.categoryid ?? c.name} value={c.id ?? c.categoryid ?? c.name}>
                  {c.name || c.fullname || c.displayname || c.shortname || c.id}
                </option>
              ))}
            </select>
          </div>
          <input
            ref={roleFieldRef}
            className="moodle-search-input"
            type="text"
            defaultValue={roleFieldVal}
            placeholder=""
            style={{ display: 'inline-block', verticalAlign: 'middle', maxWidth: 480, marginLeft: 0 }}
          />
        </div>

        <div className="filter-row">
          <label>Matriculado en cualquier curso</label>
          <div className="select-wrap select-wrapper form-inline bulk-caret">
            <select
              className="moodle-select operator-select"
              value={enrolledOp}
              onChange={(e) => setEnrolledOp(e.target.value)}
            >
              <option value="any">cualquier valor</option>
              <option value="false">No</option>
              <option value="true">Sí</option>
            </select>
          </div>
        </div>

        <div className="filter-row">
          <label>Rol del sistema</label>
          <div className="select-wrap select-wrapper form-inline bulk-caret">
            <select
              className="moodle-select operator-select system-role-select"
              style={{ width: '157.983px' }}
              value={systemRoleOp}
              onChange={(e) => setSystemRoleOp(e.target.value)}
            >
              <option value="any">cualquier rol</option>
              <option value="manager">Gestor</option>
              <option value="coursecreator">Creador de curso</option>
            </select>
          </div>
        </div>

        <div className="filter-row">
          <label>ID de la cohorte</label>
          <div className="select-wrap select-wrapper form-inline bulk-caret">
            <select className="moodle-select" value={cohortOp} onChange={(e) => setCohortOp(e.target.value)}>
              <option value="equals">es igual a</option>
              <option value="contains">contiene</option>
              <option value="notcontains">no contiene</option>
              <option value="starts">comienza con</option>
              <option value="ends">termina en</option>
              <option value="empty">está vacío</option>
            </select>
          </div>
          <input
            ref={cohortRef}
            className="moodle-search-input"
            type="text"
            defaultValue={cohortVal}
            placeholder=""
            style={{ display: 'inline-block', verticalAlign: 'middle', maxWidth: 480, marginLeft: 0 }}
          />
        </div>

        <div className="filter-row">
          <label style={{ alignSelf: 'flex-start', paddingTop: 6 }}>Primer acceso</label>
          <div style={{ display: 'flex', flexDirection: 'column', gap: 6 }}>
            <div style={{ display: 'flex', alignItems: 'center', gap: 6 }}>
              <div style={{ fontSize: 14, fontWeight: 400, color: '#333' }}>es posterior a</div>
              <input type="checkbox" checked={firstAfterEnabled} onChange={(e) => setFirstAfterEnabled(e.target.checked)} style={{ width: 14, height: 14, cursor: 'pointer', minWidth: 0 }} />
              <div style={{ marginLeft: 0, fontSize: 14, fontWeight: 400, color: '#333' }}>Habilitar</div>
              <div className="select-wrap select-wrapper form-inline bulk-caret">
                <select className="moodle-select primer-acceso-select primer-acceso-day-select" value={firstAfterDay} onChange={(e) => setFirstAfterDay(e.target.value)} disabled={!firstAfterEnabled} aria-disabled={!firstAfterEnabled}>
                  {days.map(d => (
                    <option key={d} value={String(d)}>{d}</option>
                  ))}
                </select>
              </div>
              <div className="select-wrap select-wrapper form-inline bulk-caret">
                <select className="moodle-select primer-acceso-select" value={firstAfterMonth} onChange={(e) => setFirstAfterMonth(e.target.value)} disabled={!firstAfterEnabled} aria-disabled={!firstAfterEnabled}>
                  {months.map((m, i) => (
                    <option key={i+1} value={String(i+1)}>{m}</option>
                  ))}
                </select>
              </div>
              <div className="select-wrap select-wrapper form-inline bulk-caret">
                <select className="moodle-select primer-acceso-select primer-acceso-year-select" value={firstAfterYear} onChange={(e) => setFirstAfterYear(e.target.value)} disabled={!firstAfterEnabled} aria-disabled={!firstAfterEnabled}>
                  {years.map(y => (
                    <option key={y} value={String(y)}>{y}</option>
                  ))}
                </select>
              </div>
              <i
                ref={afterCalendarIconRef}
                role="button"
                tabIndex={0}
                onKeyDown={(e) => { if (e.key === 'Enter' || e.key === ' ') { openCalendarFor('after', e.currentTarget); e.preventDefault(); } }}
                onClick={(e) => openCalendarFor('after', e.currentTarget)}
                className={`icon fa fa-calendar fa-fw calendar-icon ${firstAfterEnabled ? 'enabled' : 'disabled'}`}
                title="Calendario"
                aria-label="Calendario"
              />
            </div>
            <div style={{ display: 'flex', alignItems: 'center', gap: 6 }}>
              <div style={{ fontSize: 14, fontWeight: 400, color: '#333' }}>es anterior a</div>
              <input type="checkbox" checked={firstBeforeEnabled} onChange={(e) => setFirstBeforeEnabled(e.target.checked)} style={{ width: 14, height: 14, cursor: 'pointer', minWidth: 0 }} />
              <div style={{ marginLeft: 0, fontSize: 14, fontWeight: 400, color: '#333' }}>Habilitar</div>
              <div className="select-wrap select-wrapper form-inline bulk-caret">
                <select className="moodle-select primer-acceso-select primer-acceso-day-select" value={firstBeforeDay} onChange={(e) => setFirstBeforeDay(e.target.value)} disabled={!firstBeforeEnabled} aria-disabled={!firstBeforeEnabled}>
                  {days.map(d => (
                    <option key={d} value={String(d)}>{d}</option>
                  ))}
                </select>
              </div>
              <div className="select-wrap select-wrapper form-inline bulk-caret">
                <select className="moodle-select primer-acceso-select" value={firstBeforeMonth} onChange={(e) => setFirstBeforeMonth(e.target.value)} disabled={!firstBeforeEnabled} aria-disabled={!firstBeforeEnabled}>
                  {months.map((m, i) => (
                    <option key={i+1} value={String(i+1)}>{m}</option>
                  ))}
                </select>
              </div>
              <div className="select-wrap select-wrapper form-inline bulk-caret">
                <select className="moodle-select primer-acceso-select primer-acceso-year-select" value={firstBeforeYear} onChange={(e) => setFirstBeforeYear(e.target.value)} disabled={!firstBeforeEnabled} aria-disabled={!firstBeforeEnabled}>
                  {years.map(y => (
                    <option key={y} value={String(y)}>{y}</option>
                  ))}
                </select>
              </div>
              <i
                ref={beforeCalendarIconRef}
                role="button"
                tabIndex={0}
                onKeyDown={(e) => { if (e.key === 'Enter' || e.key === ' ') { openCalendarFor('before', e.currentTarget); e.preventDefault(); } }}
                onClick={(e) => openCalendarFor('before', e.currentTarget)}
                className={`icon fa fa-calendar fa-fw calendar-icon ${firstBeforeEnabled ? 'enabled' : 'disabled'}`}
                title="Calendario"
                aria-label="Calendario"
              />
            </div>
          </div>
        </div>

        {/* Último acceso */}
        <div className="filter-row">
          <label style={{ alignSelf: 'flex-start', paddingTop: 6 }}>Último acceso</label>
          <div style={{ display: 'flex', flexDirection: 'column', gap: 6 }}>
            <div style={{ display: 'flex', alignItems: 'center', gap: 6 }}>
              <div style={{ fontSize: 14, fontWeight: 400, color: '#333' }}>es posterior a</div>
              <input type="checkbox" checked={lastAfterEnabled} onChange={(e) => setLastAfterEnabled(e.target.checked)} style={{ width: 14, height: 14, cursor: 'pointer', minWidth: 0 }} />
              <div style={{ marginLeft: 0, fontSize: 14, fontWeight: 400, color: '#333' }}>Habilitar</div>
              <div className="select-wrap select-wrapper form-inline bulk-caret">
                <select className="moodle-select primer-acceso-select primer-acceso-day-select" value={lastAfterDay} onChange={(e) => setLastAfterDay(e.target.value)} disabled={!lastAfterEnabled} aria-disabled={!lastAfterEnabled}>
                  {days.map(d => (
                      <option key={d} value={String(d)}>{d}</option>
                    ))}
                </select>
              </div>
              <div className="select-wrap select-wrapper form-inline bulk-caret">
                <select className="moodle-select primer-acceso-select" value={lastAfterMonth} onChange={(e) => setLastAfterMonth(e.target.value)} disabled={!lastAfterEnabled} aria-disabled={!lastAfterEnabled}>
                  {months.map((m, i) => (
                      <option key={i+1} value={String(i+1)}>{m}</option>
                    ))}
                </select>
              </div>
              <div className="select-wrap select-wrapper form-inline bulk-caret">
                <select className="moodle-select primer-acceso-select primer-acceso-year-select" value={lastAfterYear} onChange={(e) => setLastAfterYear(e.target.value)} disabled={!lastAfterEnabled} aria-disabled={!lastAfterEnabled}>
                  {years.map(y => (
                      <option key={y} value={String(y)}>{y}</option>
                    ))}
                </select>
              </div>
              <i
                ref={lastAfterCalendarIconRef}
                role="button"
                tabIndex={0}
                onKeyDown={(e) => { if (e.key === 'Enter' || e.key === ' ') { openCalendarFor('last_after', e.currentTarget); e.preventDefault(); } }}
                onClick={(e) => openCalendarFor('last_after', e.currentTarget)}
                className={`icon fa fa-calendar fa-fw calendar-icon ${lastAfterEnabled ? 'enabled' : 'disabled'}`}
                title="Calendario"
                aria-label="Calendario"
              />
            </div>
            <div style={{ display: 'flex', alignItems: 'center', gap: 6 }}>
              <div style={{ fontSize: 14, fontWeight: 400, color: '#333' }}>es anterior a</div>
              <input type="checkbox" checked={lastBeforeEnabled} onChange={(e) => setLastBeforeEnabled(e.target.checked)} style={{ width: 14, height: 14, cursor: 'pointer', minWidth: 0 }} />
              <div style={{ marginLeft: 0, fontSize: 14, fontWeight: 400, color: '#333' }}>Habilitar</div>
              <div className="select-wrap select-wrapper form-inline bulk-caret">
                <select className="moodle-select primer-acceso-select primer-acceso-day-select" value={lastBeforeDay} onChange={(e) => setLastBeforeDay(e.target.value)} disabled={!lastBeforeEnabled} aria-disabled={!lastBeforeEnabled}>
                  {days.map(d => (
                      <option key={d} value={String(d)}>{d}</option>
                    ))}
                </select>
              </div>
              <div className="select-wrap select-wrapper form-inline bulk-caret">
                <select className="moodle-select primer-acceso-select" value={lastBeforeMonth} onChange={(e) => setLastBeforeMonth(e.target.value)} disabled={!lastBeforeEnabled} aria-disabled={!lastBeforeEnabled}>
                  {months.map((m, i) => (
                    <option key={i+1} value={String(i+1)}>{m}</option>
                  ))}
                </select>
              </div>
              <div className="select-wrap select-wrapper form-inline bulk-caret">
                <select className="moodle-select primer-acceso-select primer-acceso-year-select" value={lastBeforeYear} onChange={(e) => setLastBeforeYear(e.target.value)} disabled={!lastBeforeEnabled} aria-disabled={!lastBeforeEnabled}>
                  {years.map(y => (
                    <option key={y} value={String(y)}>{y}</option>
                  ))}
                </select>
              </div>
              <i
                ref={lastBeforeCalendarIconRef}
                role="button"
                tabIndex={0}
                onKeyDown={(e) => { if (e.key === 'Enter' || e.key === ' ') { openCalendarFor('last_before', e.currentTarget); e.preventDefault(); } }}
                onClick={(e) => openCalendarFor('last_before', e.currentTarget)}
                className={`icon fa fa-calendar fa-fw calendar-icon ${lastBeforeEnabled ? 'enabled' : 'disabled'}`}
                title="Calendario"
                aria-label="Calendario"
              />
            </div>
            <div style={{ display: 'flex', alignItems: 'center', gap: 8 }}>
              <input type="checkbox" checked={lastNever} onChange={(e) => setLastNever(e.target.checked)} style={{ width: 14, height: 14, cursor: 'pointer', minWidth: 0 }} />
              <div style={{ fontSize: 14, color: '#333' }}>Nunca se ha accedido</div>
            </div>
          </div>
        </div>

        {/* Tiempo de creado */}
        <div className="filter-row">
          <label style={{ alignSelf: 'flex-start', paddingTop: 6 }}>Tiempo de creado</label>
          <div style={{ display: 'flex', flexDirection: 'column', gap: 6 }}>
            <div style={{ display: 'flex', alignItems: 'center', gap: 6 }}>
              <div style={{ fontSize: 14, fontWeight: 400, color: '#333' }}>es posterior a</div>
              <input type="checkbox" checked={createdAfterEnabled} onChange={(e) => setCreatedAfterEnabled(e.target.checked)} style={{ width: 14, height: 14, cursor: 'pointer', minWidth: 0 }} />
              <div style={{ marginLeft: 0, fontSize: 14, fontWeight: 400, color: '#333' }}>Habilitar</div>
              <div className="select-wrap select-wrapper form-inline bulk-caret">
                <select className="moodle-select primer-acceso-select primer-acceso-day-select" value={createdAfterDay} onChange={(e) => setCreatedAfterDay(e.target.value)} disabled={!createdAfterEnabled} aria-disabled={!createdAfterEnabled}>
                  {days.map(d => (
                    <option key={d} value={String(d)}>{d}</option>
                  ))}
                </select>
              </div>
              <div className="select-wrap select-wrapper form-inline bulk-caret">
                <select className="moodle-select primer-acceso-select" value={createdAfterMonth} onChange={(e) => setCreatedAfterMonth(e.target.value)} disabled={!createdAfterEnabled} aria-disabled={!createdAfterEnabled}>
                  {months.map((m, i) => (
                    <option key={i+1} value={String(i+1)}>{m}</option>
                  ))}
                </select>
              </div>
              <div className="select-wrap select-wrapper form-inline bulk-caret">
                <select className="moodle-select primer-acceso-select primer-acceso-year-select" value={createdAfterYear} onChange={(e) => setCreatedAfterYear(e.target.value)} disabled={!createdAfterEnabled} aria-disabled={!createdAfterEnabled}>
                  {years.map(y => (
                    <option key={y} value={String(y)}>{y}</option>
                  ))}
                </select>
              </div>
              <i
                ref={createdAfterCalendarIconRef}
                role="button"
                tabIndex={0}
                onKeyDown={(e) => { if (e.key === 'Enter' || e.key === ' ') { openCalendarFor('created_after', e.currentTarget); e.preventDefault(); } }}
                onClick={(e) => openCalendarFor('created_after', e.currentTarget)}
                className={`icon fa fa-calendar fa-fw calendar-icon ${createdAfterEnabled ? 'enabled' : 'disabled'}`}
                title="Calendario"
                aria-label="Calendario"
              />
            </div>
            <div style={{ display: 'flex', alignItems: 'center', gap: 6 }}>
              <div style={{ fontSize: 14, fontWeight: 400, color: '#333' }}>es anterior a</div>
              <input type="checkbox" checked={createdBeforeEnabled} onChange={(e) => setCreatedBeforeEnabled(e.target.checked)} style={{ width: 14, height: 14, cursor: 'pointer', minWidth: 0 }} />
              <div style={{ marginLeft: 0, fontSize: 14, fontWeight: 400, color: '#333' }}>Habilitar</div>
              <div className="select-wrap select-wrapper form-inline bulk-caret">
                <select className="moodle-select primer-acceso-select primer-acceso-day-select" value={createdBeforeDay} onChange={(e) => setCreatedBeforeDay(e.target.value)} disabled={!createdBeforeEnabled} aria-disabled={!createdBeforeEnabled}>
                  {days.map(d => (
                    <option key={d} value={String(d)}>{d}</option>
                  ))}
                </select>
              </div>
              <div className="select-wrap select-wrapper form-inline bulk-caret">
                <select className="moodle-select primer-acceso-select" value={createdBeforeMonth} onChange={(e) => setCreatedBeforeMonth(e.target.value)} disabled={!createdBeforeEnabled} aria-disabled={!createdBeforeEnabled}>
                  {months.map((m, i) => (
                    <option key={i+1} value={String(i+1)}>{m}</option>
                  ))}
                </select>
              </div>
              <div className="select-wrap select-wrapper form-inline bulk-caret">
                <select className="moodle-select primer-acceso-select primer-acceso-year-select" value={createdBeforeYear} onChange={(e) => setCreatedBeforeYear(e.target.value)} disabled={!createdBeforeEnabled} aria-disabled={!createdBeforeEnabled}>
                  {years.map(y => (
                    <option key={y} value={String(y)}>{y}</option>
                  ))}
                </select>
              </div>
              <i
                ref={createdBeforeCalendarIconRef}
                role="button"
                tabIndex={0}
                onKeyDown={(e) => { if (e.key === 'Enter' || e.key === ' ') { openCalendarFor('created_before', e.currentTarget); e.preventDefault(); } }}
                onClick={(e) => openCalendarFor('created_before', e.currentTarget)}
                className={`icon fa fa-calendar fa-fw calendar-icon ${createdBeforeEnabled ? 'enabled' : 'disabled'}`}
                title="Calendario"
                aria-label="Calendario"
              />
            </div>
          </div>
        </div>

        {/* Última modificación */}
        <div className="filter-row">
          <label style={{ alignSelf: 'flex-start', paddingTop: 6 }}>Última modificación</label>
          <div style={{ display: 'flex', flexDirection: 'column', gap: 6 }}>
            <div style={{ display: 'flex', alignItems: 'center', gap: 6 }}>
              <div style={{ fontSize: 14, fontWeight: 400, color: '#333' }}>es posterior a</div>
              <input type="checkbox" checked={modifiedAfterEnabled} onChange={(e) => setModifiedAfterEnabled(e.target.checked)} style={{ width: 14, height: 14, cursor: 'pointer', minWidth: 0 }} />
              <div style={{ marginLeft: 0, fontSize: 14, fontWeight: 400, color: '#333' }}>Habilitar</div>
              <div className="select-wrap select-wrapper form-inline bulk-caret">
                <select className="moodle-select primer-acceso-select primer-acceso-day-select" value={modifiedAfterDay} onChange={(e) => setModifiedAfterDay(e.target.value)} disabled={!modifiedAfterEnabled} aria-disabled={!modifiedAfterEnabled}>
                    {days.map(d => (
                    <option key={d} value={String(d)}>{d}</option>
                  ))}
                </select>
              </div>
              <div className="select-wrap select-wrapper form-inline bulk-caret">
                <select className="moodle-select primer-acceso-select" value={modifiedAfterMonth} onChange={(e) => setModifiedAfterMonth(e.target.value)} disabled={!modifiedAfterEnabled} aria-disabled={!modifiedAfterEnabled}>
                  {months.map((m, i) => (
                    <option key={i+1} value={String(i+1)}>{m}</option>
                  ))}
                </select>
              </div>
              <div className="select-wrap select-wrapper form-inline bulk-caret">
                <select className="moodle-select primer-acceso-select primer-acceso-year-select" value={modifiedAfterYear} onChange={(e) => setModifiedAfterYear(e.target.value)} disabled={!modifiedAfterEnabled} aria-disabled={!modifiedAfterEnabled}>
                  {years.map(y => (
                    <option key={y} value={String(y)}>{y}</option>
                  ))}
                </select>
              </div>
              <i
                ref={modifiedAfterCalendarIconRef}
                role="button"
                tabIndex={0}
                onKeyDown={(e) => { if (e.key === 'Enter' || e.key === ' ') { openCalendarFor('modified_after', e.currentTarget); e.preventDefault(); } }}
                onClick={(e) => openCalendarFor('modified_after', e.currentTarget)}
                className={`icon fa fa-calendar fa-fw calendar-icon ${modifiedAfterEnabled ? 'enabled' : 'disabled'}`}
                title="Calendario"
                aria-label="Calendario"
              />
            </div>
            <div style={{ display: 'flex', alignItems: 'center', gap: 6 }}>
              <div style={{ fontSize: 14, fontWeight: 400, color: '#333' }}>es anterior a</div>
              <input type="checkbox" checked={modifiedBeforeEnabled} onChange={(e) => setModifiedBeforeEnabled(e.target.checked)} style={{ width: 14, height: 14, cursor: 'pointer', minWidth: 0 }} />
              <div style={{ marginLeft: 0, fontSize: 14, fontWeight: 400, color: '#333' }}>Habilitar</div>
              <div className="select-wrap select-wrapper form-inline bulk-caret">
                <select className="moodle-select primer-acceso-select primer-acceso-day-select" value={modifiedBeforeDay} onChange={(e) => setModifiedBeforeDay(e.target.value)} disabled={!modifiedBeforeEnabled} aria-disabled={!modifiedBeforeEnabled}>
                  {days.map(d => (
                    <option key={d} value={String(d)}>{d}</option>
                  ))}
                </select>
              </div>
              <div className="select-wrap select-wrapper form-inline bulk-caret">
                <select className="moodle-select primer-acceso-select" value={modifiedBeforeMonth} onChange={(e) => setModifiedBeforeMonth(e.target.value)} disabled={!modifiedBeforeEnabled} aria-disabled={!modifiedBeforeEnabled}>
                  {months.map((m, i) => (
                    <option key={i+1} value={String(i+1)}>{m}</option>
                  ))}
                </select>
              </div>
              <div className="select-wrap select-wrapper form-inline bulk-caret">
                <select className="moodle-select primer-acceso-select primer-acceso-year-select" value={modifiedBeforeYear} onChange={(e) => setModifiedBeforeYear(e.target.value)} disabled={!modifiedBeforeEnabled} aria-disabled={!modifiedBeforeEnabled}>
                  {years.map(y => (
                    <option key={y} value={String(y)}>{y}</option>
                  ))}
                </select>
              </div>
              <i
                ref={modifiedBeforeCalendarIconRef}
                role="button"
                tabIndex={0}
                onKeyDown={(e) => { if (e.key === 'Enter' || e.key === ' ') { openCalendarFor('modified_before', e.currentTarget); e.preventDefault(); } }}
                onClick={(e) => openCalendarFor('modified_before', e.currentTarget)}
                className={`icon fa fa-calendar fa-fw calendar-icon ${modifiedBeforeEnabled ? 'enabled' : 'disabled'}`}
                title="Calendario"
                aria-label="Calendario"
              />
            </div>
            <div style={{ display: 'flex', alignItems: 'center', gap: 8 }}>
              <input type="checkbox" checked={modifiedNever} onChange={(e) => setModifiedNever(e.target.checked)} style={{ width: 14, height: 14, cursor: 'pointer', minWidth: 0 }} />
              <div style={{ fontSize: 14, color: '#333' }}>Nunca se ha modificado</div>
            </div>
          </div>
        </div>

        {/* Identificación, MNET, idnumber, institución, departamento, lastip */}
        <div className="filter-row">
          <label>Identificación</label>
          <div className="select-wrap select-wrapper form-inline bulk-caret">
            <select className="moodle-select operator-select identification-select" value={identificationOp} onChange={(e) => { setIdentificationOp(e.target.value); setIdentificationTouched(true); }}>
              <option value="">cualquier valor</option>
              <option value="cas">Usar un servidor CAS (SSO)</option>
              <option value="db">Usar una base de datos externa</option>
              <option value="email">Identificación basada en Email</option>
              <option value="ldap">Usar un servidor LDAP</option>
              <option value="lti">LTI</option>
              <option value="manual">Cuentas manuales</option>
              <option value="mnet">Identificación de la Red Moodle ('Moodle Network')</option>
              <option value="nologin">No hay sesión</option>
              <option value="none">Sin identificación</option>
              <option value="oauth2">OAuth 2</option>
              <option value="oidc">Conexión OpenID</option>
              <option value="shibboleth">Shibboleth</option>
              <option value="webservice">Identificación de Servicios Web ('Web Services')</option>
            </select>
          </div>
        </div>
        <div className="filter-row">
          <label>Proveedor de ID MNET</label>
          <div className="select-wrap select-wrapper form-inline bulk-caret">
            <select className="moodle-select operator-select mnet-provider-select" value={mnetProviderOp} onChange={(e) => { setMnetProviderOp(e.target.value); setMnetProviderTouched(true); }}>
              <option value="">cualquier valor</option>
              <option value="sena_zajuna">SENA - Zajuna (local)</option>
              <option value="0">id: 0 (Error)</option>
            </select>
          </div>
        </div>
        <div className="filter-row">
          <label>Número de ID</label>
          <div className="select-wrap select-wrapper form-inline bulk-caret">
            <select className="moodle-select operator-select idnumber-select" value={idNumberOp} onChange={(e) => setIdNumberOp(e.target.value)}>
              <option value="contains">contiene</option>
              <option value="notcontains">no contiene</option>
              <option value="equals">es igual a</option>
              <option value="starts">comienza con</option>
              <option value="ends">termina en</option>
              <option value="empty">está vacío</option>
            </select>
          </div>
          <input
            ref={idNumberRef}
            className="moodle-search-input"
            type="text"
            defaultValue={idNumberVal}
            placeholder=""
            style={{ display: 'inline-block', verticalAlign: 'middle', maxWidth: 480, marginLeft: 0 }}
          />
        </div>
        <div className="filter-row">
          <label>Institución</label>
          <div className="select-wrap select-wrapper form-inline bulk-caret">
            <select className="moodle-select operator-select institution-select" value={institutionOp} onChange={(e) => setInstitutionOp(e.target.value)}>
              <option value="contains">contiene</option>
              <option value="notcontains">no contiene</option>
              <option value="equals">es igual a</option>
              <option value="starts">comienza con</option>
              <option value="ends">termina en</option>
              <option value="empty">está vacío</option>
            </select>
          </div>
          <input
            ref={institutionRef}
            className="moodle-search-input"
            type="text"
            defaultValue={institutionVal}
            placeholder=""
            style={{ display: 'inline-block', verticalAlign: 'middle', maxWidth: 480, marginLeft: 0 }}
          />
        </div>
        <div className="filter-row">
          <label>Departamento</label>
          <div className="select-wrap select-wrapper form-inline bulk-caret">
            <select className="moodle-select operator-select department-select" value={departmentOp} onChange={(e) => setDepartmentOp(e.target.value)}>
              <option value="contains">contiene</option>
              <option value="notcontains">no contiene</option>
              <option value="equals">es igual a</option>
              <option value="starts">comienza con</option>
              <option value="ends">termina en</option>
              <option value="empty">está vacío</option>
            </select>
          </div>
          <input
            ref={departmentRef}
            className="moodle-search-input"
            type="text"
            defaultValue={departmentVal}
            placeholder=""
            style={{ display: 'inline-block', verticalAlign: 'middle', maxWidth: 480, marginLeft: 0 }}
          />
        </div>
        <div className="filter-row">
          <label>Última dirección IP</label>
          <div className="select-wrap select-wrapper form-inline bulk-caret">
            <select className="moodle-select operator-select lastip-select" value={lastIpOp} onChange={(e) => setLastIpOp(e.target.value)}>
              <option value="contains">contiene</option>
              <option value="notcontains">no contiene</option>
              <option value="equals">es igual a</option>
              <option value="starts">comienza con</option>
              <option value="ends">termina en</option>
              <option value="empty">está vacío</option>
            </select>
          </div>
          <input
            ref={lastIpRef}
            className="moodle-search-input"
            type="text"
            defaultValue={lastIpVal}
            placeholder=""
            style={{ display: 'inline-block', verticalAlign: 'middle', maxWidth: 480, marginLeft: 0 }}
          />
        </div>
      </div>
    </div>
  );
}

export default React.memo(MoreUserFilters);
