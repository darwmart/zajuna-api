import React, { useEffect, useState, useCallback } from "react";
import {
  Typography,
  TextField,
  List,
  ListItem,
  ListItemText,
  CircularProgress,
  Alert,
  Pagination,
  PaginationItem,
  Box,
  Button,
  Chip,
} from "@mui/material";
import RestoreIcon from "@mui/icons-material/Restore";
import { getUsers, undeleteUsers } from "../services/usersService";
import useDebounce from "../hooks/useDebounce";

function UsersDeletedList() {
  const [users, setUsers] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [query, setQuery] = useState("");
  const [page, setPage] = useState(1);
  const [perPage] = useState(10);
  const [total, setTotal] = useState(0);

  const debouncedQuery = useDebounce(query, 400);

  const fetchUsers = useCallback(() => {
    setLoading(true);
    setError(null);
    getUsers(debouncedQuery, page, perPage, 1) // deleted = 1
      .then((data) => {
        setUsers(data.items);
        setTotal(data.total);
        setLoading(false);
      })
      .catch(() => {
        setError("No se pudieron cargar los usuarios eliminados");
        setLoading(false);
      });
  }, [debouncedQuery, page, perPage]);

  useEffect(() => {
    fetchUsers();
  }, [fetchUsers]);

  const handleQueryChange = (e) => {
    setQuery(e.target.value);
    setPage(1);
  };

  const handleRestore = async (id) => {
    try {
      await undeleteUsers([id]);
      fetchUsers(); // Refrescar lista
    } catch (error) {
      console.error("Error restaurando usuario:", error);
    }
  };

  const totalPages = Math.ceil(total / perPage);

  return (
    <Box>
      <Typography variant="h5" gutterBottom>
        Usuarios Eliminados
      </Typography>

      <TextField
        fullWidth
        variant="outlined"
        label="Filtrar por nombre, apellido o email"
        value={query}
        onChange={handleQueryChange}
        style={{ marginBottom: "20px" }}
      />

      {loading && <CircularProgress />}
      {error && <Alert severity="error">{error}</Alert>}

      {!loading && !error && (
        <>
          <Chip
            label={`Total: ${total} eliminados`}
            color="secondary"
            variant="outlined"
            style={{ marginBottom: "15px", fontWeight: "bold" }}
          />
          <List>
            {users.map((u) => (
              <ListItem key={u.id} divider>
                <ListItemText
                  primary={`ID: ${u.id} - ${u.firstname} ${u.lastname}`}
                  secondary={`${u.email} | Usuario: ${u.username}`}
                />
                <Button
                  variant="contained"
                  color="success"
                  startIcon={<RestoreIcon />}
                  onClick={() => handleRestore(u.id)}
                >
                  Restaurar
                </Button>
              </ListItem>
            ))}
          </List>

          {totalPages > 1 && (
            <Pagination
              count={totalPages}
              page={page}
              onChange={(_, value) => setPage(value)}
              color="primary"
              siblingCount={page <= 10 ? 9 : 4}
              boundaryCount={page <= 10 ? 0 : 1}
              renderItem={(item) => (
                <PaginationItem
                  {...item}
                />
              )}
              style={{
                marginTop: "20px",
                display: "flex",
                justifyContent: "center",
              }}
            />
          )}
        </>
      )}
    </Box>
  );
}

export default UsersDeletedList;
