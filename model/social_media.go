package model

import (
	"time"
)

type SocialMedia struct {
	ID             string `gorm:"primaryKey;type:varchar(255)"`
	Name           string `gorm:"not null;type:varchar(50)"`
	SocialMediaURL string `gorm:"not null;type:varchar(255)"`
	UserID         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type SocialCreateReq struct {
	Name           string `json:"name" validate:"required"`
	SocialMediaURL string `json:"social_media_url" validate:"required"`
}

type SocialUpdateReq struct {
	Name           string `json:"name" validate:"required"`
	SocialMediaURL string `json:"social_media_url" validate:"required"`
}

type SocialUpdateRes struct {
	ID string `json:"id"`
}

type SocialDeleteRes struct {
	ID string `json:"id"`
}

type SocialResponse struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	SocialMediaURL string    `json:"social_media_url"`
	UserID         string    `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
