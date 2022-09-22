package controllers

import (
	"code-paste/database"
	"code-paste/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid/v4"
)

func generateUUID() string {
	return shortuuid.New()
}

func CreatePaste(c *gin.Context) {
	var p model.Paste
	if err := c.BindJSON(&p); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	p.ExpiredAt = time.Now().AddDate(0, 0, p.ExpiredDays)
	p.ID = generateUUID()
	database.DB.Create(&p)
	c.JSON(http.StatusAccepted, gin.H{
		"status": "ok",
		"uuid":   p.ID,
	})
}
