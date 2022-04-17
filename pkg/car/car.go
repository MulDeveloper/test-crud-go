package car

import (
	"time"

	"github.com/google/uuid"
)

type Car struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Brand     string    `json:"brand,omitempty"`
	Model     string    `json:"model,omitempty"`
	Price     float64   `json:"price,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
