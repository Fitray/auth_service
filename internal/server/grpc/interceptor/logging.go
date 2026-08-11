package interceptor

import (
	"context"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func LoggingUnaryInterceptor(log *zap.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		startedAt := time.Now()

		response, err := handler(ctx, req)

		fields := []zap.Field{
			zap.String("request_id", requestIDFromContext(ctx)),
			zap.String("method", info.FullMethod),
			zap.String("code", status.Code(err).String()),
			zap.Duration("duration", time.Since(startedAt)),
		}

		if err != nil {
			log.Error("gRPC request failed", fields...)
		} else {
			log.Info("gRPC request", fields...)
		}

		return response, err
	}
}
