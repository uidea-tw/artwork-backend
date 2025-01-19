package v1

import "github.com/gin-gonic/gin"

type User struct{}

func NewUser() User {
	return User{}
}

//	@Summary	[US-001] 登入使用者
//	@Produce	json
//	@Tags		User
//
//	@Param		body	body		service.LoginAdminRequest	true	"登入請求參數"
// w
//	@Success	200		{object}	app.EmptySuccessResponse{}	"成功"
//	@Failure	400		{object}	errcode.Error				"請求失敗"
//	@Failure	500		{object}	errcode.Error				"伺服器異常"
//	@Router		/users/auth/login [post]
func (t User) Login(c *gin.Context) {}

//	@Summary	[US-002] 登出使用者
//	@Tags		User
//	@Success	200	{object}	app.EmptySuccessResponse{}	"成功"
//	@Failure	400	{object}	errcode.Error				"請求失敗"
//	@Failure	500	{object}	errcode.Error				"伺服器異常"
//	@Router		/users/auth/logout [post]
func (t User) Logout(c *gin.Context) {}

//	@Summary	[US-003] 註冊帳號
//	@Tags		User
//	@Success	200	{object}	app.EmptySuccessResponse{}	"成功"
//	@Failure	400	{object}	errcode.Error				"請求失敗"
//	@Failure	409	{object}	app.ErrorResponse{}			"資料已存在"
//	@Failure	500	{object}	errcode.Error				"伺服器異常"
//	@Router		/users/auth/signup [post]
func (t User) Signup(c *gin.Context) {}

//	@Summary	[US-004] 取得使用者資訊
//	@Tags		User
//	@Success	200	{object}	app.EmptySuccessResponse{}	"成功"
//	@Failure	400	{object}	errcode.Error				"請求失敗"
//	@Failure	500	{object}	errcode.Error				"伺服器異常"
//	@Router		/users/{id} [get]
func (t User) Get(c *gin.Context) {}

//	@Summary	[US-005] 取得使用者列表
//	@Tags		User
//	@Success	200	{object}	app.EmptySuccessResponse{}	"成功"
//	@Failure	400	{object}	errcode.Error				"請求失敗"
//	@Failure	500	{object}	errcode.Error				"伺服器異常"
//	@Router		/users [get]
func (t User) List(c *gin.Context) {}

//	@Summary	[US-006] 更新使用者資訊
//	@Tags		User
//	@Success	200	{object}	app.EmptySuccessResponse{}	"成功"
//	@Failure	400	{object}	errcode.Error				"請求失敗"
//	@Failure	500	{object}	errcode.Error				"伺服器異常"
//	@Router		/users/{id} [put]
func (t User) Update(c *gin.Context) {}

//	@Summary	[US-007] 刪除使用者
//	@Tags		User
//	@Success	200	{object}	app.EmptySuccessResponse{}	"成功"
//	@Failure	400	{object}	errcode.Error				"請求失敗"
//	@Failure	500	{object}	errcode.Error				"伺服器異常"
//	@Router		/users/{id} [delete]
func (t User) Delete(c *gin.Context) {}
