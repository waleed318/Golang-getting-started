package driver

import (
	"database/sql"
	"os"

	"github.com/lib/pq"
	"github.com/waleed318/Golang-getting-started/helper"
)

var db *sql.DB

func ConnectDB() *sql.DB {
	pgURL, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	helper.LogFatal(err)

	db, err = sql.Open("postgres", pgURL)
	helper.LogFatal(err)

	err = db.Ping()
	helper.LogFatal(err)

	return db
}
