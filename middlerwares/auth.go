package middlerwares

import (
	"crud-api/models"
	"crud-api/services"
	"crud-api/utilities"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Strips 'TOKEN ' prefix from token string
func EnforceAuthenticatedMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("currentUser")
		if exists && user.(models.User).ID != 0 {
			return
		} else {
			err, _ := c.Get("authErr")
			_ = c.AbortWithError(http.StatusUnauthorized, err.(error))
			return
		}
	}
}

func UpdateUserContext(c *gin.Context, user models.User) {
	c.Set("currentUser", user)
	c.Set("currentUserId", user.ID)
}

func UserLoaderMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer := c.Request.Header.Get("authorization")
		if bearer != "" {
			jwtParts := strings.Split(bearer, " ")
			if len(jwtParts) == 2 {
				jwtEncoded := jwtParts[1]
				authPayload, err := utilities.GetGoogleUserPayload(jwtEncoded)

				if err != nil {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
					return
				}

				user, err := services.GetUserByEmail(authPayload.Email)

				if err != nil {
					return
				}
				UpdateUserContext(c, user)
				c.Next()
			}
		}
		c.Next()
	}
}
