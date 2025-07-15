package repository

import (
	"context"
	"io"
	"log"

	"github.com/sony-nurdianto/ecommerce/ecommerce_web/internal/models"
	"github.com/sony-nurdianto/ecommerce/ecommerce_web/internal/pbgen"
)

type ProductRepo interface {
	GetListProduct(ctx context.Context, limit int, offset int) ([]models.Product, error)
}

type productRepo struct {
	service pbgen.ProductServiceClient
}

func NewProductRepo(service pbgen.ProductServiceClient) *productRepo {
	return &productRepo{
		service: service,
	}
}

func (pr productRepo) GetListProduct(ctx context.Context, limit int, offset int) (data []models.Product, _ error) {
	req := &pbgen.ListProductRequest{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	stream, err := pr.service.ListProduct(ctx, req)
	if err != nil {
		return data, err
	}

	streamCtx := stream.Context()
	for {
		select {
		case <-streamCtx.Done():
			return data, streamCtx.Err()
		default:
			res, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					log.Println("Stream is Over")
					return data, nil
				}

				return data, err
			}

			product := models.Product{
				Name:  res.GetName(),
				Price: res.GetPrice(),
			}

			data = append(data, product)
		}
	}
}
