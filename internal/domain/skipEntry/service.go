package skipEntry

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

func CreateNewSkipEntry(c *gin.Context, db *gorm.DB) {
	var newEntryRequest CreateNewEntryRequest

	if err := c.ShouldBindJSON(&newEntryRequest); err != nil {
		responses.GenericBadRequestError(c.Writer)
		return
	}

	_, err := auth.AuthenticateByHeader(c, db)
	if err != nil {
		responses.GenericUnauthorizedError(c.Writer)
		return
	}

	newEntry := models.SkipEntry{
		Reason: newEntryRequest.Reason,
	}

	userId, err := uuid.Parse(newEntryRequest.UserID)
	if err != nil {
		responses.GenericBadRequestError(c.Writer, "Invalid user ID structure")
		return
	}
	newEntry.UserID = userId

	lessonId, err := uuid.Parse(newEntryRequest.LessonID)
	if err != nil {
		responses.GenericBadRequestError(c.Writer, "Invalid lesson ID structure")
		return
	}
	newEntry.LessonID = lessonId

	amountId, err := uuid.Parse(newEntryRequest.AmountID)
	if err != nil {
		responses.GenericBadRequestError(c.Writer, "Invalid amount ID structure")
		return
	}
	newEntry.AmountID = amountId

	err = createSkipEntry(newEntry, db)
	if err != nil {
		responses.GenericInternalServerError(c.Writer)
		return
	}

	c.JSON(http.StatusCreated, newEntry)

}

func EditSkipEntry(c *gin.Context, db *gorm.DB) {
	var newEntryRequest EditSkipEntryRequest
	if err := c.ShouldBind(&newEntryRequest); err != nil {
		responses.GenericBadRequestError(c.Writer)
		return
	}

	_, err := auth.AuthenticateByHeader(c, db)
	if err != nil {
		responses.GenericUnauthorizedError(c.Writer)
		return
	}
	updateEntry := models.SkipEntry{
		Reason: newEntryRequest.Reason,
	}
	userId, err := uuid.Parse(newEntryRequest.UserID)
	if err != nil {
		responses.GenericBadRequestError(c.Writer, "Invalid user ID structure")
		return
	}
	updateEntry.UserID = userId

	lessonId, err := uuid.Parse(newEntryRequest.LessonID)
	if err != nil {
		responses.GenericBadRequestError(c.Writer, "Invalid lesson ID structure")
		return
	}
	updateEntry.LessonID = lessonId

	amountId, err := uuid.Parse(newEntryRequest.AmountID)
	if err != nil {
		responses.GenericBadRequestError(c.Writer, "Invalid amount ID structure")
		return
	}
	updateEntry.AmountID = amountId

	id, err := uuid.Parse(newEntryRequest.ID)
	if err != nil {
		responses.GenericBadRequestError(c.Writer, "Invalid ID structure")
		return
	}
	updateEntry.ID = id

	err = updateSkipEntry(&updateEntry, db)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responses.GenericNotFoundError(c.Writer)
			return
		}
		responses.GenericInternalServerError(c.Writer)
		return
	}

	c.JSON(http.StatusCreated, updateEntry)
}
