package exolve

type SendSmsRequest struct{
	Number string `json:"number" validate:"required"`
	Destination string `json:"destination" validate:"required"`
	Text string `json:"text" validate:"required"`
}
