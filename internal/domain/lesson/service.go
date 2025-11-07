package lesson

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
	c.JSON(http.StatusCreated, response)
}

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

	err = updateLesson(lesson, db)
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
	c.JSON(http.StatusOK, nil)
}

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

func GetFilteredLessonsWithSkipEntries(c *gin.Context, db *gorm.DB) {
	var filterLessonRequest FilterLessonRequest
	if err := c.ShouldBindJSON(&filterLessonRequest); err != nil {
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
