package responseTypes

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
	ClassName     string    `json:"className"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type SpecificLesson struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	StartDateTime time.Time `json:"startDateTime"`
	EndDateTime   time.Time `json:"endDateTime"`

	ClassID      uuid.UUID `json:"classId"`
	ClassName    string    `json:"className"`
	SemesterID   uuid.UUID `json:"semesterId"`
	SemesterName string    `json:"semesterName"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type FilteredLesson struct {
	ID            uuid.UUID           `json:"id"`
	Name          string              `json:"name"`
	StartDateTime time.Time           `json:"startDateTime"`
	EndDateTime   time.Time           `json:"endDateTime"`
	ClassID       uuid.UUID           `json:"classId"`
	ClassName     string              `json:"className"`
	SemesterID    uuid.UUID           `json:"semesterId"`
	SemesterName  string              `json:"semesterName"`
	SkipEntries   []SkipEntryResponse `json:"skipEntries"`
	CreatedAt     time.Time           `json:"createdAt"`
	UpdatedAt     time.Time           `json:"updatedAt"`
}
