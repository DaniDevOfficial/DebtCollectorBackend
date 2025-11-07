package models

import (
	"time"

	"github.com/google/uuid"
)

type SkipEntry struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey"`
	Reason string    `gorm:"type:text;not null"`

	UserID uuid.UUID `gorm:"type:uuid;default:"`
	User   User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	LessonID uuid.UUID `gorm:"type:uuid;not null;index"`
	Lesson   Lesson    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	AmountID uuid.UUID `gorm:"type:uuid;not null;index"`
	Amount   Amount    `gorm:"foreignKey:AmountID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`

	CreatedAt time.Time `gorm:"type:timestamptz;not null;autoCreateTime"`
	UpdatedAt time.Time `gorm:"type:timestamptz;not null;autoCreateTime;autoUpdateTime"`
}
