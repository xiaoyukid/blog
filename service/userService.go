package service

import (
	models "blog/model"
	"blog/repositories"
)

type UserService interface {
	GetUser(user *models.User) models.User
	Register(user *models.User) error
}

type userService struct {
	userRepository repositories.UserRepository
}

func (u userService) GetUser(user *models.User) models.User {
	return u.userRepository.GetUser(user)
}

func (u userService) Register(user *models.User) error {
	return u.userRepository.Register(user)
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}
