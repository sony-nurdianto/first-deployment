package storage

import (
	"context"
	"time"

	_ "github.com/lib/pq"
)

type PostgresDataStore struct {
	db PostgresDb
}

func OpenPostgres(dburi string, pgi PostgresInstance) (*PostgresDataStore, error) {
	db, err := pgi.Open("postgres", dburi)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Second)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second*10,
	)

	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return &PostgresDataStore{db: db}, nil
}

func (pds *PostgresDataStore) Database() PostgresDb {
	return pds.db
}
