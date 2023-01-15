package models

import "github.com/google/uuid"

type Exercise struct {
	Model
	Name string `gorm:"not null" json:"name" binding:"required"`
	UserID uuid.UUID `gorm:"type:uuid;not null"`
	Progressions []Progression `json:"progressions" binding:"required"`
}