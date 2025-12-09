import React, { useState } from "react";
import {
  Typography,
  TextField,
  Button,
  List,
  ListItem,
  ListItemText,
  Alert,
  Box,
  Stack,
  MenuItem,
  Paper,
  Divider,
} from "@mui/material";
import { getUsersByField } from "../services/usersService";

function UserSearchByField() {
  const [field, setField] = useState("email");
  const [value, setValue] = useState("");
  const [users, setUsers] = useState([]);
  const [error, setError] = useState(null);
  const [loading, setLoading] = useState(false);

  const handleSearch = async () => {
    if (!value.trim()) return;
    setError(null);
    setUsers([]);
    setLoading(true);

    try {
      const data = await getUsersByField(field, value);
      setUsers(data.items || []);
    } catch {
      setError("Error buscando usuarios");
    } finally {
      setLoading(false);
    }
  };

  return (
    <Paper elevation={3} sx={{ p: 3, mt: 3 }}>
      <Typography variant="h5" gutterBottom>
        Buscar usuarios por campo
      </Typography>

      <Stack direction="row" spacing={2} sx={{ mb: 2 }}>
        <TextField
          select
          label="Campo"
          value={field}
          onChange={(e) => setField(e.target.value)}
          sx={{ width: 200 }}
        >
          <MenuItem value="email">Correo electrónico</MenuItem>
          <MenuItem value="username">Nombre de usuario</MenuItem>
          <MenuItem value="firstname">Nombre</MenuItem>
          <MenuItem value="lastname">Apellido</MenuItem>
        </TextField>

        <TextField
          fullWidth
          variant="outlined"
          label={`Buscar por ${field}`}
          value={value}
          onChange={(e) => setValue(e.target.value)}
        />

        <Button
          variant="contained"
          color="primary"
          onClick={handleSearch}
          disabled={!value.trim() || loading}
        >
          {loading ? "Buscando..." : "Buscar"}
        </Button>
      </Stack>

      {error && <Alert severity="error">{error}</Alert>}

      <Divider sx={{ my: 2 }} />

      {users.length > 0 ? (
        <List>
          {users.map((u) => (
            <ListItem key={u.id} divider>
              <ListItemText
                primary={`${u.firstname} ${u.lastname}`}
                secondary={
                  <>
                    <span>{`${u.email} | Usuario: ${u.username}`}</span>
                    <br />
                    <span>{`Último acceso: ${u.lastlogin || "Nunca"}`}</span>
                  </>
                }
              />
            </ListItem>
          ))}
        </List>
      ) : (
        !loading &&
        !error && (
          <Typography variant="body2" color="text.secondary">
            No hay resultados aún.
          </Typography>
        )
      )}
    </Paper>
  );
}

export default UserSearchByField;
