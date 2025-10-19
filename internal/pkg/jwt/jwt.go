package jwt

import (
	"dept-collector/internal/models"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var secretKey = []byte("capybara") // TODO: add secret key via .env or some rotation

type NewRefreshTokenDataDB struct {
	UserId       string     `json:"userId"`
	RefreshToken string     `json:"refresh_token"`
	LifeTime     *time.Time `json:"lifeTime"`
}

func CreateToken(userData User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"UserId":   userData.UserId,
			"Username": userData.Username,
			"Exp":      time.Now().Add(time.Minute * 5).Unix(),
		})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) (isValid bool, jwtData Payload, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return false, jwtData, err
	}

	if !token.Valid {
		return false, jwtData, nil
	}

	jwtData, err = DecodeBearer(tokenString)
	if err != nil {
		if errors.Is(err, TokenIsNotValidDueToExpirationDate) {
			return false, jwtData, nil

		}
		return false, jwtData, err
	}

	return true, jwtData, nil
}

var RefreshTokenNotInDbError = errors.New("refresh token not found in database")
var TokenIsNotValidDueToExpirationDate = errors.New("token is not valid due to expiration date")

func VerifyRefreshToken(tokenString string, db *gorm.DB) (Payload, error) {

	isValid, payload, err := VerifyToken(tokenString)
	if err != nil || !isValid {
		return Payload{}, err
	}

	inDB, err := VerifyRefreshTokenInDB(tokenString, payload.UserId, db)
	if err != nil {
		return payload, err
	}
	if !inDB {
		return payload, RefreshTokenNotInDbError
	}

	return payload, nil
}

func VerifyRefreshTokenInDB(token string, userId string, db *gorm.DB) (bool, error) {

	var tokenData models.RefreshToken
	result := db.Where("refresh_token = ?", token).First(&tokenData)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}

	if (tokenData.UserID.String() == userId) || (tokenData.ExpiresAt.Before(time.Now())) {
		return false, nil
	}

	return true, nil
}

func CreateRefreshToken(userData User, isTimeBased bool, db *gorm.DB) (string, error) {
	t := time.Now().Add(time.Hour * 24 * 365)

	if isTimeBased {
		t = time.Now().Add(time.Hour * 24 * 14)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"UserId":   userData.UserId,
			"Username": userData.Username,
			"Exp":      t.Unix(),
		})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	userUUID, err := uuid.Parse(userData.UserId)
	if err != nil {
		return "", err
	}

	data := CreateTokenInput{
		UserID:       userUUID,
		RefreshToken: tokenString,
		ExpiresAt:    t,
	}

	err = PushRefreshTokenToDB(data, db)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return tokenString, nil
}

func DecodeBearer(tokenString string) (Payload, error) {
	splitToken := strings.Split(tokenString, ".")
	if len(splitToken) != 3 {
		return Payload{}, fmt.Errorf("invalid token format")
	}

	payloadSegment := splitToken[1]
	payloadBytes, err := base64.RawURLEncoding.DecodeString(payloadSegment)
	if err != nil {
		return Payload{}, fmt.Errorf("failed to decode payload: %v", err)
	}

	var payload Payload
	err = json.Unmarshal(payloadBytes, &payload)
	if err != nil {
		return Payload{}, fmt.Errorf("failed to unmarshal payload: %v", err)
	}

	if payload.Exp > 0 {
		if payload.Exp < time.Now().Unix() {
			return payload, TokenIsNotValidDueToExpirationDate
		}
	}

	return payload, nil
}

func PushRefreshTokenToDB(data CreateTokenInput, db *gorm.DB) error {
	token := models.RefreshToken{
		UserID:       data.UserID,
		RefreshToken: data.RefreshToken,
		ExpiresAt:    data.ExpiresAt,
	}

	if err := db.Create(&token).Error; err != nil {
		return err
	}

	return nil
}

func VoidRefreshTokenInDB(token string, db *gorm.DB) error {
	db.Delete(&models.RefreshToken{RefreshToken: token})
	return nil
}
