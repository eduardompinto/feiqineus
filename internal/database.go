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

func NewDatabase() *Database {
	pool, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig:     config,
		MaxConnections: 10,
	})
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return &Database{Pool: pool}
}

// Close the connection pool
func (db *Database) Close() {
	db.Pool.Close()
}
