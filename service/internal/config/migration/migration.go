package migration

import (
	"database/sql"
	"embed"

	"github.com/pressly/goose/v3"
)

//go:embed files/*.sql
var migrations embed.FS

type ConnectionOpts struct {
	DB      *sql.DB
	Dialect string
}

func MigrateUP(connOpts *ConnectionOpts) error {
	goose.SetBaseFS(migrations)
	if err := goose.SetDialect(connOpts.Dialect); err != nil {
		return err
	}
	if err := goose.Up(connOpts.DB, "files"); err != nil {
		return err
	}
	return nil
}
