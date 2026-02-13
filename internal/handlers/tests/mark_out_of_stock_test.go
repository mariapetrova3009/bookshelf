package tests

import (
	"bookshelf/internal/handlers"
	"bookshelf/internal/models"
	"bookshelf/internal/repository"
	repoMocks "bookshelf/internal/repository/mocks"
	"bookshelf/internal/services"
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestMarkOutOfStock_InvalidID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repoMocks.NewMockBooksRepository(ctrl)
	svc := services.NewBooksService(repo)

	r := chi.NewRouter()
	r.Post("/books/{id}/mark-out-of-stock", handlers.MarkOutOfStock(svc))

	w := doReq(r, http.MethodPost, "/books/abc/mark-out-of-stock")
	require.Equal(t, http.StatusBadRequest, w.Code, w.Body.String())
}

func TestMarkOutOfStock_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repoMocks.NewMockBooksRepository(ctrl)
	svc := services.NewBooksService(repo)

	repo.EXPECT().
		MarkOutOfStock(gomock.Any(), uint64(10)).
		Return(models.Book{}, repository.ErrNotFound).
		Times(1)

	r := chi.NewRouter()
	r.Post("/books/{id}/mark-out-of-stock", handlers.MarkOutOfStock(svc))

	w := doReq(r, http.MethodPost, "/books/10/mark-out-of-stock")
	require.Equal(t, http.StatusNotFound, w.Code, w.Body.String())
}

func TestMarkOutOfStock_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repoMocks.NewMockBooksRepository(ctrl)
	svc := services.NewBooksService(repo)

	repo.EXPECT().
		MarkOutOfStock(gomock.Any(), uint64(10)).
		Return(models.Book{Id: 10, OutOfStock: true}, nil).
		Times(1)

	r := chi.NewRouter()
	r.Post("/books/{id}/mark-out-of-stock", handlers.MarkOutOfStock(svc))

	w := doReq(r, http.MethodPost, "/books/10/mark-out-of-stock")
	require.Equal(t, http.StatusOK, w.Code, w.Body.String())
}
