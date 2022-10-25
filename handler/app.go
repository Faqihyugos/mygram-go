package handler

import (
	"github.com/Faqihyugos/mygram-go/auth"
	"github.com/Faqihyugos/mygram-go/comment"
	"github.com/Faqihyugos/mygram-go/config"
	"github.com/Faqihyugos/mygram-go/photo"
	"github.com/Faqihyugos/mygram-go/sosmed"
	"github.com/Faqihyugos/mygram-go/user"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartApp() {
	db := config.StartDB()
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	userHandler := NewUserHandler(userService, authService)

	photoRepository := photo.NewRepository(db)
	photoService := photo.NewService(photoRepository)
	photoHandler := NewPhotoHandler(photoService, authService)

	commentRepository := comment.NewRepository(db)
	commentService := comment.NewService(commentRepository)
	commentHandler := NewCommentHandler(commentService, authService)

	sosmedRepository := sosmed.NewRepository(db)
	sosmedService := sosmed.NewService(sosmedRepository)
	sosmedHandler := NewSosmedHandler(sosmedService, authService)

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// user
	userRouter := router.Group("/users")
	userRouter.POST("/register", userHandler.RegisterUser)
	userRouter.POST("/login", userHandler.Login)
	userRouter.PUT("", auth.Authentication(userService), userHandler.UpdateUser)
	userRouter.DELETE("", auth.Authentication(userService), userHandler.DeleteUser)

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
	commentRouter.GET("/", commentHandler.GetAllComment)
	commentRouter.PUT("/:commentId", auth.CommentAuthorization(), commentHandler.UpdateComment)
	commentRouter.DELETE("/:commentId", auth.CommentAuthorization(), commentHandler.DeleteComment)

	socialMediaRouter := router.Group("/socialmedias")
	socialMediaRouter.Use(auth.Authentication(userService))
	socialMediaRouter.POST("/", sosmedHandler.CreateSosmed)
	socialMediaRouter.GET("/", sosmedHandler.GetAllSosmed)
	socialMediaRouter.PUT("/:socialMediaId", auth.SocialMediaAuthorization(), sosmedHandler.UpdateSosmed)
	socialMediaRouter.DELETE("/:socialMediaId", auth.SocialMediaAuthorization(), sosmedHandler.DeleteSosmed)

	router.Run()
}
