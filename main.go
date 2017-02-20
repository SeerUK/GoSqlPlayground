package main

import (
	"database/sql"
	"log"

	"github.com/SeerUK/go-sql-playground/database"
	_ "github.com/SeerUK/go-sql-playground/database/versions"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.Println("==> Connecting to MySQL...")

	db, err := sql.Open("mysql", "root:gsp@tcp(127.0.0.1:3306)/gsp")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	database.Migrate(db)
}
