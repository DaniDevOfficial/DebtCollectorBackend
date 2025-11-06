package semester

import (
	"github.com/google/uuid"
	"time"
)

type NewSemesterResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
