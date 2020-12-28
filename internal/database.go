package internal

import (
	"fmt"
	"github.com/jackc/pgx"
	"log"
	"os"
	"strconv"
)

// Database This holds the database connection pool and can acquire connections
type Database struct {
	Pool *pgx.ConnPool
}

func getDefaultConfig() pgx.ConnConfig {
	dbPort, err := strconv.Atoi(EnvOrDefault("DATABASE_PORT", "5432"))
	if err != nil {
		dbPort = 5432
	}
	return pgx.ConnConfig{
		Host:     EnvOrDefault("DATABASE_HOST", "localhost"),
		Port:     uint16(dbPort),
		Database: EnvOrDefault("DATABASE_NAME", "feiqineus"),
		User:     EnvOrDefault("DATABASE_USER", "postgres"),
		Password: EnvOrDefault("DATABASE_PASSWORD", "password"),
	}
}

func NewDatabase() *Database {
	var config pgx.ConnConfig
	var err error
	dbUrl, ok := os.LookupEnv("DATABASE_URL")
	if ok {
		config, err = pgx.ParseConnectionString(dbUrl)
		if err != nil {
			log.Printf("Can't parse DATABASE_URL using DEFAULT CONFIG")
			config = getDefaultConfig()
		}
	} else {
		log.Printf("DATABASE_URL not found, using DEFAULT CONFIG")
		config = getDefaultConfig()
	}

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
