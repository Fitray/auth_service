package auth_postgres

import "github.com/Fitray/auth_service/internal/postgres"

type AuthPostgres struct {
	pool postgres.Pool
}

func NewAuthPostgres(pool postgres.Pool) *AuthPostgres {
	return &AuthPostgres{
		pool: pool,
	}
}
