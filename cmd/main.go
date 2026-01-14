package main

import (
	"fmt"
	"log"
	"simple-library-api/internal/config"
	"simple-library-api/internal/database"
	"simple-library-api/internal/handlers"
	"simple-library-api/internal/repository"
	"simple-library-api/internal/server"

	_ "simple-library-api/docs"
)

// @title Simple Library API
// @version 1.0
// @description A sample API to manage books in a library
// @termsOfService http://swagger.io/terms

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host http://127.0.0.1:8080
// @BasePath /

func main() {
	// Configure MySQL Connection
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Open database connection
	db, err := database.NewMySQLDB(cfg.DSN)
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
