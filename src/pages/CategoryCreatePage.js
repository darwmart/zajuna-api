import React, { useEffect, useState } from 'react';
import { useNavigate, useLocation } from 'react-router-dom';
import AttoToolbar from '../components/AttoToolbar/AttoToolbar';
import '../styles/moodle-theme.css';
import { coursesService } from '../services/coursesService';

export default function CategoryCreatePage() {
  const [parent, setParent] = useState('');
  const [name, setName] = useState('');
  const [idNumber, setIdNumber] = useState('');
  const [description, setDescription] = useState('');
  const [categories, setCategories] = useState([]);
  const [loading, setLoading] = useState(true);
  const [parentError, setParentError] = useState('');
  const [selectValue, setSelectValue] = useState('');
  const [creating, setCreating] = useState(false);
  const [error, setError] = useState(null);
  const navigate = useNavigate();
  const location = useLocation();

  useEffect(() => {
    let mounted = true;
    const load = async () => {
      try {
        const res = await coursesService.getCategories();
        const cats = Array.isArray(res) ? res : (res.items || res.categories || res || []);
        if (!mounted) return;
        setCategories(cats);
        // If navigation state provided a parentId (user came from Categories panel), prefer it.
        const incomingParent = location && location.state && location.state.parentId ? String(location.state.parentId) : null;
        if (incomingParent) {
          setParent(incomingParent);
        } else {
          // default new category to be a top-level ('Superior') category
          setParent('0');
        }
      } catch (e) {
        setCategories([]);
      } finally {
        if (mounted) setLoading(false);
      }
    };
    load();
    return () => { mounted = false; };
  }, []);

  const handleCreate = async (e) => {
    e.preventDefault();

    // Validaciones
    if (!parent) {
      setParentError('Debe suministrar un valor aquí');
      return;
    }

    if (!name.trim()) {
      setError('El nombre de la categoría es requerido');
      return;
    }

    setCreating(true);
    setError(null);

    try {
      // Convertir parent a número (0 para Superior, o el ID de la categoría padre)
      const parentValue = parseInt(parent) || 0;

      const result = await coursesService.createCategory({
        categories: [{
          name: name.trim(),
          parent: parentValue,
          idnumber: idNumber.trim() || undefined,
          description: description || undefined,
          descriptionformat: 1, // HTML por defecto
          theme: undefined
        }]
      });

      console.log('Categoría creada:', result);

      // Redirigir a la vista de administración de cursos con las categorías actualizadas
      // El backend devuelve la lista completa de categorías después de crear
      navigate('/courses/manage', {
        state: {
          updatedCategories: result.categories || result,
          successMessage: 'Categoría creada exitosamente'
        }
      });

    } catch (err) {
      console.error('Error creando categoría:', err);
      setError(err.response?.data?.message || err.message || 'Error al crear la categoría');
    } finally {
      setCreating(false);
    }
  };

  // current selected category name (derived from categories list)
  const selectedName = parent === '0' ? 'Superior' : (parent ? ((categories || []).find((c) => String(c.id) === String(parent)) || {}).name || '' : '');

  // Reusable required icon SVG component
  const RequiredIcon = ({ style = {} }) => (
    <svg width="18" height="18" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" style={style}>
      <circle cx="12" cy="12" r="10" fill="#c62828"/>
      <text x="12" y="14" textAnchor="middle" fontSize="14" fontWeight="700" fill="#fff">!</text>
    </svg>
  );

  // Use users-category-page so the layout fills the available view and is not boxed inside a panel
  return (
    <div className="users-category-page" style={{ padding: 0, background: '#fff', minHeight: 'calc(100vh - 120px)' }}>
      {/* Outer container replicates the screenshot: left green vertical stripe and wide content area (full-bleed, no panel) */}
      <div style={{ display: 'flex', alignItems: 'stretch' }}>
        <div style={{ flex: 1 }}>
          {/* reduce left padding to align with other content (Users table) */}
          <div className="mform-user-title" style={{ padding: '0px' }}>
            <h2 style={{ margin: 0, fontSize: 30, fontWeight: 300 }}>Crear nueva categoría</h2>
          </div>

          {/* use the standard mform-body padding so content sits closer to the left gutter */}
          <div className="mform-body" style={{ padding: '16px 0px' }}>
            {/* Mostrar error si existe - funcionalidad añadida de darwin */}
            {error && (
              <div style={{ padding: '12px', marginBottom: '16px', backgroundColor: '#f8d7da', color: '#721c24', borderRadius: '4px', border: '1px solid #f5c6cb' }}>
                {error}
              </div>
            )}

            {/* Use the same container used by other pages so horizontal gutters align */}
            <div className="moodle-side-marker" aria-hidden="true" />
            <div className="moodle-container" style={{ position: 'relative' }}>
              <form className="moodle-userform" onSubmit={handleCreate}>
                <div className="mform-row">
                  <div className="mform-label" style={{ gap: 2, width: 330 }}>
                    <span>Categoría padre</span>
                    <span className="req-icon" title="Requerido" style={{ marginLeft: 2, display: 'inline-flex', alignItems: 'center', verticalAlign: 'middle' }}>
                      <RequiredIcon />
                    </span>
                  </div>
                  <div className="mform-field">
                    {/* Stack the pill and the select vertically so the select sits below the label */}
                    <div className="field-with-icons" style={{ display: 'flex', flexDirection: 'column', alignItems: 'flex-start', gap: 8 }}>

                      {/* Dynamic pill: shows selected category and can be removed with the × button */}
                      {parent ? (
                        <div style={{ display: 'inline-flex', alignItems: 'center', gap: 8, background: '#d8d8d8', border: '3px solid var(--moodle-green)', height: 34, padding: '0 10px', borderRadius: 6, color: '#000' }}>
                          <button type="button" aria-label="Eliminar categoría padre" onClick={() => { setParent(''); setParentError('Debe suministrar un valor aquí'); }} style={{ background: 'transparent', border: 'none', color: '#000', fontWeight: 400, cursor: 'pointer', padding: 0, margin: 0 }}>×</button>
                          <span style={{ lineHeight: '34px', color: '#000', fontWeight: 700, fontSize: 15 }}>{selectedName || parent}</span>
                        </div>
                      ) : null}

                      {parentError ? (
                        <div style={{ marginBottom: 6 }}>
                          <div style={{ color: '#b00020', fontWeight: 400, fontSize: '13px' }}>- {parentError}</div>
                          <div style={{ color: '#666', marginTop: 6, fontSize: '12px' }}>No hay selección</div>
                        </div>
                      ) : null}
                      <div className="select-wrapper">
                        <select
                          className="mform-select"
                          value={selectValue}
                          onChange={(e) => {
                            const v = String(e.target.value);
                            // set parent to chosen value ('') clears selection
                            setParent(v);
                            // reset the visible select to show the placeholder 'Buscar'
                            setSelectValue('');
                            if (v) setParentError('');
                          }}
                          style={{ width: 218, height: 36.5, color: '#777' }}
                        >
                          {/* Placeholder shown in control but hidden from the dropdown */}
                          <option value="" hidden>Buscar</option>
                          {/* show Superior only if it's not currently selected */}
                          {String(parent) !== '0' && <option value="0" style={{ color: '#000' }}>Superior</option>}
                          {/* list categories except the currently selected one */}
                          {(categories || []).filter((c) => String(c.id) !== String(parent)).map((c) => (
                            <option key={c.id} value={String(c.id)} style={{ color: '#000' }}>{c.name || c.fullname || c.text}</option>
                          ))}
                        </select>
                      </div>
                    </div>
                  </div>
                </div>

              <div className="mform-row">
                <div className="mform-label" style={{ gap: 2, width: 330 }}>
                  <span>Nombre de la categoría</span>
                  <span className="req-icon" title="Requerido" style={{ marginLeft: 2, display: 'inline-flex', alignItems: 'center', verticalAlign: 'middle' }}>
                    <RequiredIcon />
                  </span>
                </div>
                <div className="mform-field">
                  <input className="mform-input" value={name} onChange={(e) => setName(e.target.value)} style={{ width: 308, height: 36.5 }} />
                </div>
              </div>

              <div className="mform-row">
                <div className="mform-label" style={{width:330}}>
                  <span>Número ID de la categoría</span>
                  <span className="help-icon" title="Ayuda" style={{ marginLeft: 8, display: 'inline-flex', alignItems: 'center', verticalAlign: 'middle' }}>
                    <svg width="18" height="18" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><circle cx="12" cy="12" r="10" fill="#17a2b8"/><text x="12" y="14" textAnchor="middle" fontSize="14" fontWeight="700" fill="#fff">?</text></svg>
                  </span>
                </div>
                <div className="mform-field">
                  <input className="mform-input" value={idNumber} onChange={(e) => setIdNumber(e.target.value)} style={{ width: 128, height: 36.5 }} />
                </div>
              </div>

              <div className="mform-row">
                <div className="mform-label" style={{width:330}}><span>Descripción</span></div>
                <div className="mform-field">
                  {/* Atto-like editor toolbar (visual replica) */}
                  <AttoToolbar value={description} onChange={(html) => setDescription(html)} placeholder="Introduce la descripción..." />
                </div>
              </div>

              {/* Actions row moved out so buttons align with inputs */}
              <div className="mform-row">
                <div className="mform-label" style={{width:330}}/>
                <div className="mform-field">
                  <div className="mform-actions">
                    <button type="submit" className="btn-primary" disabled={creating}>
                      {creating ? 'Creando...' : 'Crear categoría'}
                    </button>
                    <button
                      type="button"
                      className="btn-secondary"
                      onClick={() => {
                        // Navigate back to the admin/manage view for courses and categories
                        navigate('/courses/manage');
                      }}
                      disabled={creating}
                    >
                      Cancelar
                    </button>
                  </div>
                </div>
              </div>

              <div className="mform-hint" style={{ marginTop: 18 }}>
                <span className="req-icon" style={{ marginRight: 2, display: 'inline-flex', alignItems: 'center', verticalAlign: 'middle' }}>
                  <RequiredIcon />
                </span>
                Requerido
              </div>
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
