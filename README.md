# 🎓 Student Management API with Go and Gin (Alura)

This repository contains the project of a **RESTful API** developed in **Go** (Golang) using the **Gin** framework for routing and **GORM** as the ORM for data persistence in a **PostgreSQL** database.

The project is structured to demonstrate the creation of a complete application, focusing on:
*   **Development Best Practices:** Separation of concerns (Controllers, Models, Database, Routes).
*   **Data Validation:** Use of struct tags and the `gopkg.in/validator.v2` library to ensure data integrity.
*   **Containers:** Complete setup with `Dockerfile` and `docker-compose.yml` for an isolated development and production environment.

## 🚀 Features

The API enables full management of student records (CRUD - Create, Read, Update, Delete).

| Route | Method | Description | Controller |
| :--- | :--- | :--- | :--- |
| `/` | `GET` | Returns the API status. | N/A |
| `/ping` | `GET` | Simple health check (returns "pong"). | N/A |
| `/alunos` | `GET` | Lists all registered students. | `TodosAlunos` |
| `/alunos` | `POST` | Creates a new student. Requires validation of `Nome`, `RG` (9 digits), and `CPF` (11 digits). | `CriarNovoAluno` |
| `/alunos/:id` | `GET` | Retrieves a student by ID. | `BuscarAlunoPorID` |
| `/alunos/:id` | `PATCH` | Updates a student's data by ID. | `EditarAluno` |
| `/alunos/:id` | `DELETE` | Deletes a student by ID. | `DeletarAluno` |
| `/alunos/cpf/:cpf` | `GET` | Retrieves a student by CPF number. | `BuscaAlunoPorCPF` |
| `/alunos/saudacao/:nome` | `GET` | Example route that returns a personalized greeting. | `Saudacoes` |
| `/index` | `GET` | Displays a simple HTML page with the list of students (View). | `ExibePaginaIndex` |

## 🛠️ Technologies Used

*   **Language:** Go (Golang)
*   **Web Framework:** [Gin Gonic](https://github.com/gin-gonic/gin)
*   **ORM:** [GORM](https://gorm.io/)
*   **Database:** PostgreSQL
*   **Containers:** Docker and Docker Compose

## ⚙️ Environment Setup

The project uses **Docker Compose** to orchestrate the Go application and the PostgreSQL database, simplifying environment setup.

### Prerequisites

Make sure you have [Docker](https://www.docker.com/get-started) and [Docker Compose](https://docs.docker.com/compose/install/) installed on your machine.

### 1. Clone the Repository

```bash
git clone https://github.com/guilhermeonrails/api-go-gin.git
cd api-go-gin
```

### 2. Start the Application

Run the following command to build the images and start the containers:

```bash
docker-compose up --build
```

Docker Compose will:
1.  Build the Go application image (`app`) using the `Dockerfile`.
2.  Start the PostgreSQL container (`postgres`), waiting until it is healthy.
3.  Start the Go application container, which will connect to the database and run migrations (creating the `alunos` table).

The application will be accessible at `http://localhost:8080`.

### 3. Environment Variables

The database connection is configured through environment variables defined in `docker-compose.yml` and read by the `database/db.go` file.

| Variable | Default Value | Description |
| :--- | :--- | :--- |
| `DB_HOST` | `postgres` | Database service name in Docker Compose. |
| `DB_USER` | `root` | PostgreSQL user. |
| `DB_PASSWORD` | `root` | PostgreSQL password. |
| `DB_NAME` | `root` | Database name. |
| `DB_PORT` | `5432` | PostgreSQL port. |

## 💻 Running Locally (Without Docker)

If you prefer to run the application directly on your machine, follow these steps:

### Prerequisites

*   [Go (version 1.24 or higher)](https://golang.org/dl/)
*   A PostgreSQL server running locally.

### 1. Configure the Database

Create a PostgreSQL database and set the required environment variables for the connection (replace with your own credentials):

```bash
export DB_HOST=localhost
export DB_USER=your_username
export DB_PASSWORD=your_password
export DB_NAME=your_database
export DB_PORT=5432
```

### 2. Install Dependencies and Run

```bash
go mod tidy
go run main.go
```

The application will start on port `8080`.

## 📝 API Usage Example (POST)

To create a new student, send a `POST` request to the `/alunos` route with a JSON body in the following format:

```json
{
    "nome": "João da Silva",
    "rg": "123456789",
    "cpf": "12345678901"
}
```

**Example using `curl`:**

```bash
curl -X POST http://localhost:8080/alunos \
-H "Content-Type: application/json" \
-d '{"nome": "João da Silva", "rg": "123456789", "cpf": "12345678901"}'
```