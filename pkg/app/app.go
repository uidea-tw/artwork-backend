package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uidea/artwork-backend/pkg/errcode"
)

type Response struct {
	Ctx *gin.Context
}

type EmptyObject struct{}

type Pager struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

type EmptySuccessResponse struct {
	Code int         `json:"code" example:"0"`
	Data EmptyObject `json:"data"`
	Msg  EmptyObject `json:"msg"`
}

type ErrorResponse struct {
	Code    int      `json:"code"`
	Msg     string   `json:"msg"`
	Details []string `json:"details"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{
		Ctx: ctx,
	}
}

func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}

	response := gin.H{"code": errcode.Success.Code(), "data": data, "msg": nil}
	r.Ctx.JSON(http.StatusOK, response)
}

func (r *Response) ToResponseList(list interface{}, totalRows int) {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"list": list,
		"pager": Pager{
			Page:      GetPage(r.Ctx),
			PageSize:  GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	})
}

func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{"code": err.Code(), "msg": err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}

	r.Ctx.JSON(err.StatusCode(), response)
}
