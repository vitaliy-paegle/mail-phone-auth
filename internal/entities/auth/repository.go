package auth

import "mail-phone-auth/pkg/postgres"

type Repository struct {
	postgres *postgres.Postgres
}

func NewRepository(postgres *postgres.Postgres) *Repository {
	repository := Repository{postgres: postgres}

	return  &repository
}