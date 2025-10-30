package semester

type NewSemesterRequest struct {
	Name      string `json:"name" binding:"required"`
	StartDate string `json:"startDate" binding:"required"`
	EndDate   string `json:"endDate" binding:"required"`
}
