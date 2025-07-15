package storage

import "database/sql"

type PostgresInstance interface {
	Open(drvName string, dscName string) (PostgresDb, error)
}

type pgInstance struct{}

func NewPgInstance() *pgInstance {
	return &pgInstance{}
}

func (pgi *pgInstance) Open(drvName string, dscName string) (PostgresDb, error) {
	return sql.Open(drvName, dscName)
}
