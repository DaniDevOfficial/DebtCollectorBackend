package user

import (
	"dept-collector/internal/models"
	"dept-collector/internal/pkg/auth"
	"dept-collector/internal/pkg/frontendErrors"
	"dept-collector/internal/pkg/hashing"
	"dept-collector/internal/pkg/jwt"
	"dept-collector/internal/pkg/responses"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// SignUp godoc
// @Summary      Register a new user
// @Description  Creates a new account and returns JWT + refresh token
// @Tags         auth
// @Accept       json
// @Param        request body CreateNewUserRequest true "New account data"
// @Success      200  {string}  string  "JWT and Refresh tokens in headers"
// @Failure      400  {string}  bad Request
// @Failure      500  {string}  internal server error
// @Router       /user/signup [post]
func SignUp(c *gin.Context, db *gorm.DB) {
	var newAccountRequest CreateNewUserRequest
	if err := c.ShouldBind(&newAccountRequest); err != nil {
		responses.GenericBadRequestError(c.Writer)
		return
	}

	taken, err := isUsernameOrEmailTaken(newAccountRequest.Username, newAccountRequest.Email, db)
	if err != nil {
		responses.GenericInternalServerError(c.Writer)
		return
	}
	if taken {
		responses.HttpErrorResponse(c.Writer, http.StatusUnauthorized, frontendErrors.UsernameOrEmailAlreadyTaken, "")
		return
	}

	passwordHash, err := hashing.HashPassword(newAccountRequest.Password)
	if err != nil {
		responses.GenericInternalServerError(c.Writer)
		return
	}

	newUser := models.User{
		ID:       uuid.New(),
		Name:     newAccountRequest.Username,
		Email:    newAccountRequest.Email,
		Password: passwordHash,
	}
	err = createNewUser(&newUser, db)
	if err != nil {
		responses.GenericInternalServerError(c.Writer)
		return
	}

	jwtUserData := jwt.User{
		Username: newUser.Name,
		UserId:   newUser.ID.String(),
	}

	refreshToken, err := jwt.CreateRefreshToken(jwtUserData, false, db)
	if err != nil {
		responses.GenericInternalServerError(c.Writer)
		return
	}

	jwtToken, err := jwt.CreateToken(jwtUserData)
	if err != nil {
		log.Println(err)
		responses.GenericInternalServerError(c.Writer)
		return
	}

	c.Header("Authorization", jwtToken)
	c.Header("RefreshToken", refreshToken)

	c.JSON(http.StatusOK, "")
}

// Login godoc
// @Summary      Login to a existing account
// @Description  Login into account and returns JWT + refresh token
// @Tags         auth
// @Accept       json
// @Param        request body LoginRequest true "Login credentials"
// @Success      200  {string}  string  "JWT and Refresh tokens in headers"
// @Failure      400  {string}  bad Request
// @Failure      500  {string}  internal server error
// @Router       /user/login [post]
func Login(c *gin.Context, db *gorm.DB) {
	var loginRequest LoginRequest

	if err := c.ShouldBind(&loginRequest); err != nil {
		responses.GenericBadRequestError(c.Writer)
		return
	}

	user, err := getUserByName(loginRequest.Username, db)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responses.HttpErrorResponse(c.Writer, http.StatusUnauthorized, frontendErrors.UsernameOrPasswordAreWrong, "")
			return
		}
		responses.GenericInternalServerError(c.Writer)
		return
	}
	if !hashing.CheckHashedString(user.Password, loginRequest.Password) {
		responses.HttpErrorResponse(c.Writer, http.StatusUnauthorized, frontendErrors.UsernameOrPasswordAreWrong, "")
		return
	}

	jwtUserData := jwt.User{
		Username: user.Name,
		UserId:   user.ID.String(),
	}

	refreshToken, err := jwt.CreateRefreshToken(jwtUserData, false, db)
	if err != nil {
		responses.GenericInternalServerError(c.Writer)
		return
	}

	jwtToken, err := jwt.CreateToken(jwtUserData)
	if err != nil {
		responses.GenericInternalServerError(c.Writer)
		return
	}
	c.Header("Authorization", jwtToken)
	c.Header("RefreshToken", refreshToken)

	c.JSON(http.StatusOK, "")
}

// CheckAuth godoc
// @Summary      Check the auth status
// @Description  Checks the auth status of a refresh and jwt token "JWT and Refresh tokens in headers"
// @Tags         auth
// @Success      200  {string}  auth is valid
// @Failure      401  {string}  auth is not valid
// @Router       /user/checkAuth [get]
func CheckAuth(c *gin.Context, db *gorm.DB) {
	_, err := auth.AuthenticateByHeader(c, db)
	if err != nil {
		responses.GenericUnauthorizedError(c.Writer)
		return
	}

	c.JSON(http.StatusOK, "")
}
