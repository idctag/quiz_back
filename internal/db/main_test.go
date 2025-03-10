package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/idctag/quiz_back/util"
	"github.com/jackc/pgx/v5"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("Cannot load conf")
	}

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	defer conn.Close(ctx)

	testQueries = New(conn)

	os.Exit(m.Run())
}
