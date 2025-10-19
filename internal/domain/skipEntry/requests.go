package skipEntry

type CreateNewEntryRequest struct {
	UserID   string `json:"userId"`
	Reason   string `json:"reason"`
	LessonID string `json:"lessonId"`
	AmountID string `json:"amountId"`
}
