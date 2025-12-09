import React from "react";
import { Box } from "@mui/material";
import ZajunaLogo from "../assets/logos/zajuna-logo.svg";

// Footer: single blue stripe matching Moodle Zajuna blue
// Height requested: 120px
export default function ZajunaFooter() {
  return (
    <Box
      component="footer"
      sx={{
        width: "100%",
        height: "120px",
        backgroundColor: "var(--moodle-header-blue)",
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
      }}
    >
      <img src={ZajunaLogo} alt="Zajuna" style={{ height: 24, objectFit: "contain" }} />
    </Box>
  );
}
