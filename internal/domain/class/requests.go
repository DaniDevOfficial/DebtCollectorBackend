package class

type NewClassRequest struct {
	Name       string `json:"name" binding:"required"`
	SemesterID string `json:"semesterId" binding:"required,uuid"`
}
