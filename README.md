Zajuna API

API backend desarrollada en Go (Golang) para la gestión de usuarios, cursos y categorías del sistema Zajuna, integrable con Moodle y frontend en React.

Requisitos previos

Antes de comenzar asegúrate de tener instalado:
| Herramienta | Versión recomendada | Verificar instalación |
| ----------- | ------------------- | --------------------- |
| Go          | ≥ 1.22              | `go version`          |
| Git         | ≥ 2.30              | `git --version`       |
| PostgreSQL  | ≥ 14                | `psql --version`      |

Instalación del proyecto
Clonar el repositorio

git clone https://github.com/tu-usuario/zajuna-api.git
cd zajuna-api

Crear archivo de entorno

Copia el ejemplo y ajusta los valores según tu entorno local:
cp .env.example .env

Ejemplo de configuración:
# .env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=tu_password
DB_NAME=zajuna
SERVER_PORT=8081
ENV=development

Dependencias del proyecto

Estas son las dependencias principales que utiliza el proyecto:
| Paquete                                    | Descripción                                         | Versión  |
| ------------------------------------------ | --------------------------------------------------- | -------- |
| **github.com/gin-gonic/gin**               | Framework HTTP rápido y minimalista para APIs REST. | v1.11.0  |
| **github.com/go-playground/validator/v10** | Validador estructural de datos.                     | v10.27.0 |
| **github.com/joho/godotenv**               | Carga variables de entorno desde archivos `.env`.   | v1.5.1   |
| **github.com/lib/pq**                      | Driver oficial PostgreSQL para Go.                  | v1.10.9  |
| **github.com/stretchr/testify**            | Librería para pruebas unitarias.                    | v1.11.1  |
| **gorm.io/driver/postgres**                | Driver PostgreSQL para GORM ORM.                    | v1.6.0   |
| **gorm.io/gorm**                           | ORM (Object Relational Mapper) para Go.             | v1.31.0  |

Instalación de dependencias

Después de clonar el repositorio, ejecuta:
go mod tidy

Esto descargará e instalará todas las dependencias necesarias y limpiará las no utilizadas.

Para verificar que todo esté correcto:
go mod verify

Ejecutar el servidor

Inicia el servidor local con:
go run cmd/main.go

