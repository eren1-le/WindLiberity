/*
 * @Author: eren dengdengd1222@mail.com
 * @Date: 2026-10-08 09:22:38
 * @LastEditors: eren dengdengd1222@mail.com
 * @LastEditTime: 2026-10-10 11:09:52
 * @FilePath: /WindLiberity/internal/router/router_api_v1.go
 * @Description:
 *
 */
package router

import (
	"WindLiberity/internal/api/v1/user"

	"github.com/gin-gonic/gin"
)

func setApiv1(v1 *gin.RouterGroup) {
	// verify
	v1.POST("/reg", user.Register)
	v1.POST("/login", user.Login)
}
