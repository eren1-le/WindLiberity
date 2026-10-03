/*
 * @Author: eren dengdengd1222@mail.com
 * @Date: 2026-09-26 16:43:36
 * @LastEditors: eren dengdengd1222@mail.com
 * @LastEditTime: 2026-09-26 16:43:39
 * @FilePath: /WindLiberity/pkg/middleware/time.go
 * @Description:
 *
 */
package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

// Timeout 超时中间件
func Timeout(t time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
