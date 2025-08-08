package auth

import (
	"mail-phone-auth/pkg/postgres"
	"time"
)

type Repository struct {
	postgres *postgres.Postgres
}

func NewRepository(postgres *postgres.Postgres) *Repository {
	repository := Repository{postgres: postgres}

	return &repository
}

func (r *Repository) CreateEmailAuth(auth *Auth) error {

	result := r.postgres.DB.Create(auth)

	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}

}

func (r *Repository) ReadLastAuthByEmail(email string) *Auth {

	var authList []Auth

	r.postgres.Table("auths").
		Where("email = ?", email).
		Where("deleted_at is NULL").
		Order("created_at DESC").
		Scan(&authList)

	if len(authList) > 0 {
		return &authList[0]
	} else {
		return nil
	}

}

func (r *Repository) Delete(ID int) error {

	result := r.postgres.DB.Table("auths").
		Where("id = ?", ID).
		Update("deleted_at", time.Now())

	if result.Error != nil {
		return result.Error
	}

	return nil

}
