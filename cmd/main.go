package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/miraccan00/flashcard/application"
	"github.com/miraccan00/flashcard/infrastructure/database"
	"github.com/miraccan00/flashcard/infrastructure/httpserver"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://myuser:mypassword@postgres:5432/flashcards_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	flashcardRepo := database.NewFlashcardRepository(db)
	flashcardService := application.NewFlashcardService(flashcardRepo)
	flashcardHandler := httpserver.NewFlashcardHandler(flashcardService)

	router := httpserver.NewRouter(flashcardHandler)
	log.Fatal(http.ListenAndServe(":8081", router))
}
