package main

import (
	"code-paste/controllers"
	"code-paste/database"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type configs struct {
	Server   serverConfig
	Frontend frontendConfig
	Database databaseConfig
}

type serverConfig struct {
	Port int
}

type frontendConfig struct {
	Host string
	Port int
}

type databaseConfig struct {
	Host     string
	User     string
	Password string
}

func main() {
	var config configs
	toml.DecodeFile("./config.toml", &config)

	database.Init(config.Database.Host, config.Database.User, config.Database.Password)

	e := gin.New()

	front := fmt.Sprintf("http://%s:%d",
		config.Frontend.Host,
		config.Frontend.Port,
	)

	e.SetTrustedProxies([]string{front})

	logFile, _ := os.Create("log.txt")
	e.Use(gin.LoggerWithWriter(logFile))

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{front}
	e.Use(cors.New(corsConfig))

	e.POST("/api/create", controllers.CreatePaste)
	e.GET("/api/read/:id", controllers.ReadPaste)

	e.Run(fmt.Sprintf(":%d", config.Server.Port))
}
