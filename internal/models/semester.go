package models

import (
	"time"

	"github.com/google/uuid"
)

type Semester struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string    `gorm:"type:varchar(100);not null"`
	StartDate time.Time `gorm:"not null"`
	EndDate   time.Time `gorm:"not null"`
	Classes   []Class   `gorm:"foreignKey:SemesterID"`

	CreatedAt time.Time `gorm:"type:timestampz;not null;autoCreateTimes"`
	UpdatedAt time.Time `gorm:"type:timestampz;not null;autoCreateTimes;autoUpdateTimess"`
}
