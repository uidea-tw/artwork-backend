package v1

import "github.com/gin-gonic/gin"

type Article struct{}

func NewArticle() Article {
	return Article{}
}

// @Summary	[AC-001] 取得文章資訊
// @Tags		Aritcle
// @Success	200	{object}	app.EmptySuccessResponse{}	"成功"
// @Failure	400	{object}	errcode.Error				"請求失敗"
// @Failure	500	{object}	errcode.Error				"伺服器異常"
// @Router		/articles/{id} [get]
func (t Article) Get(c *gin.Context) {}

// @Summary	[AC-002] 取得文章列表
// @Tags		Aritcle
// @Success	200	{object}	app.EmptySuccessResponse{}	"成功"
// @Failure	400	{object}	errcode.Error				"請求失敗"
// @Failure	500	{object}	errcode.Error				"伺服器異常"
// @Router		/articles [delete]
func (t Article) List(c *gin.Context) {}

// @Summary	[AC-003] 新增文章
// @Tags		Aritcle
// @Success	200	{object}	app.EmptySuccessResponse{}	"成功"
// @Failure	400	{object}	errcode.Error				"請求失敗"
// @Failure	500	{object}	errcode.Error				"伺服器異常"
// @Router		/articles [post]
func (t Article) Create(c *gin.Context) {
	// 1. 獲取參數

}

// @Summary	[AC-004] 更新文章
// @Tags		Aritcle
// @Success	200	{object}	app.EmptySuccessResponse{}	"成功"
// @Failure	400	{object}	errcode.Error				"請求失敗"
// @Failure	500	{object}	errcode.Error				"伺服器異常"
// @Router		/articles/{id} [put]
func (t Article) Update(c *gin.Context) {}

// @Summary	[AC-005] 刪除文章
// @Tags		Aritcle
// @Success	200	{object}	app.EmptySuccessResponse{}	"成功"
// @Failure	400	{object}	errcode.Error				"請求失敗"
// @Failure	500	{object}	errcode.Error				"伺服器異常"
// @Router		/articles/{id} [delete]
func (t Article) Delete(c *gin.Context) {}
