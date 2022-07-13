package endpoints

import (
	"crud-api/middlerwares"
	"crud-api/serializers"
	"crud-api/services"
	"crud-api/utilities"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func UsersRegisterRouter(router *gin.RouterGroup) {
	router.GET("/login", UsersLogin)
	router.Use(middlerwares.UserLoaderMiddleware())
	router.Use(middlerwares.EnforceAuthenticatedMiddleware())
	{
		router.GET("/", UserList)
	}
}

func UsersLogin(c *gin.Context) {

	bearer := c.Request.Header.Get("authorization")
	jwtParts := strings.Split(bearer, " ")
	jwtEncoded := jwtParts[1]
	authPayload, err := utilities.GetGoogleUserPayload(jwtEncoded)
	if err != nil {
		c.JSON(http.StatusForbidden, serializers.ResponseError("login", errors.New("token invalid")))
		return
	}
	user, err := services.GetOrCreateUserByEmail(authPayload)
	if err != nil {
		c.JSON(http.StatusForbidden, serializers.ResponseError("login", errors.New("user not found")))
		return
	}

	middlerwares.UpdateUserContext(c, user)
	userSerializers := serializers.UserSerializer{C: c}
	c.JSON(http.StatusOK, userSerializers.UserLoginSuccessResponse())

}

func UserList(c *gin.Context) {
	pageSizeStr := c.Query("page_size")
	pageStr := c.Query("page")
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		pageSize = 5
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}
	users, totalCount, err := services.FindUserPage(pageSize, page)

	if err != nil {
		c.JSON(http.StatusNotFound, serializers.ResponseError("users", errors.New("Invalid param")))
	}
	c.JSON(http.StatusOK, serializers.CreateUserPageResponse(users, page, pageSize, totalCount))
}
