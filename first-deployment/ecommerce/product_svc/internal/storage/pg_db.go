package storage

import (
	"context"
	"database/sql"
	"time"
)

type PostgresDb interface {
	PingContext(ctx context.Context) error
	Prepare(query string) (*sql.Stmt, error)
	SetMaxOpenConns(n int)
	SetMaxIdleConns(n int)
	SetConnMaxLifetime(d time.Duration)
	Close() error
}

type postgresDb struct {
	db *sql.DB
}

func NewPostgresDb(db *sql.DB) *postgresDb {
	return &postgresDb{
		db: db,
	}
}

func (p *postgresDb) PingContext(ctx context.Context) error {
	return p.db.PingContext(ctx)
}

func (p *postgresDb) Prepare(query string) {
	p.db.Prepare(query)
}

func (p *postgresDb) SetMaxOpenConns(n int) {
	p.db.SetMaxOpenConns(n)
}

func (p *postgresDb) SetMaxIdleConns(n int) {
	p.db.SetMaxIdleConns(n)
}

func (p *postgresDb) SetConnMaxLifetime(d time.Duration) {
	p.db.SetConnMaxLifetime(d)
}

func (p *postgresDb) Close() error {
	return p.db.Close()
}
