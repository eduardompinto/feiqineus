package internal

import (
	"fmt"
	"github.com/jackc/pgx"
	"os"
)

var config = pgx.ConnConfig{
	Host:     "database",
	Port:     5432,
	Database: "feiqineus",
	User:     "postgres",
	Password: "password",
}

// Database This holds the database connection pool and can acquire connections
type Database struct {
	Pool *pgx.ConnPool
}

// Acquire a connection
func (db *Database) Initialize() {
	if db.Pool == nil {
		var err error
		db.Pool, err = pgx.NewConnPool(pgx.ConnPoolConfig{
			ConnConfig:     config,
			MaxConnections: 10,
		})
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			os.Exit(1)
		}
	}
}

// Close the connection pool
func (db *Database) Close() {
	db.Pool.Close()
}
