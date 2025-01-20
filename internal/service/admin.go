package service

import (
	"errors"
	"fmt"

	"github.com/uidea/artwork-backend/internal/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type CreateAdminRequest struct {
	Name     string `json:"name" binding:"required,min=2,max=20" example:"admin"`
	Username string `json:"username" binding:"required,min=5,max=20" example:"admin"`
	Password string `json:"password" binding:"required,min=5,max=20" example:"password"`
}

type LoginAdminRequest struct {
	Username string `json:"username" binding:"required,min=5,max=20"`
	Password string `json:"password" binding:"required,min=5,max=20"`
}

func (svc *Service) CreateAdmin(param *CreateAdminRequest) error {
	_, err := svc.dao.GetAdmin(param.Username)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		password, _ := bcrypt.GenerateFromPassword([]byte(param.Password), bcrypt.DefaultCost)
		return svc.dao.CreateAdmin(param.Name, param.Username, string(password))
	}

	if err != nil {
		return fmt.Errorf("failed to query admin: %w", err)
	}

	return errors.New("admin already exists")
}

func (svc *Service) CheckAuth(param *LoginAdminRequest) (model.Admin, error) {
	admin, err := svc.dao.GetAdmin(param.Username)
	if err != nil {
		return model.Admin{}, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(param.Password)); err != nil {
		return model.Admin{}, err
	}
	return admin, nil
}

func (svc *Service) GetAdminById(id string) (model.Admin, error) {
	admin, err := svc.dao.GetAdminByID(id)
	if err != nil {
		return model.Admin{}, err
	}
	return admin, nil
}
