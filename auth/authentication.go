package auth

import (
	"net/http"
	"strings"

	"github.com/Faqihyugos/mygram-go/user"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authentication(userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {

		authService := NewService()
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Token Not Found",
				"message": "Unauthenticated",
			})
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Token Not Valid",
				"message": "Unauthenticated",
			})
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Token Not Valid",
				"message": "Unauthenticated",
			})
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "User Not Found",
				"message": "Unauthenticated",
			})
			return
		}

		c.Set("currentUser", user)
	}
}
