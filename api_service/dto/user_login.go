package dto

type UserLoginRequest struct {
	Username string `json:"username" validate:"required,min=6,max=30"`
	Password string `json:"password" validate:"required,min=8,max=15"`
}

type UserLoginResponse struct {
	Token string `json:"token" validate:"required"`
}
