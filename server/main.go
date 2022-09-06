package main

import (
	"code-paste/controllers"
	"code-paste/cron"
	"code-paste/database"
	"fmt"
	"os"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	database.Init()

	cron.Start()

	e := gin.Default()

	e.Use(gin.LoggerWithWriter(os.Stdout))

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	e.Use(cors.New(corsConfig))

	store := persistence.NewInMemoryStore(time.Second)

	e.POST("/api/create", controllers.CreatePaste)
	e.GET("/api/read/:id", cache.CachePage(store, time.Minute, controllers.ReadPaste))

	e.Run(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
}
