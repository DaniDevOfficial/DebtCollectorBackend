package models

import (
	"time"

	"github.com/google/uuid"
)

type Class struct {
	ID         uuid.UUID `gorm:"primaryKey;type:uuid;"`
	Name       string    `gorm:"type:varchar(100);not null"`
	SemesterID uuid.UUID `gorm:"type:uuid;not null;index"`
	Semester   Semester  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time `gorm:"type:timestamptz;not null;autoCreateTime"`
	UpdatedAt time.Time `gorm:"type:timestamptz;not null;autoCreateTime;autoUpdateTimes"`
}
