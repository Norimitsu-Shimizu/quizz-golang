package parameter

type PostUserRequest struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PostLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
