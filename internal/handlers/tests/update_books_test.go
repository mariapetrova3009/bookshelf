package tests

import (
	"bookshelf/internal/handlers"
	"bookshelf/internal/models"
	"bookshelf/internal/repository"
	repoMocks "bookshelf/internal/repository/mocks"
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestUpdateBook_InvalidID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repoMocks.NewMockBooksRepository(ctrl)

	r := chi.NewRouter()
	r.Put("/books/{id}", handlers.UpdateBook(repo))

	body := []byte(`{"title":"Dune","author":"Frank Herbert","year":1965,"isbn":"x","out_of_stock":false,"read":true,"rating":9}`)
	w := doJSON(r, http.MethodPut, "/books/abc", body)

	require.Equal(t, http.StatusBadRequest, w.Code, w.Body.String())
}

func TestUpdateBook_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repoMocks.NewMockBooksRepository(ctrl)
	repo.EXPECT().
		UpdateBook(gomock.Any(), uint64(10), gomock.Any()).
		Return(models.Book{}, repository.ErrNotFound).
		Times(1)

	r := chi.NewRouter()
	r.Put("/books/{id}", handlers.UpdateBook(repo))

	body := []byte(`{"title":"Dune","author":"Frank Herbert","year":1965,"isbn":"x","out_of_stock":false,"read":true,"rating":9}`)
	w := doJSON(r, http.MethodPut, "/books/10", body)

	require.Equal(t, http.StatusNotFound, w.Code, w.Body.String())
}
