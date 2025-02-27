package routes

import (
	"Chess-Go/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	authController := &controllers.AuthController{}
	authRouter := router.Group("/auth")
	authRouter.POST("/login", authController.Login)
}
