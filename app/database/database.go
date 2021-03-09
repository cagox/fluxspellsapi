package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/cagox/fluxspellsapi/app"
)

func OpenDatabase() {
	db, err := sql.Open("mysql", databaseURI())
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	app.Config.Database = db
}

func CloseDatabase() {
	app.Config.Database.Close() //TODO: Add error handling.
}

//Here we just build the URI and return it. By using a function, we can change it in one place later.
func databaseURI() string {
	uriString := app.Config.DatabaseUserName + ":" + app.Config.DatabasePassword + "@/" + app.Config.DatabaseName + "?" + app.Config.DatabaseOptions
	return uriString
}
