package main

import (
	"github.com/Faqihyugos/mygram-go/auth"
	"github.com/Faqihyugos/mygram-go/comment"
	"github.com/Faqihyugos/mygram-go/config"
	"github.com/Faqihyugos/mygram-go/handler"
	"github.com/Faqihyugos/mygram-go/photo"
	"github.com/Faqihyugos/mygram-go/user"

	"github.com/gin-gonic/gin"
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

	commentRepository := comment.NewRepository(db)
	commentService := comment.NewService(commentRepository)
	commentHandler := handler.NewCommentHandler(commentService, authService)

	router := gin.Default()

	// user
	userRouter := router.Group("/users")
	userRouter.POST("/register", userHandler.RegisterUser)
	userRouter.POST("/login", userHandler.Login)
	userRouter.PUT("/:id", auth.Authentication(userService), userHandler.UpdateUser)
	userRouter.DELETE("/:id", auth.Authentication(userService), userHandler.DeleteUser)

	// photo
	photoRouter := router.Group("/photos")
	photoRouter.Use(auth.Authentication(userService))
	photoRouter.POST("/", photoHandler.CreatePhoto)
	photoRouter.GET("/", photoHandler.GetAllPhoto)
	photoRouter.PUT("/:photoId", auth.PhotoAuthorization(), photoHandler.UpdatePhoto)
	photoRouter.DELETE("/:photoId", auth.PhotoAuthorization(), photoHandler.DeletePhoto)

	commentRouter := router.Group("/comments")
	commentRouter.Use(auth.Authentication(userService))
	commentRouter.POST("/", commentHandler.CreateComment)

	router.Run()

}
