package auth

import (
	"mail-phone-auth/pkg/postgres"
)

type Repository struct {
	postgres *postgres.Postgres
}

func NewRepository(postgres *postgres.Postgres) *Repository {
	repository := Repository{postgres: postgres}

	return  &repository
}

func (r *Repository) CreateEmailAuth(auth *Auth) error {
	result := r.postgres.DB.Create(auth)

	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}

}