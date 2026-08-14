package authorization_postgres

import "github.com/Fitray/auth_service/internal/postgres"

type AuthorizationPostgres struct {
	pool postgres.Pool
}

func NewAuthorizationPostgres(pool postgres.Pool) *AuthorizationPostgres {
	return &AuthorizationPostgres{
		pool: pool,
	}
}
