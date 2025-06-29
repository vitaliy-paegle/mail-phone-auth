package exolve

import (
	"encoding/json"
	"fmt"
)


type Config struct {
	Number string
	Key string
}

type Exolve struct {
	config Config
	senSmsMethod string
	sendSmsPath string
}

func New(config Config) *Exolve {

	const SEND_SMS_METHOD = "POST"
	const SEND_SMS_PATH = "https://api.exolve.ru/messaging/v1/SendSMS"

	return &Exolve {
		config: config,
		senSmsMethod: SEND_SMS_METHOD,
		sendSmsPath: SEND_SMS_PATH,
	}
}


func (exolve *Exolve) SendSms(destination string, text string) {

	request := SendSmsRequest{
		Number: exolve.config.Number,
		Destination: destination,
		Text: text,
	}

	jsonRequest, err := json.Marshal(request)

	if err != nil {
		fmt.Print(err)
		return 
	}

	fmt.Print(string(jsonRequest))

}