package models

import (
	"time"

	"github.com/google/uuid"
)

type Lesson struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name          string    `gorm:"type:varchar(100);not null"`
	StartDateTime time.Time `gorm:"not null"`
	EndDateTime   time.Time `gorm:"not null"`
	ClassID       uuid.UUID `gorm:"type:uuid;not null;index"`
	Class         Class     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	SkipEntries []SkipEntry `gorm:"foreignKey:lesson_id"`
	CreatedAt   time.Time   `gorm:"type:timestampz;not null;autoCreateTimes"`
	UpdatedAt   time.Time   `gorm:"type:timestampz;not null;autoCreateTimes;autoUpdateTimess"`
}
