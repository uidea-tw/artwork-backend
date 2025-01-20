package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint32    `gorm:"primary_key" json:"id"`
	Name         string    `json:"name"`
	Nickname     string    `json:"nickname"`
	Introduction string    `json:"introduction"`
	Birth        string    `json:"birth"`
	Gender       string    `json:"gender"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (a User) TableName() string {
	return "users"
}

func (u User) Create(db *gorm.DB) error {
	return db.Create(&u).Error
}

func (u User) Get(db *gorm.DB) (User, error) {
	var user User
	db = db.Where("id = ?", u.ID)
	err := db.First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return user, err
	}

	return user, nil
}

func (u User) Delete(db *gorm.DB) error {
	if err := db.Where("id = ?", u.ID).Delete(&u).Error; err != nil {
		return err
	}
	return nil
}

func (u User) List(db *gorm.DB) ([]*User, error) {
	var users []*User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
