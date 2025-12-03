import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import './styles/index.css';

import Index from './components/pages/Index';
import Bilinguismo from './components/pages/Bilinguismo';
import Campesena from './components/pages/Campesena';
import Titulada from './components/pages/Titulada';
import Soporte from './components/pages/Soporte';

function App() {
  return (
    <BrowserRouter
      future={{
        v7_startTransition: true,
        v7_relativeSplatPath: true
      }}
    >
      <Routes>
        <Route path="/" element={<Index />} />
        <Route path="/bilinguismo" element={<Bilinguismo />} />
        <Route path="/campesena" element={<Campesena />} />
        <Route path="/titulada" element={<Titulada />} />
        <Route path="/soporte" element={<Soporte />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
