import React from 'react';

export default function AttoButton({ title, onClick, children, active, disabled, ariaLabel, className = '' }) {
  const classes = ['atto-btn', active ? 'atto-active' : '', className].filter(Boolean).join(' ');
  return (
    <button
      type="button"
      className={classes}
      data-title={title}
      aria-pressed={active ? 'true' : 'false'}
      aria-label={ariaLabel || title}
      onClick={onClick}
      onMouseDown={(e) => e.preventDefault()}
      disabled={disabled}
    >
      {children}
    </button>
  );
}
