package tests

import (
	"bookshelf/internal/handlers"
	"bookshelf/internal/models"
	"bookshelf/internal/repository"
	repoMocks "bookshelf/internal/repository/mocks"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetDetails_InvalidID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repoMocks.NewMockBooksRepository(ctrl)

	r := chi.NewRouter()
	r.Get("/books/{id}", handlers.GetDetails(repo))

	w := doReq(r, http.MethodGet, "/books/a")
	require.Equal(t, http.StatusBadRequest, w.Code, w.Body.String())
}

func TestGetDetails_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repoMocks.NewMockBooksRepository(ctrl)
	repo.EXPECT().
		GetDetails(gomock.Any(), uint64(10)).
		Return(models.Book{}, repository.ErrNotFound).
		Times(1)

	r := chi.NewRouter()
	r.Get("/books/{id}", handlers.GetDetails(repo))

	w := doReq(r, http.MethodGet, "/books/10")
	require.Equal(t, http.StatusNotFound, w.Code, w.Body.String())
}

func TestGetDetails_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repoMocks.NewMockBooksRepository(ctrl)

	expected := models.Book{
		Id:     10,
		Title:  "Dune",
		Author: "Frank Herbert",
		Year:   1965,
	}

	repo.EXPECT().
		GetDetails(gomock.Any(), uint64(10)).
		Return(expected, nil).
		Times(1)

	r := chi.NewRouter()
	r.Get("/books/{id}", handlers.GetDetails(repo))

	w := doReq(r, http.MethodGet, "/books/10")
	require.Equal(t, http.StatusOK, w.Code, w.Body.String())

	var got models.Book
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &got), w.Body.String())

	require.Equal(t, expected.Id, got.Id)
	require.Equal(t, expected.Title, got.Title)
	require.Equal(t, expected.Author, got.Author)
	require.Equal(t, expected.Year, got.Year)
}
