package model

type UserRegisterRequest struct {
	Username        string `json:"username" binding:"required,min=3,max=20"`
	Password        string `json:"password" binding:"required,min=3,max=20"`
	PasswordConfirm string `json:"password_confirm" binding:"required,min=3,max=20"`
}

type UserRegisterResponse struct {
	Token string `json:"token"`
}

type UserLoginRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=3,max=20"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

type UserProfileResponse struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	Nickname    string `json:"nickname"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
}
