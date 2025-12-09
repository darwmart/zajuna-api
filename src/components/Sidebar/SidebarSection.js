import React, { useState } from "react";
import "../../styles/moodle-theme.css";

export default function SidebarSection({ title, children, defaultOpen = false, className = '' }) {
  const [isOpen] = useState(defaultOpen);

  // allow injecting a custom class for special sections (e.g. upcoming events)
  const rootClass = `sidebar-section ${className}`.trim();

  return (
    <div className={rootClass}>
      <div className="sidebar-section-body" style={{padding: 16}}>
        <h6 className="sidebar-section-header" style={{paddingLeft: 8}}>
          <span className="sidebar-section-title">{title}</span>
        </h6>
        {isOpen && (
          <nav className="sidebar-section-content">
            {children}
          </nav>
        )}
      </div>
    </div>
  );
}
