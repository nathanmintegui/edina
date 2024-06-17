package user

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nathanmintegui/edina/internal/models"
)

type Repository interface {
	FindUserByLogin(ctx context.Context, login string) (models.User, error)
}

type RepositoryPostgres struct {
	Conn *pgxpool.Pool
}

func (r *RepositoryPostgres) FindUserByLogin(ctx context.Context, login string) (models.User, error) {
	var user models.User

	err := r.Conn.QueryRow(
		ctx,
		"SELECT * FROM users WHERE login = $1;", login,
	).Scan(&user.Id, &user.PublicId, &user.Login, &user.Password)

	if err == pgx.ErrNoRows {
		return models.User{}, errors.New("no rows where found for users")
	}

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
