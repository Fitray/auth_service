package interceptor

import (
	"context"
	"time"

	"google.golang.org/grpc"
)

func TimeoutInterceptor(timeout time.Duration) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		if timeout <= 0 {
			return handler(ctx, req)
		}

		if deadline, ok := ctx.Deadline(); ok {
			if time.Until(deadline) <= timeout {
				return handler(ctx, req)
			}
		}

		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		return handler(ctx, req)
	}
}
