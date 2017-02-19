package main

import (
	"fmt"

	"github.com/SeerUK/go-sql-playground/database"
	_ "github.com/SeerUK/go-sql-playground/database/versions"
)

func main() {
	fmt.Println("Hello, World!")

	database.Migrate()
}
