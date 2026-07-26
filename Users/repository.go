package users

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user User) (*User, error)
	GetByEmail(email string) (*User, error)
	GetById(id uint) (*User, error)
	CreateWithTx(tx *gorm.DB, user *User) error
}

type gormRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &gormRepository{
		db: db,
	}
}

func (r *gormRepository) CreateUser(user User) (*User, error) {

	result := r.db.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
func (r *gormRepository) GetByEmail(email string) (*User, error) {

	var user User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err // Yahan explicit 'nil' return karein agar error hai
	}
	return &user, nil
}
func (r *gormRepository) GetById(id uint) (*User, error) {
	var user User
	result := r.db.First(&user, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
func (r *gormRepository) CreateWithTx(tx *gorm.DB, user *User) error {
	return tx.Create(user).Error
}
