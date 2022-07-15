package main

import (
	"code-paste/controllers"
	"code-paste/cron"
	"code-paste/database"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type configs struct {
	Server   serverConfig
	Database databaseConfig
}

type serverConfig struct {
	Port int
}

type databaseConfig struct {
	Host     string
	User     string
	Password string
	Port     int
}

func main() {
	var config configs
	toml.DecodeFile("./config.toml", &config)

	logFile, _ := os.Create("log.txt")

	database.Init(logFile, config.Database.Host, config.Database.User, config.Database.Password, config.Database.Port)

	cron.Start()

	e := gin.Default()

	e.Use(gin.LoggerWithWriter(logFile))

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	e.Use(cors.New(corsConfig))

	e.POST("/api/create", controllers.CreatePaste)
	e.GET("/api/read/:id", controllers.ReadPaste)

	e.Run(fmt.Sprintf(":%d", config.Server.Port))
}
