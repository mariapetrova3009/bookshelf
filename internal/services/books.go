package services

import (
	"bookshelf/internal/models"
	"bookshelf/internal/repository"
	"context"
)

const recommendLimit = 5

type BooksService struct {
	repository.BooksRepository
}

func NewBooksService(repo repository.BooksRepository) *BooksService {
	return &BooksService{repo}
}

func (s *BooksService) Recommend(ctx context.Context) ([]models.Book, error) {
	res, err := s.BooksRepository.RecommendBooks(ctx, recommendLimit)
	return res, err
}

func (s *BooksService) MarkOutOfStock(ctx context.Context, id uint64) (models.Book, error) {
	res, err := s.BooksRepository.MarkOutOfStock(ctx, id)
	return res, err
}
