package service

import (
	"time"

	"github.com/uidea/artwork-backend/internal/model"
)

type UserRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

type CreateUserRequest struct {
	Name         string `json:"name"`
	NickName     string `json:"nickname"`
	Introduction string `json:"introduction"`
	Birth        string `json:"birth"`
	Gender       string `json:"gender"`
}

type User struct {
	ID           uint32    `json:"id"`
	Name         string    `json:"name"`
	Nickname     string    `json:"nickname"`
	Introduction string    `json:"introduction"`
	Birth        string    `json:"birth"`
	Gender       string    `json:"gender"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (svc *Service) GetUser(param *UserRequest) (*User, error) {
	user, err := svc.dao.GetUser(param.ID)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:           user.ID,
		Name:         user.Name,
		Nickname:     user.Nickname,
		Introduction: user.Introduction,
		Birth:        user.Birth,
		Gender:       user.Gender,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}, nil
}

func (svc *Service) CreateUser(param *CreateUserRequest) error {
	return svc.dao.CreateUser(
		param.Name,
		param.NickName,
		param.Introduction,
		param.Birth,
		param.Gender,
	)
}

func (svc *Service) DeleteUser(param *UserRequest) error {
	err := svc.dao.DeleteUser(param.ID)
	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) GetUserList() ([]*model.User, error) {
	return svc.dao.ListUser()
}
