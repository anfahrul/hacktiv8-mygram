package service

import (
	"errors"
	"time"

	"github.com/anfahrul/hacktiv8-mygram/helper"
	"github.com/anfahrul/hacktiv8-mygram/model"
	"github.com/anfahrul/hacktiv8-mygram/repository"
)

type SocialService interface {
	Create(socialReqData model.SocialCreateReq, userID string) (*model.SocialResponse, error)
	GetAll() ([]model.SocialResponse, error)
	GetOne(socialID string) (model.SocialResponse, error)
	UpdateSocialMedia(socialReqData model.SocialUpdateReq, userID string, socialID string) (*model.SocialResponse, error)
	Delete(socialID string, userID string) (model.SocialResponse, error)
}

type SocialServiceIml struct {
	socialRepository repository.SocialRepository
}

func NewSocialService(socialRepo repository.SocialRepository) SocialService {
	return &SocialServiceIml{
		socialRepository: socialRepo,
	}
}

func (s *SocialServiceIml) Create(socialReqData model.SocialCreateReq, userID string) (*model.SocialResponse, error) {
	socialID := helper.GenerateID()
	newSocial := model.SocialMedia{
		ID:             socialID,
		Name:           socialReqData.Name,
		SocialMediaURL: socialReqData.SocialMediaURL,
		UserID:         userID,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	err := s.socialRepository.Create(newSocial)
	if err != nil {
		return nil, err
	}

	return &model.SocialResponse{
		ID:             newSocial.ID,
		Name:           newSocial.Name,
		SocialMediaURL: newSocial.SocialMediaURL,
		UserID:         newSocial.UserID,
		CreatedAt:      newSocial.CreatedAt,
		UpdatedAt:      newSocial.UpdatedAt,
	}, nil
}

func (s *SocialServiceIml) GetAll() ([]model.SocialResponse, error) {
	photosResult, err := s.socialRepository.FindAll()
	if err != nil {
		return []model.SocialResponse{}, err
	}

	socialsResponse := []model.SocialResponse{}
	for _, socialRes := range photosResult {
		socialsResponse = append(socialsResponse, model.SocialResponse(socialRes))
	}

	return socialsResponse, nil
}

func (s *SocialServiceIml) GetOne(socialID string) (model.SocialResponse, error) {
	socialsResult, err := s.socialRepository.FindByID(socialID)
	if err != nil {
		return model.SocialResponse{}, err
	}

	return model.SocialResponse(socialsResult), nil
}

func (s *SocialServiceIml) UpdateSocialMedia(socialReqData model.SocialUpdateReq, userID string, socialID string) (*model.SocialResponse, error) {
	findSocialMediaResponse, err := s.socialRepository.FindByID(socialID)
	if err != nil {
		return nil, err
	}

	if userID != findSocialMediaResponse.UserID {
		return nil, errors.New("Unauthorized")
	}

	updatedSocialReq := model.SocialMedia{
		ID:             socialID,
		Name:           socialReqData.Name,
		SocialMediaURL: socialReqData.SocialMediaURL,
		UserID:         userID,
		UpdatedAt:      time.Now(),
	}

	err = s.socialRepository.Update(updatedSocialReq)
	if err != nil {
		return nil, err
	}

	return &model.SocialResponse{
		ID: socialID,
	}, nil
}

func (s *SocialServiceIml) Delete(socialID string, userID string) (model.SocialResponse, error) {
	findSocialResponse, err := s.socialRepository.FindByID(socialID)
	if err != nil {
		return model.SocialResponse{}, err
	}

	if userID != findSocialResponse.UserID {
		return model.SocialResponse{}, errors.New("Unauthorized")
	}

	err = s.socialRepository.Delete(model.SocialMedia{ID: socialID})
	if err != nil {
		return model.SocialResponse{}, err
	}

	return model.SocialResponse{
		ID: socialID,
	}, nil
}
