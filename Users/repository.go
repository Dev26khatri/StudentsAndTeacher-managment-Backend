package users

import "gorm.io/gorm"

type Repository interface {
}

type userRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &userRepository{
		db: db,
	}
}
