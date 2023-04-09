package repository

import (
	"github.com/anfahrul/hacktiv8-mygram/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(userReqData model.User) error
	FindByUsername(username string) (model.User, error)
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (r *UserRepositoryImpl) Create(userReqData model.User) error {
	err := r.DB.Create(&userReqData).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepositoryImpl) FindByUsername(username string) (model.User, error) {
	var user model.User
	err := r.DB.First(&user, "username = ?", username).Error

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
