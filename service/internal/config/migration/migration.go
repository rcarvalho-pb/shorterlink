package migration

import (
	"database/sql"
	"embed"

	"github.com/pressly/goose/v3"
)

//go:embed files/*.sql
var migrations embed.FS

func MigrateUP(db *sql.DB, dialect string) error {
	goose.SetBaseFS(migrations)
	if err := goose.SetDialect(dialect); err != nil {
		return err
	}
	if err := goose.Up(db, "files"); err != nil {
		return err
	}
	return nil
}
