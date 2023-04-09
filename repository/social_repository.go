package repository

import (
	"errors"

	"github.com/anfahrul/hacktiv8-mygram/model"
	"gorm.io/gorm"
)

type SocialRepository interface {
	Create(photoReqData model.SocialMedia) error
	FindAll() ([]model.SocialMedia, error)
	FindByID(socialID string) (model.SocialMedia, error)
	FindByUserID(userID string) ([]model.SocialMedia, error)
	Update(socialReqData model.SocialMedia) error
	Delete(photoReqData model.SocialMedia) error
}

type SocialRepositoryImpl struct {
	DB *gorm.DB
}

func NewSocialRepository(db *gorm.DB) SocialRepository {
	return &SocialRepositoryImpl{
		DB: db,
	}
}

func (r *SocialRepositoryImpl) Create(socialReqData model.SocialMedia) error {
	err := r.DB.Create(&socialReqData).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *SocialRepositoryImpl) FindAll() ([]model.SocialMedia, error) {
	socials := []model.SocialMedia{}

	err := r.DB.Find(&socials).Error
	if err != nil {
		return []model.SocialMedia{}, err
	}

	return socials, nil
}

func (r *SocialRepositoryImpl) FindByID(socialID string) (model.SocialMedia, error) {
	social := model.SocialMedia{}

	err := r.DB.Debug().Where("id = ?", socialID).Take(&social).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.SocialMedia{}, err
		}

		return model.SocialMedia{}, err
	}

	return social, nil
}

func (r *SocialRepositoryImpl) FindByUserID(userID string) ([]model.SocialMedia, error) {
	socials := []model.SocialMedia{}

	err := r.DB.Debug().Where("user_id = ?", userID).Find(&socials).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []model.SocialMedia{}, err
		}

		return []model.SocialMedia{}, err
	}

	return socials, nil
}

func (r *SocialRepositoryImpl) Update(socialReqData model.SocialMedia) error {
	err := r.DB.Save(&model.SocialMedia{
		ID:             socialReqData.ID,
		Name:           socialReqData.Name,
		SocialMediaURL: socialReqData.SocialMediaURL,
		UserID:         socialReqData.UserID,
		UpdatedAt:      socialReqData.UpdatedAt,
	}).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *SocialRepositoryImpl) Delete(socialReqData model.SocialMedia) error {
	err := r.DB.Delete(&socialReqData).Error
	if err != nil {
		return err
	}

	return nil
}
