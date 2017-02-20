package versions

import "github.com/SeerUK/go-sql-playground/database"

func init() {
	database.RegisterVersion(1, `
		CREATE DATABASE example (
			id int UNSIGNED NOT NULL AUTO_INCREMENT,
			message varchar(255) NOT NULL,
			last_modified timestamp,

			PRIMARY KEY (id)
		);
	`)
}
