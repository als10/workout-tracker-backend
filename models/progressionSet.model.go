package models

import "github.com/google/uuid"

type ProgressionSet struct {
	Model
	Reps uint `gorm:"not null" json:"reps" binding:"required"`
	ExerciseSetID uuid.UUID `gorm:"type:uuid;not null"`
	ProgressionID uuid.UUID `gorm:"type:uuid;not null" json:"progressionId" binding:"required"`
	Progression Progression
}
