package routes

import (
	"github.com/als10/workout-tracker-backend/controllers"
	"github.com/als10/workout-tracker-backend/middleware"
	"github.com/gin-gonic/gin"
)

type ExerciseRouteController struct {
	exerciseController controllers.ExerciseController
}

func NewExerciseRouteController(exerciseController controllers.ExerciseController) ExerciseRouteController {
	return ExerciseRouteController{exerciseController}
}

func (ec *ExerciseRouteController) ExerciseRoute(rg *gin.RouterGroup) {
	router := rg.Group("exercise")
	
	router.Use(middleware.DeserializeUser())
	router.POST("/", ec.exerciseController.CreateExercise)
	router.GET("/", ec.exerciseController.GetExercises)
	router.PUT("/:exerciseId", ec.exerciseController.UpdateExercise)
	router.DELETE("/:exerciseId", ec.exerciseController.DeleteExercise)
}

