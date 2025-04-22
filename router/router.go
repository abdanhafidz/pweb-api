package router

import (
	"log"

	"github.com/gin-gonic/gin"
	"pweb-api.abdanhafidz.com/config"
	"pweb-api.abdanhafidz.com/controller"
)

func StartService() {
	router := gin.Default()
	router.GET("/", controller.HomeController)

	AuthRoute(router)
	UserRoute(router)
	err := router.Run(config.TCP_ADDRESS)
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
