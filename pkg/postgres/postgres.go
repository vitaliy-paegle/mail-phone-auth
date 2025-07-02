package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host string `json:"host"`
	User string `json:"user"`
	Password string `json:"password"`
	DBname string `json:"dbname"`
	Port string `json:"port"`
	SSLmode string `json:"sslmode"`
}

// postgres.json:
// {
// 	"host": "localhost",
// 	"user": "main",
// 	"password": "123456",
// 	"dbname": "main",
// 	"port": "5100",
// 	"sslmode": "disable"
// }

type Postgres struct {
	*gorm.DB
}

func NewPostgres(config *Config) (*Postgres, error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", config.Host, config.User, config.Password, config.DBname, config.Port, config.SSLmode)
	postgres, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return  nil, err
	} else {
		return &Postgres{postgres}, nil
	}	

}