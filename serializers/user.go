package serializers

import (
	"crud-api/models"
	"time"

	"github.com/gin-gonic/gin"
)

type UserSerializer struct {
	C *gin.Context
}

type UserCreateResponse struct {
	Message string `json:"username"`
}

type UserLoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

func (self *UserSerializer) UserCreatedSuccessResponse() UserCreateResponse {
	user := UserCreateResponse{
		Message: "User Created success",
	}
	return user
}

func (self *UserSerializer) UserLoginSuccessResponse() UserLoginResponse {
	user := self.C.MustGet("currentUser").(models.User)
	token := user.GenerateJwtToken()
	response := UserLoginResponse{
		Message: "login success",
		Token:   token,
	}
	return response
}

func CreateUserPageResponse(users []models.User, page, pageSize, totalCount int) interface{} {
	var resources = make([]interface{}, len(users))
	for index, user := range users {
		result := map[string]interface{}{
			"id":         user.ID,
			"username":   user.Username,
			"email":      user.Email,
			"bio":        user.Bio,
			"created_at": user.CreatedAt.UTC().Format(time.RFC3339Nano),
			"updated_at": user.UpdatedAt.UTC().Format(time.RFC3339Nano),
		}

		resources[index] = result
	}
	return CreatePagedResponse(resources, "users", page, pageSize, totalCount)
}
