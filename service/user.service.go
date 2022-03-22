package service

import (
	"gca/model"
	"gca/repository"
)

type Service interface {
	GetUser() (model.User, error)
}
type service struct {
	repository repository.Repository
}

func NewServiceUser(repository repository.Repository) *service {
	return &service{repository}
}
func (s *service) GetUser() (model.User, error) {
	user, err := s.repository.FindAll()
	if err != nil {
		return user, err
	}
	return user, nil
}
