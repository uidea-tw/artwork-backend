package service

import (
	"context"

	"github.com/uidea/artwork-backend/global"
	"github.com/uidea/artwork-backend/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DBEngine)
	return svc
}
