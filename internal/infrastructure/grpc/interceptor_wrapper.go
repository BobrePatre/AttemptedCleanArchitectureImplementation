package grpc

import (
	"go.uber.org/fx"
)

const UnaryServerInterceptorGroup = `group:"unaryServerInterceptors"`

func AsUnaryServerInterceptor(f any) any {
	return fx.Annotate(
		f,
		fx.ResultTags(UnaryServerInterceptorGroup),
	)
}
