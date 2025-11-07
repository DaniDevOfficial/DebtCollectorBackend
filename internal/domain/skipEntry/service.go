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

// CreateNewSkipEntry godoc
// @Summary      Creates a new Skip entry
// @Description  Creates a new Skip entry and returns said entry
// @Tags         SkipEntries
// @Accept       json
// @Produce      json
// @Param        request body CreateNewEntryRequest true "Create new Skipentry"
// @Success      200  {string}  responseTypes.SkipEntryResponse
// @Failure      400  {string}  bad Request
// @Failure      500  {string}  internal server error
// @Router       /skips [post]
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
		ID:     uuid.New(),
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

	err = createSkipEntry(&newEntry, db)
	if err != nil {
		responses.GenericInternalServerError(c.Writer)
		return
	}

	c.JSON(http.StatusCreated, buildSkipEntryResponse(&newEntry))
}

// EditSkipEntry godoc
// @Summary      Edit a Skip entry
// @Description  Edit a skip entry based on id and returns the new entry
// @Tags         SkipEntries
// @Accept       json
// @Produce      json
// @Param        request body EditSkipEntryRequest true "Edit skip entry"
// @Success      200  {string}  responseTypes.SkipEntryResponse
// @Failure      400  {string}  bad Request
// @Failure      500  {string}  internal server error
// @Router       /skips [post]
func EditSkipEntry(c *gin.Context, db *gorm.DB) {
	var newEntryRequest EditSkipEntryRequest
	if err := c.ShouldBindJSON(&newEntryRequest); err != nil {
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

	c.JSON(http.StatusCreated, buildSkipEntryResponse(&updateEntry))
}

func DeleteSkipEntry(c *gin.Context, db *gorm.DB) {
	var deleteEntryRequest SingleIdRequest

	if err := c.ShouldBindJSON(&deleteEntryRequest); err != nil {
		responses.GenericBadRequestError(c.Writer)
		return
	}

	_, err := auth.AuthenticateByHeader(c, db)
	if err != nil {
		responses.GenericUnauthorizedError(c.Writer)
		return
	}

	entryId, err := uuid.Parse(deleteEntryRequest.ID)
	if err != nil {
		responses.GenericBadRequestError(c.Writer, "Invalid ID structure")
		return
	}

	err = deleteSkipEntry(entryId, db)
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

func GetSpecificSkipEntry(c *gin.Context, db *gorm.DB) {
	var getEntryRequest SingleIdRequest
	if err := c.ShouldBindJSON(&getEntryRequest); err != nil {
		responses.GenericBadRequestError(c.Writer)
		return
	}
	_, err := auth.AuthenticateByHeader(c, db)
	if err != nil {
		responses.GenericUnauthorizedError(c.Writer)
		return
	}
	entryId, err := uuid.Parse(getEntryRequest.ID)
	if err != nil {
		responses.GenericBadRequestError(c.Writer, "Invalid ID structure")
		return
	}
	entry, err := getSpecificEntry(entryId, db)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responses.GenericNotFoundError(c.Writer)
			return
		}
		responses.GenericInternalServerError(c.Writer)
		return
	}

	c.JSON(http.StatusOK, buildSkipEntryResponse(entry))
}

func GetFilteredSkipEntries(c *gin.Context, db *gorm.DB) {
	var getEntryRequest FilterSkipEntryRequest
	if err := c.ShouldBindJSON(&getEntryRequest); err != nil {
		responses.GenericBadRequestError(c.Writer)
		return
	}
	_, err := auth.AuthenticateByHeader(c, db)
	if err != nil {
		responses.GenericUnauthorizedError(c.Writer)
		return
	}

	entries, err := getAllEntries(getEntryRequest, db)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responses.GenericNotFoundError(c.Writer)
			return
		}
		responses.GenericInternalServerError(c.Writer)
		return
	}

	c.JSON(http.StatusOK, buildSkipEntriesResponse(entries))
}
