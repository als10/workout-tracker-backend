package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Model
	Email string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Workouts []Workout
	Exercises []Exercise
}

type SignUpInput struct {
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
}

type SignInInput struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
