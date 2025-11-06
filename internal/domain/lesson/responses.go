package lesson

import (
	"github.com/google/uuid"
	"time"
)

type LessonResponse struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	StartDateTime time.Time `json:"startDateTime"`
	EndDateTime   time.Time `json:"endDateTime"`
	ClassID       uuid.UUID `json:"classId"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
