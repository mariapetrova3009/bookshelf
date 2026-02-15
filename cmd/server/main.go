package main

import (
	"bookshelf/internal/config"
	"bookshelf/internal/migrations"
	"bookshelf/internal/repository"
	"bookshelf/internal/routes"
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	cfg := config.MustLoad()

	db, err := sql.Open("pgx", cfg.DSN())
	if err != nil {
		log.Fatal("failed to connect to storage: ", err)
	}
	defer db.Close()

	var migErr error
	for i := 0; i < 30; i++ {
		migErr = migrations.Run("migrations", cfg.MigrateDSN())
		if migErr == nil {
			break
		}
		log.Println("migrations not ready yet:", migErr)
		time.Sleep(1 * time.Second)
	}
	if migErr != nil {
		log.Fatal("failed to run migrations: ", migErr)
	}

	repo := repository.NewRepo(db)
	router := routes.NewRouter(repo)

	srv := http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: router,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("listen: ", err)
	}
}
