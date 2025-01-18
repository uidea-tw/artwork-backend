package v1

import "github.com/gin-gonic/gin"

type Author struct{}

func NewAuthor() Author {
	return Author{}
}

//	@Summary	[AM-001] 登入後台管理員
//	@Produce	json
//	@Tags		Author
//
//	@Param		body	body		service.LoginAdminRequest	true	"登入請求參數"
//
//	@Success	200		{object}	app.EmptySuccessResponse{}	"成功"
//	@Failure	400		{object}	errcode.Error				"請求失敗"
//	@Failure	500		{object}	errcode.Error				"伺服器異常"
//	@Router		/api/v1/admins/login [post]
func (t Author) Login(c *gin.Context) {}

func (t Author) Logout(c *gin.Context)   {}
func (t Author) Register(c *gin.Context) {}
func (t Author) Get(c *gin.Context)      {}
func (t Author) List(c *gin.Context)     {}
func (t Author) Create(c *gin.Context)   {}
func (t Author) Update(c *gin.Context)   {}
func (t Author) Delete(c *gin.Context)   {}
