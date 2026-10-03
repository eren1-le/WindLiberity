/*
 * @Author: eren dengdengd1222@mail.com
 * @Date: 2026-09-22 15:02:01
 * @LastEditors: eren dengdengd1222@mail.com
 * @LastEditTime: 2026-09-26 15:47:04
 * @FilePath: /WindLiberity/pkg/middleware/limit.go
 * @Description:
 *
 */
package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"WindLiberity/pkg/app"
)

//see: https://github.com/aviddiviner/gin-limit/blob/master/limit.go

// MaxLimiter 限制同时最大请求数
func MaxLimiter(n int) gin.HandlerFunc {
	sem := make(chan struct{}, n)
	acquire := func() { sem <- struct{}{} }
	release := func() { <-sem }
	return func(c *gin.Context) {
		if n == 0 {
			c.Next()
			return
		}
		acquire()       // before request
		defer release() // after request
		c.Next()
	}
}

// IPLimiter ip限制
func IPLimiter(limit int, expire time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		if app.Redis == nil || limit == 0 {
			c.Next()
			return
		}
		key := fmt.Sprint("ip-limit:", c.ClientIP())

		count, _ := app.Redis.Get(c.Request.Context(), key).Int()

		if count >= limit {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"code": http.StatusTooManyRequests,
				"msg":  "too many request",
			})
			return
		}

		c.Next()
		pipe := app.Redis.Pipeline()
		pipe.Incr(c.Request.Context(), key)
		pipe.Expire(c.Request.Context(), key, expire)
		pipe.Exec(c.Request.Context())
	}

}
