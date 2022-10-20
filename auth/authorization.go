package auth

import (
	"net/http"
	"strconv"

	"github.com/Faqihyugos/mygram-go/config"
	"github.com/Faqihyugos/mygram-go/helper"
	"github.com/Faqihyugos/mygram-go/photo"
	"github.com/Faqihyugos/mygram-go/user"
	"github.com/gin-gonic/gin"
)

func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := config.StartDB()
		photoId, err := strconv.Atoi(c.Param("photoId"))
		if err != nil {
			response := helper.ApiResponse("Invalid parameter", http.StatusBadRequest, "Bad Request", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		currentUser := c.MustGet("currentUser").(user.User)
		userID := int(currentUser.ID)
		photo := photo.Photo{}

		err = db.Select("user_id").First(&photo, int(photoId)).Error
		if err != nil {
			response := helper.ApiResponse("Data doesn't exist", http.StatusNotFound, "Data Not Found", nil)
			c.JSON(http.StatusBadRequest, response)
			return
		}

		if photo.UserID != userID {
			response := helper.ApiResponse("You are not allowed to access this data", http.StatusUnauthorized, "Unauthorized", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Next()
	}
}
