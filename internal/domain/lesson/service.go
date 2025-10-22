package lesson

import (
	"dept-collector/internal/models"
	"dept-collector/internal/pkg/auth"
	"dept-collector/internal/pkg/responses"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateNewLesson(c *gin.Context, db *gorm.DB) {
	var newLessonRequest NewLessonRequest
	if err := c.ShouldBind(&newLessonRequest); err != nil {
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
	err = createNewLesson(lesson, db)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responses.GenericNotFoundError(c.Writer)
			return
		}
		responses.GenericInternalServerError(c.Writer)
		return
	}

	c.JSON(http.StatusCreated, lesson)
}
