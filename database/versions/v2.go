package versions

import "github.com/SeerUK/go-sql-playground/database"

func init() {
	database.RegisterVersion(&V2{})
}

type V2 struct{}

func (v V2) Migration() string {
	return `
		CREATE DATABASE example (
			id int UNSIGNED NOT NULL AUTO_INCREMENT,
			message varchar(255) NOT NULL,
			last_modified timestamp,

			PRIMARY KEY (id)
		)
	`
}

func (v V2) Number() int {
	return 2
}
