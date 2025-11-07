package semester

type NewSemesterRequest struct {
	Name      string `json:"name" binding:"required"`
	StartDate string `json:"startDate" binding:"required"`
	EndDate   string `json:"endDate" binding:"required"`
}

type EditSemesterRequest struct {
	ID        string `json:"id" binding:"required,uuid"`
	Name      string `json:"name" binding:"required"`
	StartDate string `json:"startDate" binding:"required"`
	EndDate   string `json:"endDate" binding:"required"`
}

type SemesterIdRequest struct {
	ID string `json:"id" binding:"required,uuid"`
}
