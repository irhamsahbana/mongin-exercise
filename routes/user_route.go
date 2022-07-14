package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/irhamsahbana/mongin-exercise/controllers"
)

func UserRoute(router *gin.Engine) {
	router.GET("/users", controllers.GetUsers())
	router.POST("/users", controllers.CreateUSer())
	router.GET("/users/:userId", controllers.GetUser())
	router.PUT("/users/:userId", controllers.EditUser())
	router.DELETE("/users/:userId", controllers.DeleteUser())
}