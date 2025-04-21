package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"message"`
}

func ResponseWrap() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		status := ctx.Writer.Status()
		var resp Response
		resp.Code = status
		switch status {
		case http.StatusOK:
			resp.Msg = "success"
		default:
			resp.Msg = "error"
		}

		if data, exist := ctx.Get("data"); exist {
			resp.Data = data
		}

		ctx.JSON(status, resp)
	}
}
