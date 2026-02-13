package migrations

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Run(path string, dsn string) error {
	m, err := migrate.New("file://"+path, dsn)
	if err != nil {
		return fmt.Errorf("migrate.New: %w", err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			return nil
		}
		return fmt.Errorf("migrate.Up: %w", err)
	}
	return nil
}
