package car

import (
	"context"

	"github.com/google/uuid"
)

type CarRepository interface {
	GetAll(ctx context.Context) ([]Car, error)
	GetOne(ctx context.Context, id uuid.UUID) (Car, error)
	Create(ctx context.Context, post *Car) error
	Update(ctx context.Context, id uuid.UUID, post Car) error
	Delete(ctx context.Context, id uuid.UUID) error
}
