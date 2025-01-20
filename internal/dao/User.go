package dao

import (
	"github.com/uidea/artwork-backend/internal/model"
)

func (d *Dao) CreateUser(
	Name string,
	Nickname string,
	Introduction string,
	Birth string,
	Gender string,
) error {
	user := model.User{
		Name:         Name,
		Nickname:     Nickname,
		Introduction: Introduction,
		Birth:        Birth,
		Gender:       Gender,
	}
	return user.Create(d.engine)
}

func (d *Dao) GetUser(id uint32) (model.User, error) {
	user := model.User{ID: id}
	return user.Get(d.engine)
}

func (d *Dao) DeleteUser(id uint32) error {
	user := model.User{ID: id}
	return user.Delete(d.engine)
}

func (d *Dao) ListUser() ([]*model.User, error) {
	user := model.User{}
	return user.List(d.engine)
}
