package utils

import (
	"github.com/gin-gonic/gin"
	"pweb-api.abdanhafidz.com/models"
)

func GetAccount(c *gin.Context) models.AccountData {
	cParam, _ := c.Get("accountData")
	return cParam.(models.AccountData)
}
