package handlers

import (
	"bookshelf/internal/services"
	"net/http"
)

func RecommendBooks(svc *services.BooksService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		books, err := svc.Recommend(r.Context())

		if err != nil {
			writeError(w, http.StatusInternalServerError, "failed to load recommendations")
			return
		}
		writeJSON(w, http.StatusOK, books)
	}
}
