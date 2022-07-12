package endpoints

import (
	"crud-api/models"
	"crud-api/serializers"
	"crud-api/services"
	"crud-api/validators"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UsersRegisterRouter(router *gin.RouterGroup) {
	router.POST("/", UsersRegistration)
	router.POST("/login", UsersLogin)
}

func UsersRegistration(c *gin.Context) {
	userDto := validators.UserSignUpRequest{}

	if err := userDto.Bind(c); err != nil {
		c.JSON(http.StatusBadRequest, validators.ValidatorError(err))
		return
	}
	if err := services.CreateOneUser(&userDto); err != nil {
		c.JSON(http.StatusUnprocessableEntity, serializers.ResponseError("database", err))
	}
	userSerializer := serializers.UserSerializer{c}
	c.JSON(http.StatusOK, userSerializer.Response())
}

func UsersLogin(c *gin.Context) {

	var json validators.LoginRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, validators.ValidatorError(err))
		return
	}

	user, err := services.FindOneUser(&models.User{Username: json.Username})

	if err != nil {
		c.JSON(http.StatusForbidden, dtos.CreateDetailedErrorDto("login_error", err))
		return
	}

	if user.IsValidPassword(json.Password) != nil {
		c.JSON(http.StatusForbidden, dtos.CreateDetailedErrorDto("login", errors.New("invalid credentials")))
		return
	}

	c.JSON(http.StatusOK, dtos.CreateLoginSuccessful(&user))

}
