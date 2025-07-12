package exolve

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)


type Config struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
	Key string `json:"key" validate:"required"`
	AdminPhoneNumber string `json:"admin_phone_number" validate:"omitempty"`
}
	// exolve.json
	// "phone_number": "79581235656",
	// "key": "efgRi65jdfefffERR",
	// "admin_phone_number": "79584567890"

type Exolve struct {
	config *Config
	path string 
}

func New(config *Config, sendTestMessage bool) *Exolve {

	exolve := Exolve{config: config}
	exolve.path = "https://api.exolve.ru/messaging/v1/SendSMS"

	if exolve.config.AdminPhoneNumber != "" && sendTestMessage {
		exolve.SendSms(exolve.config.AdminPhoneNumber, "Test Message")
	}

	return &exolve
}


func (exolve *Exolve) SendSms(to string, text string) error {
	
	client := &http.Client{}

	requestData := SendSmsRequestData{
		Number: exolve.config.PhoneNumber,
		Destination: to,
		Text: text,
	}

	data, err := json.Marshal(requestData)

	if err != nil {
		log.Println("EXOLVE ERR: ", err)
		return err
	}

	buffer := bytes.NewBuffer(data)

	req, err := http.NewRequest("POST", exolve.path, buffer)

	if err != nil {
		log.Println("EXOLVE ERR: ", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
  req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", exolve.config.Key))

	resp, err := client.Do(req)	

	if err != nil {
		log.Println("EXOLVE ERR: ", err)
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println("EXOLVE ERR: ", err)
		return err
	}

	log.Println(string(body))
	return  nil

}


func (exolve *Exolve) SendCode(phone_number string, code string) error {
	to := strings.ReplaceAll(phone_number, "+", "")
	
	template := 
	`
		Запрос на авторизацию.
		Код подтверждения: {code}
	`

	message := strings.ReplaceAll(template, "{code}", code)

	err := exolve.SendSms(to, message)

	if err != nil {
		log.Println("EXOLVE ERR: ", err)
		return err
	}

	return nil
}