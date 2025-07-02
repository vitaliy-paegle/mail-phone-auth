package user

type UserCreateRequest struct {
	Name string `json:"name" validate:"required"`
	Phone string `json:"phone" validate:"e164"`
	Email string `json:"email" validate:"email"`
}