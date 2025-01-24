package service

import (
	"time"

	"github.com/uidea/artwork-backend/internal/model"
)

type CreateArtworkRequest struct {
	Title       string `json:"title"`
	Cover       string `json:"cover"`
	CoverBlur   string `json:"cover_blur"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Author      string `json:"author"`
	Year        uint32 `json:"year"`
	Length      uint32 `json:"length"`
	Width       uint32 `json:"width"`
	Height      uint32 `json:"height"`
	Price       uint32 `json:"price"`
}

type ArtworRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

type UpdateArtworRequest struct {
	ID          uint32 `form:"id" binding:"required,gte=1"`
	Title       string `json:"title"`
	Cover       string `json:"cover"`
	CoverBlur   string `json:"cover_blur"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Author      string `json:"author"`
	Year        uint32 `json:"year"`
	Length      uint32 `json:"length"`
	Width       uint32 `json:"width"`
	Height      uint32 `json:"height"`
	Price       uint32 `json:"price"`
}

type Artwork struct {
	ID          uint32    `json:"id"`
	Title       string    `json:"title"`
	Cover       string    `json:"cover"`
	CoverBlur   string    `json:"cover_blur"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	Author      string    `json:"author"`
	Year        uint32    `json:"year"`
	Length      uint32    `json:"length"`
	Width       uint32    `json:"width"`
	Height      uint32    `json:"height"`
	Price       uint32    `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (svc *Service) CreateArtWork(param *CreateArtworkRequest) error {
	return svc.dao.CreateArtwork(
		param.Title,
		param.Cover,
		param.CoverBlur,
		param.Description,
		param.Content,
		param.Author,
		param.Year,
		param.Length,
		param.Width,
		param.Height,
		param.Price,
	)
}

func (svc *Service) GetArtwork(param *ArtworRequest) (*Artwork, error) {
	artwork, err := svc.dao.GetArtwork(param.ID)
	if err != nil {
		return nil, err
	}
	return &Artwork{
		ID:          artwork.ID,
		Title:       artwork.Title,
		Cover:       artwork.Cover,
		CoverBlur:   artwork.CoverBlur,
		Description: artwork.Description,
		Content:     artwork.Content,
		Author:      artwork.Author,
		Year:        artwork.Year,
		Length:      artwork.Length,
		Width:       artwork.Width,
		Height:      artwork.Height,
		Price:       artwork.Price,
		CreatedAt:   artwork.CreatedAt,
		UpdatedAt:   artwork.UpdatedAt,
	}, nil
}

func (svc *Service) DeleteArtwork(param *ArtworRequest) error {
	err := svc.dao.DeleteArtwork(param.ID)
	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) GetArtWorkList() ([]*model.Artwork, error) {
	return svc.dao.ListArtWork()
}

func (svc *Service) UpdateArtWork(param *UpdateArtworRequest) error {
	return svc.dao.UpdateArtwork(
		param.ID,
		param.Title,
		param.Cover,
		param.CoverBlur,
		param.Description,
		param.Content,
		param.Author,
		param.Year,
		param.Length,
		param.Width,
		param.Height,
		param.Price,
	)
}
