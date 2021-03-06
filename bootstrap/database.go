package bootstrap

import (
	"database/sql"

	"github.com/anishmgoyal/webstart/constants"
)

// GetDatabaseConnection initializes a connection to the database
// and verifies that it works.
func GetDatabaseConnection() *sql.DB {
	db, err := sql.Open("postgres",
		"postgres://"+constants.DatabaseUsername+":"+
			constants.DatabasePassword+"@"+constants.DatabaseHost+"/"+
			constants.DatabaseName+constants.DatabaseExtraArgs)

	if err != nil {
		panic("Failed to connect to DB: " + err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic("Failed to connect to DB: " + err.Error())
	}

	return db
}
