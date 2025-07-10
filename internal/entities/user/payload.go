package user

type UserCreateRequest struct {
	Name string `json:"name"`
	Phone string `json:"phone" validate:"e164"`
	Email string `json:"email" validate:"required,email"`
}

type UserUpdateRequest struct {
	Name string `json:"name"`
	Phone string `json:"phone" validate:"e164"`
}


type UserAllResponse struct {
	Users []User `json:"users"`
	Count int `json:"count"`
}