package handlers

import (
	"bookshelf/internal/repository"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func GetDetails(repo repository.BooksRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			writeError(w, http.StatusBadRequest, "invalid id")
			return
		}

		book, err := repo.GetDetails(r.Context(), id)
		if err != nil {
			if errors.Is(err, repository.ErrNotFound) {
				writeError(w, http.StatusNotFound, "book not found")
				return
			}
			writeError(w, http.StatusInternalServerError, "failed to get book")
			return
		}
		writeJSON(w, http.StatusOK, book)
	}
}
