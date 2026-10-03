/*
 * @Author: eren dengdengd1222@mail.com
 * @Date: 2026-09-29 15:32:30
 * @LastEditors: eren dengdengd1222@mail.com
 * @LastEditTime: 2026-09-29 15:32:32
 * @FilePath: /WindLiberity/storage/redis/config.go
 * @Description:
 *
 */
package redis
 
import "time"

type Config struct {
	Url          string
	Addr         string
	Username     string
	Password     string
	DB           int
	MinIdleConn  int
	PoolSize     int
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PoolTimeout  time.Duration
	Trace        bool
}
