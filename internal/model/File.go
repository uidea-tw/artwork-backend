package model

import (
	"time"

	"gorm.io/gorm"
)

type File struct {
	ID          string    `gorm:"type:uuid" json:"id"`
	Filename    string    `json:"filename"`
	StorageName string    `json:"storage_name"`
	Bucket      string    `json:"bucket"`
	ContentType string    `json:"content_type"`
	Size        int64     `json:"size"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (f File) Get(db *gorm.DB) (File, error) {
	var file File
	err := db.Where("id = ?", f.ID).First(&file).Error
	return file, err
}

func (f File) TableName() string {
	return "files"
}

func (f *File) Create(db *gorm.DB) error {
	return db.Create(&f).Error
}

func (f *File) Delete(db *gorm.DB) error {
	return db.Delete(&f).Error
}
