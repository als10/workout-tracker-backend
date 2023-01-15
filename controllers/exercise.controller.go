package controllers

import (
	"net/http"

	"github.com/als10/workout-tracker-backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ExerciseController struct {
	DB *gorm.DB
}

func NewExerciseController(DB *gorm.DB) ExerciseController {
	return ExerciseController{DB}
}

func (ec *ExerciseController) GetExercises(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)

	var exercises []models.Exercise
	results := ec.DB.Where("user_id = ?", currentUser.ID).Preload("Progressions").Find(&exercises)

	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(exercises), "data": exercises})
}

func (ec *ExerciseController) CreateExercise(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var payload *models.Exercise

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newExercise := models.Exercise{
		UserID: currentUser.ID,
		Name: payload.Name,
		Progressions: payload.Progressions,
	}

	result := ec.DB.Create(&newExercise)
	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newExercise})
}

func (ec *ExerciseController) UpdateExercise(ctx *gin.Context) {
	exerciseId := ctx.Param("exerciseId")
	currentUser := ctx.MustGet("currentUser").(models.User)

	var payload *models.Exercise
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	var exerciseToUpdate models.Exercise
	result := ec.DB.First(&exerciseToUpdate, "id = ?", exerciseId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Exercise with given ID does not exist!"})
		return
	}
	if exerciseToUpdate.UserID != currentUser.ID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You can only update exercises that you have created!"})
		return
	}

	updatedExercise := models.Exercise{
		UserID: currentUser.ID,
		Name: payload.Name,
		Progressions: payload.Progressions,
	}

	ec.DB.Model(&exerciseToUpdate).Updates(updatedExercise)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedExercise})
}

func (ec *ExerciseController) DeleteExercise(ctx *gin.Context) {
	exerciseId := ctx.Param("exerciseId")
	currentUser := ctx.MustGet("currentUser").(models.User)

	var exerciseToDelete models.Exercise
	result := ec.DB.First(&exerciseToDelete, "id = ?", exerciseId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Exercise with given ID does not exist!"})
		return
	}
	if exerciseToDelete.UserID != currentUser.ID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You can only delete exercises that you have created!"})
		return
	}

	ec.DB.Delete(&exerciseToDelete)
	ctx.JSON(http.StatusNoContent, nil)
}
