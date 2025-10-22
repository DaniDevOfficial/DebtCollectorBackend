package models

import (
	"time"

	"github.com/google/uuid"
)

type Class struct {
	ID         string    `gorm:"primaryKey;type:uuid;"`
	Name       string    `gorm:"type:varchar(100);not null"`
	SemesterID uuid.UUID `gorm:"type:uuid;not null;index"`
	Semester   Semester  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time `gorm:"type:timestampz;not null;autoCreateTimes"`
	UpdatedAt time.Time `gorm:"type:timestampz;not null;autoCreateTimes;autoUpdateTimess"`
}
