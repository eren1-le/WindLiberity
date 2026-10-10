package router

import  (

	"WindLiberity/internal/ws"

	"github.com/binbinly/pkg/logger"
	"WindLiberity/transport/websocket"
)

// NewWsRouter
func NewWsRouter() *websocket.Engine {
	r := websocket.NewEngine()
	r.Use(func(c *websocket.Context) {
		logger.Debugf("[ws] event: %v", c.Req.Event())
		c.Next()
	})
	r.AddRoute("ping", websocket.HandlerChain{Ping})

	return r
}

// Ping
func Ping(c *websocket.Context) {
	if err := c.Req.Conn().Send(c, 0, ws.Pack("pong", "")); err != nil {
		logger.Warnf("[ws.ping] err: %v", err)
	}
}

