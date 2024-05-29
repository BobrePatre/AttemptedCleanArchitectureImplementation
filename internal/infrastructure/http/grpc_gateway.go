package http

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func NewGatewayServer() *runtime.ServeMux {
	gatewayMux := runtime.NewServeMux()
	return gatewayMux
}
