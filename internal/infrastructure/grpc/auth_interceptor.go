package grpc

import (
	"context"
	"google.golang.org/grpc"
)

// NewAuthInterceptor TODO: Add logic to invoke an authentication provider to verify routes
func NewAuthInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		return nil, nil
	}
}
