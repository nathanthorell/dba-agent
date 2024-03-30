package db

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

func NewDBConnection(connectionString string) (*sql.DB, error) {
	// Open a connection to the database using the provided connection string
	db, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		return nil, err
	}

	// Check if the connection is successful
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
