package main

import (
	"context"
	"fmt"
	"log"

	"github.com/devnickxyz/gotodo/cmd"
	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool
var ctx = context.Background()

func DB() (*pgxpool.Pool, error) {
	var err error
	pool, err = pgxpool.New(ctx, "postgresql://postgres:mysecretpassword@localhost:5432/gotodo")
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
