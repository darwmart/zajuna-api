import React, { useState, useRef, useEffect } from 'react';

// Small, self-contained FiltersPanel component for EnrolledUsersView.
// Props:
// - filters: array of filter objects { id, field, operator, value }
// - onAdd: () => void
// - onChange: (id, changed) => void
// - onRemove: (id) => void
// - onApply: () => void
// - onClear: () => void
// - joinType: 'any'|'all', onJoinChange
// - extraActions: node to render at the right (e.g. Matricular usuarios button)

export default function FiltersPanel({ filters, onAdd, onChange, onRemove, onApply, onClear, joinType, onJoinChange, filterTypes = [] }) {
  // helper to find filterType metadata
  const findType = (name) => filterTypes.find(t => t.name === name) || null;

  // When a row's type is set to 'accesssince' and no value has been provided yet,
  // initialize it to the first available option (1 día) so the UI shows a sensible default.
  useEffect(() => {
    if (!filters || !filters.length) return;
    filters.forEach((f) => {
      // Only auto-initialize accesssince when value is truly undefined and the
      // user hasn't explicitly cleared it (_cleared flag). This allows the user
      // to remove the default tag and keep the field empty without it being
      // immediately reinitialized.
      if (f && f.field === 'accesssince' && typeof f.value === 'undefined' && !f._cleared) {
        const meta = findType('accesssince');
        if (meta && Array.isArray(meta.values) && meta.values.length) {
          try {
            onChange && onChange(f.id, { ...f, value: meta.values[0].value });
          } catch (e) {
            // ignore any errors to avoid breaking render
          }
        }
      }
    });
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [filters, filterTypes]);

  // Small combobox component: controlled input with dropdown suggestions.
  function Combobox({ id, value, onChange, options = [], placeholder }) {
    const [query, setQuery] = useState('');
    const [open, setOpen] = useState(false);
    const [highlight, setHighlight] = useState(0);
    const wrapperRef = useRef(null);

    // Keep query in sync with external value (show title when value set)
    useEffect(() => {
      const selected = options.find(o => o.value === value);
      setQuery(selected ? selected.title : '');
    }, [value, options]);

    // Close on outside click
    useEffect(() => {
      function onDocClick(e) {
        if (wrapperRef.current && !wrapperRef.current.contains(e.target)) {
          setOpen(false);
        }
      }
      document.addEventListener('click', onDocClick);
      return () => document.removeEventListener('click', onDocClick);
    }, []);

    const filtered = options.filter(o => {
      if (!query) return true;
      return (o.title || '').toLowerCase().includes(query.toLowerCase()) || (o.value || '').toLowerCase().includes(query.toLowerCase());
    });

    function selectOption(opt) {
      onChange(opt.value);
      setQuery(opt.title);
      setOpen(false);
    }

    function onInputKeyDown(e) {
      if (e.key === 'ArrowDown') {
        e.preventDefault();
        setOpen(true);
        setHighlight(h => Math.min(h + 1, Math.max(0, filtered.length - 1)));
      } else if (e.key === 'ArrowUp') {
        e.preventDefault();
        setOpen(true);
        setHighlight(h => Math.max(h - 1, 0));
      } else if (e.key === 'Enter') {
        e.preventDefault();
        if (open && filtered[highlight]) {
          selectOption(filtered[highlight]);
        }
      } else if (e.key === 'Escape') {
        setOpen(false);
      }
    }

    return (
      <div className="combobox" ref={wrapperRef} style={{ position: 'relative' }}>
        <input
          id={id}
          className="moodle-input combobox-input"
          value={query}
          placeholder={placeholder}
          onChange={(e) => { setQuery(e.target.value); setOpen(true); setHighlight(0); }}
          onKeyDown={onInputKeyDown}
          onFocus={() => setOpen(true)}
          autoComplete="off"
          aria-expanded={open}
          aria-haspopup="listbox"
        />
        {open && (
          <ul role="listbox" className="combobox-list" style={{ listStyle: 'none', margin: 0, padding: 0 }}>
            {filtered && filtered.length ? (
              filtered.map((opt, idx) => (
                <li
                  role="option"
                  aria-selected={value === opt.value}
                  key={opt.value}
                  className={`combobox-item ${idx === highlight ? 'highlight' : ''}`}
                  onMouseEnter={() => setHighlight(idx)}
                  onMouseDown={(e) => { e.preventDefault(); selectOption(opt); }}
                  style={{ padding: '8px 12px', cursor: 'pointer' }}
                >
                  {opt.title}
                </li>
              ))
            ) : (
              <li className="combobox-noresults" style={{ padding: '8px 12px', color: '#666' }}>No hay sugerencias</li>
            )}
          </ul>
        )}
      </div>
    );
  }

  // Multi-select combobox: allows typing to filter, selecting items adds them to the array
  function MultiCombobox({ id, value = [], onChange, options = [], placeholder }) {
    const [query, setQuery] = useState('');
    const [open, setOpen] = useState(false);
    const [highlight, setHighlight] = useState(0);
    const wrapperRef = useRef(null);

    useEffect(() => {
      // clear query when value changes externally
      setQuery('');
    }, [value]);

    useEffect(() => {
      function onDocClick(e) {
        if (wrapperRef.current && !wrapperRef.current.contains(e.target)) {
          setOpen(false);
        }
      }
      document.addEventListener('click', onDocClick);
      return () => document.removeEventListener('click', onDocClick);
    }, []);

    const filtered = options.filter(o => {
      // exclude already selected
      if (Array.isArray(value) && value.indexOf(o.value) !== -1) return false;
      if (!query) return true;
      return (o.title || '').toLowerCase().includes(query.toLowerCase()) || (o.value || '').toLowerCase().includes(query.toLowerCase());
    });

    function addOption(opt) {
      const next = Array.isArray(value) ? [...value, opt.value] : [opt.value];
      onChange(next);
      setQuery('');
      setOpen(false);
    }

    function removeOption(val) {
      const next = (Array.isArray(value) ? value : []).filter(x => x !== val);
      onChange(next);
    }

    function onInputKeyDown(e) {
      if (e.key === 'ArrowDown') {
        e.preventDefault();
        setOpen(true);
        setHighlight(h => Math.min(h + 1, Math.max(0, filtered.length - 1)));
      } else if (e.key === 'ArrowUp') {
        e.preventDefault();
        setOpen(true);
        setHighlight(h => Math.max(h - 1, 0));
      } else if (e.key === 'Enter') {
        e.preventDefault();
        if (open && filtered[highlight]) {
          addOption(filtered[highlight]);
        }
      } else if (e.key === 'Escape') {
        setOpen(false);
      }
    }

    return (
      // Multi combobox: keep input fixed width like other controls and render
      // selected chips to the right (absolute) so they don't change input width.
      <div className="combobox multi" ref={wrapperRef} style={{ position: 'relative', display: 'inline-block' }}>
        <input
          id={id}
          className="moodle-input combobox-input"
          value={query}
          placeholder={placeholder}
          onChange={(e) => { setQuery(e.target.value); setOpen(true); setHighlight(0); }}
          onKeyDown={onInputKeyDown}
          onFocus={() => setOpen(true)}
          autoComplete="off"
          aria-expanded={open}
          aria-haspopup="listbox"
          style={{ width: 210, boxSizing: 'border-box' }}
        />

  <div className="tag-list" style={{ position: 'absolute', left: 'calc(100% + 8px)', top: '50%', transform: 'translateY(-50%)', display: 'flex', gap: 8, alignItems: 'center', flexWrap: 'nowrap', overflowX: 'auto', zIndex: 10 }}>
          {(Array.isArray(value) ? value : []).map((val) => {
            const title = (options.find(o => o.value === val) || {}).title || val;
            // Use a non-breaking space for display so labels like "1 día" don't break
            const displayTitle = (String(title || '')).replace(/\s+/g, '\u00A0');
            return (
              <div key={val} className="chip tag-chip" style={{ display: 'inline-flex', alignItems: 'center', gap: 6, padding: '6px 8px', borderRadius: 8 }}>
                <span style={{ fontSize: 13 }}>{displayTitle}</span>
                <button type="button" className="chip-close" onClick={() => removeOption(val)} aria-label={`Quitar ${title}`} style={{ border: 'none', background: 'transparent', cursor: 'pointer' }}>×</button>
              </div>
            );
          })}
        </div>

        {open && (
          <ul role="listbox" className="combobox-list" style={{ listStyle: 'none', margin: 0, padding: 0 }}>
            {filtered && filtered.length ? (
              filtered.map((opt, idx) => (
                <li
                  role="option"
                  aria-selected={false}
                  key={opt.value}
                  className={`combobox-item ${idx === highlight ? 'highlight' : ''}`}
                  onMouseEnter={() => setHighlight(idx)}
                  onMouseDown={(e) => { e.preventDefault(); addOption(opt); }}
                  style={{ padding: '8px 12px', cursor: 'pointer' }}
                >
                  {opt.title}
                </li>
              ))
            ) : (
              <li className="combobox-noresults" style={{ padding: '8px 12px', color: '#666' }}>No hay sugerencias</li>
            )}
          </ul>
        )}
      </div>
    );
  }

  // Tag input for free-text multi values (used by 'keywords').
  // Allows typing and pressing Enter / comma / blur to create tags. Displays chips with
  // a green outline, light grey background and black text, with a close '×' button.
  function TagInput({ id, value = [], onChange, placeholder }) {
    const [query, setQuery] = useState('');
    const wrapperRef = useRef(null);

    useEffect(() => {
      // clear query when value changes externally
      setQuery('');
    }, [value]);

    useEffect(() => {
      function onDocClick(e) {
        if (wrapperRef.current && !wrapperRef.current.contains(e.target)) {
          // on blur create tag from remaining query
          if (query && query.trim()) {
            addTag(query.trim());
          }
        }
      }
      document.addEventListener('click', onDocClick);
      return () => document.removeEventListener('click', onDocClick);
    }, [query]);

    function addTag(t) {
      if (!t) return;
      const trimmed = String(t).trim();
      if (!trimmed) return;
      const next = Array.isArray(value) ? [...value] : [];
      if (next.indexOf(trimmed) !== -1) return; // avoid duplicates
      next.push(trimmed);
      onChange(next);
      setQuery('');
    }

    function removeTag(val) {
      const next = (Array.isArray(value) ? value : []).filter(x => x !== val);
      onChange(next);
    }

    function onKeyDown(e) {
      if (e.key === 'Enter' || e.key === ',' ) {
        e.preventDefault();
        if (query && query.trim()) addTag(query.trim());
      } else if (e.key === 'Escape') {
        setQuery('');
      }
    }

    return (
      // Keep the input a fixed width to match other value controls and position
      // the tags absolutely to the right so they don't change the input width.
      <div className="combobox tag-input" ref={wrapperRef} style={{ position: 'relative', display: 'inline-block' }}>
        <input
          id={id}
          className="moodle-input combobox-input"
          value={query}
          placeholder={placeholder}
          onChange={(e) => { setQuery(e.target.value); }}
          onKeyDown={onKeyDown}
          onBlur={() => { if (query && query.trim()) addTag(query.trim()); }}
          autoComplete="off"
          style={{ width: 210, boxSizing: 'border-box' }}
        />

  <div className="tag-list" style={{ position: 'absolute', left: 'calc(100% + 8px)', top: '50%', transform: 'translateY(-50%)', display: 'flex', gap: 8, alignItems: 'center', flexWrap: 'nowrap', overflowX: 'auto', zIndex: 10 }}>
          {(Array.isArray(value) ? value : []).map((val) => {
            const displayVal = (String(val || '')).replace(/\s+/g, '\u00A0');
            return (
              <div key={val} className="chip tag-chip" style={{ display: 'inline-flex', alignItems: 'center', gap: 8, padding: '6px 10px', borderRadius: 8, color: '#000', fontWeight: 600 }}>
                <span style={{ fontSize: 13 }}>{displayVal}</span>
                <button type="button" className="chip-close" onClick={() => removeTag(val)} aria-label={`Quitar ${val}`} style={{ border: 'none', background: 'transparent', cursor: 'pointer', fontSize: 14 }}>×</button>
              </div>
            );
          })}
        </div>
      </div>
    );
  }
  
  // compute which filter types are already used (selected) in other rows
  const usedTypes = (filters || []).map(f => f.field).filter(Boolean);
  // Map the external joinType prop ('all'|'any'|'none') to the select values used in the UI
  const headerJoinValue = (joinType === 'all') ? '2' : (joinType === 'any') ? '1' : (joinType === 'none') ? '0' : '2';

  function handleHeaderJoinChange(e) {
    const v = String(e.target.value);
    const mapped = v === '2' ? 'all' : (v === '1' ? 'any' : 'none');
    onJoinChange && onJoinChange(mapped);
  }

  // Show the header only when the user has clicked "Añadir condición". It will
  // be hidden again if all filters are removed.
  const [headerVisible, setHeaderVisible] = useState(false);

  // Keep a local header select value so the connector updates immediately
  // (stores the select raw value: '2'|'1'|'0'). Initialize from the prop-derived value.
  const [headerJoinLocal, setHeaderJoinLocal] = useState(headerJoinValue);
  useEffect(() => { setHeaderJoinLocal(headerJoinValue); }, [headerJoinValue]);

  function handleAddClick(e) {
    // mark header visible before adding the row so the UI updates immediately
    // and ensure the header select defaults to 'Todos' when it appears.
    setHeaderJoinLocal('2');
    onJoinChange && onJoinChange('all');
    setHeaderVisible(true);
    if (onAdd) onAdd(e);
  }

  function handleHeaderSelectChange(e) {
    const v = String(e.target.value);
    setHeaderJoinLocal(v);
    const mapped = v === '2' ? 'all' : (v === '1' ? 'any' : 'none');
    onJoinChange && onJoinChange(mapped);
  }

  // If header is visible but there is only one (or zero) filter row, hide
  // the header again. The header should only be shown when there are at
  // least two rows to combine.
  useEffect(() => {
    if (headerVisible && (!filters || filters.length <= 1)) {
      setHeaderVisible(false);
    }
  }, [filters, headerVisible]);
  

  return (
    <div id="core-filter" className="filter-group my-2 p-2 bg-light border-radius border filters-panel" data-filterverb={joinType === 'all' ? 2 : 1} style={{ position: 'relative' }}>
      {/* Header removed: only functional filter rows are shown below */}

      {/* Header: show global join selector only when at least one condition exists */}
      {headerVisible ? (
        <div className="filters-header" style={{ display: 'flex', alignItems: 'center', gap: 0, padding: '0px 0px', marginBottom: 8 }}>
          <div style={{ minWidth: 70, fontSize: 16, color: '#333' }}>Coincidir</div>
          <div className="filter-control">
            <div className="select-wrap select-wrapper">
              <select className="moodle-select small" value={headerJoinLocal} onChange={handleHeaderSelectChange} aria-label="Coincidir">
                <option value="2">Todos</option>
                <option value="1">Cualquiera</option>
                <option value="0">Ninguno</option>
              </select>
            </div>
          </div>
          <div style={{ marginLeft: 8, fontSize: 16 }}>de los siguientes:</div>
        </div>
      ) : null}

    <div data-filterregion="filters" style={{ marginTop: 10 }}>
  {filters && filters.length ? filters.map((f, idx) => (
          <React.Fragment key={f.id}>
          {headerVisible && idx > 0 ? (
            <div className="filters-connector" style={{ display: 'flex', alignItems: 'center', padding: '8px 0' }}>
              <span className="connector-letter" style={{ fontWeight: 600, fontSize: 14, color: '#333' }}>{headerJoinLocal === '1' ? 'O' : 'Y'}</span>
            </div>
          ) : null}
          <div data-filterregion="filter" data-filter-type={f.field} style={{ marginBottom: 8 }}>
            <fieldset>
              <legend className="sr-only">Filtro {idx + 1}</legend>
              <div className="filter-row" style={{ alignItems: 'center' }}>
                <label className="filter-label">Coincidir</label>

                <div className="filter-control">
                  <div className="select-wrap select-wrapper">
                    {/* The join selector is global (how rows are combined). Use the provided
                        joinType prop and call onJoinChange when it changes. We keep the
                        selector visually per-row for compatibility with the original UI. */}
                    <select className="moodle-select small" data-filterfield="join" id={`core-filter_row-jointype-${f.id}`} value={f.join ?? '1'} onChange={(e) => onChange(f.id, { ...f, join: e.target.value })}>
                      <option value="0">Ninguno</option>
                      <option value="1">Cualquiera</option>
                      <option value="2">Todos</option>
                    </select>
                  </div>
                </div>

                <div className="filter-control">
                  <div className="select-wrap select-wrapper">
                    {/* Once a type has been selected for a row it becomes locked (disabled).
                        Also do not show types that are already chosen in other rows. */}
                    <select
                      className="moodle-select large"
                      data-filterfield="type"
                      id={`core-filter_row-filtertype-${f.id}`}
                      value={f.field}
                      onChange={(e) => {
                        const newField = e.target.value;
                        // If the user selects the 'accesssince' filter, and there's no
                        // value yet, initialize it to the first available option (1 día).
                        let changed = { ...f, field: newField };
                        if (newField === 'accesssince' && !changed._cleared && (!changed.value && changed.value !== 0)) {
                          const meta = findType('accesssince');
                          if (meta && Array.isArray(meta.values) && meta.values.length) {
                            // When the user explicitly selects the accesssince type we want
                            // to initialize the value to the first option (1 día). Also
                            // reset the _cleared flag because this is an explicit user action.
                            changed = { ...changed, value: meta.values[0].value, _cleared: false };
                          }
                        }
                        onChange(f.id, changed);
                      }}
                      disabled={!!f.field}
                    >
                      <option value="">Seleccionar</option>
                      {filterTypes
                        .filter(ft => ft.name === f.field || !usedTypes.includes(ft.name))
                        .map(ft => (<option key={ft.name} value={ft.name}>{ft.title}</option>))}
                    </select>
                  </div>
                </div>

                <div className="filter-control input" style={{ flex: 1 }}>
                  {(() => {
                    const meta = findType(f.field);
                    // Do not show any input when no type is selected
                    if (!meta) {
                      return null;
                    }

                    // Determine placeholder: 'Escriba...' for keywords, otherwise
                    // 'Escriba o seleccione...' as requested.
                    const placeholder = (meta && meta.name === 'keywords') ? 'Escriba...' : 'Escriba o seleccione...';

                    // If the type allows free text (custom) or has no predefined values
                    // (e.g. 'keywords' / Palabra clave) show a text input or tag input.
                    if (meta && meta.name === 'keywords') {
                      // keywords: allow multiple free-text tags
                      return (
                        <TagInput
                          id={`core-filter_row-value-${f.id}`}
                          value={Array.isArray(f.value) ? f.value : (f.value ? [f.value] : [])}
                          onChange={(vals) => onChange(f.id, { ...f, value: vals })}
                          placeholder={placeholder}
                        />
                      );
                    }

                    if (meta.allowCustom || !(meta.values && meta.values.length)) {
                      return <input className="moodle-input" value={f.value || ''} onChange={(e) => onChange(f.id, { ...f, value: e.target.value })} placeholder={placeholder} />;
                    }

                    // Multiple-choice lists (with values) should render a multi-select + chips
                    if (meta.allowMultiple && meta.values && meta.values.length) {
                      // Use the MultiCombobox so users can type to filter and add multiple values (chips)
                      const optsMulti = meta.values.map(v => ({ value: v.value, title: v.title }));
                      return (
                        <MultiCombobox
                          id={`core-filter_row-value-${f.id}`}
                          value={Array.isArray(f.value) ? f.value : []}
                          onChange={(vals) => onChange(f.id, { ...f, value: vals })}
                          options={optsMulti}
                          placeholder={placeholder}
                        />
                      );
                    }

                    // Single select with provided values
                    if (meta.values && meta.values.length) {
                      // Render our custom Combobox for single-choice lists so the user
                      // can type to filter and we can show "No hay sugerencias" when empty.
                      const opts = meta.values.map(v => ({ value: v.value, title: v.title }));

                        // If this is a single-choice select, render the Combobox as the input
                        // and show the selected value as a tag to the right (so the input
                        // width doesn't change and the dropdown of options remains accessible).
                        const selected = opts.find(o => String(o.value) === String(f.value));
                        // If the type is accesssince and there's no value yet, prefer the first option
                        // as default in the parent state. This will make the tag appear immediately.
                        if (meta.name === 'accesssince' && (typeof f.value === 'undefined') && !f._cleared) {
                          const first = opts[0] && opts[0].value;
                          if (first) {
                            // set parent state so the tag appears; wrap in try/catch to avoid render problems
                            try { onChange && onChange(f.id, { ...f, value: first }); } catch (e) {}
                          }
                        }

                        return (
                          <div style={{ position: 'relative', display: 'inline-block' }}>
                            {/* Pass an empty value to Combobox when a selection exists so the input
                                shows the placeholder and doesn't filter options by the selected title. */}
                            <Combobox
                              id={`core-filter_row-value-${f.id}`}
                              value={selected ? '' : (f.value || '')}
                              onChange={(val) => onChange(f.id, { ...f, value: val })}
                              options={opts}
                              placeholder={placeholder}
                            />

                            {selected ? (
                              <div className="tag-list" style={{ position: 'absolute', left: 'calc(100% + 8px)', top: '50%', transform: 'translateY(-50%)', display: 'flex', gap: 8, alignItems: 'center', zIndex: 10 }}>
                                <div className="chip tag-chip" style={{ display: 'inline-flex', alignItems: 'center', gap: 8, padding: '6px 10px', borderRadius: 8 }}>
                                  <span style={{ fontSize: 13 }}>{(String(selected.title || '')).replace(/\s+/g, '\u00A0')}</span>
                                  <button type="button" className="chip-close" onClick={() => onChange(f.id, { ...f, value: '', _cleared: true })} aria-label={`Quitar ${selected.title}`} style={{ border: 'none', background: 'transparent', cursor: 'pointer', fontSize: 14 }}>×</button>
                                </div>
                              </div>
                            ) : null}
                          </div>
                        );
                    }

                    // Fallback: show a text input
                    return <input className="moodle-input" value={f.value || ''} onChange={(e) => onChange(f.id, { ...f, value: e.target.value })} placeholder={placeholder} />;
                  })()}
                </div>

                <button data-filteraction="remove" onClick={() => onRemove(f.id)} className="filters-remove" aria-label="Eliminar filtro de fila" type="button">✕</button>
              </div>

            </fieldset>
          </div>
          </React.Fragment>
  )) : null}
      </div>

      <div className="d-flex" data-filterregion="actions" style={{ display: 'flex', alignItems: 'center', marginTop: 8 }}>
  <button type="button" className="btn btn-link text-reset" data-filteraction="add" onClick={handleAddClick} disabled={(filters || []).length >= 5} title={(filters || []).length >= 5 ? 'Máximo 5 condiciones' : ''} style={{ display: 'inline-flex', alignItems: 'center', gap: 8 }}>
          <span className="add-plus" style={{ fontSize: 22, lineHeight: 1, fontWeight: 700, display: 'inline-block' }}>+</span>
          <span className="add-text" style={{ fontSize: 13, lineHeight: 1, fontWeight: 400 }}>Añadir condición</span>
        </button>
        <div style={{ marginLeft: 'auto', display: 'flex', gap: 8, alignItems: 'center' }}>
          <button data-filteraction="reset" type="button" className="btn btn-secondary" onClick={onClear}>Limpiar filtros</button>
          <button data-filteraction="apply" type="button" className="btn btn-primary" onClick={onApply}>Aplicar filtros</button>
        </div>
      </div>
    </div>
  );
}
