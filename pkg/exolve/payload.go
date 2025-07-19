package exolve

type SendSmsRequestData struct {
	Number      string `json:"number"`
	Destination string `json:"destination"`
	Text        string `json:"text"`
}
