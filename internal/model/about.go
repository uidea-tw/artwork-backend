package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type About struct {
	ID        uint32    `gorm:"primary_key" json:"id"`
	Content   string    `json:"content"`
	Cover     string    `json:"cover"`
	CoverBlur string    `json:"cover_blur"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (a About) TableName() string {
	return "abouts"
}

func (a About) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

func (a About) Get(db *gorm.DB) (*About, error) {
	var about About
	err := db.First(&about).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &about, nil
}

func (a About) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&a).Where("id = ?", a.ID).Updates(values).Error
}

func (a About) Save(db *gorm.DB, about *About) error {
	return db.Save(about).Error
}

func (a *About) First(db *gorm.DB) (*About, error) {
	var about About
	err := db.First(a).Error
	if err != nil {
		return nil, err
	}

	return &about, nil
}
