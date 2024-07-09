package httpserver

import (
	"github.com/gorilla/mux"
)

func NewRouter(handler *FlashcardHandler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/flashcards", handler.CreateFlashcard).Methods("POST")
	router.HandleFunc("/flashcards", handler.GetAllFlashcards).Methods("GET")
	router.HandleFunc("/flashcards/{id:[0-9]+}", handler.GetFlashcardByID).Methods("GET")
	router.HandleFunc("/flashcards/{id:[0-9]+}", handler.UpdateFlashcard).Methods("PUT")
	router.HandleFunc("/flashcards/{id:[0-9]+}", handler.DeleteFlashcard).Methods("DELETE")

	return router
}
