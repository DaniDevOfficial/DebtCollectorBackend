package semester

import (
	"dept-collector/internal/models"
	"dept-collector/internal/pkg/auth"
	"dept-collector/internal/pkg/responses"
	"dept-collector/internal/responseTypes"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateNewSemester godoc
// @Summary      Creates a new semester
// @Description  Creates a new semester and returns it
// @Tags         Semesters
// @Accept       json
// @Produce      json
// @Param        request body NewSemesterRequest true "Create new semester"
// @Success      201  {object}  responseTypes.SemesterResponse
// @Failure      400  {string}  bad request
// @Failure      401  {string}  unauthorized
// @Failure      500  {string}  internal server error
// @Router       /semesters [post]
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

	response := responseTypes.SemesterResponse{
		ID:        newSemester.ID,
		Name:      newSemester.Name,
		StartDate: newSemester.StartDate,
		EndDate:   newSemester.EndDate,
		CreatedAt: newSemester.CreatedAt,
		UpdatedAt: newSemester.UpdatedAt,
	}

	c.JSON(http.StatusCreated, response)
}

// EditSemester godoc
// @Summary      Edit an existing semester
// @Description  Updates semester information
// @Tags         Semesters
// @Accept       json
// @Produce      json
// @Param        request body EditSemesterRequest true "Edit semester"
// @Success      200  {object}  responseTypes.SemesterResponse
// @Failure      400  {string}  bad request
// @Failure      401  {string}  unauthorized
// @Failure      404  {string}  not found
// @Failure      500  {string}  internal server error
// @Router       /semesters [put]
func EditSemester(c *gin.Context, db *gorm.DB) {
	var req EditSemesterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		responses.GenericBadRequestError(c.Writer)
		return
	}

	_, err := auth.AuthenticateByHeader(c, db)
	if err != nil {
		responses.GenericUnauthorizedError(c.Writer)
		return
	}

	startDate, err := time.Parse(time.RFC3339, req.StartDate)
	if err != nil {
		responses.GenericBadRequestError(c.Writer, "Invalid start date format, must be RFC3339")
		return
	}

	endDate, err := time.Parse(time.RFC3339, req.EndDate)
	if err != nil {
		responses.GenericBadRequestError(c.Writer, "Invalid end date format, must be RFC3339")
		return
	}

	id, err := uuid.Parse(req.ID)
	if err != nil {
		responses.GenericBadRequestError(c.Writer, "Invalid ID structure")
		return
	}

	semester := models.Semester{
		ID:        id,
		Name:      req.Name,
		StartDate: startDate,
		EndDate:   endDate,
	}

	if err := updateSemester(&semester, db); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responses.GenericNotFoundError(c.Writer)
			return
		}
		responses.GenericInternalServerError(c.Writer)
		return
	}

	response := responseTypes.SemesterResponse{
		ID:        semester.ID,
		Name:      semester.Name,
		StartDate: semester.StartDate,
		EndDate:   semester.EndDate,
		CreatedAt: semester.CreatedAt,
		UpdatedAt: semester.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

// GetSpecificSemester godoc
// @Summary      Get a specific semester
// @Description  Returns semester details by Id
// @Tags         Semesters
// @Accept       json
// @Produce      json
// @Param        request body SemesterIdRequest true "Get semester"
// @Success      200  {object}  responseTypes.SemesterResponse
// @Failure      400  {string}  bad request
// @Failure      401  {string}  unauthorized
// @Failure      404  {string}  not found
// @Failure      500  {string}  internal server error
// @Router       /semesters [get]
func GetSpecificSemester(c *gin.Context, db *gorm.DB) {
	var request SemesterIdRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		responses.GenericBadRequestError(c.Writer)
		return
	}

	_, err := auth.AuthenticateByHeader(c, db)
	if err != nil {
		responses.GenericUnauthorizedError(c.Writer)
		return
	}

	id, err := uuid.Parse(request.ID)
	if err != nil {
		responses.GenericBadRequestError(c.Writer, "Invalid ID")
		return
	}

	semester, err := getSemester(id, db)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responses.GenericNotFoundError(c.Writer)
			return
		}
		responses.GenericInternalServerError(c.Writer)
		return
	}

	response := responseTypes.SemesterResponse{
		ID:        semester.ID,
		Name:      semester.Name,
		StartDate: semester.StartDate,
		EndDate:   semester.EndDate,
		CreatedAt: semester.CreatedAt,
		UpdatedAt: semester.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

// DeleteSemester godoc
// @Summary      Delete a semester
// @Description  Deletes a semester by ID
// @Tags         Semesters
// @Accept       json
// @Produce      json
// @Param        request body SemesterIdRequest true "Delete semester"
// @Success      204  {string}  no content
// @Failure      400  {string}  bad request
// @Failure      401  {string}  unauthorized
// @Failure      404  {string}  not found
// @Failure      500  {string}  internal server error
// @Router       /semesters [delete]
func DeleteSemester(c *gin.Context, db *gorm.DB) {
	var req SemesterIdRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		responses.GenericBadRequestError(c.Writer)
		return
	}

	_, err := auth.AuthenticateByHeader(c, db)
	if err != nil {
		responses.GenericUnauthorizedError(c.Writer)
		return
	}

	id, err := uuid.Parse(req.ID)
	if err != nil {
		responses.GenericBadRequestError(c.Writer, "Invalid ID structure")
		return
	}

	if err := deleteSemester(id, db); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responses.GenericNotFoundError(c.Writer)
			return
		}
		responses.GenericInternalServerError(c.Writer)
		return
	}

	c.Status(http.StatusNoContent)
}
