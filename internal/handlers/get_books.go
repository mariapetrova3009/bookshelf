package handlers

import (
	"bookshelf/internal/repository"
	"net/http"
	"strconv"
)

func GetBooks(repo repository.BooksRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		page := 1
		limit := 10

		if val := r.URL.Query().Get("page"); val != "" {
			p, err := strconv.Atoi(val)
			if err != nil || p < 1 {
				writeError(w, http.StatusBadRequest, "invalid page")
				return
			}
			page = p
		}

		if val := r.URL.Query().Get("limit"); val != "" {
			l, err := strconv.Atoi(val)
			if err != nil || l < 1 {
				writeError(w, http.StatusBadRequest, "invalid limit")
				return
			}
			limit = l
			if limit > 100 {
				limit = 100
			}

		}

		offset := (page - 1) * limit

		books, err := repo.GetBooks(r.Context(), limit, offset)
		if err != nil {
			writeError(w, http.StatusInternalServerError, "failed to get books")
			return
		}
		writeJSON(w, http.StatusOK, books)
	}
}
