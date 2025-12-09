import api from "../api/apiClientUsers";

// Buscar usuarios con query y paginación
// Backend soporta filtros: firstname, lastname, username, email, page, limit
export const getUsers = async (q = "", page = 1, perPage = 15, filters = {}) => {
  const params = {
    page,
    limit: perPage,
    ...filters, // Puede incluir: firstname, lastname, username, email
  };

  const res = await api.get("/users", { params });

  // Normalizar respuesta del backend al formato esperado por el frontend
  const response = res.data;
  // El backend devuelve: { data: [...], pagination: {...} }
  const usersList = response.data || response.users || [];

  return {
    items: usersList,
    users: usersList,
    total: response.pagination?.total || 0,
    page: response.pagination?.page || page,
    limit: response.pagination?.limit || perPage,
  };
};

// Buscar usuarios por campo (helper que usa getUsers con filtros específicos)
export const getUsersByField = async (field, values) => {
  const filters = {};
  // Mapear el campo al formato esperado por el backend
  if (field === "email") filters.email = values;
  else if (field === "username") filters.username = values;
  else if (field === "firstname") filters.firstname = values;
  else if (field === "lastname") filters.lastname = values;

  const res = await api.get("/users", { params: filters });

  // Normalizar respuesta del backend al formato esperado por el frontend
  const response = res.data;
  // El backend devuelve: { data: [...], pagination: {...} }
  const usersList = response.data || response.users || [];

  return {
    items: usersList,
    users: usersList,
    total: response.pagination?.total || 0,
  };
};

// Obtener un usuario por ID (helper)
export const getUserById = async (id) => {
  try {
    const response = await getUsers("", 1, 1000); // obtener muchos usuarios
    // getUsers ahora devuelve un objeto normalizado con items/users
    const list = response.items || response.users || [];

    // Buscar el usuario por ID usando comparación estricta
    const user = list.find((u) => String(u.id) === String(id) || String(u.userid) === String(id));

    return user || null;
  } catch (error) {
    console.error('Error en getUserById:', error);
    throw error;
  }
};

// Crear nuevo usuario
// TODO: Este endpoint aún no está implementado en el backend
export const createUser = async (user) => {
  const res = await api.post("/users", { users: [user] });
  return res.data;
};

// Actualizar usuario existente
export const updateUser = async (user) => {
  const res = await api.put("/users/update", { users: [user] });
  return res.data;
};

// Suspender usuario - cambia el estado a suspended=1 usando DELETE
export const suspendUser = async (userId) => {
  const res = await api.delete("/users", {
    data: { userids: [userId] }
  });
  return res.data;
};

// Activar usuario - cambia el estado a suspended=0 usando PUT
export const activateUser = async (userId) => {
  const res = await api.put("/users/update", {
    users: [{ id: userId, suspended: 0 }]
  });
  return res.data;
};

// Toggle user status - cambia automáticamente el estado (0 <-> 1)
export const toggleUserStatus = async (userId) => {
  const res = await api.put(`/users/${userId}/toggle-status`);
  return res.data;
};

// Eliminar usuario(s) - Suspende el usuario cambiando su estado
export const deleteUsers = async (ids) => {
  const res = await api.delete("/users", {
    data: { userids: ids },
  });
  return res.data;
};


// Restaurar usuarios eliminados (deleted = 0)
// TODO: Este endpoint aún no está implementado en el backend
export const undeleteUsers = async (userids = []) => {
  const res = await api.patch("/users/restore", { userids });
  return res.data;
};

