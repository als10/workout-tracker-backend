package routes

import (
	"github.com/als10/workout-tracker-backend/controllers"
	"github.com/als10/workout-tracker-backend/middleware"
	"github.com/gin-gonic/gin"
)

type WorkoutRouteController struct {
	workoutController controllers.WorkoutController
}

func NewWorkoutRouteController(workoutController controllers.WorkoutController) WorkoutRouteController {
	return WorkoutRouteController{workoutController}
}

func (wc *WorkoutRouteController) WorkoutRoute(rg *gin.RouterGroup) {
	router := rg.Group("workout")
	
	router.Use(middleware.DeserializeUser())
	router.POST("/", wc.workoutController.CreateWorkout)
	router.GET("/", wc.workoutController.GetWorkouts)
	router.PUT("/:workoutId", wc.workoutController.UpdateWorkout)
	router.DELETE("/:workoutId", wc.workoutController.DeleteWorkout)
}

