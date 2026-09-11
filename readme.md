# School Management System API

A lightweight, high-performance RESTful API built with **Go (Golang)**, **Echo framework**, **GORM**, and **PostgreSQL**.

---

## 🛠 Features

* **RESTful Architecture:** Expressive routing powered by Echo.
* **Database Management:** Object-Relational Mapping (ORM) and automatic schema migrations with GORM.
* **Data Validation:** Request payload validation via `go-playground/validator`.
* **UUID Primary Keys:** Uses PostgreSQL `uuid-ossp` extension for secure record IDs.
* **CLI Utility Flags:** Built-in command flags for database migration, configuration generation, and service checks.
* **Container Ready:** Dockerfile and Docker Compose included for instant deployment.

---

## 📁 Project Structure

```text
school-api/
├── app/            # Application bootstrappers and background routines
├── configs/        # YAML configurations and loading logic
├── controllers/    # API routes definition
├── db/             # Database initialization and migration models
├── handlers/       # Request logic for endpoints (CRUD operations)
├── models/         # Struct definitions and GORM models
├── server/         # Echo server instance and middleware setup
├── store/          # Shared database pointers and persistence layer
├── docker-compose.yml
├── Dockerfile
├── go.mod
└── main.go