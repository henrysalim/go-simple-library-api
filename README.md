# Simple Library API 📚

A powerful, robust, and clean RESTful API for managing a library system, built with Go and MySQL.

This project demonstrates **Standard Go Project Layout** and modern **Best Practices** including Dependency Injection,
the Repository Pattern, and Go 1.22+ routing.

## 🚀 Features

* **CRUD Operations**: Create, Read, Update, and Delete books.
* **Clean Architecture**: Separation of concerns (Handlers vs. Repositories vs. Models).
* **Dependency Injection**: Modular and testable code structure.
* **Modern Routing**: Uses Go 1.22's `net/http` path value matching (no external router needed).
* **MySQL Integration**: robust data persistence using `go-sql-driver/mysql`.

## 🏃🏻Run the Project

To run the project, simply run the following commands:

### 1. Download/Sync all dependencies

`go mod tidy`

### 2. Run the main file from the root

`go run cmd/main.go`

Then the server will run on http://127.0.0.1:8080

To access the API documentation: http://127.0.0.1:8080/swagger

## 🛠️ Tech Stack

* **Language**: Go (Golang)
* **Database**: MySQL
* **Standard Lib**: `net/http`, `database/sql`, `encoding/json`, `joho/godotenv`

## 📂 Project Structure

```text
SimpleLibraryAPI/
├── cmd/
│   └── main.go           # Application entry point & wiring
├── internal/
│   ├── config/           # Configuration logic for database connection
│   ├── database/         # DB connection & migrations
│   ├── handlers/         # HTTP Controllers (Requests/Responses)
│   ├── model/            # Data structures (structs)
│   ├── repository/       # Database logic (SQL queries)
│   └── server/           # HTTP Server configuration
├── .env.example          # Example of .env file 
└── go.mod                # Go module definition 
