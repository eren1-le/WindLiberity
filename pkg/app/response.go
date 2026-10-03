/*
 * @Author: eren dengdengd1222@mail.com
 * @Date: 2026-09-23 10:49:15
 * @LastEditors: eren dengdengd1222@mail.com
 * @LastEditTime: 2026-09-23 11:09:39
 * @FilePath: /WindLiberity/pkg/app/response.go
 * @Description: 
 * 
 */
package app

import (
	"net/http"

	"github.com/binbinly/pkg/errno"
	"github.com/binbinly/pkg/util"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int 	`json:"code"`
	Msg string 	`json:"msg"`
	Data any 	`json:"data"`
}


func Success (c *gin.Context, data any) {
	if data == nil {
		data = gin.H{}
	}

	c.AbortWithStatusJSON(http.StatusOK, Response {
		Code: errno.Success.Code(),
		Msg:  errno.Success.Msg(),
		Data: data,
	}) 
}

func SuccessNil(c *gin.Context) {
	Success(c, nil)
}

func Error(c *gin.Context, err *errno.Error) {
	code, msg := errno.DecodeErr(err)
	c.AbortWithStatusJSON(http.StatusOK, Response{
		Code: code,
		Msg: msg,
		Data: gin.H{},
	})
}

func ErrorParamInvalid(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusOK, Response{
		Code: errno.ErrInvalidParam.Code(),
		Msg: err.Error(),
		Data: gin.H{},
	})
}

func RouteNotFound(c *gin.Context) {
	c.String(http.StatusNotFound, "not found")
}

type healthCheckResponse struct {
	Status   string `json:"status"`
	Hostname string `json:"hostname"`
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, healthCheckResponse{Status: "UP", Hostname: util.Hostname()})
}
