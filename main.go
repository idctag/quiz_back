package main

import (
	"context"
	"log"

	"github.com/idctag/quiz_back/internal/db"
	"github.com/idctag/quiz_back/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()

	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("Cannot load conf")
	}

	dbpool, err := pgxpool.New(ctx, config.DBSource)
	if err != nil {
		log.Fatalf("Cannot connect to database: %v\n", err)
	}
	err = dbpool.Ping(ctx)
	if err != nil {
		log.Fatalf("Unable to ping database: %v\n", err)
	}

	store := db.NewStore(dbpool)
	defer store.Close()
}
