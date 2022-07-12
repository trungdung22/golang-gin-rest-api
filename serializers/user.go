package serializers

import (
	"crud-api/models"

	"github.com/gin-gonic/gin"
)

type UserSerializer struct {
	c *gin.Context
}

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Bio      string `json:"bio"`
	Token    string `json:"token"`
}

func (self *UserSerializer) Response() UserResponse {
	myUserModel := self.c.MustGet("currUser").(models.UserModel)
	user := UserResponse{
		Username: myUserModel.Username,
		Email:    myUserModel.Email,
		Bio:      myUserModel.Bio,
	}
	return user
}
