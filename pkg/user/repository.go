package user

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	GetAll(ctx context.Context) ([]User, error)
	GetOne(ctx context.Context, id uuid.UUID) (User, error)
	GetByUsername(ctx context.Context, username string) (User, error)
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, id uuid.UUID, user User) error
	Delete(ctx context.Context, id uuid.UUID) error
}
