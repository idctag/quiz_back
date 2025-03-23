package sqlc

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/idctag/quiz_back/util"
	"github.com/jackc/pgx/v5"
)

var (
	testQueries *Queries
	ctx         = context.Background()
)

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatalf("Cannot load conf: %v", err)
	}

	conn, err := pgx.Connect(ctx, config.DBSource)
	if err != nil {
		log.Fatalf("Cannot connect to DB: %v", err)
	}

	defer conn.Close(ctx)

	testQueries = New(conn)

	os.Exit(m.Run())
}
