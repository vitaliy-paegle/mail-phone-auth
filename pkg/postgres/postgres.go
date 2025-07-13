package postgres

import (
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	Host     string `json:"host" validate:"required"`
	User     string `json:"user" validate:"required"`
	Password string `json:"password" validate:"required"`
	DBname   string `json:"dbname" validate:"required"`
	Port     string `json:"port" validate:"required"`
	SSLmode  string `json:"sslmode" validate:"required"`
}

// postgres.json
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

	postgres, err := gorm.Open(sqlite.Open("mail_phone_auth.db"), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to database:", err)
		return nil, err
	}

	log.Println("Connected to SQLite database")
	return &Postgres{postgres}, nil
}
