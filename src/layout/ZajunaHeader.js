import React from "react";
import { Box } from "@mui/material";
import "../styles/moodle-theme.css";
import ZajunaLogo from "../assets/logos/zajuna-logo.svg";
import SenaLogo from "../assets/logos/sena-logo.svg";

function ZajunaHeader() {
  return (
    <Box sx={{ position: "sticky", top: 0, left: 0, width: "100%", zIndex: 1200 }}>
      {/* Franja superior blanca (más gruesa) */}
      <Box
        sx={{
          width: "100%",
          height: "84px",
          backgroundColor: "#FFFFFF",
          display: "flex",
          alignItems: "center",
        }}
      >
        <Box sx={{
          width: "100%",
          display: "flex",
          alignItems: "center",
          justifyContent: "space-between",
          px: 3,
          boxSizing: "border-box",
        }}>
          <img src={ZajunaLogo} alt="Zajuna" style={{ height: 34, objectFit: "contain" }} />
          <img src={SenaLogo} alt="SENA" style={{ height: 64, objectFit: "contain" }} />
        </Box>
      </Box>

      {/* Franja azul (Moodle), menos gruesa que la blanca */}
      <Box
        sx={{
          width: "100%",
          height: "68px",
          backgroundColor: "var(--moodle-header-blue)",
          display: "flex",
          alignItems: "center",
          px: 3,
        }}
      >
        <a 
          href="https://oferta.senasofiaplus.edu.co/sofia-oferta/" 
          target="_blank"
          rel="noopener noreferrer"
          style={{
            color: "#FFFFFF",
            fontSize: "14px",
            fontFamily: "Arial, sans-serif",
            textDecoration: "none",
            fontWeight: "400",
          }}
        >
          Accede a Sofía
        </a>
      </Box>

      {/* Franja verde (SENA), menos gruesa que la azul */}
      <Box
        sx={{
          width: "100%",
          height: "32px",
          backgroundColor: "var(--moodle-green)",
          display: "flex",
          alignItems: "center",
          justifyContent: "center",
          color: "#FFFFFF",
          fontSize: "14px",
          fontWeight: "600",
          fontFamily: "Arial, sans-serif",
        }}
      >
        SENA - Zajuna
      </Box>
    </Box>
  );
}

export default ZajunaHeader;
