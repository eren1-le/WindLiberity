package server

import (
	"context"
	"net/http"

	"WindLiberity/pkg/app"
	"WindLiberity/transport/websocket"

	"WindLiberity/transport/websocket"

	"github.com/binbinly/pkg/logger"
	"github.com/binbinly/pkg/transport/ws"

	"github.com/rs/xid"
)

//New Server websocket
func NewWsServer(conf *app.ServerConfig) websocket.Server {
	s := websocket.NewServer()
	s.Init(websocket.WithID(xid.New().String()),
		websocket.WithAddr(conf.Addr),
		websocket.WithWriteWait(conf.WriteTimeout)
		websocket.WithRouter())
}