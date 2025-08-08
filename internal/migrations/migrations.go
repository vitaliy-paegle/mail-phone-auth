package main

import (
	"log"
	"mail-phone-auth/internal/app/files"
	"mail-phone-auth/internal/entity/auth"
	"mail-phone-auth/internal/entity/role"
	"mail-phone-auth/internal/entity/user"
	"mail-phone-auth/pkg/postgres"
)

func main() {
	const postgresCongigFilePath = "./config/postgres.json"

	postgresConfig, err := files.InitConfig[postgres.Config](postgresCongigFilePath)
	if err != nil {
		log.Fatal(err)
	}

	postgres, err := postgres.NewPostgres(postgresConfig)
	if err != nil {
		log.Fatal(err)
	}

	postgres.AutoMigrate(
		user.User{},
		auth.Auth{},
		role.Role{},
	)

}
