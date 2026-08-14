package app_errors

import "google.golang.org/grpc/codes"

type Code struct {
	Name     string
	GRPCCode codes.Code
}

var (
	InternalServerErr = Code{
		Name:     "internal_server_error",
		GRPCCode: codes.Internal,
	}

	BadRequest = Code{
		Name:     "bad_request",
		GRPCCode: codes.InvalidArgument,
	}

	Unauthorized = Code{
		Name:     "unauthorized",
		GRPCCode: codes.Unauthenticated,
	}

	Forbidden = Code{
		Name:     "forbidden",
		GRPCCode: codes.PermissionDenied,
	}

	NotFound = Code{
		Name:     "not_found",
		GRPCCode: codes.NotFound,
	}

	AlreadyExists = Code{
		Name:     "already_exists",
		GRPCCode: codes.AlreadyExists,
	}

	UnprocessableEntity = Code{
		Name:     "unprocessable_entity",
		GRPCCode: codes.FailedPrecondition,
	}

	TooManyRequests = Code{
		Name:     "too_many_requests",
		GRPCCode: codes.ResourceExhausted,
	}

	ServiceUnavailable = Code{
		Name:     "service_unavailable",
		GRPCCode: codes.Unavailable,
	}

	GatewayTimeout = Code{
		Name:     "gateway_timeout",
		GRPCCode: codes.DeadlineExceeded,
	}

	RequestCancelled = Code{
		Name:     "request_cancelled",
		GRPCCode: codes.Canceled,
	}

	PayloadTooLarge = Code{
		Name:     "payload_too_large",
		GRPCCode: codes.ResourceExhausted,
	}
)
