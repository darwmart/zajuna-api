import React from "react";
import { Box } from "@mui/material";
import ZajunaHeader from "./ZajunaHeader";
import ZajunaSidebar from "./ZajunaSidebar";
import ZajunaFooter from "./ZajunaFooter";

function ZajunaLayout({ children }) {
  return (
    <Box sx={{ display: "flex", flexDirection: "column", height: "auto", overflow: "hidden" }}>
      {/* Encabezados */}
      <ZajunaHeader />

      {/* Franja blanca separadora */}
      <Box sx={{ 
        width: "100%", 
        height: "120px", 
        backgroundColor: "#FFFFFF",
        borderBottom: "1px solid var(--moodle-green)",
        boxShadow: "0 2px 4px rgba(0,0,0,0.1)"
      }} />

      {/* Cuerpo principal */}
      <Box sx={{ display: "flex", backgroundColor: "#E9ECEF", minHeight: "400px" }}>
        {/* Sidebar - comienza al nivel del contenido */}
          <Box sx={{ 
            /* use CSS variable so layout adapts automatically when --sidebar-width changes */
            width: "var(--sidebar-width)",
            minWidth: "var(--sidebar-width)",
            backgroundColor: "#FFFFFF"
          }}>
          <ZajunaSidebar />
        </Box>

        {/* Contenido principal */}
        <Box
          sx={{
            flex: 1,
            minWidth: 0,
            pt: "12px", /* reduced top padding */
            pr: "31px", /* right padding set to 31px as requested */
            pb: 2,      /* reduced bottom padding */
            pl: 2,      /* reduced left padding */
            backgroundColor: "white",
            borderLeft: "1px solid var(--moodle-green)",
            fontFamily: "Arial, sans-serif",
            boxSizing: "border-box",
          }}
        >
          {children}
        </Box>
      </Box>

      {/* Separador blanco antes del pie de página */}
      <Box sx={{ 
        width: "100%", 
        height: "28px", 
        backgroundColor: "#FFFFFF",
        borderTop: "1px solid var(--moodle-green)"
      }} />

      {/* Pie de página: franja azul */}
      <ZajunaFooter />
    </Box>
  );
}

export default ZajunaLayout;
