package routes

import (
	"bookshelf/internal/handlers"
	"bookshelf/internal/repository"
	"bookshelf/internal/services"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(repo repository.BooksRepository) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.RequestID) // для трейсинга
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	svc := services.NewBooksService(repo)
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
	router.Route("/books", func(r chi.Router) {

		r.Get("/recommend", handlers.RecommendBooks(svc))
		r.Post("", handlers.CreateBook(repo))
		r.Get("", handlers.GetBooks(repo))
		r.Get("/{id}", handlers.GetDetails(repo))
		r.Put("/{id}", handlers.UpdateBook(repo))
		r.Delete("/{id}", handlers.DeleteBook(repo))

		r.Post("/{id}/mark-out-of-stock", handlers.MarkOutOfStock(svc))
	})
	return router
}

func notImplemented(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "not implemented", http.StatusNotImplemented)
}
