package dto

import "app/internal/domain"

type UserResponse struct {
	User struct {
		UserName string `json:"username"`
		Email    string `json:"email"`
		Token    string `json:"token"`
	} `json:"user"`
}

func ToUserResponse(u *domain.User, token string) *UserResponse {
	return &UserResponse{
		User: struct {
			UserName string `json:"username"`
			Email    string `json:"email"`
			Token    string `json:"token"`
		}{
			UserName: u.UserName,
			Email:    u.Email,
			Token:    token,
		},
	}
	// user := new(UserResponse)
	// user.User.Email = u.Email
	// user.User.Token = token
	// user.User.UserName = u.UserName
	// return user
}
