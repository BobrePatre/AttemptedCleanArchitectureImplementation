package grpc

import (
	"context"
	"google.golang.org/grpc"
	"log/slog"
)

func NewLoggingInterceptor(logger *slog.Logger) grpc.UnaryServerInterceptor {
	logger.Info("Initializing logging interceptor")
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		logger.Info(
			"New Grpc Req",
			"Method", info.FullMethod,
			"Server", info.Server,
		)
		return handler(ctx, req)
	}
}
