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
  Chip,
  Button,
  Stack,
} from "@mui/material";

import { getUsers } from "../services/usersService";
import useDebounce from "../hooks/useDebounce";
import { deleteUsers, undeleteUsers } from "../services/usersService";
import { useNavigate } from "react-router-dom";


function UsersList() {
  const navigate = useNavigate();
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

    getUsers(debouncedQuery, page, perPage)
      .then((data) => {
        setUsers(data.items);
        setTotal(data.total);
        setLoading(false);
      })
      .catch(() => {
        setError("No se pudieron cargar los usuarios");
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

  const totalPages = Math.ceil(total / perPage);

  return (
    <Box>
      <Typography variant="h5" gutterBottom>
        Usuarios
      </Typography>

      <TextField
        fullWidth
        variant="outlined"
        label="Filtrar por nombre completo, email o usuario"
        value={query}
        onChange={handleQueryChange}
        style={{ marginBottom: "20px" }}
      />

      {loading && <CircularProgress />}
      {error && <Alert severity="error">{error}</Alert>}

      {!loading && !error && (
        <>
            <Chip
            label={`Total: ${total} usuarios`}
            color="primary"
            variant="outlined"
            style={{ marginBottom: "15px", fontWeight: "bold" }}
            />
          <List>
            {users.map((u) => (
              <ListItem
                key={u.id}
                divider
                secondaryAction={
                  <Stack direction="row" spacing={1}>
                    <Button
                      variant="outlined"
                      color="primary"
                      size="small"
                      onClick={() => navigate(`/users/${u.id}/edit`)}
                    >
                      Editar
                    </Button>

                    <Button
                      variant="outlined"
                      color="error"
                      size="small"
                      onClick={async () => {
                        if (window.confirm(`¿Eliminar a ${u.firstname}?`)) {
                          await deleteUsers([u.id]);
                          fetchUsers();
                        }
                      }}
                    >
                      Eliminar
                    </Button>
                  </Stack>
                }
              >
                <ListItemText
                  primary={`ID: ${u.id} - ${u.firstname} ${u.lastname}`}
                  secondary={
                    <>
                      <span>{`${u.email} | Usuario: ${u.username}`}</span>
                      <br />
                      <span>{`${u.city || "Sin ciudad"}, ${u.country || "Sin país"}`}</span>
                      <br />
                      <span>{`Último acceso: ${u.lastlogin || "Nunca"}`}</span>
                    </>
                  }
                />
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

export default UsersList;
