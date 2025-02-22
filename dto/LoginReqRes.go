package dto

type LoginRequest struct {
	UserName string `json:"userName"  validate:"required,min=5"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginResponse struct {
	Status string `json:"status"`
	Token  string `json:"token"`
}
