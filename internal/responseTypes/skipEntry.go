package responseTypes

import (
	"github.com/google/uuid"
	"time"
)

type SkipEntryResponse struct {
	ID     uuid.UUID `json:"id"`
	Reason string    `json:"reason"`

	UserID   uuid.UUID `json:"userId"`
	UserName string    `json:"userName"`

	LessonID   uuid.UUID `json:"lessonId"`
	LessonName string    `json:"lessonName"`

	Amount float64 `json:"amount"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
