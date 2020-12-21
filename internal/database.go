package internal

import (
	"context"
	"fmt"
	"github.com/jackc/pgx"
	"os"
)

var config = pgx.ConnConfig{
	Host:     "localhost",
	Port:     5432,
	Database: "feiqineus",
	User:     "postgres",
	Password: "password",
}

// Database This holds the database connection pool and can acquire connections
type Database struct {
	pool *pgx.ConnPool
}

// Acquire a connection
func (db *Database) Acquire() (*pgx.Conn, error) {
	if db.pool == nil {
		var err error
		db.pool, err = pgx.NewConnPool(pgx.ConnPoolConfig{
			ConnConfig:     config,
			MaxConnections: 10,
		})
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			os.Exit(1)
		}
	}
	return db.pool.AcquireEx(context.Background())
}

// Close the connection pool
func (db *Database) Close() {
	db.pool.Close()
}
