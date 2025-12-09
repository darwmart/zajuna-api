import React, { useEffect, useState, useMemo } from 'react';
import { CircularProgress } from '@mui/material';
import { coursesService } from '../services/coursesService';
import FiltersPanel from './FiltersPanel';

// Render enrolled users for a given course in a Moodle-like table
export default function EnrolledUsersView({ course, onClose }) {
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [users, setUsers] = useState([]);
  const [courseDetail, setCourseDetail] = useState(null);
  const [query, setQuery] = useState('');
  // Letter filters for first name and last name ("Todos" means no filter)
  const [nameLetter, setNameLetter] = useState('Todos');
  const [surnameLetter, setSurnameLetter] = useState('Todos');
  const [filters, setFilters] = useState([]);
  const [appliedFilters, setAppliedFilters] = useState([]);
  const [joinType, setJoinType] = useState('any');
  // Columns collapsed state (persisted to localStorage)
  const [collapsedCols, setCollapsedCols] = useState(() => {
    try {
      const raw = window.localStorage.getItem('enrolled:collapsedCols');
      return raw ? JSON.parse(raw) : {};
    } catch (e) { return {}; }
  });

  // store measured header widths so we can preserve column positions when collapsed
  const thRefs = React.useRef({});
  const [colWidths, setColWidths] = useState({});

  // selection state for checkboxes (selected user ids)
  const [selectedIds, setSelectedIds] = useState(() => new Set());
  const selectAllRef = React.useRef(null);
  const [headerToggled, setHeaderToggled] = useState(false);

  const toggleSelectOne = (id) => {
    setSelectedIds((prev) => {
      const next = new Set(prev);
      if (next.has(id)) next.delete(id);
      else next.add(id);
      // clicking an individual row should not mark the header as the source
      setHeaderToggled(false);
      return next;
    });
  };
  // toggleSelectAll and indeterminate effect are defined after `filtered`
  let toggleSelectAll;

  // Measure header cell widths after first render and on resize so collapsed columns
  // can keep their original size (avoids shifting other columns when hiding content).
  useEffect(() => {
    const measure = () => {
      const next = {};
      Object.keys(thRefs.current).forEach((k) => {
        const el = thRefs.current[k];
        if (el && typeof el.clientWidth === 'number') next[k] = el.clientWidth;
      });
      setColWidths(next);
    };
    // measure on next tick to ensure DOM is ready
    const id = setTimeout(measure, 30);
    window.addEventListener('resize', measure);
    return () => { clearTimeout(id); window.removeEventListener('resize', measure); };
  }, [loading, users.length]);

  const toggleCol = (key) => {
    setCollapsedCols((prev) => {
      const next = { ...(prev || {}), [key]: !prev?.[key] };
      try { window.localStorage.setItem('enrolled:collapsedCols', JSON.stringify(next)); } catch (e) {}
      return next;
    });
  };

  useEffect(() => {
    let mounted = true;
    const load = async () => {
      if (!course) return;
      setLoading(true);
      setError(null);
      try {
        const data = await coursesService.getEnrolledUsers(course.id);
  const arr = Array.isArray(data) ? data : (data.items || data.users || []);
  if (!mounted) return;
  // debug log to inspect payload shape and roles
  console.debug('EnrolledUsersView: raw payload', data);
  console.debug('EnrolledUsersView: users array', arr && arr.length ? arr.slice(0,5) : arr);
  // additional serialized debug: print role-related fields for first users so it's easy to copy/paste
  try {
    const sample = (arr && arr.length) ? arr.slice(0,5).map(u => ({ id: u.id, roles: u.roles, Role: u.Role, role: u.role, role_assignments: u.role_assignments, roleassignments: u.roleassignments })) : [];
    console.debug('EnrolledUsersView: roles fields (serialized)', sample);
    if (arr && arr.length) console.debug('EnrolledUsersView: first user JSON', JSON.stringify(arr[0]));
  } catch (err) {
    console.debug('EnrolledUsersView: failed to serialize sample users', err);
  }
  setUsers(arr);
      // additionally try to fetch course detail (to obtain enrolment methods etc.)
      try {
        const detail = await coursesService.getCourseDetail(course.id);
        if (detail) setCourseDetail(detail);
      } catch (err) {
        // ignore silently, fallback will use enrolments found on users
        console.debug('EnrolledUsersView: failed to load course detail', err);
      }
      } catch (e) {
        console.error('Failed to load enrolled users', e);
        setError('No se pudieron cargar los usuarios matriculados');
        if (mounted) setUsers([]);
      } finally {
        if (mounted) setLoading(false);
      }
    };
    load();
    return () => { mounted = false; };
  }, [course]);

  const letters = useMemo(() => ['Todos', 'A','B','C','D','E','F','G','H','I','J','K','L','M','N','Ñ','O','P','Q','R','S','T','U','V','W','X','Y','Z'], []);

  // Filter types data adapted from the Moodle HTML snippet
  const filterTypes = [
    { name: 'keywords', title: 'Palabra clave', allowCustom: true, allowMultiple: true, values: [] },
    { name: 'status', title: 'Estatus', allowCustom: false, allowMultiple: true, values: [ { value: '0', title: 'Activo' }, { value: '1', title: 'Inactivo' } ] },
    { name: 'roles', title: 'Roles', allowCustom: false, allowMultiple: true, values: [ { value: '-1', title: 'No hay roles' }, { value: '1', title: 'Gestor' }, { value: '2', title: 'Creador de curso' }, { value: '3', title: 'Instructor' }, { value: '4', title: 'Instructor sin permiso de edición' }, { value: '5', title: 'Aprendiz' }, { value: '6', title: 'Invitado' }, { value: '7', title: 'Usuario identificado' }, { value: '8', title: 'Usuario identificado en la página principal del sitio' }, { value: '9', title: 'Soporte técnico' } ] },
    { name: 'enrolments', title: 'Métodos de matriculación', allowCustom: false, allowMultiple: true, values: [ { value: '12', title: 'Matriculación manual' } ] },
    { name: 'accesssince', title: 'Mostrar usuarios que han estado inactivos durante más de', allowCustom: false, allowMultiple: false, values: [
      { value: '1762750800', title: '1 día' },{ value: '1762664400', title: '2 días' },{ value: '1762578000', title: '3 días' },{ value: '1762491600', title: '4 días' },{ value: '1762405200', title: '5 días' },{ value: '1762318800', title: '6 días' },{ value: '1762232400', title: '1 semana' },{ value: '1761627600', title: '2 semanas' },{ value: '1761022800', title: '3 semanas' },{ value: '1760418000', title: '4 semanas' },{ value: '1759813200', title: '5 semanas' },{ value: '1759208400', title: '6 semanas' },{ value: '1758603600', title: '7 semanas' },{ value: '1757998800', title: '8 semanas' },{ value: '1757394000', title: '9 semanas' },{ value: '1756789200', title: '10 semanas' },{ value: '1760158800', title: '1 mes' },{ value: '1757566800', title: '2 meses' },{ value: '1754888400', title: '3 meses' },{ value: '1752210000', title: '4 meses' },{ value: '1749618000', title: '5 meses' },{ value: '1746939600', title: '6 meses' },{ value: '1744347600', title: '7 meses' },{ value: '1741669200', title: '8 meses' },{ value: '1739250000', title: '9 meses' },{ value: '1736571600', title: '10 meses' },{ value: '1733893200', title: '11 meses' },{ value: '1731301200', title: '1 año' },{ value: '-1', title: 'Nunca' }
    ] }
  ];

  // Build dynamic filterTypes based on loaded users (roles, groups, enrolments, keywords suggestions)
  const dynamicFilterTypes = useMemo(() => {
    const rolesSet = new Set();
    const groupsSet = new Set();
    const enrolmentsSet = new Set();
    const keywordsSet = new Set();

    (users || []).forEach((u) => {
      const fullname = `${u.firstname || ''} ${u.lastname || ''}`.trim();
      if (fullname) keywordsSet.add(fullname);
      if (u.email) keywordsSet.add(u.email);
      if (u.username) keywordsSet.add(u.username);

      // collect roles from multiple possible shapes
      const pushRole = (r) => {
        if (!r) return;
        if (typeof r === 'string') rolesSet.add(r);
        else if (typeof r === 'object') {
          if (r.name) rolesSet.add(r.name);
          else if (r.shortname) rolesSet.add(r.shortname);
          else if (r.rolename) rolesSet.add(r.rolename);
        }
      };
      if (Array.isArray(u.roles)) u.roles.forEach(pushRole);
      if (Array.isArray(u.Role)) u.Role.forEach(pushRole);
      if (Array.isArray(u.role)) u.role.forEach(pushRole);

      // groups
      if (Array.isArray(u.groups)) {
        u.groups.forEach((g) => {
          if (!g) return;
          if (typeof g === 'string') groupsSet.add(g);
          else if (g.name) groupsSet.add(g.name);
          else groupsSet.add(String(g.id || JSON.stringify(g)));
        });
      }

      // enrolments (some backends include an 'enrolments' array)
      if (Array.isArray(u.enrolments)) {
        u.enrolments.forEach((e) => {
          if (!e) return;
          if (e.type) enrolmentsSet.add(e.type);
          else if (e.name) enrolmentsSet.add(e.name);
          else if (e.plugin) enrolmentsSet.add(e.plugin);
        });
      }
      // new: some backends return aggregated enrol methods as `enrol_methods`
      if (Array.isArray(u.enrol_methods)) {
        u.enrol_methods.forEach((m) => {
          if (!m) return;
          enrolmentsSet.add(m);
        });
      }
    });

    // Map sets into the {value,title} shape expected by FiltersPanel
    const map = (s) => Array.from(s).map(v => ({ value: String(v), title: String(v) }));
  

    return filterTypes.map((ft) => {
      if (ft.name === 'roles') return { ...ft, values: map(rolesSet), allowMultiple: true };
      if (ft.name === 'enrolments') {
        // Merge enrolment methods found on users with any declared on the course detail
        const courseEnrols = (courseDetail && Array.isArray(courseDetail.enrolments)) ?
          courseDetail.enrolments.map(e => (e.type || e.name || e.plugin || e.id || String(e))).filter(Boolean) : [];
        courseEnrols.forEach(e => enrolmentsSet.add(e));
        return { ...ft, values: map(enrolmentsSet), allowMultiple: true };
      }
      // KEEP keywords without dynamic values so the UI shows a free text input (no suggestions)
      if (ft.name === 'keywords') return { ...ft, values: [], allowMultiple: true };
      if (ft.name === 'groups') return { ...ft, values: map(groupsSet), allowMultiple: true };
      return ft;
    });
  }, [users, filterTypes, courseDetail]);

  // Helper: determine if user is active in this course. Prefer enrol_statuses (array)
  // returned from the backend; if not present, fall back to the aggregated `active` boolean.
  const isActiveInCourse = (u) => {
    if (!u) return false;
    // prefer enrol_statuses: if any status === 0 -> active
    if (Array.isArray(u.enrol_statuses) && u.enrol_statuses.length) {
      try {
        return u.enrol_statuses.some(s => Number(s) === 0);
      } catch (e) {
        // fallback
      }
    }
    // fallback to active boolean from backend
    if (typeof u.active === 'boolean') return !!u.active;
    // last resort: use existing isActive heuristic (suspended/deleted)
    return isActive(u);
  };

  // Extract role labels from user object. Declared as a function here so it's
  // available (hoisted) before any use in memoized matchers.
  function extractRoles(u) {
    if (!u) return '-';
    // Prefer when backend returns an array under 'roles' (lowercase)
    if (Array.isArray(u.roles) && u.roles.length) return u.roles.join(', ');
    // If backend provided a simple string
    if (typeof u.roles === 'string' && u.roles.trim()) return u.roles;
    // support backend returning Roles as array in different capitalizations (Role, role)
    if (Array.isArray(u.Role) && u.Role.length) return u.Role.join(', ');
    if (Array.isArray(u.role) && u.role.length) return u.role.join(', ');
    if (typeof u.roleshort === 'string' && u.roleshort.trim()) return u.roleshort + (u.rolename ? `, ${u.rolename}` : '');
    if (typeof u.rolename === 'string' && u.rolename.trim()) return u.rolename;

    // Try multiple possible fields that may contain role arrays
    const candidates = [u.roles, u.Role, u.role, u.role_assignments, u.roleassignments, u.roleAssignments, u.roles_array];
    let arr = null;
    for (const c of candidates) {
      if (Array.isArray(c) && c.length > 0) { arr = c; break; }
    }

    if (!arr) return '-';

    const labels = arr.map((r) => {
      if (!r) return null;
      if (typeof r === 'string') return r;
      // try common keys
      return r.shortname || r.roleshort || r.name || r.rolename || r.role || r.role_short || r.role_shortname || (r.roleid ? String(r.roleid) : null) || null;
    }).filter(Boolean);

    return labels.length ? labels.join(', ') : '-';
  }

  // Ensure there is one editable filter row when the view opens so the user
  // can immediately select a type. Add only if there are no filters yet.
  useEffect(() => {
    if ((!filters || filters.length === 0) && dynamicFilterTypes && dynamicFilterTypes.length) {
      const id = `f-${Date.now()}`;
      setFilters([{ id, field: '', operator: 'contains', value: '' }]);
    }
    // only run when dynamicFilterTypes becomes available or filters change
  }, [dynamicFilterTypes]);

  const filtered = useMemo(() => {
    // Apply base filters (letter + quick query)
    const base = users.filter((u) => {
      // Apply initial-letter filters for first name and/or surname if set
      if (nameLetter && nameLetter !== 'Todos') {
        const firstName = (u.firstname || '').trim();
        if (!firstName) return false;
        const first = firstName.charAt(0).toUpperCase();
        if (nameLetter === 'Ñ') {
          if (first !== 'Ñ') return false;
        } else if (first !== nameLetter) return false;
      }
      if (surnameLetter && surnameLetter !== 'Todos') {
        const lastName = (u.lastname || '').trim();
        if (!lastName) return false;
        const first = lastName.charAt(0).toUpperCase();
        if (surnameLetter === 'Ñ') {
          if (first !== 'Ñ') return false;
        } else if (first !== surnameLetter) return false;
      }
      if (!query) return true;
      const q = query.toLowerCase();
      const hay = ((u.firstname || '') + ' ' + (u.lastname || '') + ' ' + (u.email || '') + ' ' + (u.username || '')).toLowerCase();
      return hay.includes(q);
    });

    if (!appliedFilters || appliedFilters.length === 0) return base;

    const matchCondition = (user, cond) => {
      if (!cond.field) return true;
      // normalize values to array of strings (lowercased)
      const raw = cond.value ?? '';
      const vals = Array.isArray(raw) ? raw.map(v => String(v).toLowerCase()) : String(raw).split(',').map(s => s.trim()).filter(Boolean).map(s => s.toLowerCase());

      const check = (target) => {
        const t = String(target || '').toLowerCase();
        if (cond.operator === 'contains') return vals.some(v => t.includes(v));
        if (cond.operator === 'equals') return vals.some(v => t === v);
        if (cond.operator === 'startsWith') return vals.some(v => t.startsWith(v));
        return false;
      };

      if (cond.field === 'name') {
        const name = `${user.firstname || ''} ${user.lastname || ''}`;
        return check(name);
      }
      if (cond.field === 'email') {
        return check(user.email || '');
      }
      if (cond.field === 'roles') {
        const r = (extractRoles(user) || '').toString();
        return check(r);
      }
      if (cond.field === 'groups') {
        const groups = Array.isArray(user.groups) ? user.groups.map(g => (g.name || g).toString()).join('||') : '';
        return check(groups);
      }
      if (cond.field === 'status') {
        // Use enrol-status semantics: prefer enrol_statuses (array) -> if any === 0 => active for this course.
        const active = isActiveInCourse(user);
        if (vals.length === 0) return false;
        // Accept numeric codes ('0' active, '1' inactive) or textual 'activo'/'inactivo'
        return vals.some(v => {
          if (!v) return false;
          if (v === '0' || v === 'activo' || v === 'activo') return active;
          if (v === '1' || v === 'inactivo' || v === 'suspendido') return !active;
          if (v === 'true') return active;
          if (v === 'false') return !active;
          // fallback: strict equality
          return v === String(active);
        });
      }
      if (cond.field === 'keywords') {
        const hay = ((user.firstname || '') + ' ' + (user.lastname || '') + ' ' + (user.email || '') + ' ' + (user.username || '')).toLowerCase();
        return vals.some(v => hay.includes(v));
      }
      if (cond.field === 'accesssince') {
        // values are timestamps (as strings) or '-1' for 'Nunca'
        if (!vals || vals.length === 0) return false;
        // consider only the first value (UI provides a single selection)
        const raw = vals[0];
        if (!raw) return false;
        if (raw === '-1') {
          // 'Nunca' => match users with no last access or lastaccess === 0
          const v = user.lastaccess ?? user.lastlogin ?? user.lastseen ?? user.lastcourseaccess;
          return (v === undefined || v === null || v === 0 || v === '0');
        }
        const cutoff = Number(raw);
        if (Number.isNaN(cutoff)) return false;
        let last = user.lastaccess ?? user.lastlogin ?? user.lastseen ?? user.lastcourseaccess;
        if (last === undefined || last === null) return false;
        if (typeof last === 'string') {
          // try numeric string first
          const asNum = Number(last);
          if (!Number.isNaN(asNum)) last = asNum;
          else {
            // try parsing an ISO/human date string
            const parsed = Date.parse(last);
            if (!Number.isNaN(parsed)) {
              // Date.parse returns ms since epoch; convert to seconds to match cutoff
              last = Math.floor(parsed / 1000);
            } else {
              return false;
            }
          }
        }
  // treat 0 as 'Nunca' but include those users when applying a numeric cutoff
  // (i.e. someone who never accessed should be considered 'inactive' for any positive cutoff)
  if (last === 0) return true;
  // match users whose last access is older (<=) than the cutoff
  return Number(last) <= cutoff;
      }
      return true;
    };

    // Determine effective join mode.
    // Prefer the global `joinType` (set by header) so the header controls the
    // logical combination of the applied filters. Fall back to per-row join
    // when global joinType is not provided.
    const effectiveJoin = (() => {
      if (!appliedFilters || appliedFilters.length === 0) return 'any';
      if (joinType) return joinType; // 'all'|'any'|'none'
      // legacy: derive from per-row join values if global not present
      const first = appliedFilters[0].join ?? '1';
      const allSame = appliedFilters.every(f => (f.join ?? '1') === first);
      const map = (v) => (v === '2' ? 'all' : (v === '0' ? 'none' : 'any'));
      if (allSame) return map(first);
      return 'any';
    })();

    return base.filter((u) => {
      const results = appliedFilters.map(f => matchCondition(u, f));
      if (effectiveJoin === 'all') return results.every(Boolean);
      if (effectiveJoin === 'none') return results.every(r => !r);
      // default 'any'
      return results.some(Boolean);
    });
  }, [users, query, nameLetter, surnameLetter, appliedFilters, joinType]);

  // Now that `filtered` is declared we can implement select-all logic and
  // update the indeterminate state of the header checkbox.
  toggleSelectAll = () => {
    const ids = (filtered || []).map(u => u.id);
    const allSelected = ids.length > 0 && ids.every(id => selectedIds.has(id));
    if (allSelected) {
      setSelectedIds(new Set());
      setHeaderToggled(true);
    } else {
      setSelectedIds(new Set(ids));
      setHeaderToggled(true);
    }
  };

  useEffect(() => {
    const el = selectAllRef.current;
    if (!el) return;
    // Only show indeterminate/checked visual state when the header checkbox
    // was the source of the selection (headerToggled). If selection was done
    // by clicking individual rows, we don't want the header checkbox to be
    // visually highlighted.
    if (!headerToggled) {
      el.indeterminate = false;
      return;
    }

    const ids = (filtered || []).map(u => u.id);
    const allSelected = ids.length > 0 && ids.every(id => selectedIds.has(id));
    const someSelected = ids.some(id => selectedIds.has(id));
    el.indeterminate = !allSelected && someSelected;
  }, [filtered, selectedIds]);

  const formatLastAccess = (u) => {
    if (!u) return '-';
    // backend may provide lastaccess or lastlogin as timestamp or human string
    const v = u.lastaccess ?? u.lastlogin ?? u.lastseen ?? u.lastcourseaccess;
    if (!v && v !== 0) return 'Nunca';
    // treat 0 as 'Nunca' as well
    if (v === 0) return 'Nunca';
    // if it's numeric timestamp in seconds
    if (typeof v === 'number') {
      try {
        const d = new Date(v * 1000);
        return d.toLocaleString();
      } catch (e) { return String(v); }
    }
    return String(v);
  };

  const isActive = (u) => {
    if (!u) return false;
    if (u.suspended || u.deleted) return false;
    return true;
  };

  const [formAction, setFormAction] = useState('');

  const handleFormAction = (e) => {
    const val = e.target.value || '';
    if (!val) return setFormAction('');
    // small placeholder action: log selected bulk action.
    // Later: implement navigation, modal or API call depending on value.
    console.debug('EnrolledUsersView: bulk action selected', val);
    // reset to placeholder option
    setFormAction('');
  };


  // Filter handlers for FiltersPanel
  const addCondition = () => {
    const id = String(Date.now()) + Math.floor(Math.random() * 1000);
    setFilters((s) => [...s, { id, field: '', operator: 'contains', value: '' }]);
  };

  const updateCondition = (id, changed) => {
    setFilters((s) => s.map(f => (f.id === id ? changed : f)));
  };

  const removeCondition = (id) => {
    setFilters((s) => s.filter(f => f.id !== id));
  };

  const applyFiltersHandler = () => {
    // copy current filters into appliedFilters (we'll use these in the memo)
    setAppliedFilters(filters.filter(f => f.field));
  };

  const clearFiltersHandler = () => {
    setFilters([]);
    setAppliedFilters([]);
  };

  const onJoinChange = (type) => {
    setJoinType(type);
  };

  return (
    <div style={{ width: '100%', boxSizing: 'border-box', display: 'flex', flexDirection: 'column' }} className="users-enrolled-panel">
  <div style={{ display: 'flex', alignItems: 'center', gap: 40, marginBottom: 30, paddingLeft: 0 }}>
            <div style={{ display: 'flex', alignItems: 'center', gap: 6, color: '#111', fontWeight: 700, fontSize: 20 }}>
              <span style={{ color: 'inherit' }}>Usuarios matriculados</span>
              <svg width="10" height="7" viewBox="0 0 14 10" xmlns="http://www.w3.org/2000/svg" aria-hidden="true" focusable="false" style={{ display: 'inline-block', verticalAlign: 'middle', alignSelf: 'center' }}>
                <path d="M1 1l6 6 6-6" fill="none" stroke="currentColor" strokeWidth="1.2" strokeLinecap="round" strokeLinejoin="round" />
              </svg>
            </div>
            <button
              type="button"
              className="btn-moodle btn-primary"
              style={{ background: 'var(--moodle-green)', color: '#fff', border: 'none', padding: '8px 14px', borderRadius: 12, fontWeight: 700, marginTop: 8, marginBottom: 12 }}
            >
              Matricular usuarios
            </button>
            
          </div>

  <FiltersPanel
        filters={filters}
        onAdd={addCondition}
        onChange={(id, changed) => updateCondition(id, changed)}
        onRemove={removeCondition}
        onApply={applyFiltersHandler}
        onClear={clearFiltersHandler}
        joinType={joinType}
        onJoinChange={onJoinChange}
    filterTypes={dynamicFilterTypes}
      />
  {/* Participants count line: shows 0, 1 or N participantes encontrados */}
  <div
    className="participants-count"
    role="status"
    aria-live="polite"
    style={{ marginTop: 0, marginBottom: 12, color: '#333', paddingLeft: 0, fontSize: 14, fontWeight: 400, lineHeight: '1.2' }}
  >
    {filtered.length === 1 ? '1 participante encontrado' : `${filtered.length} participantes encontrados`}
  </div>
  {/* Initial-letter filters for First name and Surname (placed below participants count) */}
  <div className="initial-letter" style={{ paddingLeft: 0 }}>
    {(nameLetter !== 'Todos' || surnameLetter !== 'Todos' || Object.values(collapsedCols || {}).some(Boolean)) && (
      <div className="reset-table-prefs-wrapper">
        <button
          type="button"
          className="reset-table-prefs"
          onClick={() => {
            setNameLetter('Todos');
            setSurnameLetter('Todos');
            setCollapsedCols({});
            try { window.localStorage.removeItem('enrolled:collapsedCols'); } catch (e) {}
          }}
        >
          Restablecer preferencias de tabla
        </button>
      </div>
    )}
    <div className="initial-row">
      <div className="initial-label">Nombre</div>
      <div className="initial-buttons">
        <button
          type="button"
          className={`initial-btn todos ${nameLetter === 'Todos' ? 'active' : ''}`}
          onClick={() => setNameLetter('Todos')}
          aria-pressed={nameLetter === 'Todos'}
        >
          Todos
        </button>
        {letters.slice(1).map((l) => (
          <button
            key={`name-${l}`}
            type="button"
            className={`initial-btn letter ${nameLetter === l ? 'active' : ''}`}
            onClick={() => setNameLetter(prev => (prev === l ? 'Todos' : l))}
            aria-pressed={nameLetter === l}
          >
            {l}
          </button>
        ))}
      </div>
    </div>

    <div className="initial-row">
      <div className="initial-label">Apellido(s)</div>
      <div className="initial-buttons">
        <button
          type="button"
          className={`initial-btn todos ${surnameLetter === 'Todos' ? 'active' : ''}`}
          onClick={() => setSurnameLetter('Todos')}
          aria-pressed={surnameLetter === 'Todos'}
        >
          Todos
        </button>
        {letters.slice(1).map((l) => (
          <button
            key={`surname-${l}`}
            type="button"
            className={`initial-btn letter ${surnameLetter === l ? 'active' : ''}`}
            onClick={() => setSurnameLetter(prev => (prev === l ? 'Todos' : l))}
            aria-pressed={surnameLetter === l}
          >
            {l}
          </button>
        ))}
      </div>
    </div>
  </div>
  <div className="enrolled-table-container" style={{ background: '#fff', borderRadius: 6, paddingLeft: 0 }}>
        {loading ? (
          <div style={{ padding: 18, textAlign: 'center' }}><CircularProgress /></div>
        ) : error ? (
          <div style={{ color: '#a00', padding: 8 }}>{error}</div>
        ) : (
          <>
            {filtered && filtered.length === 0 ? (
              <div style={{ padding: 18, paddingLeft: 0, textAlign: 'left' }}>
                <h2 className="h2">Nada que mostrar</h2>
              </div>
            ) : (
              <>
                <table className="table-moodle" style={{ width: '100%', borderCollapse: 'collapse' }}>
                  <thead>
                    <tr>
                      <th style={{ width: 40, paddingBottom: 0 }}>
                        <div style={{ display: 'flex', flexDirection: 'column', alignItems: 'flex-start' }}>
                          <div />
                          <div className="th-minimize">
                            <input
                              ref={selectAllRef}
                              type="checkbox"
                              title="Seleccionar todos"
                              onChange={toggleSelectAll}
                              checked={headerToggled && (filtered || []).length > 0 && (filtered || []).every(u => selectedIds.has(u.id))}
                            />
                          </div>
                        </div>
                      </th>
                      <th data-col="name" ref={(el) => { thRefs.current.name = el; }} style={collapsedCols.name && colWidths.name ? { width: `${colWidths.name}px`, minWidth: `${colWidths.name}px` } : undefined}>
                        <div style={{ display: 'flex', flexDirection: 'column', alignItems: 'flex-start' }}>
                          <div>{collapsedCols.name ? '' : 'Nombre / Apellido(s)'}</div>
                          <div className="th-minimize">
                            <button type="button" className="col-toggle" title={collapsedCols.name ? 'Mostrar Nombre completo' : 'Ocultar Nombre completo'} onClick={() => toggleCol('name')}>{collapsedCols.name ? '+' : '−'}</button>
                          </div>
                        </div>
                      </th>
                      <th data-col="email" ref={(el) => { thRefs.current.email = el; }} style={collapsedCols.email && colWidths.email ? { width: `${colWidths.email}px`, minWidth: `${colWidths.email}px` } : undefined}>
                        <div style={{ display: 'flex', flexDirection: 'column', alignItems: 'flex-start' }}>
                          <div>{collapsedCols.email ? '' : 'Dirección de correo'}</div>
                          <div className="th-minimize">
                            <button type="button" className="col-toggle" title={collapsedCols.email ? 'Mostrar Dirección de correo' : 'Ocultar Dirección de correo'} onClick={() => toggleCol('email')}>{collapsedCols.email ? '+' : '−'}</button>
                          </div>
                        </div>
                      </th>
                      <th data-col="roles" ref={(el) => { thRefs.current.roles = el; }} style={collapsedCols.roles && colWidths.roles ? { width: `${colWidths.roles}px`, minWidth: `${colWidths.roles}px` } : undefined}>
                        <div style={{ display: 'flex', flexDirection: 'column', alignItems: 'flex-start' }}>
                          <div style={{ color: '#000' }}>{collapsedCols.roles ? '' : 'Roles'}</div>
                          <div className="th-minimize">
                            <button type="button" className="col-toggle" title={collapsedCols.roles ? 'Mostrar Roles' : 'Ocultar Roles'} onClick={() => toggleCol('roles')}>{collapsedCols.roles ? '+' : '−'}</button>
                          </div>
                        </div>
                      </th>
                      <th data-col="groups" ref={(el) => { thRefs.current.groups = el; }} style={collapsedCols.groups && colWidths.groups ? { width: `${colWidths.groups}px`, minWidth: `${colWidths.groups}px` } : undefined}>
                        <div style={{ display: 'flex', flexDirection: 'column', alignItems: 'flex-start' }}>
                          <div style={{ color: '#000' }}>{collapsedCols.groups ? '' : 'Grupos'}</div>
                          <div className="th-minimize">
                            <button type="button" className="col-toggle" title={collapsedCols.groups ? 'Mostrar Grupos' : 'Ocultar Grupos'} onClick={() => toggleCol('groups')}>{collapsedCols.groups ? '+' : '−'}</button>
                          </div>
                        </div>
                      </th>
                      <th data-col="lastaccess" ref={(el) => { thRefs.current.lastaccess = el; }} style={collapsedCols.lastaccess && colWidths.lastaccess ? { width: `${colWidths.lastaccess}px`, minWidth: `${colWidths.lastaccess}px` } : undefined}>
                        <div style={{ display: 'flex', flexDirection: 'column', alignItems: 'flex-start' }}>
                          <div>{collapsedCols.lastaccess ? '' : 'Último acceso al curso'}</div>
                          <div className="th-minimize">
                            <button type="button" className="col-toggle" title={collapsedCols.lastaccess ? 'Mostrar Último acceso al curso' : 'Ocultar Último acceso al curso'} onClick={() => toggleCol('lastaccess')}>{collapsedCols.lastaccess ? '+' : '−'}</button>
                          </div>
                        </div>
                      </th>
                      <th data-col="status" ref={(el) => { thRefs.current.status = el; }} style={collapsedCols.status && colWidths.status ? { width: `${colWidths.status}px`, minWidth: `${colWidths.status}px` } : undefined}>
                        <div style={{ display: 'flex', flexDirection: 'column', alignItems: 'flex-start' }}>
                          <div style={{ color: '#000' }}>{collapsedCols.status ? '' : 'Estatus'}</div>
                          <div className="th-minimize">
                            <button type="button" className="col-toggle" title={collapsedCols.status ? 'Mostrar Estatus' : 'Ocultar Estatus'} onClick={() => toggleCol('status')}>{collapsedCols.status ? '+' : '−'}</button>
                          </div>
                        </div>
                      </th>
                    </tr>
                  </thead>
                  <tbody>
                    {filtered.map((u) => (
                      <tr key={u.id}>
                        <td>
                          <input type="checkbox" checked={selectedIds.has(u.id)} onChange={() => toggleSelectOne(u.id)} />
                        </td>
                        <td className="user-name" data-col="name" style={collapsedCols.name ? { width: colWidths.name ? `${colWidths.name}px` : undefined, minWidth: colWidths.name ? `${colWidths.name}px` : undefined } : {}}>
                          <div className="col-content" style={collapsedCols.name ? { visibility: 'hidden' } : {}}>
                            <div style={{ display: 'flex', alignItems: 'center', gap: 10 }}>
                              <span className="userinitials size-35">{((u.firstname||'')[0]||'')+((u.lastname||'')[0]||'')}</span>
                              <div>
                                <div style={{ color: 'var(--moodle-green)', fontWeight: 600 }}>{u.firstname} {u.lastname}</div>
                              </div>
                            </div>
                          </div>
                        </td>
                        <td data-col="email" style={collapsedCols.email ? { width: colWidths.email ? `${colWidths.email}px` : undefined, minWidth: colWidths.email ? `${colWidths.email}px` : undefined } : {}}>
                          <div className="col-content" style={collapsedCols.email ? { visibility: 'hidden' } : {}}>{u.email || '-'}</div>
                        </td>
                        <td data-col="roles" style={collapsedCols.roles ? { width: colWidths.roles ? `${colWidths.roles}px` : undefined, minWidth: colWidths.roles ? `${colWidths.roles}px` : undefined } : {}}>
                          <div className="col-content" style={collapsedCols.roles ? { visibility: 'hidden' } : {}}>{extractRoles(u)}</div>
                        </td>
                        <td data-col="groups" style={collapsedCols.groups ? { width: colWidths.groups ? `${colWidths.groups}px` : undefined, minWidth: colWidths.groups ? `${colWidths.groups}px` : undefined } : {}}>
                          <div className="col-content" style={collapsedCols.groups ? { visibility: 'hidden' } : {}}>{Array.isArray(u.groups) && u.groups.length ? u.groups.map(g => g.name || g).join(', ') : 'No hay grupos'}</div>
                        </td>
                        <td data-col="lastaccess" style={collapsedCols.lastaccess ? { width: colWidths.lastaccess ? `${colWidths.lastaccess}px` : undefined, minWidth: colWidths.lastaccess ? `${colWidths.lastaccess}px` : undefined } : {}}>
                          <div className="col-content" style={collapsedCols.lastaccess ? { visibility: 'hidden' } : {}}>{formatLastAccess(u)}</div>
                        </td>
                        <td data-col="status" style={collapsedCols.status ? { width: colWidths.status ? `${colWidths.status}px` : undefined, minWidth: colWidths.status ? `${colWidths.status}px` : undefined } : {}}>
                          <div className="col-content" style={collapsedCols.status ? { visibility: 'hidden' } : {}}>
                            <div style={{ display: 'inline-flex', alignItems: 'center', gap: 8 }}>
                              {isActiveInCourse(u) ? (
                                <span className="badge badge-success">Activo</span>
                              ) : (
                                <span className="badge badge-warning">Suspendido</span>
                              )}
                              <span className="status-icons" aria-hidden="true">
                                <i className="icon fa fa-info-circle fa-fw" title="Matriculación manual" role="img" aria-label="Matriculación manual"></i>
                                <i className="icon fa fa-cog fa-fw" title="Editar matrícula" role="img" aria-label="Editar matrícula"></i>
                                <i className="icon fa fa-trash fa-fw" title="Dar de baja" role="img" aria-label="Dar de baja"></i>
                              </span>
                            </div>
                          </div>
                        </td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              </>
            )}
          </>
        )}
      </div>
      {/* Bulk action control placed at panel level and aligned to page start */}
      <div style={{ paddingLeft: 0 }}>
        <div className="userlist buttons" style={{ marginTop: 12, padding: '8px 0 18px 0' }}>
          <label className="form-inline" style={{ fontSize: 14, color: '#333', paddingLeft: 0 }}>
            Con los usuarios seleccionados...
            <select
              id="formactionid"
              className="select custom-select ml-2"
              name="formaction"
              value={formAction}
              onChange={(e) => { setFormAction(e.target.value); handleFormAction(e); }}
              disabled={(selectedIds && selectedIds.size === 0)}
              style={{ display: 'inline-block', verticalAlign: 'middle', maxWidth: 480, marginLeft: '15px' }}
            >
              <option value="" key="opt-default">Elegir...</option>
              <option value="#messageselect">Enviar un mensaje</option>
              <option value="#addgroupnote">Agregar una nueva anotación</option>
              <optgroup label="Descargar datos de tabla como">
                <option value="bulkchange.php?operation=download_participants&dataformat=csv">Valores separados por comas (.csv)</option>
                <option value="bulkchange.php?operation=download_participants&dataformat=excel">Microsoft Excel (.xlsx)</option>
                <option value="bulkchange.php?operation=download_participants&dataformat=html">Tabla HTML</option>
                <option value="bulkchange.php?operation=download_participants&dataformat=json">Javascript Object Notation (.json)</option>
                <option value="bulkchange.php?operation=download_participants&dataformat=ods">OpenDocument (.ods)</option>
                <option value="bulkchange.php?operation=download_participants&dataformat=pdf">Portable Document Format (.pdf)</option>
              </optgroup>
              <optgroup label="Matriculación manual">
                <option value="bulkchange.php?plugin=manual&operation=editselectedusers">Editar las matrículas de usuario seleccionadas</option>
                <option value="bulkchange.php?plugin=manual&operation=deleteselectedusers">Eliminar las matrículas de usuario seleccionadas</option>
              </optgroup>
              <optgroup label="Auto-matriculación">
                <option value="bulkchange.php?plugin=self&operation=editselectedusers">Editar las matrículas de usuario seleccionadas</option>
                <option value="bulkchange.php?plugin=self&operation=deleteselectedusers">Eliminar las matrículas de usuario seleccionadas</option>
              </optgroup>
            </select>
          </label>
        </div>
        {/* New action button placed one line under the bulk-action selector and aligned right */}
        <div style={{ marginTop: 6, display: 'flex', justifyContent: 'flex-end', alignItems: 'center' }}>
          <button
            type="button"
            className="btn-moodle btn-primary"
            onClick={() => {
              // placeholder behaviour: open the same flow as the top-level matricular
              console.debug('Bulk: Matricular usuarios clicked', Array.from(selectedIds));
            }}
            disabled={!(selectedIds && selectedIds.size > 0)}
            style={{ background: 'var(--moodle-green)', color: '#fff', border: 'none', padding: '8px 14px', borderRadius: 12, fontWeight: 700 }}
          >
            Matricular usuarios
          </button>
        </div>
      </div>
    </div>
  );
}
