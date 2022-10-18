package user

import "time"

type UserFormatter struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Age      uint   `json:"age"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

type UserLoginFormatter struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

type UserUpdateFormatter struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Age       uint      `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FormatLogin(user User, token string) UserLoginFormatter {
	formatterLogin := UserLoginFormatter{
		ID:       user.ID,
		Username: user.Username,
		Token:    token,
	}
	return formatterLogin
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

func FormatUpdaeUser(user User) UserUpdateFormatter {
	formatter := UserUpdateFormatter{
		ID:        user.ID,
		Username:  user.Username,
		Age:       user.Age,
		Email:     user.Email,
		UpdatedAt: user.UpdatedAt,
	}
	return formatter
}
