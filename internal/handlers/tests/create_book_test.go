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

func TestCreateBook_InvalidJSON(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repoMocks.NewMockBooksRepository(ctrl)

	r := chi.NewRouter()
	r.Post("/books", handlers.CreateBook(repo))

	w := doJSON(r, http.MethodPost, "/books", []byte(`{"title":`))
	require.Equal(t, http.StatusBadRequest, w.Code, w.Body.String())
}

func TestCreateBook_Validation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repoMocks.NewMockBooksRepository(ctrl)

	r := chi.NewRouter()
	r.Post("/books", handlers.CreateBook(repo))

	w := doJSON(r, http.MethodPost, "/books",
		[]byte(`{"title":"","author":"","year":2020,"isbn":"x","out_of_stock":false,"read":false,"rating":5}`),
	)
	require.Equal(t, http.StatusBadRequest, w.Code, w.Body.String())
}

func TestCreateBook_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repoMocks.NewMockBooksRepository(ctrl)
	repo.EXPECT().
		CreateBook(gomock.Any(), gomock.Any()).
		Return(models.Book{Id: 1, Title: "Dune", Author: "Frank Herbert", Year: 1965}, nil).
		Times(1)

	r := chi.NewRouter()
	r.Post("/books", handlers.CreateBook(repo))

	body := []byte(`{"title":"Dune","author":"Frank Herbert","year":1965,"isbn":"9780441172719","out_of_stock":false,"read":true,"rating":9}`)
	w := doJSON(r, http.MethodPost, "/books", body)

	require.Equal(t, http.StatusCreated, w.Code, w.Body.String())
}
