
package websocket

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"

	"github.com/binbinly/pkg/logger"
	"github.com/binbinly/pkg/util"
	"github.com/gorilla/websocket"
)

var (
	// ErrConnNotFound 连接未找到
	ErrConnNotFound = errors.New("connection not found")
	// ErrConnNotFinish 连接未完成，不可以发送消息
	ErrConnNotFinish = errors.New("connection not finish when send msg")
)

type ConnHandlerFunc func(cin uint64, conn Connection) error

// Server is a simple micro server abstraction
type Server interface {
	// Init Initialise options
	Init(...Option)
	// Options Retrieve the options
	Options() *Options
	// Start the server
	Start(ctx context.Context) error
	// Stop the server
	Stop(ctx context.Context) error
	// Endpoint return a real address to registry endpoint.
	Endpoint() (*url.URL, error)
	// Range 遍历所有连接
	Range(f ConnHandlerFunc)
	// Total 服务器连接总数
	Total() int
}

// wsServer
type wsServer struct {
	managers []*Manager
	handler  *Handler
	opts     *Options
	lis      net.Listener
	endpoint *url.URL
	upgrader *websocket.Upgrader
}
