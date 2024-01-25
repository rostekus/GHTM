package handler

type UserRegisterRequest struct {
	Username       string `json:"username" validate:"required,min=3"`
	Email          string `json:"email" validate:"required,email"`
	Password       string `json:"password1" validate:"required,min=6"`
	PasswordSecond string `json:"password2" validate:"required,min=6"`
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}
