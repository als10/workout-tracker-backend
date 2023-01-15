package controllers

import (
	"net/http"

	"github.com/als10/workout-tracker-backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type WorkoutController struct {
	DB *gorm.DB
}

func NewWorkoutController(DB *gorm.DB) WorkoutController {
	return WorkoutController{DB}
}

func (wc *WorkoutController) GetWorkouts(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)

	var workouts []models.Workout
	results := wc.DB.Where("user_id = ?", currentUser.ID)
	results = results.Preload("ExerciseSets.Exercise")
	results = results.Preload("ExerciseSets.ProgressionSets.Progression")
	results = results.Find(&workouts)

	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(workouts), "data": workouts})
}

func (wc *WorkoutController) CreateWorkout(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var payload *models.Workout

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newWorkout := models.Workout{
		UserID: currentUser.ID,
		ExerciseSets: payload.ExerciseSets,
	}

	result := wc.DB.Create(&newWorkout)
	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newWorkout})
}

func (wc *WorkoutController) UpdateWorkout(ctx *gin.Context) {
	workoutId := ctx.Param("workoutId")
	currentUser := ctx.MustGet("currentUser").(models.User)

	var payload *models.Workout
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	var workoutToUpdate models.Workout
	result := wc.DB.First(&workoutToUpdate, "id = ?", workoutId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Workout with given ID does not exist!"})
		return
	}
	if workoutToUpdate.UserID != currentUser.ID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You can only update workouts that you have created!"})
		return
	}

	updatedWorkout := models.Workout{
		UserID: currentUser.ID,
		ExerciseSets: payload.ExerciseSets,
	}

	wc.DB.Model(&workoutToUpdate).Updates(updatedWorkout)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedWorkout})
}

func (wc *WorkoutController) DeleteWorkout(ctx *gin.Context) {
	workoutId := ctx.Param("workoutId")
	currentUser := ctx.MustGet("currentUser").(models.User)

	var workoutToDelete models.Workout
	result := wc.DB.First(&workoutToDelete, "id = ?", workoutId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Workout with given ID does not exist!"})
		return
	}
	if workoutToDelete.UserID != currentUser.ID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You can only delete workouts that you have created!"})
		return
	}

	wc.DB.Delete(&workoutToDelete)
	ctx.JSON(http.StatusNoContent, nil)
}
