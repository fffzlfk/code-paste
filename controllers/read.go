package controllers

import (
	"code-paste/database"
	"code-paste/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadPaste(c *gin.Context) {
	id := c.Param("id")
	var p model.Paste
	res := database.DB.First(&p, "id = ?", id)
	if res.Error != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusAccepted, p)
}
