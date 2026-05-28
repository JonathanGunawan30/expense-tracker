package request

type UserRegisterRequest struct {
	Name     string `json:"name" validate:"required,max=200"`
	Email    string `json:"email" validate:"required,max=255,email"`
	Password string `json:"password" validate:"required,min=6,max=72"`
}
