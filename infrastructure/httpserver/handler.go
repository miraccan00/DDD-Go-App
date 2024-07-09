package httpserver

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/miraccan00/flashcard/application"
	"github.com/miraccan00/flashcard/domain"

	"github.com/gorilla/mux"
)

type FlashcardHandler struct {
	service *application.FlashcardService
}

func NewFlashcardHandler(service *application.FlashcardService) *FlashcardHandler {
	return &FlashcardHandler{service: service}
}

func (h *FlashcardHandler) CreateFlashcard(w http.ResponseWriter, r *http.Request) {
	var flashcard domain.Flashcard
	if err := json.NewDecoder(r.Body).Decode(&flashcard); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.CreateFlashcard(flashcard); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *FlashcardHandler) GetAllFlashcards(w http.ResponseWriter, r *http.Request) {
	flashcards, err := h.service.GetAllFlashcards()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(flashcards)
}

func (h *FlashcardHandler) GetFlashcardByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	flashcard, err := h.service.GetFlashcardByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(flashcard)
}

func (h *FlashcardHandler) UpdateFlashcard(w http.ResponseWriter, r *http.Request) {
	var flashcard domain.Flashcard
	if err := json.NewDecoder(r.Body).Decode(&flashcard); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateFlashcard(flashcard); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *FlashcardHandler) DeleteFlashcard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteFlashcard(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
