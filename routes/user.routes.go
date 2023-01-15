package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/als10/workout-tracker-backend/controllers"
	"github.com/als10/workout-tracker-backend/middleware"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewRouteUserController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(rg *gin.RouterGroup) {
	router := rg.Group("users")

	router.GET("/me", middleware.DeserializeUser(), uc.userController.GetMe)
}

