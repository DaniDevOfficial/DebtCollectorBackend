package lesson

import "time"

type NewLessonRequest struct {
	Name      string    `json:"name" binding:"required"`
	StartDate time.Time `json:"startDate" binding:"required"`
	EndDate   time.Time `json:"endDate" binding:"required"`
	ClassId   string    `json:"classId" binding:"required,uuid"`
}

type EditLessonRequest struct {
	ID        string    `json:"id" binding:"required,uuid"`
	Name      string    `json:"name" binding:"required"`
	StartDate time.Time `json:"startDate" binding:"required,date"`
	EndDate   time.Time `json:"endDate" binding:"required,date"`
	ClassId   string    `json:"classId" binding:"required,uuid"`
}

type SpecificLessonRequest struct {
	ID string `json:"id" binding:"required,uuid"`
}

type FilterLessonRequest struct {
	Name      *string    `json:"reason" binding:"omitempty"`
	StartDate *time.Time `json:"startDate" binding:"omitempty,dateTime"`
	EndDate   *time.Time `json:"endDate" binding:"omitempty,dateTime"`
	ClassID   *string    `json:"class_id" binding:"omitempty,uuid"`
}
