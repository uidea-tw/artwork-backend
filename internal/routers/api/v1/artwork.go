package v1

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/uidea/artwork-backend/global"
	"github.com/uidea/artwork-backend/internal/service"
	"github.com/uidea/artwork-backend/pkg/app"
	"github.com/uidea/artwork-backend/pkg/convert"
	"github.com/uidea/artwork-backend/pkg/errcode"
)

type Artwork struct{}

func NewArtwork() Artwork {
	return Artwork{}
}

// @Summary	[AC-001] 取得作品資訊
// @Tags		Artwork
// @Success	200	{object}	app.EmptySuccessResponse{}	"成功"
// @Failure	400	{object}	errcode.Error				"請求失敗"
// @Failure	500	{object}	errcode.Error				"伺服器異常"
// @Router		/artworks/{id} [get]
func (t Artwork) Get(c *gin.Context) {
	param := service.ArtworRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)

	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	artwork, err := svc.GetArtwork(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetArtwork err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetArtworkFail)
		return
	}
	response.ToResponse(artwork)
}

// @Summary	[AC-002] 取得作品列表
// @Tags		Artwork
// @Success	200	{object}	app.EmptySuccessResponse{}	"成功"
// @Failure	400	{object}	errcode.Error				"請求失敗"
// @Failure	500	{object}	errcode.Error				"伺服器異常"
// @Router		/artworks [delete]
func (t Artwork) List(c *gin.Context) {
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	artworks, err := svc.GetArtWorkList()
	if err != nil {
		global.Logger.Errorf(c, "svc.GetTagList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}
	response.ToResponse(artworks)
}

// @Summary	[AC-003] 新增作品
// @Tags		Artwork
// @Success	200	{object}	app.EmptySuccessResponse{}	"成功"
// @Failure	400	{object}	errcode.Error				"請求失敗"
// @Failure	500	{object}	errcode.Error				"伺服器異常"
// @Router		/artworks [post]
func (t Artwork) Create(c *gin.Context) {
	param := service.CreateArtworkRequest{}
	response := app.NewResponse(c)
	vaild, errs := app.BindAndValid(c, &param)
	fmt.Printf("errs: %v \n", errs)
	fmt.Printf("vaild: %v \n", vaild)
	if !vaild {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateArtWork(&param)

	if err != nil {
		global.Logger.Errorf(c, "svc.CreateArtwork err: %v", err)
		response.ToErrorResponse(errcode.AlreadyExistsRecord.WithDetails(err.Error()))
		return
	}

	response.ToResponse(gin.H{})
}

// @Summary	[AC-004] 更新作品
// @Tags		Artwork
// @Success	200	{object}	app.EmptySuccessResponse{}	"成功"
// @Failure	400	{object}	errcode.Error				"請求失敗"
// @Failure	500	{object}	errcode.Error				"伺服器異常"
// @Router		/artworks/{id} [put]
func (t Artwork) Update(c *gin.Context) {
	param := service.UpdateArtworRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.UpdateArtWork(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.UpdateArtwork err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}

	response.ToResponse(gin.H{})
}

// @Summary	[AC-005] 刪除作品
// @Tags		Artwork
// @Success	200	{object}	app.EmptySuccessResponse{}	"成功"
// @Failure	400	{object}	errcode.Error				"請求失敗"
// @Failure	500	{object}	errcode.Error				"伺服器異常"
// @Router		/artworks/{id} [delete]
func (t Artwork) Delete(c *gin.Context) {
	param := service.ArtworRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)

	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteArtwork(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetArtwork err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteArtworkFail)
		return
	}
	response.ToResponse(gin.H{})
}
