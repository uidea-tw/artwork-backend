package model

import (
	"fmt"

	"github.com/uidea/artwork-backend/pkg/setting"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}

type config struct {
	user   string
	pass   string
	adrr   string
	port   string
	dbname string
}

func NewDBEngine(database *setting.DatabaseSettingS) (*gorm.DB, error) {
	conf := &config{
		user:   database.UserName,
		pass:   database.Password,
		adrr:   "127.0.0.1",
		port:   "5432",
		dbname: database.DBName,
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable TimeZone=%s", conf.adrr, conf.user, conf.pass, conf.dbname, conf.port, "Asia/Taipei")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 打印所有sql
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}
