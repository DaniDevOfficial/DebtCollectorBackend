package responseTypes

import (
	"time"

	"github.com/google/uuid"
)

type ClassResponse struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	SemesterID uuid.UUID `json:"semesterId"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
