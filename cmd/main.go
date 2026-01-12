package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"simple-library-api/internal/database"
	"simple-library-api/internal/handlers"
	"simple-library-api/internal/repository"
	"simple-library-api/internal/server"

	"github.com/go-sql-driver/mysql"
)

func main() {
	// Configure MySQL Connection
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "library_api_go",
		AllowNativePasswords: true,
	}

	// Open database connection
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalf("Error opening db connection: %v", err)
	}
	defer db.Close()

	// Run migrations
	if err := database.Migrate(db); err != nil {
		log.Fatal(err)
	}

	// Test database connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	fmt.Println("Connected to MySQL database!")

	//	init layers (dependency injection)
	bookRepo := repository.NewBookRepository(db)
	bookHandler := handlers.NewBookHandler(bookRepo)

	server := server.NewServer("127.0.0.1:8080", bookHandler)

	fmt.Println("Server is running on port 8080")
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
