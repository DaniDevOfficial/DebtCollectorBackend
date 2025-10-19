package models

import (
	"time"

	"github.com/google/uuid"
)

type SkipEntry struct {
	ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID uuid.UUID `gorm:"type:uuid;default:"`

	UpdatedAt time.Time `gorm:"type:timestamptz;not null;autoCreateTime;autoUpdateTime"`
}
