package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Artwork struct {
	ID          uint32    `gorm:"primary_key" json:"id"`
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

func (a Artwork) TableName() string {
	return "artworks"
}

func (a Artwork) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

func (a Artwork) Get(db *gorm.DB) (*Artwork, error) {
	var artwork Artwork
	db = db.Where("id = ?", a.ID)
	err := db.First(&artwork).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &artwork, nil
}

func (a Artwork) Delete(db *gorm.DB) error {
	if err := db.Where("id = ?", a.ID).Delete(&a).Error; err != nil {
		return err
	}
	return nil
}

func (a Artwork) List(db *gorm.DB) ([]*Artwork, error) {
	var artworks []*Artwork
	if err := db.Find(&artworks).Error; err != nil {
		return nil, err
	}
	return artworks, nil
}

func (a Artwork) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&a).Where("id = ?", a.ID).Updates(values).Error
}
