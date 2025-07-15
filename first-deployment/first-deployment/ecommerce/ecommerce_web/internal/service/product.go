package service

import (
	"context"
	"time"

	"github.com/sony-nurdianto/ecommerce/ecommerce_web/internal/models"
	"github.com/sony-nurdianto/ecommerce/ecommerce_web/internal/repository"
)

type ProductService interface {
	Products(limit int, offset int) (data []models.Product, _ error)
}

type productService struct {
	repo repository.ProductRepo
}

func NewProductService(repo repository.ProductRepo) *productService {
	return &productService{repo: repo}
}

func (pu *productService) Products(limit int, offset int) (data []models.Product, _ error) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		2*time.Second,
	)

	defer cancel()

	lp, err := pu.repo.GetListProduct(ctx, limit, offset)
	if err != nil {
		return data, err
	}

	return lp, nil
}
