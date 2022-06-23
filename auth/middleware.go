package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"survorest/helper"
	"survorest/user"
)

func AuthMiddleware(authService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			response := helper.ApiResponse("Header not provided", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		extractedToken := strings.Split(clientToken, "Bearer ")
		if len(extractedToken) == 2 {
			clientToken = strings.TrimSpace(extractedToken[1])
		} else {
			response := helper.ApiResponse("Invalid Authorization Format", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		jwtWrapper := JwtWrapper{
			SecretKey: "survosecret",
			Issuer: "AuthService",
		}
		claims,err := jwtWrapper.ValidateToken(clientToken)
		if err != nil {
			response := helper.ApiResponse("Unauthorize", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		userID := claims.UserID
		user , _ := authService.GetUserByID(userID)
		c.Set("claims",user)
	}
}

func AuthAdminMiddleware() gin.HandlerFunc{
	return func(c *gin.Context) {
		session := sessions.Default(c)

		userIDSession := session.Get("userID")

		if userIDSession == nil {
			c.Redirect(http.StatusFound,"/dashboard")
			return
		}
	}
}
