package main

import database "github.com/rcarvalho-pb/shortlinker-go/internal/adapter/db/sqlite"

func main() {
	_ = database.InitDB("teste")
}
