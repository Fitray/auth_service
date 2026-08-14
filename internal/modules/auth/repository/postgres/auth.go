package auth_postgres

import (
	"context"

	"github.com/Fitray/auth_service/internal/entities/dto"
)

func (r *AuthPostgres) RegisterUser(
	ctx context.Context,
	registerReq dto.RegisterRequest,
) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.GetTimeout())
	defer cancel()

	row := r.pool.QueryRow(
		ctx,
		`INSERT INTO auth.users
		(login, pass_hash, phone_number)
		VALUES ($1, $2, $3)
		RETURNING id`,
		registerReq.Login,
		registerReq.Password,
		registerReq.PhoneNumber,
	)

	var userId int64
	if err := row.Scan(
		&userId,
	); err != nil {
		return 0, mapCreateUserError(err)
	}
	return userId, nil
}
