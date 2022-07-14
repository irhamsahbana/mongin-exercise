package main

import (
	"github.com/gin-gonic/gin"
	"github.com/irhamsahbana/mongin-exercise/configs"
	"github.com/irhamsahbana/mongin-exercise/routes"
)


func main() {
	router := gin.Default()

	// run database
	configs.ConnectDB()

	// routes
	routes.UserRoute(router)

	router.Run(":4000")
}