package router

import (
	"github.com/gin-gonic/gin"
	AuthController "pweb-api.abdanhafidz.com/controller/auth"
)

func AuthRoute(router *gin.Engine) {
	routerGroup := router.Group("/api/v1/auth")
	{
		routerGroup.POST("/login", AuthController.Login)
		routerGroup.POST("/register", AuthController.Register)

	}
}
