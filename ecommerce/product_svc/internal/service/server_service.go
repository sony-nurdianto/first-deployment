package service

import "github.com/sony-nurdianto/ecommerce/product_svc/internal/pbgen"

type ProductService struct {
	pbgen.UnimplementedProductServiceServer
}

func (ps *ProductService) ListProduct(
	req *pbgen.ListProductRequest,
	stream pbgen.ProductService_ListProductServer,
) error {
	resp := []*pbgen.ListProductResponse{
		{
			Name:  "Mangga",
			Price: 5.0,
		},
		{
			Name:  "Pisang",
			Price: 3.0,
		},
		{
			Name:  "Peach",
			Price: 6.0,
		},
	}

	ctx := stream.Context()
	for _, v := range resp {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if err := stream.Send(v); err != nil {
				return err
			}
		}
	}

	return nil
}
