package application

import "github.com/miraccan00/flashcard/domain"

type FlashcardService struct {
	repo domain.FlashcardRepository
}

func NewFlashcardService(repo domain.FlashcardRepository) *FlashcardService {
	return &FlashcardService{repo: repo}
}

func (s *FlashcardService) CreateFlashcard(flashcard domain.Flashcard) error {
	return s.repo.Create(flashcard)
}

func (s *FlashcardService) GetAllFlashcards() ([]domain.Flashcard, error) {
	return s.repo.GetAll()
}

func (s *FlashcardService) GetFlashcardByID(id int) (domain.Flashcard, error) {
	return s.repo.GetByID(id)
}

func (s *FlashcardService) UpdateFlashcard(flashcard domain.Flashcard) error {
	return s.repo.Update(flashcard)
}

func (s *FlashcardService) DeleteFlashcard(id int) error {
	return s.repo.Delete(id)
}
