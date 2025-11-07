package skipEntry

import "time"

type CreateNewEntryRequest struct {
	UserID   string `json:"userId" binding:"required"`
	Reason   string `json:"reason" binding:"required"`
	LessonID string `json:"lessonId" binding:"required"`
	AmountID string `json:"amountId" binding:"required"`
}

type EditSkipEntryRequest struct {
	ID       string `json:"id" binding:"required,uuid"`
	UserID   string `json:"userId" binding:"required,uuid"`
	Reason   string `json:"reason" binding:"required,max=255"`
	LessonID string `json:"lessonId" binding:"required,uuid"`
	AmountID string `json:"amountId" binding:"required,uuid"`
}

type SingleIdRequest struct {
	ID string `json:"id" binding:"required,uuid"`
}

type FilterSkipEntryRequest struct {
	UserID     *string    `json:"userId" binding:"omitempty,uuid"`
	LessonID   *string    `json:"lessonId" binding:"omitempty,uuid"`
	AmountID   *string    `json:"amountId" binding:"omitempty,uuid"`
	Reason     *string    `json:"reason" binding:"omitempty"`
	StartDate  *time.Time `json:"startDate" binding:"omitempty"`
	EndDate    *time.Time `json:"endDate" binding:"omitempty"`
	SemesterID *string    `json:"semesterId" binding:"omitempty,uuid"`
	ClassID    *string    `json:"classId" binding:"omitempty,uuid"`
}
