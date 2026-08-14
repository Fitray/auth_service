package interceptor

import (
	"context"
	"runtime/debug"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func RecoveryUnaryInterceptor(log *zap.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp any, err error) {
		defer func() {
			if recovered := recover(); recovered != nil {
				log.Error("gRPC panic recovered",
					zap.Any("panic", recovered),
					zap.String("method", info.FullMethod),
					zap.String("request_id", requestIDFromContext(ctx)),
					zap.ByteString("stack", debug.Stack()),
				)

				err = status.Error(
					codes.Internal,
					"internal server error",
				)
			}
		}()

		return handler(ctx, req)
	}
}
