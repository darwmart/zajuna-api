import React, { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import UserForm from "../components/UserForm";
import { getUserById } from "../services/usersService";

export default function UserEditPage() {
  const navigate = useNavigate();
  const { id } = useParams();
  const [user, setUser] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    let active = true;
    (async () => {
      try {
        console.log('Cargando usuario con id:', id);
        const u = await getUserById(id);
        console.log('Usuario recibido:', u);
        if (active) {
          if (u) {
            setUser(u);
          } else {
            setError("No se encontró el usuario con ID: " + id);
          }
        }
      } catch (e) {
        console.error('Error al cargar usuario:', e);
        if (active) setError("No se pudo cargar el usuario: " + (e.message || 'error desconocido'));
      } finally {
        if (active) setLoading(false);
      }
    })();
    return () => { active = false; };
  }, [id]);

  if (loading) return <div>Cargando...</div>;
  if (error) return <div style={{ color: '#b00020' }}>{error}</div>;
  if (!user) return <div>No se encontró el usuario.</div>;

  return (
    <div>
      <UserForm
        user={user}
        onCancel={() => navigate("/")}
        onSuccess={() => navigate("/")}
      />
    </div>
  );
}
