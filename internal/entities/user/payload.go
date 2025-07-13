package user

// UserCreateRequest - запрос на создание пользователя
type UserCreateRequest struct {
	Name  string `json:"name" validate:"required,min=2,max=50" example:"Иван Иванов"`
	Phone string `json:"phone" validate:"required,e164" example:"+79124852012"`
	Email string `json:"email" validate:"required,email" example:"user@example.com"`
}

// UserUpdateRequest - запрос на обновление пользователя
type UserUpdateRequest struct {
	Name  string `json:"name" validate:"required,min=2,max=50" example:"Иван Иванов"`
	Phone string `json:"phone" validate:"required,e164" example:"+79124852012"`
}

// UserAllResponse - ответ на запрос списка пользователей
type UserAllResponse struct {
	Users []User `json:"users"`
	Count int    `json:"count" example:"10"`
}

// ErrorResponse ответ с ошибкой
type ErrorResponse struct {
	Error string `json:"error" example:"Произошла ошибка"`
}
