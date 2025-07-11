package jino_mail

import (
	"crypto/tls"
	"fmt"
	"log"

	"gopkg.in/gomail.v2"
)

type Config struct {
	Host string `json:"host" validate:"required"`
	Port int `json:"port" validate:"required"`
	User string `json:"user" validate:"required"`
	Password string `json:"password" validate:"required"`
	AdminAddress string `json:"admin_address" validate:"omitempty,email"`
}

// jino_mail.json
// {
// 	"host": "smtp.jino.ru",
// 	"port": "465",
// 	"user": "name@mail.ru",
// 	"password": "abc123
//  "admin_address": "admin@mail.ru"
// }

type JinoMail struct {
	config *Config
}

func New(config *Config) *JinoMail{

	jinoMail := JinoMail{config: config}

	if jinoMail.config.AdminAddress != "" {
		go jinoMail.SendMail(jinoMail.config.AdminAddress, " server is running", "test message")
	}

	return &jinoMail
}


func (jmail *JinoMail) SendMail(address string, messageHTML string, subject string) error {
	
	msg := gomail.NewMessage()
	msg.SetHeader("From", jmail.config.User)
	msg.SetHeader("To", jmail.config.AdminAddress)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", messageHTML)

	dialer := gomail.NewDialer(jmail.config.Host, jmail.config.Port, jmail.config.User, jmail.config.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	err := dialer.DialAndSend(msg)

	if (err != nil) {
		log.Println("JINO MAIL ERROR: ", err)
		return  err
	}	
	return  nil
}

func (jmail *JinoMail) SendCode(address string, code string) error {
	msg := "Код подтверждения: " + fmt.Sprintf("<b>%s</b>", code)
	subject := "Запрос на авторизацию"
	return jmail.SendMail(address, msg, subject)
}

