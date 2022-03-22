package repository

import (
	"gca/model"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() (model.User, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() (model.User, error) {
	var user model.User
	err := r.db.Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
