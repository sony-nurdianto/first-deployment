package repository

import "github.com/sony-nurdianto/ecommerce/product_svc/internal/storage"

type postgresRepo struct {
	db storage.PostgresDb
}

func NewPostgresRepo(db storage.PostgresDb) *postgresRepo {
	return &postgresRepo{
		db: db,
	}
}

func (r *postgresRepo) GetProductList(limit int, offset int) {
}
