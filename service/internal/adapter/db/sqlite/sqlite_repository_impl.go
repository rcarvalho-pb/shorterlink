package database

import (
	"database/sql"

	"github.com/rcarvalho-pb/shortlinker-go/internal/domain/entity"
)

type DB struct {
	Database *sql.DB
}

func InitDB(dbURI string) *DB {
	return &DB{}
}

func (db *DB) Add(link *entity.Link) error {
	return nil
}

func (db *DB) Get(slug string) (*entity.Link, error) {
	return nil, nil
}
