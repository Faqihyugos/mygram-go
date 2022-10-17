package user

type RegisterUserInput struct {
	Age      uint   `json:"age" binding:"required"`
	Email    string `json:"email" binding:"required,email" validate:"unique"`
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required" validate:"unique"`
}
