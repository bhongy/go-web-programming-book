/*
Package `data` interfaces the underlying SQL database.

No code should interact with the database directly
without going through this package.
*/
package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// make "postgres" driver available for `sql.Open`
	_ "github.com/lib/pq"
)

// Db is the singleton SQL db connection
var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	))
	if err != nil {
		log.Fatal(err)
	}

	err = Db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
