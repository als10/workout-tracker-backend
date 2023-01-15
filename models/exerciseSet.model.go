package models

import "github.com/google/uuid"

type ExerciseSet struct {
	Model
	WorkoutID uuid.UUID `gorm:"type:uuid;not null"`
	ExerciseID uuid.UUID `gorm:"type:uuid;not null" json:"exerciseId" binding:"required"`
	Exercise Exercise
	ProgressionSets []ProgressionSet `json:"progressionSets" binding:"required"`
}
