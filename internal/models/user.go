package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name      string    `gorm:"type:varchar(100);not null"`
	Email     string    `gorm:"type:varchar(150);uniqueIndex;not null"`
	Password  string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"type:timestamptz;not null;autoCreateTime"`
	UpdatedAt time.Time `gorm:"type:timestamptz;not null;autoCreateTime;autoUpdateTime"`
}
