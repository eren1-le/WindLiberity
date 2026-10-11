/*
 * @Author: eren dengdengd1222@mail.com
 * @Date: 2026-10-08 09:32:00
 * @LastEditors: eren dengdengd1222@mail.com
 * @LastEditTime: 2026-10-10 11:08:04
 * @FilePath: /WindLiberity/internal/service/service.go
 * @Description:
 *
 */
package service

import (

	"WindLiberity/internal/repository"
	"WindLiberity/internal/ws"
	"WindLiberity/pkg/app"
	"WindLiberity/transport/websocket"

	"github.com/binbinly/pkg/cache"

	"github.com/redis/go-redis/v9"
)

const (
	msgFriendCreate = "你们已经是好友了，可以开始聊天啦"
	msgKickOut      = "账号已在其他地方登录了!"
)

// 用于触发编译期的接口的合理性检查机制
var _ IService = (*Service)(nil)

type IService interface {
	User
}

var Svc IService

type Service struct {
	opts options
	repo repository.IRepo
	rdb  *redis.Client
	ws   *ws.Server
}

// New init service
func New(ws1 websocket.Server, opts ...Option) (s *Service) {
	rdb := app.InitRedis()
	s = &Service{
		opts: newOptions(opts...),
		repo: repository.New(app.InitDB(), cache.NewRedisCache(rdb)),
		rdb:  rdb,
		ws:   ws.New(ws1, rdb),
	}

	Svc = s
	return s
	
}
// Close service
func (s *Service) Close() error {
	return s.rdb.Close()
}