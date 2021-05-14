package config

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// createDBConnection function for creating database connection
func createDBConnection(descriptor string) *sql.DB {
	db, err := sql.Open("postgres", descriptor)
	if err != nil {
		defer db.Close()
		return db
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	return db
}

// LoadPostgres function for creating database connection Postgres Database
func LoadPostgres() *sql.DB {
	// return createDBConnection("postgres://admin:secret@host.docker.internal:5433/panelharga?sslmode=disable")
	return createDBConnection("postgres://user_panelharga:P@n3lH4rg4BKP@10.18.11.21:5432/panelhargabkp?sslmode=disable")
}
