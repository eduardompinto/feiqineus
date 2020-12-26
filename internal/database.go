package internal

import (
	"fmt"
	"github.com/jackc/pgx"
	"os"
)

func getDBHost() string {
	h, ok := os.LookupEnv("DATABASE_HOST")
	if ok {
		return h
	}
	return "localhost"
}

var config = pgx.ConnConfig{
	Host:     getDBHost(),
	Port:     5432,
	Database: "feiqineus",
	User:     "postgres",
	Password: "password",
}

// Database This holds the database connection pool and can acquire connections
type Database struct {
	Pool *pgx.ConnPool
}

// Initialize the db connection
func (db *Database) Initialize() {
	if db.Pool != nil {
		return
	}
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

// Close the connection pool
func (db *Database) Close() {
	db.Pool.Close()
}
