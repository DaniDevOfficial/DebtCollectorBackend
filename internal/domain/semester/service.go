package semester

import (
	"dept-collector/internal/models"
	"dept-collector/internal/pkg/auth"
	"dept-collector/internal/pkg/responses"
	"dept-collector/internal/responseTypes"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

// CreateNewSemester godoc
// @Summary      Creates a new semester
// @Description  Creates a new semester and returns it
// @Tags         Semesters
// @Accept       json
// @Produce      json
// @Param        request body NewSemesterRequest true "Create new semester"
// @Success      201  {object}  NewSemesterResponse
// @Failure      400  {string}  bad request
// @Failure      401  {string}  unauthorized
// @Failure      500  {string}  internal server error
// @Router       /semesters/create [post]
func CreateNewSemester(c *gin.Context, db *gorm.DB) {
	var newSemesterRequest NewSemesterRequest

	if err := c.ShouldBindJSON(&newSemesterRequest); err != nil {
		log.Println(err)
		responses.GenericBadRequestError(c.Writer)
		return
	}

	_, err := auth.AuthenticateByHeader(c, db)
	if err != nil {
		log.Println(err)
		responses.GenericUnauthorizedError(c.Writer)
		return
	}

	startDate, err := time.Parse(time.RFC3339, newSemesterRequest.StartDate)
	if err != nil {
		responses.GenericBadRequestError(c.Writer, "Invalid start date format, must be RFC3339 (e.g. 2025-09-01T00:00:00Z)")
		return
	}

	endDate, err := time.Parse(time.RFC3339, newSemesterRequest.EndDate)
	if err != nil {
		responses.GenericBadRequestError(c.Writer, "Invalid end date format, must be RFC3339 (e.g. 2025-12-20T23:59:59Z)")
		return
	}

	newSemester := models.Semester{
		ID:        uuid.New(),
		Name:      newSemesterRequest.Name,
		StartDate: startDate,
		EndDate:   endDate,
	}

	err = createSemester(&newSemester, db)
	if err != nil {
		responses.GenericInternalServerError(c.Writer)
		return
	}

	response := responseTypes.NewSemesterResponse{
		ID:        newSemester.ID,
		Name:      newSemester.Name,
		StartDate: newSemester.StartDate,
		EndDate:   newSemester.EndDate,
		CreatedAt: newSemester.CreatedAt,
		UpdatedAt: newSemester.UpdatedAt,
	}

	c.JSON(http.StatusCreated, response)
}
