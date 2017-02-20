package database

import (
	"database/sql"
	"fmt"
	"log"
	"sort"
)

var versions []int
var migrations map[int]string = make(map[int]string)

func RegisterVersion(versionNo int, migration string) {
	versions = append(versions, versionNo)
	migrations[versionNo] = migration
}

// @todo: Needs sql.DB
func Migrate(db *sql.DB) {
	sort.Ints(versions)

	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// @todo: Some kind of config...
	table := "migration_versions"

	if !hasTable(db, "gsp", table) {
		log.Println("==> Migration versions table doesn't exist! Creating...")

		createMigrationVerionsTable(db, table)
	}

	versionNo, err := getMigrationVersion(db, table)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(versionNo)

	// @todo: Check status of migrations table.
	// @todo: 1. Table exists.
	// @todo: 2. Current version.

	for _, migration := range migrations {
		fmt.Println(migration)
	}
}

func hasTable(db *sql.DB, schema string, table string) bool {
	stmt, err := db.Prepare(`
		SELECT
			COUNT(1) AS hasTable
		FROM
			information_schema.tables
		WHERE
			table_schema = ?
		AND
			table_name = ?
		LIMIT 1
	`)

	if err != nil {
		log.Fatal(err)
	}

	var hasTable bool

	err = stmt.QueryRow(schema, table).Scan(&hasTable)
	if err != nil {
		log.Fatal(err)
	}

	return hasTable
}

func createMigrationVerionsTable(db *sql.DB, table string) error {
	format := `
		CREATE TABLE %s (
			version INT UNSIGNED NOT NULL,
			last_modified TIMESTAMP,

			PRIMARY KEY (version)
		) COMMENT="The migration versions table.";
	`

	_, err := db.Exec(fmt.Sprintf(format, table))

	return err
}

func getMigrationVersion(db *sql.DB, table string) (int, error) {
	var version int

	format := "SELECT version FROM %s ORDER BY version DESC LIMIT 1;"
	err := db.QueryRow(fmt.Sprintf(format, table)).Scan(&version)

	if err == sql.ErrNoRows {
		return 0, nil
	}

	return version, err
}
