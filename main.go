package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/devnickxyz/gotodo/cmd"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var pool *pgxpool.Pool
var ctx = context.Background()

type ConnDB struct {
	URL      string
	User     string
	Password string
}

func DB() (*pgxpool.Pool, error) {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var db ConnDB
	db.URL = os.Getenv("PSQL_URL")
	db.User = os.Getenv("PSQL_USER")
	db.Password = os.Getenv("PSQL_PASSWORD")
	connDB := fmt.Sprintf("postgresql://%s:%s@%s", db.User, db.Password, db.URL)

	config, err := pgxpool.ParseConfig(connDB)

	// Set pool configuration
	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnLifetime = time.Hour
	config.MaxConnIdleTime = 30 * time.Minute

	// pool, err = pgxpool.New(ctx, connDB)
	pool, err = pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatal("Unable to connect to database aka creating conection pool %w\n", err)
	}
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		log.Fatal("Unable to ping db:", err)
	}
	fmt.Println("Connected to DB")

	return pool, nil
}

func main() {
	var err error
	pool, err = DB()
	if err != nil {
		pool.Close()
		// fmt.Errorf("%w\n", err)
	}
	cmd.Ctx = ctx
	cmd.Pool = pool
	cmd.Execute()
}
