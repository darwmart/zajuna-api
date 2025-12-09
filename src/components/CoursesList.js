import React, { useEffect, useState } from "react";
import {
  Container,
  Typography,
  List,
  ListItem,
  ListItemText,
  CircularProgress,
  Alert,
  Button,
  TextField,
  Dialog,
  DialogTitle,
  DialogContent,
} from "@mui/material";
import { coursesService } from "../services/coursesService";

export default function CoursesList() {
  const [courses, setCourses] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const [selectedCourse, setSelectedCourse] = useState(null);
  const [enrolledUsers, setEnrolledUsers] = useState([]);

  const [userId, setUserId] = useState("");
  const [userCourses, setUserCourses] = useState([]);

  // Cargar cursos
  useEffect(() => {
    const loadCourses = async () => {
      try {
        const data = await coursesService.getAll();
        setCourses(data.items);
      } catch {
        setError("No se pudieron cargar los cursos");
      } finally {
        setLoading(false);
      }
    };
    loadCourses();
  }, []);

  // Ver usuarios matriculados
  const handleShowEnrolled = async (course) => {
    try {
      setSelectedCourse(course);
      const data = await coursesService.getEnrolledUsers(course.id);
      setEnrolledUsers(data.items || []);
    } catch {
      alert("Error cargando usuarios del curso");
    }
  };

  // Consultar cursos de un usuario
  const handleSearchUserCourses = async () => {
    if (!userId) return;
    try {
      const data = await coursesService.getUserCourses(userId);
      setUserCourses(data.items || []);
    } catch {
      alert("Error cargando cursos del usuario");
    }
  };

  return (
    <Container maxWidth="md" sx={{ mt: 4 }}>
      <Typography variant="h4" gutterBottom>
        Cursos disponibles
      </Typography>

      {loading && <CircularProgress />}
      {error && <Alert severity="error">{error}</Alert>}

      {!loading && !error && (
        <List>
          {courses.map((c) => (
            <ListItem
              key={c.id}
              divider
              secondaryAction={
                <Button
                  variant="contained"
                  onClick={() => handleShowEnrolled(c)}
                >
                  Ver inscritos
                </Button>
              }
            >
              <ListItemText
                primary={`${c.fullname}`}
                secondary={`ID: ${c.id} | Visible: ${c.visible ? "Sí" : "No"}`}
              />
            </ListItem>
          ))}
        </List>
      )}

      {/* Diálogo con usuarios matriculados */}
      <Dialog open={!!selectedCourse} onClose={() => setSelectedCourse(null)}>
        <DialogTitle>
          Usuarios matriculados en: {selectedCourse?.fullname}
        </DialogTitle>
        <DialogContent>
          {enrolledUsers.length === 0 ? (
            <Typography variant="body2">No hay usuarios matriculados</Typography>
          ) : (
            <List>
              {enrolledUsers.map((u) => (
                <ListItem key={u.id}>
                  <ListItemText
                    primary={`${u.firstname} ${u.lastname}`}
                    secondary={u.email}
                  />
                </ListItem>
              ))}
            </List>
          )}
        </DialogContent>
      </Dialog>

      {/* Consultar cursos por usuario */}
      <Typography variant="h5" sx={{ mt: 4 }}>
        Buscar cursos por usuario (ID)
      </Typography>
      <TextField
        label="ID del usuario"
        value={userId}
        onChange={(e) => setUserId(e.target.value)}
        sx={{ mr: 2 }}
      />
      <Button variant="outlined" onClick={handleSearchUserCourses}>
        Buscar
      </Button>

      {userCourses.length > 0 && (
        <List sx={{ mt: 2 }}>
          {userCourses.map((uc) => (
            <ListItem key={uc.id}>
              <ListItemText
                primary={uc.fullname}
                secondary={`Rol: ${uc.roles}`}
              />
            </ListItem>
          ))}
        </List>
      )}
    </Container>
  );
}
