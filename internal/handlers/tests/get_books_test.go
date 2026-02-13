package tests

import (
	"bookshelf/internal/handlers"
	"bookshelf/internal/models"
	repoMocks "bookshelf/internal/repository/mocks"
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetBooks_InvalidPage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repoMocks.NewMockBooksRepository(ctrl)

	r := chi.NewRouter()
	r.Get("/books", handlers.GetBooks(repo))

	w := doReq(r, http.MethodGet, "/books?page=0")
	require.Equal(t, http.StatusBadRequest, w.Code, w.Body.String())
}

func TestGetBooks_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repoMocks.NewMockBooksRepository(ctrl)

	repo.EXPECT().
		GetBooks(gomock.Any(), 10, 10).
		Return([]models.Book{}, nil).
		Times(1)

	r := chi.NewRouter()
	r.Get("/books", handlers.GetBooks(repo))

	w := doReq(r, http.MethodGet, "/books?page=2&limit=10")
	require.Equal(t, http.StatusOK, w.Code, w.Body.String())
}
