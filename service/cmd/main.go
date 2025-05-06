package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	database "github.com/rcarvalho-pb/shortlinker-go/internal/adapter/db/sqlite"
	"github.com/rcarvalho-pb/shortlinker-go/internal/config/migration"
	"github.com/rcarvalho-pb/shortlinker-go/internal/domain/utils"
)

func main() {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	exePath, err = filepath.EvalSymlinks(exePath)
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Join(filepath.Dir(filepath.Dir(filepath.Dir(exePath))), "data", "link.db")
	dbURI := utils.GetEnv("DB_CONN_STRING", dir)
	db := database.InitDB(dbURI)
	connOpts := &migration.ConnectionOpts{
		DB:      db.Database,
		Dialect: "sqlite",
	}
	migration.MigrateUP(connOpts)
}
