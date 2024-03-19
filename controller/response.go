package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code MyCode 		`json:"code"`
	Msg  string 		`json:"msg"`
	Data interface{} 	`json:"data,omitempty”`
}

func ResponseError(ctx *gin.Context, c MyCode){
	rd := &ResponseData{
		Code:c,
		Msg: c.Msg(),
		Data: nil,
	}
	ctx.JSON(http.StatusOK,rd)
}

func ResponseSuccess(ctx *gin.Context,data interface{}){
	rd := &ResponseData{
		Code:CodeSuccess,
		Msg: CodeSuccess.Msg(),
		Data: data,
	}
	ctx.JSON(http.StatusOK,rd)
}