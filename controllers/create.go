package controllers

import (
	"code-paste/database"
	"code-paste/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func generateUUID() string {
	return uuid.New().String()
}

func CreatePaste(c *gin.Context) {
	var p model.Paste
	if err := c.BindJSON(&p); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	p.ExpireAt = time.Now().Add(time.Hour * 12)
	p.ID = generateUUID()
	database.DB.Create(&p)
	c.JSON(http.StatusAccepted, gin.H{
		"status": "ok",
		"uuid":   p.ID,
	})
}
