package domain

type Flashcard struct {
	ID              int    `json:"id"`
	EnglishWord     string `json:"english_word"`
	Translation     string `json:"translation"`
	ExampleSentence string `json:"example_sentence"`
}

type FlashcardRepository interface {
	Create(flashcard Flashcard) error
	GetAll() ([]Flashcard, error)
	GetByID(id int) (Flashcard, error)
	Update(flashcard Flashcard) error
	Delete(id int) error
}
