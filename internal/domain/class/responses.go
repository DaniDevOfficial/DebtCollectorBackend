package class

import (
	"github.com/google/uuid"
	"time"
)

type NewClassResponse struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	SemesterID uuid.UUID `json:"semesterId"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
