package handlers

import (
	"bookshelf/internal/models"
	"bookshelf/internal/repository"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func CreateBook(repo repository.BooksRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.BookRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if errors.Is(err, io.EOF) {
			writeError(w, http.StatusBadRequest, "empty json")
			return
		}

		if err != nil {
			writeError(w, http.StatusBadRequest, "invalid json")
			return
		}

		if req.Author == "" {
			writeError(w, http.StatusBadRequest, "author is required")
			return
		}
		if req.Title == "" {
			writeError(w, http.StatusBadRequest, "title is required")
			return
		}
		if req.Year < 0 {
			writeError(w, http.StatusBadRequest, "year must be >= 0")
			return
		}
		if req.Rating < 0 || req.Rating > 10 {
			writeError(w, http.StatusBadRequest, "rating must be between 0 and 10")
			return
		}

		book, err := repo.CreateBook(r.Context(), req)
		if err != nil {
			writeError(w, http.StatusInternalServerError, "failed to create book")
			return
		}
		writeJSON(w, http.StatusCreated, book)
	}
}
