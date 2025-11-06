package class

import (
	"dept-collector/internal/models"
	"dept-collector/internal/pkg/auth"
	"dept-collector/internal/pkg/responses"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

// CreateNewClass godoc
// @Summary      Creates a new class
// @Description  Creates a new class and returns the created record
// @Tags         Classes
// @Accept       json
// @Produce      json
// @Param        request body NewClassRequest true "Create new class"
// @Success      201  {object}  NewClassResponse
// @Failure      400  {string}  bad request
// @Failure      401  {string}  unauthorized
// @Failure      500  {string}  internal server error
// @Router       /classes/create [post]
func CreateNewClass(c *gin.Context, db *gorm.DB) {
	var newClassRequest NewClassRequest

	if err := c.ShouldBindJSON(&newClassRequest); err != nil {
		responses.GenericBadRequestError(c.Writer)
		return
	}

	_, err := auth.AuthenticateByHeader(c, db)
	if err != nil {
		responses.GenericUnauthorizedError(c.Writer)
		return
	}

	semesterId, err := uuid.Parse(newClassRequest.SemesterID)
	if err != nil {
		responses.GenericBadRequestError(c.Writer, "Invalid semester ID structure")
		return
	}

	newClass := models.Class{
		ID:         uuid.New().String(),
		Name:       newClassRequest.Name,
		SemesterID: semesterId,
	}

	err = createClass(&newClass, db)
	if err != nil {
		if errors.Is(err, ErrSemesterNotFound) {
			responses.GenericBadRequestError(c.Writer, "Semester does not exist")
			return
		}
		responses.GenericInternalServerError(c.Writer)
		return
	}
	response := NewClassResponse{
		ID:         newClass.ID,
		Name:       newClass.Name,
		SemesterID: newClass.SemesterID,
		CreatedAt:  newClass.CreatedAt,
		UpdatedAt:  newClass.UpdatedAt,
	}

	c.JSON(http.StatusCreated, response)
}
