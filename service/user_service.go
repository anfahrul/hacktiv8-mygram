package service

import (
	"errors"
	"fmt"

	"github.com/anfahrul/hacktiv8-mygram/helper"
	"github.com/anfahrul/hacktiv8-mygram/model"
	"github.com/anfahrul/hacktiv8-mygram/repository"
	"github.com/go-playground/validator/v10"
)

type UserService interface {
	Register(userReqData model.UserRegisterReq) (*model.UserRegisterRes, error)
	Login(userReqData model.UserLoginReq) (*string, error)
}

type UserServiceIml struct {
	userRepository repository.UserRepository
	validate       *validator.Validate
}

func NewUserService(userRepo repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceIml{
		userRepository: userRepo,
		validate:       validate,
	}
}

func (s *UserServiceIml) Register(userReqData model.UserRegisterReq) (*model.UserRegisterRes, error) {
	userID := helper.GenerateID()
	hashedPassword, err := helper.Hash(userReqData.Password)
	if err != nil {
		return nil, err
	}

	newUser := model.User{
		UserID:   userID,
		Username: userReqData.Username,
		Email:    userReqData.Email,
		Password: hashedPassword,
		Age:      userReqData.Age,
	}

	err = s.userRepository.Create(newUser)
	if err != nil {
		return nil, err
	}

	return &model.UserRegisterRes{
		UserID:   newUser.UserID,
		Username: newUser.Username,
		Email:    newUser.Email,
		Password: newUser.Password,
		Age:      newUser.Age,
	}, nil
}

func (s *UserServiceIml) Login(userReqData model.UserLoginReq) (*string, error) {
	userResponse, err := s.userRepository.FindByUsername(userReqData.Username)
	if err != nil {
		return nil, err
	}

	isMatch := helper.PasswordIsMatch(userReqData.Password, userResponse.Password)
	if isMatch == false {
		return nil, errors.New(fmt.Sprintf("Invalid username or password"))
	}

	token, err := helper.GenerateToken(userResponse)
	if err != nil {
		return nil, err
	}

	return &token, nil
}
