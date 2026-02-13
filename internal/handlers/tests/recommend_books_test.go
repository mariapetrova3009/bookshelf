package tests

import (
	"bookshelf/internal/handlers"
	"bookshelf/internal/models"
	repoMocks "bookshelf/internal/repository/mocks"
	"bookshelf/internal/services"
	"errors"
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRecommendBooks_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repoMocks.NewMockBooksRepository(ctrl)
	svc := services.NewBooksService(repo)

	repo.EXPECT().
		RecommendBooks(gomock.Any(), 5).
		Return([]models.Book{}, nil).
		Times(1)

	r := chi.NewRouter()
	r.Get("/books/recommend", handlers.RecommendBooks(svc))

	w := doReq(r, http.MethodGet, "/books/recommend")
	require.Equal(t, http.StatusOK, w.Code, w.Body.String())
}

func TestRecommendBooks_RepoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repoMocks.NewMockBooksRepository(ctrl)
	svc := services.NewBooksService(repo)

	repo.EXPECT().
		RecommendBooks(gomock.Any(), 5).
		Return(nil, errors.New("db error")).
		Times(1)

	r := chi.NewRouter()
	r.Get("/books/recommend", handlers.RecommendBooks(svc))

	w := doReq(r, http.MethodGet, "/books/recommend")
	require.Equal(t, http.StatusInternalServerError, w.Code, w.Body.String())
}
