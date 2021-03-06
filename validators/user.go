package validators

import (
	"crud-api/utilities"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserSignUpRequest struct {
	Username string `json:"username" binding:"required,alphanum,min=4,max=255"`
	Email    string `json:"email" validate:"required,email"`
	Bio      string `json:"bio" validate:"max=1024"`
	Image    string `json:"image" validate:"omitempty,url"`
}

type LoginRequest struct {
	// Username string `form:"username" json:"username" xml:"username" binding:"exists,username"`
	Username string `form:"username" json:"username" xml:"username" validate:"required"`
	Password string `form:"password" json:"password" validate:"required,min=8,max=255"`
}

func SignupValidator(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user UserSignUpRequest
		if err := c.ShouldBindJSON(&user); err == nil {
			validate := validator.New()
			if err := validate.Struct(&user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}

func (self *UserSignUpRequest) Bind(c *gin.Context) error {
	err := utilities.Bind(c, self)
	if err != nil {
		return err
	}
	return nil
}

func (self *LoginRequest) Bind(c *gin.Context) error {
	err := utilities.Bind(c, self)
	if err != nil {
		return err
	}
	return nil
}
