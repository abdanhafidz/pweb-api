package user

import (
	"github.com/gin-gonic/gin"
	"pweb-api.abdanhafidz.com/controller"
	"pweb-api.abdanhafidz.com/models"
	"pweb-api.abdanhafidz.com/services"
)

func Profile(c *gin.Context) {
	userProfile := services.UserProfileService{}
	userProfileController := controller.Controller[any, models.AccountDetails, models.UserProfileResponse]{
		Service: &userProfile.Service,
	}
	userProfileController.HeaderParse(c, func() {
		userProfileController.Service.Constructor.AccountId = uint(userProfileController.AccountData.UserID)
		userProfile.Retrieve()
		userProfileController.Response(c)
	},
	)
}
