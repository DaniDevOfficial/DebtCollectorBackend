package lesson

import (
	"dept-collector/internal/models"
	"dept-collector/internal/pkg/auth"
	"dept-collector/internal/pkg/responses"
	"dept-collector/internal/responseTypes"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateNewLesson godoc
// @Summary      Creates a new lesson
// @Description  Creates a new lesson and returns it
// @Tags         Lessons
// @Accept       json
// @Produce      json
// @Param        request body NewLessonRequest true "Create new lesson"
// @Success      201  {object}  responseTypes.LessonResponse
// @Failure      400  {string}  bad request
// @Failure      401  {string}  unauthorized
// @Failure      500  {string}  internal server error
// @Router       /lesson [post]
func CreateNewLesson(c *gin.Context, db *gorm.DB) {
	var newLessonRequest NewLessonRequest
	if err := c.ShouldBindJSON(&newLessonRequest); err != nil {
		responses.GenericBadRequestError(c.Writer)
		return
	}
	_, err := auth.AuthenticateByHeader(c, db)
	if err != nil {
		responses.GenericUnauthorizedError(c.Writer)
		return
	}

	classId, err := uuid.Parse(newLessonRequest.ClassId)

	if err != nil {
		responses.GenericBadRequestError(c.Writer, "Invalid classId")
		return
	}

	lesson := models.Lesson{
		ID:            uuid.New(),
		ClassID:       classId,
		Name:          newLessonRequest.Name,
		StartDateTime: newLessonRequest.StartDate,
		EndDateTime:   newLessonRequest.EndDate,
	}
	err = createNewLesson(&lesson, db)
	if err != nil {
		responses.GenericInternalServerError(c.Writer)
		return
	}

	response := responseTypes.LessonResponse{
		ID:            lesson.ID,
		Name:          lesson.Name,
		StartDateTime: lesson.StartDateTime,
		EndDateTime:   lesson.EndDateTime,
		ClassName:     lesson.Class.Name,
		ClassID:       lesson.ClassID,
		CreatedAt:     lesson.CreatedAt,
		UpdatedAt:     lesson.UpdatedAt,
	}
	c.JSON(http.StatusCreated, response)
}

// EditLesson godoc
// @Summary      Edit a lesson
// @Description  Edit a lesson and returns it
// @Tags         Lessons
// @Accept       json
// @Produce      json
// @Param        request body EditLessonRequest true "Edit Lesson"
// @Success      201  {object}  responseTypes.LessonResponse
// @Failure      400  {string}  bad request
// @Failure      401  {string}  unauthorized
// @Failure      404  {string}  not found
// @Failure      500  {string}  internal server error
// @Router       /lesson [put]
func EditLesson(c *gin.Context, db *gorm.DB) {
	var editLessonRequest EditLessonRequest
	if err := c.ShouldBindJSON(&editLessonRequest); err != nil {
		responses.GenericBadRequestError(c.Writer)
		return
	}

	_, err := auth.AuthenticateByHeader(c, db)
	if err != nil {
		responses.GenericUnauthorizedError(c.Writer)
		return
	}

	lesson := models.Lesson{
		Name:          editLessonRequest.Name,
		StartDateTime: editLessonRequest.StartDate,
		EndDateTime:   editLessonRequest.EndDate,
	}

	id, err := uuid.Parse(editLessonRequest.ID)
	if err != nil {
		responses.GenericBadRequestError(c.Writer, "Invalid id")
		return
	}
	lesson.ID = id

	classId, err := uuid.Parse(editLessonRequest.ClassId)
	if err != nil {
		responses.GenericBadRequestError(c.Writer, "Invalid classId")
		return
	}
	lesson.ClassID = classId

	err = updateLesson(&lesson, db)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responses.GenericNotFoundError(c.Writer)
			return
		}
		responses.GenericInternalServerError(c.Writer)
		return
	}
	response := responseTypes.LessonResponse{
		ID:            lesson.ID,
		Name:          lesson.Name,
		StartDateTime: lesson.StartDateTime,
		EndDateTime:   lesson.EndDateTime,
		ClassName:     lesson.Class.Name,
		ClassID:       lesson.ClassID,
		CreatedAt:     lesson.CreatedAt,
		UpdatedAt:     lesson.UpdatedAt,
	}
	c.JSON(http.StatusOK, response)
}

// DeleteLesson godoc
// @Summary      Deletes a lesson
// @Tags         Lessons
// @Accept       json
// @Produce      json
// @Param        request body SpecificLessonRequest true "Lesson id to delete"
// @Success      204  {string}  no content
// @Failure      400  {string}  bad request
// @Failure      401  {string}  unauthorized
// @Failure      404  {string}  not found
// @Failure      500  {string}  internal server error
// @Router       /lesson [delete]
func DeleteLesson(c *gin.Context, db *gorm.DB) {
	var lessonToDelete SpecificLessonRequest

	if err := c.ShouldBindJSON(&lessonToDelete); err != nil {
		responses.GenericBadRequestError(c.Writer)
		return
	}

	_, err := auth.AuthenticateByHeader(c, db)
	if err != nil {
		responses.GenericUnauthorizedError(c.Writer)
		return
	}

	id, err := uuid.Parse(lessonToDelete.ID)
	if err != nil {
		responses.GenericBadRequestError(c.Writer, "Invalid id")
		return
	}

	err = deleteLesson(id, db)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responses.GenericNotFoundError(c.Writer)
			return
		}
		responses.GenericInternalServerError(c.Writer)
		return
	}
	c.JSON(http.StatusOK, "deleted")
}

// GetSpecificLesson godoc
// @Summary      Get Specific lesson
// @Description  Get Specific lesson
// @Tags         Lessons
// @Accept       json
// @Produce      json
// @Param        request body SpecificLessonRequest true "Edit Lesson"
// @Success      201  {object}  responseTypes.SpecificLesson
// @Failure      400  {string}  bad request
// @Failure      401  {string}  unauthorized
// @Failure      404  {string}  not found
// @Failure      500  {string}  internal server error
// @Router       /lesson [get]
func GetSpecificLesson(c *gin.Context, db *gorm.DB) {
	var lessonToGet SpecificLessonRequest
	if err := c.ShouldBindJSON(&lessonToGet); err != nil {
		responses.GenericBadRequestError(c.Writer)
		return
	}

	_, err := auth.AuthenticateByHeader(c, db)
	if err != nil {
		responses.GenericUnauthorizedError(c.Writer)
		return
	}

	id, err := uuid.Parse(lessonToGet.ID)
	if err != nil {
		responses.GenericBadRequestError(c.Writer, "Invalid id")
		return
	}

	lesson, err := getLesson(id, db)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responses.GenericNotFoundError(c.Writer)
			return
		}
		responses.GenericInternalServerError(c.Writer)
		return
	}
	response := responseTypes.SpecificLesson{
		ID:            lesson.ID,
		Name:          lesson.Name,
		StartDateTime: lesson.StartDateTime,
		EndDateTime:   lesson.EndDateTime,
		ClassID:       lesson.Class.ID,
		ClassName:     lesson.Class.Name,
		SemesterID:    lesson.Class.Semester.ID,
		SemesterName:  lesson.Class.Semester.Name,
		CreatedAt:     lesson.CreatedAt,
		UpdatedAt:     lesson.UpdatedAt,
	}
	c.JSON(http.StatusOK, response)
}

// GetFilteredLessonsWithSkipEntries godoc
// @Summary      Get filtered lessons
// @Description  Get filtered lesson based on a lot of optional filters
// @Tags         Lessons
// @Accept       json
// @Produce      json
// @Param        request body FilterLessonRequest true "Edit Lesson"
// @Success      201  {object}  []responseTypes.FilteredLesson
// @Failure      400  {string}  bad request
// @Failure      401  {string}  unauthorized
// @Failure      500  {string}  internal server error
// @Router       /lesson/filtered [get]
func GetFilteredLessonsWithSkipEntries(c *gin.Context, db *gorm.DB) {
	var filterLessonRequest FilterLessonRequest
	if err := c.ShouldBindJSON(&filterLessonRequest); err != nil {
		log.Println(err)
		responses.GenericBadRequestError(c.Writer)
		return
	}

	_, err := auth.AuthenticateByHeader(c, db)
	if err != nil {
		responses.GenericUnauthorizedError(c.Writer)
		return
	}

	lessons, err := getAllLessonsWithSkipEntries(filterLessonRequest, db)
	if err != nil {
		responses.GenericInternalServerError(c.Writer)
		return
	}

	c.JSON(http.StatusOK, buildFilteredLessonsResponse(lessons))
}
