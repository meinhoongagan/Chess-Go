package main

import (
	"Chess-Go/config"
	"Chess-Go/database"
	"Chess-Go/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadConfig()

	database.Connect()
	// database.MigrateDB()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	routes.AuthRoutes(r)
	r.Run(":8000") // listen and serve on 0.0.0.0:8000
	fmt.Println("hello")
}
