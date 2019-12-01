package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
)

// Pool of pgx connections
var Pool *pgxpool.Pool

func getConnectionString() string {
	username := os.Getenv("DATABASE_USERNAME")
	password := os.Getenv("DATABASE_PASSWORD")
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	database := os.Getenv("DATABASE_NAME")
	return "postgresql://" + username + ":" + password + "@" + host + ":" + port + "/" + database
}

// GetPool returns pgxpool of connections
func GetPool() (*pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(context.Background(), getConnectionString())
	if err != nil {
		return nil, err
	}
	return pool, nil
}

// Init db connection pool
func Init() (*pgxpool.Pool, error) {
	pool, err := GetPool()
	if err != nil {
		return nil, err
	}
	Pool = pool
	return Pool, nil
}
