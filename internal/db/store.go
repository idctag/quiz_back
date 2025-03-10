package db

import (
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/net/context"
)

// Store all db functions
type Store struct {
	*Queries
	db *pgxpool.Pool
}

// Create new store
func NewStore(db *pgxpool.Pool) *Store {
	return &Store{
		Queries: New(db),
		db:      db,
	}
}

func (s *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return err
	}
	q := New(tx) // create new querries with transaction
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			log.Printf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit(ctx)
}

func (s *Store) Close() {
	s.db.Close()
}
