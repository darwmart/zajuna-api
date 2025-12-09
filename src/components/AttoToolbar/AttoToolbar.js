import React, { useRef, useEffect, useState, useCallback } from 'react';
import AttoButton from './AttoButton';
import { uploadFile } from '../../services/uploadAdapter';
// Using Font Awesome icons instead of inline SVG components for toolbar glyphs

export default function AttoToolbar({ value, onChange, placeholder }) {
  const editorRef = useRef(null);
  const selectionRef = useRef(null);
  const fileInputRef = useRef(null);
  const attachInputRef = useRef(null);
  const [expanded, setExpanded] = useState(false);
  const [activeStates, setActiveStates] = useState({});
  const [toggled, setToggled] = useState({});

  const exec = useCallback((command, valueArg) => {
    try {
      // ensure editor has focus and a selection before executing
      if (editorRef.current && typeof editorRef.current.focus === 'function') editorRef.current.focus();
      const sel = window.getSelection();
      if (sel && sel.rangeCount === 0 && editorRef.current) {
        const r = document.createRange();
        r.selectNodeContents(editorRef.current);
        r.collapse(false);
        sel.removeAllRanges();
        sel.addRange(r);
      }
      document.execCommand(command, false, valueArg);
    } catch (e) {
      // no-op in browsers that deprecate execCommand
    }
    if (editorRef.current && typeof editorRef.current.focus === 'function') editorRef.current.focus();
  }, []);

  const toggleCommand = useCallback((name, command, valueArg) => {
    try {
      if (editorRef.current && typeof editorRef.current.focus === 'function') editorRef.current.focus();
      const sel = window.getSelection();
      if (sel && sel.rangeCount === 0 && editorRef.current) {
        const r = document.createRange();
        r.selectNodeContents(editorRef.current);
        r.collapse(false);
        sel.removeAllRanges();
        sel.addRange(r);
      }
      document.execCommand(command, false, valueArg);
    } catch (e) {
      // ignore
    }
    setToggled((s) => ({ ...s, [name]: !s[name] }));
    // attempt to sync activeStates after toggle
    setTimeout(() => {
      if (document.queryCommandState) {
        setActiveStates((prev) => ({ ...prev, [name]: document.queryCommandState(name) }));
      }
    }, 0);
    if (editorRef.current && typeof editorRef.current.focus === 'function') editorRef.current.focus();
  }, []);

  const syncValue = useCallback(() => {
    if (onChange && editorRef.current) onChange(editorRef.current.innerHTML);
  }, [onChange]);

  const handleFileSelect = async (e, type) => {
    const f = e.target.files && e.target.files[0];
    if (!f) return;
    try {
      const url = await uploadFile(f);
      if (type === 'image') {
        const img = document.createElement('img');
        img.src = url;
        insertNodeAtCursor(img);
        syncValue();
      }
    } catch (err) {
      // ignore upload errors for now
    }
  };

  const handleLink = () => {
    // try to preserve selection across the prompt
    try {
      const sel = window.getSelection();
      if (sel && sel.rangeCount > 0) selectionRef.current = sel.getRangeAt(0).cloneRange();
    } catch (e) {
      selectionRef.current = null;
    }
    const url = window.prompt('Introduce la URL');
    if (!url) return;
    // restore saved selection if present
    try {
      const sel = window.getSelection();
      sel.removeAllRanges();
      if (selectionRef.current) {
        sel.addRange(selectionRef.current);
        selectionRef.current = null;
      } else if (editorRef.current) {
        const r = document.createRange(); r.selectNodeContents(editorRef.current); r.collapse(false); sel.addRange(r);
      }
    } catch (e) {
      // ignore
    }
    exec('createLink', url);
    syncValue();
  };

  const handleUnlink = () => {
    exec('unlink');
    syncValue();
  };

  const handleH5P = () => {
    alert('Insertar/editar H5P (stub)');
  };

  const insertNodeAtCursor = (node) => {
    const sel = window.getSelection();
    if (!sel || sel.rangeCount === 0) {
      if (editorRef.current) editorRef.current.appendChild(node);
      if (editorRef.current && typeof editorRef.current.focus === 'function') editorRef.current.focus();
      return;
    }
    const range = sel.getRangeAt(0);
    range.deleteContents();
    range.insertNode(node);
    range.setStartAfter(node);
    range.collapse(true);
    sel.removeAllRanges();
    sel.addRange(range);
    if (editorRef.current && typeof editorRef.current.focus === 'function') editorRef.current.focus();
  };

  const handleInput = () => {
    syncValue();
  };

  useEffect(() => {
    const updateActive = () => {
      // Only read command state when the selection/caret is inside our editor
      const sel = window.getSelection && window.getSelection();
      let inEditor = false;
      try {
        // If the focused element is inside the editor, treat as in-editor
        const activeEl = document.activeElement;
        if (activeEl && editorRef.current && editorRef.current.contains(activeEl)) {
          inEditor = true;
        } else if (sel && sel.rangeCount > 0 && editorRef.current) {
          let node = sel.anchorNode;
          while (node) {
            if (node === editorRef.current) {
              inEditor = true;
              break;
            }
            node = node.parentNode;
          }
        }
      } catch (e) {
        inEditor = false;
      }

      if (!inEditor) {
        // Clear transient active states coming from document context when outside editor
        setActiveStates({});
        return;
      }

      const states = {
        bold: document.queryCommandState && document.queryCommandState('bold'),
        italic: document.queryCommandState && document.queryCommandState('italic'),
        unordered: document.queryCommandState && document.queryCommandState('insertUnorderedList'),
        ordered: document.queryCommandState && document.queryCommandState('insertOrderedList'),
        underline: document.queryCommandState && document.queryCommandState('underline'),
      };
      setActiveStates(states);
      // Clear any local toggles that no longer match the real command state
      setToggled((prev) => {
        const next = { ...prev };
        ['bold', 'italic', 'underline'].forEach((k) => {
          if (!states[k]) next[k] = false;
        });
        return next;
      });
    };
    document.addEventListener('selectionchange', updateActive);
    return () => document.removeEventListener('selectionchange', updateActive);
  }, []);

  useEffect(() => {
    if (editorRef.current && value !== undefined && value !== editorRef.current.innerHTML) {
      editorRef.current.innerHTML = value;
    }
  }, [value]);

  return (
    <div className="atto-toolbar" style={{ marginBottom: 8 }}>
      <div className="atto-toolbar-row" role="toolbar" aria-label="Atto toolbar">
        {/* 1. Mostrar más (Expandir) */}
        <AttoButton title={expanded ? 'Mostrar menos' : 'Mostrar más'} onClick={() => setExpanded((s) => !s)} ariaLabel="more" className={`atto-more-btn ${expanded ? 'expanded' : ''}`} active={expanded}>
          <i className="icon fa fa-level-down fa-fw" aria-hidden="true" />
        </AttoButton>

        {/* Main controls (column 2 of the grid) */}
        <div className="atto-toolbar-main">
          {/* Format group */}
          <div className="atto-format-group" role="group" aria-label="format controls">
            <AttoButton title="Formato" onClick={() => {}} ariaLabel="format" className="atto-format-btn">
              <span className="atto-icon">A</span>
              <span className="caret">▾</span>
            </AttoButton>
            <AttoButton title="Negrita" onClick={() => toggleCommand('bold','bold')} ariaLabel="bold" active={toggled.bold || activeStates.bold}>
              <span className="atto-letter atto-bold">B</span>
            </AttoButton>
            <AttoButton title="Cursiva" onClick={() => toggleCommand('italic','italic')} ariaLabel="italic" active={toggled.italic || activeStates.italic}>
              <span className="atto-letter atto-italic">I</span>
            </AttoButton>
          </div>

          {/* List controls */}
          <div className="atto-list-group" role="group" aria-label="list controls">
            <AttoButton title="Viñetas" onClick={() => exec('insertUnorderedList')} ariaLabel="unordered-list" active={activeStates.unordered}>
              <i className="icon fa fa-list-ul fa-fw" aria-hidden="true" />
            </AttoButton>
            <AttoButton title="Lista numerada" onClick={() => exec('insertOrderedList')} ariaLabel="ordered-list" active={activeStates.ordered}>
              <i className="icon fa fa-list-ol fa-fw" aria-hidden="true" />
            </AttoButton>
            <AttoButton title="Sangría inversa" onClick={() => exec('outdent')} ariaLabel="outdent" active={activeStates.outdent}>
              <i className="icon fa fa-outdent fa-fw" aria-hidden="true" />
            </AttoButton>
            <AttoButton title="Sangría" onClick={() => exec('indent')} ariaLabel="indent" active={activeStates.indent}>
              <i className="icon fa fa-indent fa-fw" aria-hidden="true" />
            </AttoButton>
          </div>

          {/* Link group */}
          <div className="atto-link-group" role="group" aria-label="link controls">
            <AttoButton title="Enlace" onClick={handleLink} ariaLabel="link" active={activeStates.link}><i className="icon fa fa-link fa-fw" aria-hidden="true" /></AttoButton>
            <AttoButton title="Quitar enlace" onClick={handleUnlink} ariaLabel="unlink"><i className="icon fa fa-unlink fa-fw" aria-hidden="true" /></AttoButton>
          </div>

          {/* Media group */}
          <div className="atto-media-group" role="group" aria-label="media controls">
            <AttoButton title="Emoticonos" onClick={() => {
              try { const sel = window.getSelection(); if (sel && sel.rangeCount > 0) selectionRef.current = sel.getRangeAt(0).cloneRange(); } catch (e) { selectionRef.current = null; }
              const emoji = window.prompt('Inserta emoji');
              if (!emoji) return;
              try { const sel2 = window.getSelection(); sel2.removeAllRanges(); if (selectionRef.current) { sel2.addRange(selectionRef.current); selectionRef.current = null; } else if (editorRef.current) { const r = document.createRange(); r.selectNodeContents(editorRef.current); r.collapse(false); sel2.addRange(r); } } catch (e) {}
              exec('insertText', emoji);
              syncValue();
            }} ariaLabel="emoji"><i className="icon fa fa-smile-o fa-fw" aria-hidden="true" /></AttoButton>
            <AttoButton title="Imagen" onClick={() => fileInputRef.current && fileInputRef.current.click()} ariaLabel="image"><i className="icon fa fa-picture-o fa-fw" aria-hidden="true" /><input ref={fileInputRef} type="file" accept="image/*" style={{ display: 'none' }} onChange={(e) => handleFileSelect(e, 'image')} /></AttoButton>
            <AttoButton title="Insertar audio/video" onClick={() => alert('Insertar/editar audio/video (stub)')} ariaLabel="media"><i className="icon fa fa-file-video-o fa-fw" aria-hidden="true" /></AttoButton>
            <AttoButton title="Grabar audio" onClick={() => alert('Se abriría el grabador (stub)')} ariaLabel="record-audio"><i className="icon fa fa-microphone fa-fw" aria-hidden="true" /></AttoButton>
            <AttoButton title="Grabar vídeo" onClick={() => alert('Grabar vídeo (stub)')} ariaLabel="record-video"><i className="icon fa fa-video-camera fa-fw" aria-hidden="true" /></AttoButton>
            <AttoButton title="Gestionar ficheros" onClick={() => attachInputRef.current && attachInputRef.current.click()} ariaLabel="manage-files"><i className="icon fa fa-files-o fa-fw" aria-hidden="true" /><input ref={attachInputRef} type="file" style={{ display: 'none' }} onChange={(e) => handleFileSelect(e, 'attachment')} /></AttoButton>
            <AttoButton title="H5P" onClick={handleH5P} ariaLabel="h5p" className="atto-h5p-button">
              <span className="icon atto-h5p-icon" aria-hidden="true">H5P</span>
            </AttoButton>
          </div>

          {/* Accessibility group */}
          <div className="atto-access-group" role="group" aria-label="accessibility controls">
            <AttoButton title="Comprobaciones de accesibilidad" onClick={() => alert('Comprobaciones de accesibilidad (stub)')} ariaLabel="accessibility"><i className="icon fa fa-universal-access fa-fw" aria-hidden="true" /></AttoButton>
            <AttoButton title="Ayudante lector de pantalla" onClick={() => alert('Ayudante de lector de pantalla (stub)')} ariaLabel="screen-reader"><i className="icon fa fa-braille fa-fw" aria-hidden="true" /></AttoButton>
          </div>
        </div>

        {/* Secondary row: shown when expanded. Rendered as sibling so it occupies grid column 2 */}
        {expanded && (
          <div className="atto-toolbar-row atto-toolbar-row-secondary" role="toolbar" aria-label="more toolbar controls">
            <div className="atto-script-group" role="group" aria-label="script controls">
                <AttoButton title="Subrayar" onClick={() => toggleCommand('underline','underline')} ariaLabel="underline" active={toggled.underline || activeStates.underline}><i className="icon fa fa-underline fa-fw" aria-hidden="true" /></AttoButton>
              <AttoButton title="Tachado" onClick={() => exec('strikeThrough')} ariaLabel="strikethrough"><i className="icon fa fa-strikethrough fa-fw" aria-hidden="true" /></AttoButton>
              <AttoButton title="Subíndice" onClick={() => exec('subscript')} ariaLabel="subscript">
                <i className="icon fa fa-subscript fa-fw" aria-hidden="true" />
              </AttoButton>
              <AttoButton title="Sobrescrito" onClick={() => exec('superscript')} ariaLabel="superscript">
                <i className="icon fa fa-superscript fa-fw" aria-hidden="true" />
              </AttoButton>
            </div>

            <div className="atto-justify-group" role="group" aria-label="justify controls">
              <AttoButton title="Alinear a la izquierda" onClick={() => exec('justifyLeft')} ariaLabel="justify-left"><i className="icon fa fa-align-left fa-fw" aria-hidden="true" /></AttoButton>
              <AttoButton title="Centrar" onClick={() => exec('justifyCenter')} ariaLabel="justify-center"><i className="icon fa fa-align-center fa-fw" aria-hidden="true" /></AttoButton>
              <AttoButton title="Alinear a la derecha" onClick={() => exec('justifyRight')} ariaLabel="justify-right"><i className="icon fa fa-align-right fa-fw" aria-hidden="true" /></AttoButton>
            </div>

            <div className="atto-insert-group" role="group" aria-label="insert controls">
              <AttoButton title="Editor de ecuaciones" onClick={() => {
                const latex = window.prompt('Introduce LaTeX (ej: x^2 + y^2)');
                if (latex) {
                  const span = document.createElement('span');
                  span.className = 'atto-equation';
                  span.textContent = `\\(${latex}\\)`;
                  insertNodeAtCursor(span);
                  syncValue();
                }
              }} ariaLabel="equation">
                <i className="icon fa fa-calculator fa-fw" aria-hidden="true" />
              </AttoButton>
              <AttoButton title="Insertar carácter" onClick={() => { try { const sel = window.getSelection(); if (sel && sel.rangeCount > 0) selectionRef.current = sel.getRangeAt(0).cloneRange(); } catch (e) { selectionRef.current = null; } const ch = window.prompt('Introduce el carácter a insertar'); if (!ch) return; try { const sel2 = window.getSelection(); sel2.removeAllRanges(); if (selectionRef.current) { sel2.addRange(selectionRef.current); selectionRef.current = null; } else if (editorRef.current) { const r = document.createRange(); r.selectNodeContents(editorRef.current); r.collapse(false); sel2.addRange(r); } } catch (e) {} exec('insertText', ch); syncValue(); }} ariaLabel="insert-char">
                <i className="icon fa fa-pencil-square-o fa-fw" aria-hidden="true" />
              </AttoButton>
              <AttoButton title="Insertar tabla" onClick={() => {
                const rows = parseInt(window.prompt('Filas', '2'), 10) || 2;
                const cols = parseInt(window.prompt('Columnas', '2'), 10) || 2;
                let html = '<table class="atto-table" border="1" cellpadding="4">';
                for (let r = 0; r < rows; r++) { html += '<tr>'; for (let c = 0; c < cols; c++) html += '<td>&nbsp;</td>'; html += '</tr>'; }
                html += '</table>';
                insertNodeAtCursor((() => { const d = document.createElement('div'); d.innerHTML = html; return d; })());
                syncValue();
              }} ariaLabel="table"><i className="icon fa fa-table fa-fw" aria-hidden="true" /></AttoButton>
              <AttoButton title="Limpiar formato" onClick={() => { exec('removeFormat'); }} ariaLabel="remove-format">
                <i className="icon fa fa-i-cursor fa-fw" aria-hidden="true" />
              </AttoButton>
            </div>

            <div className="atto-undo-group" role="group" aria-label="undo controls">
              <AttoButton title="Deshacer" onClick={() => exec('undo')} ariaLabel="undo"><i className="icon fa fa-undo fa-fw" aria-hidden="true" /></AttoButton>
              <AttoButton title="Rehacer" onClick={() => exec('redo')} ariaLabel="redo"><i className="icon fa fa-repeat fa-fw" aria-hidden="true" /></AttoButton>
            </div>

            <AttoButton title="HTML" onClick={() => {
              if (!editorRef.current) return;
              const current = editorRef.current.innerHTML;
              const edited = window.prompt('HTML del editor (reemplazar):', current);
              if (edited !== null) { editorRef.current.innerHTML = edited; syncValue(); }
            }} ariaLabel="html"><i className="icon fa fa-code fa-fw" aria-hidden="true" /></AttoButton>
          </div>
        )}
      </div>

      <div
        ref={editorRef}
        className="mform-input atto-textarea"
        contentEditable
        onInput={handleInput}
        data-placeholder={placeholder || ''}
        style={{ whiteSpace: 'pre-wrap', height: 308, minHeight: 308 }}
        suppressContentEditableWarning
      />
    </div>
  );

}

