package model

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID        uint32    `gorm:"primary_key" json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (a Admin) TableName() string {
	return "admins"
}

func (a Admin) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

func (a Admin) Get(db *gorm.DB) (Admin, error) {
	var admin Admin
	err := db.Where("username = ?", a.Username).First(&admin).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return Admin{}, err
	}
	return admin, nil
}
