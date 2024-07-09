package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/miraccan00/flashcard/application"
	"github.com/miraccan00/flashcard/infrastructure/database"
	"github.com/miraccan00/flashcard/infrastructure/httpserver"

	_ "github.com/lib/pq"
)

type DbConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	Name     string
}

func main() {
	// Retrieve the value of ENVIRONMENT variable
	env := os.Getenv("ENVIRONMENT")

	// Log the environment for debugging purposes
	log.Printf("Running in environment: %s", env)

	// Load environment-specific configurations based on 'env'
	switch env {
	case "local":
		// Load .env-local configurations
		log.Println("Loading .env-local configurations...")
		// Implement your logic to load .env-local variables
		// Load environment variables from .env-local file
		if err := godotenv.Load(".env-local"); err != nil {
			log.Fatalf("Error loading .env-local file: %v", err)
		}
	case "prod":
		// Load .env-prod configurations
		log.Println("Loading .env-prod configurations...")
		// Implement your logic to load .env-prod variables
		if err := godotenv.Load(".env-prod"); err != nil {
			log.Fatalf("Error loading .env-prod file: %v", err)
		}
	default:
		log.Fatalf("Unsupported environment: %s", env)
	}

	// Parse environment variables
	dbConfig := DbConfig{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Host:     os.Getenv("DB_HOST"),
		Port:     parseInt(os.Getenv("DB_PORT")),
		Name:     os.Getenv("DB_NAME"),
	}

	// Construct database connection string
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
	)

	// Open database connection
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize application components
	flashcardRepo := database.NewFlashcardRepository(db)
	flashcardService := application.NewFlashcardService(flashcardRepo)
	flashcardHandler := httpserver.NewFlashcardHandler(flashcardService)

	// Create HTTP server and start listening
	router := httpserver.NewRouter(flashcardHandler)
	log.Fatal(http.ListenAndServe(":8081", router))
}

// Helper function to convert port string to int
func parseInt(portStr string) int {
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Error parsing DB_PORT: %v", err)
	}
	return port
}
