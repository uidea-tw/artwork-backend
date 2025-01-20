package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/uidea/artwork-backend/global"
	"github.com/uidea/artwork-backend/internal/service"
	"github.com/uidea/artwork-backend/pkg/app"
	"github.com/uidea/artwork-backend/pkg/errcode"
)

type Admin struct{}

func NewAdmin() Admin {
	return Admin{}
}

// @Summary	[AM-001] 登入後台管理員
// @Produce	json
// @Tags		Admin
// @Param		body	body		service.LoginAdminRequest	true	"登入請求參數"
// @Success	200		{object}	app.EmptySuccessResponse{}	"成功"
// @Failure	400		{object}	errcode.Error				"請求失敗"
// @Failure	500		{object}	errcode.Error				"伺服器異常"
// @Router		/admins/auth/login [post]
func (a Admin) Login(c *gin.Context) {
	param := service.LoginAdminRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	admin, err := svc.CheckAuth(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CheckAuth err: %v", err)
		response.ToErrorResponse(errcode.NotFoundRegisterAccount)
		return
	}

	token, err := app.GenerateAccessToken(param.Username, param.Password, admin.ID)
	if err != nil {
		global.Logger.Errorf(c, "app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	c.SetCookie("login_token", token, 3600, "/", "", false, true)

	response.ToResponse(gin.H{})
}

// @Summary	[AM-002] 登出後台管理員
// @Tags		Admin
// @Success	200	{object}	app.EmptySuccessResponse{}	"成功"
// @Failure	400	{object}	app.ErrorResponse{}		"請求失敗"
// @Failure	500	{object}	app.ErrorResponse{}		"伺服器異常"
// @Router		/admins/auth/logout [post]
func (a Admin) Logout(c *gin.Context) {
	c.SetCookie("login_token", "", -1, "/", "", false, true)
	response := app.NewResponse(c)
	response.ToResponse(gin.H{})
}

// @Summary	[AM-003] 建立後台管理員
// @Tags		Admin
// @Param		body	body	service.CreateAdminRequest	true	"新增後台管理員參數"
// @Accept		json
// @Success	200	{object}	app.EmptySuccessResponse{}	"成功"
// @Failure	400	{object}	app.ErrorResponse{}			"請求失敗"
// @Failure	409	{object}	app.ErrorResponse{}			"資料已存在"
// @Failure	500	{object}	app.ErrorResponse{}			"伺服器異常"
// @Router		/admins [post]ㄆ
func (a Admin) Create(c *gin.Context) {
	param := service.CreateAdminRequest{}
	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateAdmin(&param)

	if err != nil {
		global.Logger.Errorf(c, "svc.CreateAdmin err: %v", err)
		response.ToErrorResponse(errcode.AlreadyExistsRecord.WithDetails(err.Error()))
		return
	}

	response.ToResponse(gin.H{})
}

func (a Admin) Get(c *gin.Context) {
	response := app.NewResponse(c)

	cookie, err := c.Cookie("login_token")
	if err != nil {
		global.Logger.Errorf(c, "login token err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist.WithDetails(err.Error()))
		return
	}
	claims, _ := app.ParseToken(cookie)
	id, _ := claims.GetSubject()
	svc := service.New(c.Request.Context())
	admin, err := svc.GetAdminById(id)
	if err != nil {
		global.Logger.Errorf(c, "svc.GetAdminProfile err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenError.WithDetails(err.Error()))
		return
	}
	response.ToResponse(admin)
}
