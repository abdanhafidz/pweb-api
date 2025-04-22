package user

import (
	"github.com/gin-gonic/gin"
	"pweb-api.abdanhafidz.com/controller"
	"pweb-api.abdanhafidz.com/models"
	"pweb-api.abdanhafidz.com/services"
)

func UpdateProfile(c *gin.Context) {
	userProfile := services.UserProfileService{}
	userUpdateProfileController := controller.Controller[models.AccountDetails, models.AccountDetails, models.UserProfileResponse]{
		Service: &userProfile.Service,
	}

	userUpdateProfileController.RequestJSON(c, func() {
		userUpdateProfileController.Service.Constructor = userUpdateProfileController.Request
		userUpdateProfileController.HeaderParse(c, func() {
			userUpdateProfileController.Service.Constructor.AccountId = uint(userUpdateProfileController.AccountData.UserID)

		})
		userProfile.Update()
	},
	)
}
