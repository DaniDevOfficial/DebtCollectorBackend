package jwt

import (
	"time"

	"github.com/google/uuid"
)

type CreateTokenInput struct {
	UserID       uuid.UUID
	RefreshToken string
	ExpiresAt    time.Time
}
