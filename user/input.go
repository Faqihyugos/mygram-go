package user

type RegisterUserInput struct {
	Age      int    `json:"age" binding:"required"`
	Email    string `json:"email" binding:"required,email,unique"`
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required,unique"`
}
