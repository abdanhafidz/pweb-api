package auth

import (
	"github.com/gin-gonic/gin"
	"pweb-api.abdanhafidz.com/controller"
	"pweb-api.abdanhafidz.com/models"
	"pweb-api.abdanhafidz.com/services"
)

func Login(c *gin.Context) {
	authentication := services.AuthenticationService{}
	loginController := controller.Controller[models.LoginRequest, models.Account, models.AuthenticatedUser]{
		Service: &authentication.Service,
	}
	loginController.RequestJSON(c, func() {
		loginController.Service.Constructor.Email = loginController.Request.Email
		loginController.Service.Constructor.Password = loginController.Request.Password
		authentication.Authenticate()
	})
}
