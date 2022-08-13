package user

import (
	"context"
)

type Storage interface {
	Create(ctx context.Context, user *User) error
	FindAll(ctx context.Context) (u []User, err error)
	FindOne(ctx context.Context, AuthID string) (User, error)
	UpdateProfile(ctx context.Context, user User) error
	UpdateAvatar(ctx context.Context, user User) error
	Delete(ctx context.Context, id string) error
}
