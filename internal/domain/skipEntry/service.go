package skipEntry

import (
	"dept-collector/internal/models"
	"dept-collector/internal/pkg/responses"
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
