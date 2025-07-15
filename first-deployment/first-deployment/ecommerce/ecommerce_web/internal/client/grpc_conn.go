package client

import (
	"google.golang.org/grpc"
)

type ConnClientGRPC interface {
	NewClient(target string, opts ...grpc.DialOption) (conn *grpc.ClientConn, err error)
}

type grpcClientConn struct{}

func NewClientConnGRPC() *grpcClientConn {
	return &grpcClientConn{}
}

func (c *grpcClientConn) NewClient(target string, opts ...grpc.DialOption) (conn *grpc.ClientConn, err error) {
	return grpc.NewClient(target, opts...)
}
