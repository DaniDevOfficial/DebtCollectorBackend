package models

import (
	"time"

	"github.com/google/uuid"
)

type Amount struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey"`
	Value  float64   `gorm:"type:numeric(10,2);not null"`
	Name   string    `gorm:"type:varchar(255);not null"`
	Reason string    `gorm:"type:text;not null"`

	CreatedAt time.Time `gorm:"type:timestamptz;autoCreateTime"`
	UpdatedAt time.Time `gorm:"type:timestamptz;autoUpdateTime;autoUpdateTime"`
}
