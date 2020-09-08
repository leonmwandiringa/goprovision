package handlers

import (
	"github.com/gin-gonic/gin"
)

func IsApplicationWorking(c *gin.Context){
	c.JSON(200, gin.H{ 
		"error": nil,
		"status":   "sucess",
		"message" : "Indeed its is working",
	})   
	return
}
