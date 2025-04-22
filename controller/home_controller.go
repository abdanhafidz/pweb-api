package controller

import "github.com/gin-gonic/gin"

func HomeController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "PWEB API 2025 by Abdan Hafidz!",
	})
}
