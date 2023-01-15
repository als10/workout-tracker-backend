package models

import "github.com/google/uuid"

type Progression struct {
	Model
	Name string `gorm:"not null" json:"name" binding:"required"`
	Rank uint `gorm:"not null" json:"rank" binding:"required"`
	ExerciseID uuid.UUID `gorm:"type:uuid;not null"`
}