package handler

import (
	"net/http"

	"github.com/Faqihyugos/mygram-go/auth"
	"github.com/Faqihyugos/mygram-go/helper"
	"github.com/Faqihyugos/mygram-go/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

// RegisterUser godoc
// @Summary     Register User
// @Description Register User
// @Tags        users
// @Accept      json
// @Produce     json
// @Param		data body user.RegisterUserInput true "User Data"
// @Success     201 {object} user.UserFormatter
// @Failure     400 {object} helper.ApiError
// @Failure     422 {object} helper.ApiError
// @Router      /users/register [post]
func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "Failed to register user",
			"message": errors,
		})
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		helper.ApiResponseError(c, http.StatusBadRequest, "Failed to register user", err.Error())
		return
	}

	response := user.FormatUser(newUser)
	c.JSON(http.StatusCreated, response)
}

// LoginUser godoc
// @Summary     Login User
// @Description Login User
// @Tags        Users
// @Accept      json
// @Produce     json
// @Body		request body user.LoginInput true "User Data"
// @Success     200 {object} user.UserLoginFormatter
// @Failure     400 {object} helper.ApiError
// @Failure     422 {object} helper.ApiError
// @Router      /users/login [post]
func (h *userHandler) Login(c *gin.Context) {

	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		//cek validation
		errors := helper.FormatValidationError(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "Login account failed",
			"message": errors,
		})
		return
	}

	loginUser, err := h.userService.Login(input)
	if err != nil {
		helper.ApiResponseError(c, http.StatusBadRequest, "Login account failed", err.Error())
		return
	}

	token, err := h.authService.GenerateToken(loginUser.ID)
	if err != nil {
		helper.ApiResponseError(c, http.StatusBadRequest, "Login account failed", err.Error())
		return
	}

	// response := helper.ApiResponse("Login success", http.StatusOK, "succes", formatter)
	response := user.FormatLogin(token)
	c.JSON(http.StatusOK, response)

}

// UpdateUser godoc
// @Summary     Update User
// @Description Update User
// @Tags        Users
// @Accept      json
// @Produce     json
// @Body		request body user.UpdateUserInput true "User Data"
// @Success     200 {object} user.UserUpdateFormatter
// @Failure     400 {object} helper.ApiError
// @Failure     422 {object} helper.ApiError
// @Router      /users [put]
func (h *userHandler) UpdateUser(c *gin.Context) {

	var inputData user.UpdateUserInput
	err := c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "Failed to update user",
			"message": errors,
		})
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	inputData.User = currentUser
	id := int(currentUser.ID)

	updatedUser, err := h.userService.UpdateUser(id, inputData)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "Failed to update user",
			"message": err.Error(),
		})
		return
	}

	// response := helper.ApiResponse("Success to update user", http.StatusOK, "success", formatter)
	response := user.FormatUpdateUser(updatedUser)
	c.JSON(http.StatusOK, response)
}

// DeleteUser godoc
// @Summary     Delete User
// @Description Delete User
// @Tags        Users
// @Accept      json
// @Produce     json
// @Success     200 {object} nil
// @Failure     400 {object} helper.ApiError
// @Router      /users [delete]
func (h *userHandler) DeleteUser(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	id := int(currentUser.ID)

	_, err := h.userService.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Your account has been succefully deleted",
	})
}
