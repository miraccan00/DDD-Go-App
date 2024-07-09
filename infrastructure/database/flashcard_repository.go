package database

import (
	"database/sql"

	"github.com/miraccan00/flashcard/domain"
)

type FlashcardRepository struct {
	db *sql.DB
}

func NewFlashcardRepository(db *sql.DB) *FlashcardRepository {
	return &FlashcardRepository{db: db}
}

func (r *FlashcardRepository) Create(flashcard domain.Flashcard) error {
	_, err := r.db.Exec("INSERT INTO flashcards.flashcard (english_word, translation, example_sentence) VALUES ($1, $2, $3)",
		flashcard.EnglishWord, flashcard.Translation, flashcard.ExampleSentence)
	return err
}

func (r *FlashcardRepository) GetAll() ([]domain.Flashcard, error) {
	rows, err := r.db.Query("SELECT id, english_word, translation, example_sentence FROM flashcards.flashcard")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var flashcards []domain.Flashcard
	for rows.Next() {
		var flashcard domain.Flashcard
		if err := rows.Scan(&flashcard.ID, &flashcard.EnglishWord, &flashcard.Translation, &flashcard.ExampleSentence); err != nil {
			return nil, err
		}
		flashcards = append(flashcards, flashcard)
	}
	return flashcards, nil
}

func (r *FlashcardRepository) GetByID(id int) (domain.Flashcard, error) {
	var flashcard domain.Flashcard
	err := r.db.QueryRow("SELECT id, english_word, translation, example_sentence FROM flashcards WHERE id = $1", id).Scan(
		&flashcard.ID, &flashcard.EnglishWord, &flashcard.Translation, &flashcard.ExampleSentence)
	return flashcard, err
}

func (r *FlashcardRepository) Update(flashcard domain.Flashcard) error {
	_, err := r.db.Exec("UPDATE flashcards SET english_word = $1, translation = $2, example_sentence = $3 WHERE id = $4",
		flashcard.EnglishWord, flashcard.Translation, flashcard.ExampleSentence, flashcard.ID)
	return err
}

func (r *FlashcardRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM flashcards WHERE id = $1", id)
	return err
}
