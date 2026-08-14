package auth_postgres

import (
	"context"
	stdErrors "errors"
	"fmt"
	"strings"

	app_errors "github.com/Fitray/auth_service/internal/errors"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	postgresUniqueViolation = "23505"
	usersLoginConstraint    = "users_login_key"
)

func mapCreateUserError(err error) error {
	switch {
	case stdErrors.Is(err, context.Canceled):
		return app_errors.NewError(
			app_errors.RequestCancelled,
			"request cancelled",
			err,
		)
	case stdErrors.Is(err, context.DeadlineExceeded):
		return app_errors.NewError(
			app_errors.GatewayTimeout,
			"request deadline exceeded",
			err,
		)
	}

	var pgErr *pgconn.PgError
	if !stdErrors.As(err, &pgErr) {
		return fmt.Errorf("create user: %w", err)
	}

	switch {
	case pgErr.Code == postgresUniqueViolation &&
		pgErr.ConstraintName == usersLoginConstraint:
		return app_errors.NewError(
			app_errors.AlreadyExists,
			"login is already taken",
			err,
		)
	case pgErr.Code == postgresUniqueViolation:
		return app_errors.NewError(
			app_errors.AlreadyExists,
			"user already exists",
			err,
		)
	case strings.HasPrefix(pgErr.Code, "08") ||
		pgErr.Code == "57P01" ||
		pgErr.Code == "57P02" ||
		pgErr.Code == "57P03":
		return app_errors.NewError(
			app_errors.ServiceUnavailable,
			"database is temporarily unavailable",
			err,
		)
	default:
		return fmt.Errorf("create user: %w", err)
	}
}
