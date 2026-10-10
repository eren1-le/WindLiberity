package server

import (

	"WindLiberity/internal/router"
	"WindLiberity/pkg/app"
	"WindLiberity/transport/websocket"


	"github.com/rs/xid"
)

//New Server websocket
func NewWsServer(conf *app.ServerConfig) websocket.Server {
	s := websocket.NewServer()
	s.Init(websocket.WithID(xid.New().String()),
		websocket.WithAddr(conf.Addr),
		websocket.WithWriteWait(conf.WriteTimeout),
		websocket.WithRouter(router.NewWsRouter()),
	
	)
	return s
}