#!/bin/bash
# Script para levantar todos los servidores del proyecto Zajuna

echo "🚀 Iniciando servidores de Zajuna..."

# Verificar que estamos en el directorio correcto
if [ ! -d "zajuna-api" ] || [ ! -d "Zajuna" ] || [ ! -d "Landing-Lms-Zajuna" ]; then
    echo "❌ Error: Ejecuta este script desde la raíz del proyecto"
    exit 1
fi

# Función para manejar Ctrl+C
cleanup() {
    echo ""
    echo "🛑 Deteniendo servidores..."
    kill 0
    exit
}

trap cleanup SIGINT SIGTERM

# Función para detener procesos en un puerto específico
kill_port() {
    local port=$1
    local pid=$(lsof -ti:$port 2>/dev/null)
    if [ ! -z "$pid" ]; then
        echo "⚠️  Deteniendo proceso existente en puerto $port (PID: $pid)..."
        kill -9 $pid 2>/dev/null
        sleep 1
    fi
}

# Detener procesos existentes en los puertos
echo "🧹 Limpiando puertos..."
kill_port 8080
kill_port 3000
kill_port 5173

echo ""

# Verificar e instalar dependencias si es necesario
echo "📦 Verificando dependencias..."

if [ ! -d "Zajuna/node_modules" ]; then
    echo "⚙️  Instalando dependencias de Zajuna Frontend..."
    cd Zajuna && npm install && cd ..
fi

if [ ! -d "Landing-Lms-Zajuna/node_modules" ]; then
    echo "⚙️  Instalando dependencias de Landing..."
    cd Landing-Lms-Zajuna && npm install && cd ..
fi

echo ""

# 1. Iniciar API (Go)
echo "📡 Iniciando zajuna-api..."
cd zajuna-api

# Verificar si el binario existe, si no, compilarlo
if [ ! -f "zajunaApi" ]; then
    echo "⚙️  Compilando zajuna-api por primera vez..."
    go build -o zajunaApi ./cmd/server
fi

./zajunaApi &
API_PID=$!
cd ..

# Esperar un poco para que la API inicie
sleep 2

# 2. Iniciar Landing (Vite)
echo "🌐 Iniciando Landing page..."
cd Landing-Lms-Zajuna
BROWSER=true npm run dev &
LANDING_PID=$!
cd ..

# 3. Iniciar Frontend (React)
echo "⚛️  Iniciando Zajuna Frontend (React)..."
cd Zajuna
BROWSER=none npm start &
FRONTEND_PID=$!
cd ..

echo ""
echo "✅ Todos los servidores están corriendo:"
echo "   - API: http://localhost:8080"
echo "   - Landing: http://localhost:5173 (Vite)"
echo "   - Frontend: http://localhost:3000 (React)"
echo ""
echo "Presiona Ctrl+C para detener todos los servidores"

# Esperar a que todos los procesos terminen
wait
