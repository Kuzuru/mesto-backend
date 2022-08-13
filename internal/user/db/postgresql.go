package user

import (
	"context"
	"errors"

	"mesto/internal/user"
	"mesto/pkg/db/postgresql"

	"github.com/jackc/pgconn"
	"github.com/rs/zerolog/log"
)

type repository struct {
	client postgresql.Client
}

func (r *repository) Create(ctx context.Context, user *user.User) error {
	query := `INSERT INTO public.users (name, about, avatar) VALUES ($1, $2, $3) RETURNING id;`

	if err := r.client.QueryRow(ctx, query, user.Name, user.About, user.Avatar).Scan(&user.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.Is(err, pgErr) {
			log.Error().Msgf("An error occurred while creating new user: [%s] %s\n", pgErr.SQLState(), pgErr.Message)
			return nil
		}

		return err
	}

	return nil
}

func (r *repository) FindAll(ctx context.Context) (u []user.User, err error) {
	query := `SELECT id, auth_id, name, about, avatar FROM users;`

	rows, err := r.client.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	users := make([]user.User, 0)

	for rows.Next() {
		var u user.User

		err = rows.Scan(&u.ID, &u.AuthID, &u.Name, &u.About, &u.Avatar)
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *repository) FindOne(ctx context.Context, authId string) (user.User, error) {
	query := `SELECT id, auth_id, name, about, avatar FROM public.users WHERE auth_id = $1;`

	var u user.User

	err := r.client.QueryRow(ctx, query, authId).Scan(&u.ID, &u.AuthID, &u.Name, &u.About, &u.Avatar)
	if err != nil {
		return user.User{}, err
	}

	return u, nil
}

func (r *repository) UpdateProfile(ctx context.Context, user user.User) error {
	query := `UPDATE users SET name = $1, about = $2 WHERE auth_id = $3`

	_, err := r.client.Exec(ctx, query, user.Name, user.About, user.AuthID)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) UpdateAvatar(ctx context.Context, user user.User) error {
	//TODO implement me
	panic("implement me")
}

func NewRepository(client postgresql.Client) user.Storage {
	return &repository{
		client: client,
	}
}
