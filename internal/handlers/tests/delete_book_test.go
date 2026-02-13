package tests

import (
	"bookshelf/internal/handlers"
	"bookshelf/internal/repository"
	repoMocks "bookshelf/internal/repository/mocks"
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestDeleteBook_InvalidID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repoMocks.NewMockBooksRepository(ctrl)

	r := chi.NewRouter()
	r.Delete("/books/{id}", handlers.DeleteBook(repo))

	w := doReq(r, http.MethodDelete, "/books/abc")
	require.Equal(t, http.StatusBadRequest, w.Code, w.Body.String())
}

func TestDeleteBook_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := repoMocks.NewMockBooksRepository(ctrl)
	repo.EXPECT().
		DeleteBook(gomock.Any(), uint64(10)).
		Return(repository.ErrNotFound).
		Times(1)

	r := chi.NewRouter()
	r.Delete("/books/{id}", handlers.DeleteBook(repo))

	w := doReq(r, http.MethodDelete, "/books/10")
	require.Equal(t, http.StatusNotFound, w.Code, w.Body.String())
}
