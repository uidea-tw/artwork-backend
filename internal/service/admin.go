package service

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type CreateAdminRequest struct {
	Name     string `json:"name" binding:"required,min=2,max=20"`
	Username string `json:"username" binding:"required,min=5,max=20"`
	Password string `json:"password" binding:"required,min=5,max=20"`
}

type LoginAdminRequest struct {
	Username string `json:"username" binding:"required,min=5,max=20"`
	Password string `json:"password" binding:"required,min=5,max=20"`
}

func (svc *Service) CreateAdmin(param *CreateAdminRequest) error {
	_, err := svc.dao.GetAdmin(param.Username)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			password, _ := bcrypt.GenerateFromPassword([]byte(param.Password), bcrypt.DefaultCost)
			return svc.dao.CreateAdmin(param.Name, param.Username, string(password))
		} else {
			return fmt.Errorf("failed to query admin: %w", err)
		}
	} else {
		return errors.New("admin already exists")
	}
}

func (svc *Service) CheckAuth(param *LoginAdminRequest) error {
	admin, err := svc.dao.GetAdmin(param.Username)
	if err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(param.Password)); err != nil {
		return err
	}
	return nil
}
