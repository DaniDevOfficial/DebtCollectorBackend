package models

import (
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	RefreshToken string    `gorm:"type:text;not null;unique"`
	UserID       uuid.UUID `gorm:"type:uuid;not null;index"` // FK field
	User         User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ExpiresAt    time.Time `gorm:"type:timestamptz;not null"`
	CreatedAt    time.Time `gorm:"type:timestamptz;autoCreateTime"`
}
