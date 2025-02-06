package service

import (
	"log"
	"os"
	"time"

	"github.com/uidea/artwork-backend/pkg/app"
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
	ID        uint32     `json:"id"`
	Content   string     `json:"content"`
	Cover     FileFormat `json:"cover"`
	CoverBlur FileFormat `json:"cover_blur"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type FileFormat struct {
	Filename string `json:"filename"`
	Url      string `json:"url"`
	ID       string `json:"id"`
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

	appEnv := os.Getenv("APP_ENV")
	var bucketName string
	if appEnv == "production" {
		bucketName = os.Getenv("MINIO_BUCKETNAME")
	} else {
		bucketName = "testbucket"
	}

	coverPresignedURL, err := app.GetMinioPresignedURL(bucketName, about.CoverFile.StorageName)
	if err != nil {
		log.Println("❌ 無法產生 Presigned URL:", err)
		return nil, err
	}

	coverBlurFilepresignedURL, err := app.GetMinioPresignedURL(bucketName, about.CoverBlurFile.StorageName)
	if err != nil {
		log.Println("❌ 無法產生 Presigned URL:", err)
		return nil, err
	}

	return &About{
		ID:      about.ID,
		Content: about.Content,
		Cover: FileFormat{
			Url:      coverPresignedURL.String(),
			Filename: about.CoverFile.Filename,
			ID:       about.CoverFile.ID,
		},
		CoverBlur: FileFormat{
			Url:      coverBlurFilepresignedURL.String(),
			Filename: about.CoverBlurFile.Filename,
			ID:       about.CoverBlurFile.ID,
		},
		CreatedAt: about.CreatedAt,
		UpdatedAt: about.UpdatedAt,
	}, nil
}
