package models

import "github.com/google/uuid"

type Workout struct {
	Model
	UserID uuid.UUID `gorm:"type:uuid;not null"`
	ExerciseSets []ExerciseSet `json:"exerciseSets" binding:"required"`
}
