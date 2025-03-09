package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgresql://root:secret@localhost:5432/quiz?sslmode=disable")
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	defer conn.Close(ctx)

	testQueries = New(conn)

	os.Exit(m.Run())
}
