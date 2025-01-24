package service

import (
	"time"
)

type UpsertAboutRequest struct {
	Content   string `json:"content"`
	Cover     string `json:"cover"`
	CoverBlur string `json:"cover_blur"`
}

type UpdateAboutRequest struct {
	ID        uint32 `form:"id" binding:"required,gte=1"`
	Content   string `json:"content"`
	Cover     string `json:"cover"`
	CoverBlur string `json:"cover_blur"`
}

type About struct {
	ID        uint32    `json:"id"`
	Content   string    `json:"content"`
	Cover     string    `json:"cover"`
	CoverBlur string    `json:"cover_blur"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (svc *Service) UpsertAbout(param *UpsertAboutRequest) error {
	return svc.dao.UpsertAbout(
		param.Content,
		param.Cover,
		param.CoverBlur,
	)
}

func (svc *Service) GetAbout() (*About, error) {
	about, err := svc.dao.GetAbout()
	if err != nil {
		return nil, err
	}
	return &About{
		ID:        about.ID,
		Content:   about.Content,
		Cover:     about.Cover,
		CoverBlur: about.CoverBlur,
		CreatedAt: about.CreatedAt,
		UpdatedAt: about.UpdatedAt,
	}, nil
}

func (svc *Service) UpdateAbout(param *UpdateAboutRequest) error {
	return svc.dao.UpdateAbout(
		param.ID,
		param.Content,
		param.Cover,
		param.CoverBlur,
	)
}
