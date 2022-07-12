package endpoints

import (
	"crud-api/common"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UsersRegister(router *gin.RouterGroup) {
	router.POST("/", UsersRegistration)
	router.POST("/login", UsersLogin)
}

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
