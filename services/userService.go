package services

import (
	"crud-api/models"
	"crud-api/serializers"
	"crud-api/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UsersRegistration(c *gin.Context) {
	userModelValidator := validators.NewUserModelValidator()
	if err := userModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, validators.ValidatorError(err))
		return
	}
	if err := models.Handler.Create(userModelValidator.userModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, serializers.ResponseError("database", err))
	}
	c.JSON(http.StatusOK, userModelValidator.userModel)
}
