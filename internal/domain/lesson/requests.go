package lesson

import "time"

type NewLessonRequest struct {
	Name      string    `json:"name" binding:"required"`
	StartDate time.Time `json:"startDate" binding:"required,date"`
	EndDate   time.Time `json:"endDate" binding:"required,date"`
	ClassId   string    `json:"classId" binding:"required,uuid"`
}
