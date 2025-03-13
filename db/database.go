package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/idctag/quiz_back/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() {
	conf, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	dbpool, err := pgxpool.New(context.Background(), conf.DBSource)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	DB = dbpool

	fmt.Println("Connected to database")
}
