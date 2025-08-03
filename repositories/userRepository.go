package repositories

import (
	models "blog/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUser(user *models.User) models.User
	Register(user *models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) GetUser(user *models.User) models.User {
	var userModel models.User
	u.db.Where("username = ?", user.Username).First(&userModel)
	return userModel
}

func (u *userRepository) Register(user *models.User) error {
	return u.db.Create(user).Error
}
