package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/als10/workout-tracker-backend/controllers"
	"github.com/als10/workout-tracker-backend/initializers"
	"github.com/als10/workout-tracker-backend/routes"
)

var (
	server *gin.Engine

	AuthController controllers.AuthController
	AuthRouteController routes.AuthRouteController

	UserController controllers.UserController
	UserRouteController routes.UserRouteController

	WorkoutController controllers.WorkoutController
	WorkoutRouteController routes.WorkoutRouteController

	ExerciseController controllers.ExerciseController
	ExerciseRouteController routes.ExerciseRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	AuthController = controllers.NewAuthController(initializers.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(initializers.DB)
	UserRouteController = routes.NewRouteUserController(UserController)

	WorkoutController = controllers.NewWorkoutController(initializers.DB)
	WorkoutRouteController = routes.NewWorkoutRouteController(WorkoutController)

	ExerciseController = controllers.NewExerciseController(initializers.DB)
	ExerciseRouteController = routes.NewExerciseRouteController(ExerciseController)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")

	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})
	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router)
	WorkoutRouteController.WorkoutRoute(router)
	ExerciseRouteController.ExerciseRoute(router)

	log.Fatal(server.Run(":" + config.ServerPort))
}
