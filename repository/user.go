package repository

import (
	"crud/entity"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func NewUser(dbCrud *gorm.DB) User {
	return User{
		db: dbCrud,
	}

}

type UserInterfaceRepo interface {
	CreateUser(user *entity.User) (*entity.User, error)
	GetUserById(id uint) (entity.User, error)
}

func (repo User) CreateUser(user *entity.User) (*entity.User, error) {
	err := repo.db.Model(&entity.User{}).Create(user).Error
	return user, err
}

func (repo User) GetUserById(id uint) (entity.User, error) {
	var user entity.User
	repo.db.First(&user, "id = ? ", id)
	return user, nil
}
