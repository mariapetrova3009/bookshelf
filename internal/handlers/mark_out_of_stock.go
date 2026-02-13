package handlers

import (
	"bookshelf/internal/repository"
	"bookshelf/internal/services"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func MarkOutOfStock(svc *services.BooksService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			writeError(w, http.StatusBadRequest, "invalid id")
			return
		}
		book, err := svc.MarkOutOfStock(r.Context(), id)
		if err != nil {
			if errors.Is(err, repository.ErrNotFound) {
				writeError(w, http.StatusNotFound, "not found")
				return
			}
			writeError(w, http.StatusInternalServerError, "failed to mark out of stock")
			return
		}
		writeJSON(w, http.StatusOK, book)
	}
}
