package main

import (
	"net/http"
	"strings"

	"github.com/Faqihyugos/mygram-go/auth"
	"github.com/Faqihyugos/mygram-go/config"
	"github.com/Faqihyugos/mygram-go/handler"
	"github.com/Faqihyugos/mygram-go/helper"
	"github.com/Faqihyugos/mygram-go/photo"
	"github.com/Faqihyugos/mygram-go/user"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func main() {
	db := config.StartDB()
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	userHandler := handler.NewUserHandler(userService, authService)

	photoRepository := photo.NewRepository(db)
	photoService := photo.NewService(photoRepository)
	photoHandler := handler.NewPhotoHandler(photoService, authService)

	router := gin.Default()

	// user
	userRouter := router.Group("/users")
	userRouter.POST("/register", userHandler.RegisterUser)
	userRouter.POST("/login", userHandler.Login)
	userRouter.PUT("/:id", authMiddleware(authService, userService), userHandler.UpdateUser)
	userRouter.DELETE("/:id", authMiddleware(authService, userService), userHandler.DeleteUser)

	// photo
	photoRouter := router.Group("/photos")
	photoRouter.Use(authMiddleware(authService, userService))
	photoRouter.POST("/", photoHandler.CreatePhoto)
	photoRouter.GET("/", photoHandler.GetAllPhoto)
	photoRouter.PUT("/:photoId", auth.PhotoAuthorization(), photoHandler.UpdatePhoto)

	router.Run()

}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}
