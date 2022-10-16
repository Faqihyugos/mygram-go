package user

type UserFormatter struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:       user.ID,
		Username: user.Username,
		Age:      user.Age,
		Email:    user.Email,
		Token:    token,
	}
	return formatter
}
