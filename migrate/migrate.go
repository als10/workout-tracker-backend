package main

import (
	"fmt"
	"log"

	"github.com/als10/workout-tracker-backend/initializers"
	"github.com/als10/workout-tracker-backend/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func Migrate() {
	initializers.DB.AutoMigrate(
		&models.User{},
		&models.Exercise{},
		&models.Progression{},
		&models.Workout{},
		&models.ExerciseSet{},
		&models.ProgressionSet{},
	)
}

func main() {
	Migrate()
	fmt.Println("? Migration complete")
}
