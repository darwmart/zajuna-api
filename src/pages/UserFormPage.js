import React from "react";
import { useNavigate } from "react-router-dom";
import UserForm from "../components/UserForm";

export default function UserFormPage() {
  const navigate = useNavigate();
  return (
    <div>
      <UserForm
        onCancel={() => navigate("/")}
        onSuccess={() => navigate("/")}
      />
    </div>
  );
}
