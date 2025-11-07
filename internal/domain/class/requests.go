package class

import "time"

type NewClassRequest struct {
	Name       string `json:"name" binding:"required"`
	SemesterID string `json:"semesterId" binding:"required,uuid"`
}

type EditClassRequest struct {
	ID         string `json:"id" binding:"required,uuid"`
	Name       string `json:"name" binding:"required"`
	SemesterID string `json:"semesterId" binding:"required,uuid"`
}

type ClassIdRequest struct {
	ID string `json:"id" binding:"required,uuid"`
}

type FilterClassRequest struct {
	Name               *string    `json:"name" binding:"omitempty"`
	SemesterID         *string    `json:"semesterId" binding:"omitempty,uuid"`
	SemesterStartAfter *time.Time `json:"semesterStartAfter" binding:"omitempty"`
	SemesterEndBefore  *time.Time `json:"semesterEndBefore" binding:"omitempty"`
}
