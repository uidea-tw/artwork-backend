package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/uidea/artwork-backend/global"
	"github.com/uidea/artwork-backend/internal/service"
	"github.com/uidea/artwork-backend/pkg/app"
	"github.com/uidea/artwork-backend/pkg/convert"
	"github.com/uidea/artwork-backend/pkg/errcode"
)

type User struct{}

func NewUser() User {
	return User{}
}

// @Summary	[US-001] 新增使用者
// @Tags		User
// @Success	200	{object}	app.EmptySuccessResponse{}	"成功"
// @Failure	400	{object}	errcode.Error				"請求失敗"
// @Failure	409	{object}	app.ErrorResponse{}			"資料已存在"
// @Failure	500	{object}	errcode.Error				"伺服器異常"
// @Router		/users/create [post]
func (t User) Create(c *gin.Context) {
	params := service.CreateUserRequest{}
	response := app.NewResponse(c)
	vaild, errs := app.BindAndValid(c, &params)

	if !vaild {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateUser(&params)
	if err != nil {
		global.Logger.Errorf(c, "svc.CreateUser err: %v", err)
		response.ToErrorResponse(errcode.ErrorUserFail)
		return
	}

	response.ToResponse(gin.H{})
}

// @Summary	[US-002] 取得使用者資訊
// @Tags		User
// @Success	20	{object}	app.EmptySuccessResponse{}	"成功"
// @Failure	400	{object}	errcode.Error				"請求失敗"
// @Failure	500	{object}	errcode.Error				"伺服器異常"
// @Router		/users/{id} [get]
func (t User) Get(c *gin.Context) {
	param := service.UserRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)

	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	user, err := svc.GetUser(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetUserFail)
		return
	}
	response.ToResponse(user)
}

// @Summary	[US-003] 取得使用者列表
// @Tags		User
// @Success	200	{object}	app.EmptySuccessResponse{}	"成功"
// @Failure	400	{object}	errcode.Error				"請求失敗"
// @Failure	500	{object}	errcode.Error				"伺服器異常"
// @Router		/users [get]
func (t User) List(c *gin.Context) {
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	users, err := svc.GetUserList()
	if err != nil {
		global.Logger.Errorf(c, "svc.GetTagList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}
	response.ToResponse(users)
}

// @Summary	[US-004] 更新使用者資訊
// @Tags		User
// @Success	200	{object}	app.EmptySuccessResponse{}	"成功"
// @Failure	400	{object}	errcode.Error				"請求失敗"
// @Failure	500	{object}	errcode.Error				"伺服器異常"
// @Router		/users/{id} [put]
func (t User) Update(c *gin.Context) {

}

// @Summary	[US-005] 刪除使用者
// @Tags		User
// @Success	200	{object}	app.EmptySuccessResponse{}	"成功"
// @Failure	400	{object}	errcode.Error				"請求失敗"
// @Failure	500	{object}	errcode.Error				"伺服器異常"
// @Router		/users/{id} [delete]
func (t User) Delete(c *gin.Context) {
	param := service.UserRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.DeleteUser(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.DeleteArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteArticleFail)
		return
	}
	response.ToResponse(gin.H{})
}
