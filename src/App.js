import React, { useEffect } from "react";
import UsersTableView from "./components/UsersTableView";
import ZajunaLayout from "./layout/ZajunaLayout";
import { Routes, Route, useNavigate, useLocation } from "react-router-dom";
import UserFormPage from "./pages/UserFormPage";
import UserEditPage from "./pages/UserEditPage";
import UsersPage from "./pages/UsersPage";
import AccountsPage from "./pages/AccountsPage";
import CoursesPage from "./pages/CoursesPage";
import CoursesManagePage from "./pages/CoursesManagePage";
import CategoryCreatePage from "./pages/CategoryCreatePage";
import CourseCreatePage from "./pages/CourseCreatePage";
import EnrolledUsersPage from "./pages/EnrolledUsersPage";
import StudentDashboard from "./pages/StudentDashboard";

function App() {
  const navigate = useNavigate();
  const location = useLocation();

  useEffect(() => {
    // Capturar parámetros de la URL cuando se redirige desde el Landing
    const params = new URLSearchParams(location.search);
    const token = params.get('token');
    const userStr = params.get('user');
    const isAdmin = params.get('isAdmin');
    const canAccessDashboard = params.get('canAccessDashboard');

    // Si hay token en la URL, guardarlo en localStorage
    if (token && userStr) {
      console.log('🔑 Token recibido desde Landing, guardando en localStorage...');

      localStorage.setItem('zajuna_token', token);
      localStorage.setItem('zajuna_user', userStr);

      if (isAdmin) {
        localStorage.setItem('zajuna_is_admin', isAdmin);
      }

      if (canAccessDashboard) {
        localStorage.setItem('zajuna_can_access_dashboard', canAccessDashboard);
      }

      // Limpiar los parámetros de la URL para que no queden visibles
      const cleanPath = location.pathname;
      navigate(cleanPath, { replace: true });

      console.log('✅ Datos de sesión guardados correctamente');
    }
  }, [location, navigate]);

  return (
    <Routes>
      {/* Nueva ruta para estudiantes - preservando estructura de main */}
      <Route path="/student-dashboard" element={<StudentDashboard />} />

      {/* Rutas de administración - manteniendo estructura original de main */}
      <Route path="/" element={<ZajunaLayout><UsersTableView /></ZajunaLayout>} />
      <Route path="/users" element={<ZajunaLayout><UsersPage /></ZajunaLayout>} />
      <Route path="/users/accounts" element={<ZajunaLayout><AccountsPage /></ZajunaLayout>} />
      <Route path="/users/new" element={<ZajunaLayout><UserFormPage /></ZajunaLayout>} />
      <Route path="/users/:id/edit" element={<ZajunaLayout><UserEditPage /></ZajunaLayout>} />
      <Route path="/courses" element={<ZajunaLayout><CoursesPage /></ZajunaLayout>} />
      <Route path="/courses/manage" element={<ZajunaLayout><CoursesManagePage /></ZajunaLayout>} />
      <Route path="/courses/categories/new" element={<ZajunaLayout><CategoryCreatePage /></ZajunaLayout>} />
      <Route path="/courses/new" element={<ZajunaLayout><CourseCreatePage /></ZajunaLayout>} />
      <Route path="/courses/:id/enrolled" element={<ZajunaLayout><EnrolledUsersPage /></ZajunaLayout>} />
    </Routes>
  );
}

export default App;

