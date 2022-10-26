package auth

import (
	"net/http"
	"strconv"

	"github.com/Faqihyugos/mygram-go/comment"
	"github.com/Faqihyugos/mygram-go/config"
	"github.com/Faqihyugos/mygram-go/photo"
	"github.com/Faqihyugos/mygram-go/sosmed"
	"github.com/Faqihyugos/mygram-go/user"
	"github.com/gin-gonic/gin"
)

// create func Authoization middleware for 3 path photo, comment, and social media
func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := config.StartDB()
		//get param dinamic from path
		//example: /photos/1
		//get 1 from path
		id := c.Param("id")
		//convert string to int
		idInt, _ := strconv.Atoi(id)
		//get user id from token
		userID := c.MustGet("currentUser").(user.User).ID
		//get path from request
		path := c.Request.URL.Path
		//check path
		switch path {
		case "/photos":
			//get photo from db by id
			photo := photo.Photo{}
			err := db.Select("user_id").First(&photo, int(idInt)).Error
			//check if error
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"error":   "Not Found",
					"message": "Data doesn't exist",
				})
				c.Abort()
				return
			}
			//check if user id not equal with user id from token
			if photo.UserID != userID {
				c.JSON(http.StatusForbidden, gin.H{
					"error":   "Forbidden",
					"message": "You are not allowed to access this resource",
				})
				c.Abort()
				return
			}
		case "/comments":
			//get comment by id
			comment := comment.Comment{}
			err := db.Select("user_id").First(&comment, int(idInt)).Error
			//check if error
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"error":   "Not Found",
					"message": "Data doesn't exist",
				})
				c.Abort()
				return
			}
			//check if user id not equal with user id from token
			if comment.UserID != userID {
				c.JSON(http.StatusForbidden, gin.H{
					"error":   "Forbidden",
					"message": "You are not allowed to access this resource",
				})
				c.Abort()
				return
			}
		case "/socialmedias":
			//get sosmed by id
			social := sosmed.Sosmed{}
			err := db.Select("user_id").First(&social, int(idInt)).Error
			//check if error
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"error":   "Not Found",
					"message": "Data doesn't exist",
				})
				c.Abort()
				return
			}
			//check if user id not equal with user id from token
			if social.UserID != userID {
				c.JSON(http.StatusForbidden, gin.H{
					"error":   "Forbidden",
					"message": "You are not allowed to access this resource",
				})
				c.Abort()
				return
			}
		}
	}
}
