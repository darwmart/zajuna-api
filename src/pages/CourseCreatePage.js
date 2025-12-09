import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { FaQuestionCircle, FaExclamationCircle, FaCalendarAlt } from 'react-icons/fa';
import AttoToolbar from '../components/AttoToolbar/AttoToolbar';
import '../styles/moodle-theme.css';

export default function CourseCreatePage() {
  const navigate = useNavigate();
  const [collapsedAll, setCollapsedAll] = useState(false);
  const [sectionsCollapsed, setSectionsCollapsed] = useState({
    general: false,
    description: false,
    format: true,
    appearance: true,
    files: true,
    completion: true,
    groups: true,
    tags: true,
  });

  const [fullname, setFullname] = useState('');
  const [shortname, setShortname] = useState('');
  const [category, setCategory] = useState('1');
  const [summary, setSummary] = useState('');

  const toggleSection = (key) => setSectionsCollapsed((s) => ({ ...s, [key]: !s[key] }));
  const toggleAll = () => {
    const next = !collapsedAll;
    setCollapsedAll(next);
    setSectionsCollapsed((prev) => Object.keys(prev).reduce((acc, k) => ({ ...acc, [k]: next }), {}));
  };

  const handleSave = (e) => {
    e && e.preventDefault();
    console.log('Guardar curso (simulado)', { fullname, shortname, category, summary });
    navigate('/courses/manage');
  };

  return (
    <div className="moodle-userform course-create-page">
      <div className="mform-user-title" style={{ paddingLeft: '0px' }}>
        <h2 style={{ margin: 0, fontSize: 32, fontWeight: 400 }}>Crear un nuevo curso</h2>
      </div>

      <div className="mform-topbar">
        <div className="mform-topbar-right">
          <button type="button" className="link-btn" onClick={toggleAll}>
            {collapsedAll ? 'Expandir todo' : 'Colapsar todo'}
          </button>
        </div>
      </div>

      <form onSubmit={handleSave} className="mform-body" autoComplete="off" style={{paddingLeft: '0px'}}>
        {/* General section */}
        <div className="mform-section-row">
          <button
            type="button"
            className="mform-section-toggle"
            onClick={() => toggleSection('general')}
            aria-expanded={!sectionsCollapsed.general}
          >
            <span className="moodle-chevron-circle">
              <svg className="chevron" width="30" height="30" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg"
                style={{ transform: sectionsCollapsed.general ? 'rotate(-90deg)' : 'rotate(0deg)', transition: 'transform 0.2s' }}>
                <path d="M6 8L10 12L14 8" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round"/>
              </svg>
            </span>
            <span className="section-title">General</span>
          </button>
        </div>

        {!sectionsCollapsed.general && (
          <>
            <div className="mform-row">
              <label className="mform-label"><span>Nombre completo del curso</span></label>
              <div className="mform-field">
                <div className="field-with-icons">
                  <span className="field-icons"><FaExclamationCircle color="#c62828" /></span>
                  <input
                    type="text"
                    className="mform-input"
                    value={fullname}
                    onChange={(e) => setFullname(e.target.value)}
                    placeholder=""
                    style={{ width: '520px', maxWidth: '100%' }}
                  />
                  <span style={{ marginLeft: 8 }} className="help-icon"><FaQuestionCircle color="#0D47A1" /></span>
                </div>
              </div>
            </div>

            <div className="mform-row">
              <label className="mform-label"><span>Nombre corto del curso</span></label>
              <div className="mform-field">
                <div className="field-with-icons">
                  <span className="field-icons"><FaExclamationCircle color="#c62828" /></span>
                  <input type="text" className="mform-input" value={shortname} onChange={(e) => setShortname(e.target.value)} style={{ width: 220 }} />
                  <span style={{ marginLeft: 8 }} className="help-icon"><FaQuestionCircle color="#0D47A1" /></span>
                </div>
              </div>
            </div>

            <div className="mform-row">
              <label className="mform-label"><span>Categoría de cursos</span></label>
              <div className="mform-field">
                <div style={{ display: 'flex', flexDirection: 'column', gap: 8 }}>
                  <div style={{ display: 'flex', alignItems: 'center', gap: 8 }}>
                    <span style={{ color: '#c62828' }} className="field-icons"><FaExclamationCircle /></span>
                    <div style={{ display: 'inline-flex', alignItems: 'center' }}>
                      <span style={{ border: '2px solid #3fa236', borderRadius: 6, padding: '4px 8px', background: '#fff', color: '#0b3b0e', fontWeight: 600 }}>× Categoría 1</span>
                    </div>
                    <span style={{ marginLeft: 8 }} className="help-icon"><FaQuestionCircle color="#0D47A1" /></span>
                  </div>
                  <select className="mform-select" value={category} onChange={(e) => setCategory(e.target.value)} style={{ width: 260 }}>
                    <option value="1">Buscar</option>
                    <option value="2">Semillas_Titulada</option>
                    <option value="3">Formación Complementaria</option>
                  </select>
                </div>
              </div>
            </div>

            <div className="mform-row">
              <label className="mform-label"><span>Visibilidad del curso</span></label>
              <div className="mform-field">
                <div className="field-with-icons">
                  <span className="field-icons"><FaQuestionCircle color="#0D47A1" /></span>
                  <select className="mform-select" style={{ width: 140 }}>
                    <option>Mostrar</option>
                    <option>Ocultar</option>
                  </select>
                </div>
              </div>
            </div>

            <div className="mform-row">
              <label className="mform-label"><span>Fecha de inicio del curso</span></label>
              <div className="mform-field">
                <div style={{ display: 'flex', alignItems: 'center', gap: 8 }}>
                  <select className="mform-select" style={{ width: 70 }} defaultValue={27}><option>27</option></select>
                  <select className="mform-select" style={{ width: 120 }} defaultValue={'noviembre'}><option>noviembre</option></select>
                  <select className="mform-select" style={{ width: 90 }} defaultValue={2025}><option>2025</option></select>
                  <select className="mform-select" style={{ width: 60 }} defaultValue={'00'}><option>00</option></select>
                  <select className="mform-select" style={{ width: 60 }} defaultValue={'00'}><option>00</option></select>
                  <FaCalendarAlt color="#3fa236" style={{ marginLeft: 6 }} />
                </div>
              </div>
            </div>

            <div className="mform-row">
              <label className="mform-label"><span>Fecha de finalización del curso</span></label>
              <div className="mform-field">
                <div style={{ display: 'flex', alignItems: 'center', gap: 8 }}>
                  <label style={{ display: 'flex', alignItems: 'center', gap: 6 }}><input type="checkbox" defaultChecked /> Habilitar</label>
                  <select className="mform-select" style={{ width: 70 }} defaultValue={27}><option>27</option></select>
                  <select className="mform-select" style={{ width: 120 }} defaultValue={'noviembre'}><option>noviembre</option></select>
                  <select className="mform-select" style={{ width: 90 }} defaultValue={2026}><option>2026</option></select>
                  <select className="mform-select" style={{ width: 60 }} defaultValue={'00'}><option>00</option></select>
                  <select className="mform-select" style={{ width: 60 }} defaultValue={'00'}><option>00</option></select>
                  <FaCalendarAlt color="#3fa236" style={{ marginLeft: 6 }} />
                </div>
              </div>
            </div>

            <div className="mform-row">
              <label className="mform-label"><span>Número ID del curso</span></label>
              <div className="mform-field">
                <input type="text" className="mform-input" style={{ width: 180 }} />
              </div>
            </div>
          </>
        )}

        {/* Divider between General and Description */}
        <div style={{ borderTop: '1px solid #e6e6e6', margin: '12px 0' }} />

        {/* Description */}
        <div className="mform-section-row" style={{ marginTop: 18 }}>
          <button type="button" className="mform-section-toggle" onClick={() => toggleSection('description')} aria-expanded={!sectionsCollapsed.description}>
            <span className="moodle-chevron-circle">
              <svg className="chevron" width="30" height="30" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg"
                style={{ transform: sectionsCollapsed.description ? 'rotate(-90deg)' : 'rotate(0deg)', transition: 'transform 0.2s' }}>
                <path d="M6 8L10 12L14 8" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round"/>
              </svg>
            </span>
            <span className="section-title">Descripción</span>
          </button>
        </div>

        {!sectionsCollapsed.description && (
          <>
            <div className="mform-row">
              <label className="mform-label"><span>Resumen del curso</span></label>
              <div className="mform-field">
                <AttoToolbar value={summary} onChange={(v) => setSummary(v)} />
              </div>
            </div>

            <div className="mform-row">
              <label className="mform-label"><span>Archivos del resumen del curso</span></label>
              <div className="mform-field">
                <div className="file-drop-placeholder" style={{ border: '1px dashed #cfcfcf', padding: 28, borderRadius: 6 }}>
                  <div style={{ textAlign: 'center', color: '#888' }}>Puede arrastrar y soltar archivos aquí para añadirlos</div>
                </div>
              </div>
            </div>
          </>
        )}

        {/* Divider between Description and Format */}
        <div style={{ borderTop: '1px solid #e6e6e6', margin: '12px 0' }} />

        {/* Format */}
        <div className="mform-section-row" style={{ marginTop: 18 }}>
          <button type="button" className="mform-section-toggle" onClick={() => toggleSection('format')} aria-expanded={!sectionsCollapsed.format}>
            <span className="moodle-chevron-circle">
              <svg className="chevron" width="30" height="30" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg"
                style={{ transform: sectionsCollapsed.format ? 'rotate(-90deg)' : 'rotate(0deg)', transition: 'transform 0.2s' }}>
                <path d="M6 8L10 12L14 8" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round"/>
              </svg>
            </span>
            <span className="section-title">Formato de curso</span>
          </button>
        </div>

        {!sectionsCollapsed.format && (
          <div className="mform-row">
            <label className="mform-label"><span>Formato</span></label>
            <div className="mform-field">
              <select className="mform-select"><option>Formato de Secciones Flexibles</option></select>
            </div>
          </div>
        )}

        <div style={{ height: 12 }} />

        {/* Divider */}
        <div style={{ borderTop: '1px solid #e6e6e6', margin: '16px 0' }} />

        {/* Apariencia */}
        <div className="mform-section-row" style={{ marginTop: 6 }}>
          <button type="button" className="mform-section-toggle" onClick={() => toggleSection('appearance')} aria-expanded={!sectionsCollapsed.appearance}>
            <span className="moodle-chevron-circle">
              <svg className="chevron" width="30" height="30" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg"
                style={{ transform: sectionsCollapsed.appearance ? 'rotate(-90deg)' : 'rotate(0deg)', transition: 'transform 0.2s' }}>
                <path d="M6 8L10 12L14 8" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round"/>
              </svg>
            </span>
            <span className="section-title">Apariencia</span>
          </button>
        </div>

        {!sectionsCollapsed.appearance && (
          <>
            <div className="mform-row">
              <label className="mform-label"><span>Forzar idioma</span></label>
              <div className="mform-field">
                <select className="mform-select"><option>No forzar</option><option>Español</option></select>
              </div>
            </div>

            <div className="mform-row">
              <label className="mform-label"><span>Número de anuncios</span></label>
              <div className="mform-field">
                <input type="number" className="mform-input" defaultValue={5} />
              </div>
            </div>
          </>
        )}

        <div style={{ borderTop: '1px solid #e6e6e6', margin: '16px 0' }} />

        {/* Archivos y subida */}
        <div className="mform-section-row" style={{ marginTop: 6 }}>
          <button type="button" className="mform-section-toggle" onClick={() => toggleSection('files')} aria-expanded={!sectionsCollapsed.files}>
            <span className="moodle-chevron-circle">
              <svg className="chevron" width="30" height="30" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg"
                style={{ transform: sectionsCollapsed.files ? 'rotate(-90deg)' : 'rotate(0deg)', transition: 'transform 0.2s' }}>
                <path d="M6 8L10 12L14 8" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round"/>
              </svg>
            </span>
            <span className="section-title">Archivos y subida</span>
          </button>
        </div>

        {!sectionsCollapsed.files && (
          <div className="mform-row">
            <label className="mform-label"><span>Tamaño máximo para archivos cargados por usuarios</span></label>
            <div className="mform-field">
              <select className="mform-select"><option>Sitio límite de subida (1 GB)</option></select>
            </div>
          </div>
        )}

        <div style={{ borderTop: '1px solid #e6e6e6', margin: '16px 0' }} />

        {/* Rastreo de finalización */}
        <div className="mform-section-row" style={{ marginTop: 6 }}>
          <button type="button" className="mform-section-toggle" onClick={() => toggleSection('completion')} aria-expanded={!sectionsCollapsed.completion}>
            <span className="moodle-chevron-circle">
              <svg className="chevron" width="30" height="30" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg"
                style={{ transform: sectionsCollapsed.completion ? 'rotate(-90deg)' : 'rotate(0deg)', transition: 'transform 0.2s' }}>
                <path d="M6 8L10 12L14 8" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round"/>
              </svg>
            </span>
            <span className="section-title">Rastreo de finalización</span>
          </button>
        </div>

        {!sectionsCollapsed.completion && (
          <div className="mform-row">
            <label className="mform-label"><span>Habilitar seguimiento del grado de finalización</span></label>
            <div className="mform-field">
              <select className="mform-select"><option>Sí</option><option>No</option></select>
            </div>
          </div>
        )}

        <div style={{ borderTop: '1px solid #e6e6e6', margin: '16px 0' }} />

        {/* Grupos */}
        <div className="mform-section-row" style={{ marginTop: 6 }}>
          <button type="button" className="mform-section-toggle" onClick={() => toggleSection('groups')} aria-expanded={!sectionsCollapsed.groups}>
            <span className="moodle-chevron-circle">
              <svg className="chevron" width="30" height="30" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg"
                style={{ transform: sectionsCollapsed.groups ? 'rotate(-90deg)' : 'rotate(0deg)', transition: 'transform 0.2s' }}>
                <path d="M6 8L10 12L14 8" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round"/>
              </svg>
            </span>
            <span className="section-title">Grupos</span>
          </button>
        </div>

        {!sectionsCollapsed.groups && (
          <>
            <div className="mform-row">
              <label className="mform-label"><span>Modo de grupo</span></label>
              <div className="mform-field">
                <select className="mform-select"><option>No hay grupos</option></select>
              </div>
            </div>

            <div className="mform-row">
              <label className="mform-label"><span>Forzar el modo de grupo</span></label>
              <div className="mform-field">
                <select className="mform-select"><option>No</option><option>Sí</option></select>
              </div>
            </div>
          </>
        )}

        <div style={{ borderTop: '1px solid #e6e6e6', margin: '16px 0' }} />

        {/* Marcas */}
        <div className="mform-section-row" style={{ marginTop: 6 }}>
          <button type="button" className="mform-section-toggle" onClick={() => toggleSection('tags')} aria-expanded={!sectionsCollapsed.tags}>
            <span className="moodle-chevron-circle">
              <svg className="chevron" width="30" height="30" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg"
                style={{ transform: sectionsCollapsed.tags ? 'rotate(-90deg)' : 'rotate(0deg)', transition: 'transform 0.2s' }}>
                <path d="M6 8L10 12L14 8" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round"/>
              </svg>
            </span>
            <span className="section-title">Marcas</span>
          </button>
        </div>

        {!sectionsCollapsed.tags && (
          <div className="mform-row">
            <label className="mform-label"><span>Marcas</span></label>
            <div className="mform-field">
              <div style={{ color: '#2e8b35', marginBottom: 6 }}>Administrar marcas estándar</div>
              <input type="text" className="mform-input" placeholder="Introduzca etiquetas..." />
            </div>
          </div>
        )}

        {/* Divider between Marcas and Actions */}
        <div style={{ borderTop: '1px solid #e6e6e6', margin: '16px 0' }} />

        {/* Actions */}
        <div className="mform-row" style={{ marginTop: 24 }}>
          <div className="mform-label" />
          <div className="mform-field">
            <div className="mform-actions">
              <button type="submit" className="btn-primary" style={{ background: 'var(--moodle-green)', borderColor: 'var(--moodle-green)' }}>Guardar y volver</button>
              <button type="button" className="btn-primary" onClick={handleSave} style={{ background: 'var(--moodle-green)', borderColor: 'var(--moodle-green)' }}>Guardar cambios y mostrar</button>
              <button type="button" className="btn-secondary" onClick={() => navigate('/courses/manage')}>Cancelar</button>
            </div>
          </div>
        </div>

        <div className="fdescription required" style={{ marginTop: 12 }}>
          <FaExclamationCircle className="text-danger" /> Requerido
        </div>
      </form>
    </div>
  );
}
