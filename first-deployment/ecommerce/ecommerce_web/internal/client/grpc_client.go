package client

import (
	"github.com/sony-nurdianto/ecommerce/ecommerce_web/internal/pbgen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type clientGRPC struct {
	Conn    *grpc.ClientConn
	Service pbgen.ProductServiceClient
}

func InitClientGRPC(addr string, grpcClient ConnClientGRPC) (*clientGRPC, error) {
	cred := grpc.WithTransportCredentials(
		insecure.NewCredentials(),
	)

	conn, err := grpcClient.NewClient(addr, cred)
	if err != nil {
		return nil, err
	}

	client := pbgen.NewProductServiceClient(conn)
	return &clientGRPC{Conn: conn, Service: client}, nil
}
