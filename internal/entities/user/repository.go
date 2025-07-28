package user

import (
	"mail-phone-auth/pkg/postgres"
)

type Repository struct {
	postgres *postgres.Postgres
}

func NewRepository(postgres *postgres.Postgres) *Repository {
	repository := Repository{
		postgres: postgres,
	}

	return &repository
}

func (r *Repository) Read(ID int) (*User, error) {
	var user User

	result := r.postgres.DB.First(&user, ID)

	if result.Error != nil {
		return nil, result.Error
	} else {
		return &user, nil
	}
}

func (r *Repository) ReadByEmail(email string) (*User, error) {
	var user User

	result := r.postgres.DB.Table("users").
	Where("email = ?", email).
	First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *Repository) ReadAll(limit int, offset int) ([]User, error) {
	var users []User

	result := r.postgres.DB.Table("users").
	Where("deleted_at is NULL").
	Order("id ASC").
	Limit(limit).
	Offset(offset).
	Scan(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (r *Repository) Create(user *User) error {
	result := r.postgres.DB.Create(user)

	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (r *Repository) Update(user *User) error {

	result := r.postgres.DB.Updates(user)

	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func (r *Repository) Delete(ID int) error {
	result := r.postgres.DB.Delete(&User{}, ID)

	if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}
