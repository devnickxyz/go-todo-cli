package main

import (
	"context"
	"fmt"
	"log"
	"os"

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

	pool, err = pgxpool.New(ctx, connDB)
	if err != nil {
		log.Fatal("Unable to connect to database", err)
	}
	if err := pool.Ping(ctx); err != nil {
		log.Fatal("Unable to ping db:", err)
	}
	fmt.Println("Connected to DB")

	return pool, err
}

func main() {
	DB()
	cmd.Ctx = ctx
	cmd.Pool = pool
	cmd.Execute()
}
