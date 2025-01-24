package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/uidea/artwork-backend/global"
	"github.com/uidea/artwork-backend/internal/service"
	"github.com/uidea/artwork-backend/pkg/app"
	"github.com/uidea/artwork-backend/pkg/errcode"
)

type About struct{}

func NewAbout() About {
	return About{}
}

func (a About) Upsert(c *gin.Context) {
	param := service.UpsertAboutRequest{}
	response := app.NewResponse(c)
	vaild, errs := app.BindAndValid(c, &param)

	if !vaild {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpsertAbout(&param)

	if err != nil {
		global.Logger.Errorf(c, "svc.UpsertAbout err: %v", err)
		response.ToErrorResponse(errcode.AlreadyExistsRecord.WithDetails(err.Error()))
		return
	}

	response.ToResponse(gin.H{})
}

func (t About) Get(c *gin.Context) {
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())
	artwork, err := svc.GetAbout()
	if err != nil {
		global.Logger.Errorf(c, "svc.GetArtwork err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetArtworkFail)
		return
	}
	response.ToResponse(artwork)
}
