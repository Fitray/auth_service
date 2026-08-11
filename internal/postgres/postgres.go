package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/Fitray/auth_service/internal/config"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ConnectionPool struct {
	*pgxpool.Pool
	Timeout time.Duration
}

type Pool interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Close()
	GetTimeout() time.Duration
}

func NewConnectionPool(
	postgresConfig config.PostgresConfig,
) (*ConnectionPool, error) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		postgresConfig.Timeout,
	)
	defer cancel()

	pingCtx, pingCancel := context.WithTimeout(
		context.Background(),
		postgresConfig.Timeout,
	)
	defer pingCancel()

	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		postgresConfig.User,
		postgresConfig.Password,
		postgresConfig.Host,
		postgresConfig.Port,
		postgresConfig.DB,
	)

	pgxConf, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	pgxConf.MaxConns = postgresConfig.MaxConns
	pgxConf.MinConns = postgresConfig.MinConns
	pgxConf.MaxConnLifetime = postgresConfig.MaxConnLifetime
	pgxConf.MaxConnIdleTime = postgresConfig.MaxConnIdleTime

	pool, err := pgxpool.NewWithConfig(ctx, pgxConf)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(pingCtx); err != nil {
		pool.Close()
		return nil, err
	}

	return &ConnectionPool{
		Pool:    pool,
		Timeout: postgresConfig.Timeout,
	}, nil
}

func (c ConnectionPool) GetTimeout() time.Duration {
	return c.Timeout
}

func (c *ConnectionPool) Close() {
	c.Pool.Close()
}
