package helpers

import (
	"oosa_rewild/internal/models"

	"github.com/gin-gonic/gin"
)

func GetAuthUser(c *gin.Context) models.Users {
	user, exists := c.Get("user")
	if exists {
		userDetail := user.(*models.Users)
		return *userDetail
	} else {
		return models.Users{}
	}
}
