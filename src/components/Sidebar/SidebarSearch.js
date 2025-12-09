import React, { useState } from "react";
import { FaSearch } from "react-icons/fa";
import "../../styles/moodle-theme.css";

export default function SidebarSearch({ placeholder = "Ajustes de búsqueda" }) {
  const [searchValue, setSearchValue] = useState("");

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log("Buscar:", searchValue);
    // Aquí puedes agregar la lógica de búsqueda
  };

  return (
    <form className="sidebar-search" onSubmit={handleSubmit}>
      <div className="sidebar-search-wrapper">
        <input
          type="text"
          value={searchValue}
          onChange={(e) => setSearchValue(e.target.value)}
          placeholder={placeholder}
          className="sidebar-search-input"
        />
        <button type="submit" className="sidebar-search-btn" aria-label="Buscar">
          <FaSearch size={14} />
        </button>
      </div>
    </form>
  );
}
