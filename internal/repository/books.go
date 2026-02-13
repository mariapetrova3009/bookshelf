package repository

import (
	"bookshelf/internal/models"
	"context"
)

//go:generate mockgen -source=books.go -destination=mocks/books_repository_mock.go -package=mocks
type BooksRepository interface {
	CreateBook(ctx context.Context, in models.BookRequest) (models.Book, error)
	UpdateBook(ctx context.Context, id uint64, in models.BookRequest) (models.Book, error)

	GetBooks(ctx context.Context, limit, offset int) ([]models.Book, error)
	GetDetails(ctx context.Context, id uint64) (models.Book, error)
	DeleteBook(ctx context.Context, id uint64) error

	RecommendBooks(ctx context.Context, limit int) ([]models.Book, error)
	MarkOutOfStock(ctx context.Context, id uint64) (models.Book, error)
}
