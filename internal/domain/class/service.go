package class

import (
	"dept-collector/internal/models"
	"dept-collector/internal/pkg/auth"
	"dept-collector/internal/pkg/responses"
	"dept-collector/internal/responseTypes"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateNewClass godoc
// @Summary      Creates a new class
// @Description  Creates a new class and returns the created record
// @Tags         Classes
// @Accept       json
// @Produce      json
// @Param        request body NewClassRequest true "Create new class"
// @Success      201  {object}  responseTypes.ClassResponse
// @Failure      400  {string}  bad request
// @Failure      401  {string}  unauthorized
// @Failure      500  {string}  internal server error
// @Router       /classes [post]
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
		ID:         uuid.New(),
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
	response := responseTypes.ClassResponse{
		ID:         newClass.ID,
		Name:       newClass.Name,
		SemesterID: newClass.SemesterID,
		CreatedAt:  newClass.CreatedAt,
		UpdatedAt:  newClass.UpdatedAt,
	}

	c.JSON(http.StatusCreated, response)
}

// EditClass godoc
// @Summary      Edit an existing class
// @Description  Updates class details like name or semester
// @Tags         Classes
// @Accept       json
// @Produce      json
// @Param        request body EditClassRequest true "Edit existing class"
// @Success      200  {object}  responseTypes.ClassResponse
// @Failure      400  {string}  bad request
// @Failure      401  {string}  unauthorized
// @Failure      404  {string}  not found
// @Failure      500  {string}  internal server error
// @Router       /classes [put]
func EditClass(c *gin.Context, db *gorm.DB) {
	var editClassRequest EditClassRequest

	if err := c.ShouldBindJSON(&editClassRequest); err != nil {
		responses.GenericBadRequestError(c.Writer)
		return
	}

	_, err := auth.AuthenticateByHeader(c, db)
	if err != nil {
		responses.GenericUnauthorizedError(c.Writer)
		return
	}

	classId, err := uuid.Parse(editClassRequest.ID)
	if err != nil {
		responses.GenericBadRequestError(c.Writer, "Invalid class ID structure")
		return
	}

	semesterId, err := uuid.Parse(editClassRequest.SemesterID)
	if err != nil {
		responses.GenericBadRequestError(c.Writer, "Invalid semester ID structure")
		return
	}

	class, err := getClass(classId, db)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responses.GenericNotFoundError(c.Writer)
			return
		}
		responses.GenericInternalServerError(c.Writer)
		return
	}

	class.Name = editClassRequest.Name
	class.SemesterID = semesterId

	if err := updateClass(&class, db); err != nil {
		responses.GenericInternalServerError(c.Writer)
		return
	}

	response := responseTypes.ClassResponse{
		ID:         class.ID,
		Name:       class.Name,
		SemesterID: class.SemesterID,
		CreatedAt:  class.CreatedAt,
		UpdatedAt:  class.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

// GetClass godoc
// @Summary      Get a specific class
// @Description  Retrieves a class by its ID (from request body)
// @Tags         Classes
// @Accept       json
// @Produce      json
// @Param        request body ClassIdRequest true "Get class by ID"
// @Success      200  {object}  responseTypes.NewClassResponse
// @Failure      400  {string}  bad request
// @Failure      401  {string}  unauthorized
// @Failure      404  {string}  not found
// @Failure      500  {string}  internal server error
// @Router       /classes [post]
func GetClass(c *gin.Context, db *gorm.DB) {
	var request ClassIdRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		responses.GenericBadRequestError(c.Writer)
		return
	}

	_, err := auth.AuthenticateByHeader(c, db)
	if err != nil {
		responses.GenericUnauthorizedError(c.Writer)
		return
	}

	classId, err := uuid.Parse(request.ID)
	if err != nil {
		responses.GenericBadRequestError(c.Writer, "Invalid class ID")
		return
	}

	class, err := getClass(classId, db)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responses.GenericNotFoundError(c.Writer)
			return
		}
		responses.GenericInternalServerError(c.Writer)
		return
	}

	response := responseTypes.ClassResponse{
		ID:         class.ID,
		Name:       class.Name,
		SemesterID: class.SemesterID,
		CreatedAt:  class.CreatedAt,
		UpdatedAt:  class.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

// DeleteClass godoc
// @Summary      Delete a class
// @Description  Deletes a class by its ID (from request body)
// @Tags         Classes
// @Accept       json
// @Produce      json
// @Param        request body ClassIdRequest true "Delete class request"
// @Success      204  {string}  no content
// @Failure      400  {string}  bad request
// @Failure      401  {string}  unauthorized
// @Failure      404  {string}  not found
// @Failure      500  {string}  internal server error
// @Router       /classes [delete]
func DeleteClass(c *gin.Context, db *gorm.DB) {
	var request ClassIdRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		responses.GenericBadRequestError(c.Writer)
		return
	}

	_, err := auth.AuthenticateByHeader(c, db)
	if err != nil {
		responses.GenericUnauthorizedError(c.Writer)
		return
	}

	classId, err := uuid.Parse(request.ID)
	if err != nil {
		responses.GenericBadRequestError(c.Writer, "Invalid class ID")
		return
	}

	err = deleteClass(classId, db)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responses.GenericNotFoundError(c.Writer)
			return
		}
		responses.GenericInternalServerError(c.Writer)
		return
	}

	c.Status(http.StatusNoContent)
}

// GetFilteredClasses godoc
// @Summary      Filter classes
// @Description  Filters classes by name, semester ID, or semester start/end date
// @Tags         Classes
// @Accept       json
// @Produce      json
// @Param        request body FilterClassRequest true "Filter classes"
// @Success      200  {array}  responseTypes.ClassResponse
// @Failure      400  {string}  bad request
// @Failure      401  {string}  unauthorized
// @Failure      500  {string}  internal server error
// @Router       /classes/filter [post]
func GetFilteredClasses(c *gin.Context, db *gorm.DB) {
	var filterRequest FilterClassRequest

	if err := c.ShouldBindJSON(&filterRequest); err != nil {
		responses.GenericBadRequestError(c.Writer)
		return
	}

	_, err := auth.AuthenticateByHeader(c, db)
	if err != nil {
		responses.GenericUnauthorizedError(c.Writer)
		return
	}

	classes, err := getFilteredClasses(filterRequest, db)
	if err != nil {
		responses.GenericInternalServerError(c.Writer)
		return
	}

	var response []responseTypes.ClassResponse
	for _, class := range classes {
		response = append(response, responseTypes.ClassResponse{
			ID:         class.ID,
			Name:       class.Name,
			SemesterID: class.SemesterID,
			CreatedAt:  class.CreatedAt,
			UpdatedAt:  class.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, response)
}
