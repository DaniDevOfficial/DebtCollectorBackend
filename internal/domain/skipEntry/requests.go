package skipEntry

type CreateNewEntryRequest struct {
	UserID   string `json:"userId"`
	Reason   string `json:"reason"`
	LessonID string `json:"lessonId"`
	AmountID string `json:"amountId"`
}

type EditSkipEntryRequest struct {
	ID       string `json:"id"`
	UserID   string `json:"userId"`
	Reason   string `json:"reason"`
	LessonID string `json:"lessonId"`
	AmountID string `json:"amountId"`
}
