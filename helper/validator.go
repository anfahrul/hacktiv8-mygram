package helper

import (
	"github.com/anfahrul/hacktiv8-mygram/model"
	"github.com/go-playground/validator/v10"
)

func UserRegisterValidator(requestData model.UserRegisterReq) []error {
	validate = validator.New()

	err := validate.Struct(requestData)
	if err != nil {
		errResponse := []error{}
		for _, err := range err.(validator.ValidationErrors) {
			errResponse = append(errResponse, ErrorRequestMessages(err.Field(), err.ActualTag(), err.Param()))
		}

		return errResponse
	}

	return nil
}

func UserLoginValidator(requestData model.UserLoginReq) []error {
	validate = validator.New()

	err := validate.Struct(requestData)
	if err != nil {
		errResponse := []error{}
		for _, err := range err.(validator.ValidationErrors) {
			errResponse = append(errResponse, ErrorRequestMessages(err.Field(), err.ActualTag(), err.Param()))
		}

		return errResponse
	}

	return nil
}

func PhotoCreateValidator(requestData model.PhotoCreateReq) []error {
	validate = validator.New()

	err := validate.Struct(requestData)
	if err != nil {
		errResponse := []error{}
		for _, err := range err.(validator.ValidationErrors) {
			errResponse = append(errResponse, ErrorRequestMessages(err.Field(), err.ActualTag(), err.Param()))
		}

		return errResponse
	}

	return nil
}

func PhotoUpdateValidator(requestData model.PhotoUpdateReq) []error {
	validate = validator.New()

	err := validate.Struct(requestData)
	if err != nil {
		errResponse := []error{}
		for _, err := range err.(validator.ValidationErrors) {
			errResponse = append(errResponse, ErrorRequestMessages(err.Field(), err.ActualTag(), err.Param()))
		}

		return errResponse
	}

	return nil
}