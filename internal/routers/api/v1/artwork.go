package v1

import "github.com/gin-gonic/gin"

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
func (t Artwork) Get(c *gin.Context) {}

// @Summary	[AC-002] 取得作品列表
// @Tags		Artwork
// @Success	200	{object}	app.EmptySuccessResponse{}	"成功"
// @Failure	400	{object}	errcode.Error				"請求失敗"
// @Failure	500	{object}	errcode.Error				"伺服器異常"
// @Router		/artworks [delete]
func (t Artwork) List(c *gin.Context) {}

// @Summary	[AC-003] 新增作品
// @Tags		Artwork
// @Success	200	{object}	app.EmptySuccessResponse{}	"成功"
// @Failure	400	{object}	errcode.Error				"請求失敗"
// @Failure	500	{object}	errcode.Error				"伺服器異常"
// @Router		/artworks [post]
func (t Artwork) Create(c *gin.Context) {}

// @Summary	[AC-004] 更新作品
// @Tags		Artwork
// @Success	200	{object}	app.EmptySuccessResponse{}	"成功"
// @Failure	400	{object}	errcode.Error				"請求失敗"
// @Failure	500	{object}	errcode.Error				"伺服器異常"
// @Router		/artworks/{id} [put]
func (t Artwork) Update(c *gin.Context) {}

// @Summary	[AC-005] 刪除作品
// @Tags		Artwork
// @Success	200	{object}	app.EmptySuccessResponse{}	"成功"
// @Failure	400	{object}	errcode.Error				"請求失敗"
// @Failure	500	{object}	errcode.Error				"伺服器異常"
// @Router		/artworks/{id} [delete]
func (t Artwork) Delete(c *gin.Context) {}
