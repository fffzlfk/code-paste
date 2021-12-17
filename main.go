package main

import (
	"code-paste/controllers"
	"code-paste/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()

	e := gin.New()
	e.Use(gin.Logger())
	e.Use(cors.Default())

	e.POST("/api/create", controllers.CreatePaste)
	e.GET("/api/read/:id", controllers.ReadPaste)

	e.Run(":8080")
}
