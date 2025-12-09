import React from "react";
import { FaHome, FaGlobe, FaBook, FaCog, FaChevronRight, FaUsers, FaUserPlus, FaUsersCog, FaTable, FaUser, FaUserFriends } from "react-icons/fa";
import { useNavigate } from "react-router-dom";
import "../../styles/moodle-theme.css";

const iconMap = {
  home: FaHome,
  site: FaGlobe,
  book: FaBook,
  admin: FaCog,
  chevron: FaChevronRight,
  users: FaUsers,
  "user-plus": FaUserPlus,
  "users-cog": FaUsersCog,
  cog: FaCog, // Moodle settings icon
  table: FaTable, // Grades/Calificaciones icon
  user: FaUser, // Single user icon
  group: FaUserFriends, // Group icon
};

export default function SidebarNavItem({ 
  icon, 
  text, 
  level = 0, 
  href = "#", 
  active = false,
  onClick,
  hasChildren = false,
  isExpanded = false,
  onToggle,
  showBullet = false,
}) {
  const IconComponent = icon ? iconMap[icon] : null;
  const navigate = useNavigate();
  
  const handleChevronClick = (e) => {
    e.preventDefault();
    e.stopPropagation();
    if (hasChildren && onToggle) {
      onToggle();
    }
  };

  const handleTextClick = (e) => {
    // If there's a custom onClick handler, call it first (priority)
    if (onClick) {
      e.preventDefault();
      onClick(e);
      return;
    }

    // If a navigation href is provided, prefer navigating when the text is clicked.
    // The chevron still toggles expand/collapse. This preserves the common UX where
    // the text opens the linked view while the chevron only expands the subtree.
    if (href && href !== "#") {
      // If the item both has children and a href, do BOTH: toggle the subtree and navigate.
      // Toggle first so the UI expands immediately, then navigate programmatically.
      if (hasChildren && onToggle) {
        e.preventDefault();
        try {
          onToggle();
        } catch (err) {
          // ignore toggle errors
        }
        try {
          navigate(href);
        } catch (err) {
          if (typeof window !== 'undefined') {
            window.location.hash = href;
          }
        }
        return;
      }

      // Otherwise just navigate
      try {
        navigate(href);
      } catch (err) {
        if (typeof window !== 'undefined') {
          window.location.hash = href;
        }
      }
      return;
    }
  };

  return (
    <div className={`sidebar-nav-item level-${level} ${active ? 'active' : ''}`}>
      <div className="sidebar-nav-link">
        {(hasChildren || showBullet) && (
          <span className="nav-item-chevron" onClick={handleChevronClick}>
            <FaChevronRight 
              size={10} 
              style={{ 
                transform: hasChildren && isExpanded ? 'rotate(90deg)' : 'rotate(0deg)',
                transition: 'transform 0.2s'
              }} 
            />
          </span>
        )}
        {IconComponent && (
          <span className="nav-item-icon">
            <IconComponent size={14} />
          </span>
        )}
        {/* Render as a real link so hash navigation and "open in new tab" work reliably. */}
        {href && href !== "#" ? (
          <a className="nav-item-text" href={`#${href}`} onClick={handleTextClick}>{text}</a>
        ) : (
          <span className="nav-item-text" onClick={handleTextClick}>{text}</span>
        )}
      </div>
    </div>
  );
}
