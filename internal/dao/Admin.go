package dao

import "github.com/uidea/artwork-backend/internal/model"

func (d *Dao) CreateAdmin(Name string, Username string, Password string) error {
	admin := model.Admin{
		Name:     Name,
		Username: Username,
		Password: Password,
	}
	return admin.Create(d.engine)
}

func (d *Dao) GetAdmin(username string) (model.Admin, error) {
	admin := model.Admin{Username: username}
	return admin.Get(d.engine)
}
