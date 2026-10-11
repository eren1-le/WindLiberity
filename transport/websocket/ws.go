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
	"github.com/rs/xid"
	"github.com/zhenjl/cityhash"
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
	// GetManager 所有连接管理
	GetManager(cid uint64) *Manager
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

// NewServer
func NewServer() Server {
	return &wsServer{
		opts: defOptions,
	}
}

// Options
func (s *wsServer) Options() *Options {
	return s.opts
}

// Init
func (s *wsServer) Init(opts ...Option) {
	for _, o := range opts {
		o(s.opts)
	}
	if s.opts.ID == "" {
		s.opts.ID = xid.New().String()
	}
	//Initial Connection manager
	s.managers = make([]*Manager, s.opts.ManagerSize)
	for i := 0; i < s.opts.ManagerSize; i++ {
		s.managers[i] = NewManager()
	}

	// Initial Message Handler
	s.handler = NewHandler(s.opts.WorkerPoolSize, s.opts.Router)
	s.upgrader = &websocket.Upgrader{
		ReadBufferSize: s.opts.ReadBufferSize,
		WriteBufferSize: s.opts.WriteBufferSize,
		CheckOrigin: func(r *http.Request) bool { return true },
	}
}

// Start Server
func (s *wsServer) Start(ctx context.Context) error {
	s.handler.Init(s.opts.MaxWorkerTaskLen)
	return s.Listen()
}
// Stop
func (s *wsServer) Stop(ctx context.Context) error {
	log.Print(("[Websocket] server is stopping"))
	// shutdown listening and connection
	err := s.lis.Close()
	for _, manager := range s.managers {
		manager.Clear()
	}
	return err
}

func (s *wsServer) Listen() error {
	var cid uint64 = 1
	lis, err := net.Listen("tcp", s.opts.Addr)
	if err != nil {
		return err
	}
	s.lis = lis

	if _, err = s.Endpoint(); err != nil {
		return err
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//set max nums of connection, reject if exceed
		if s.Total() >= s.Options().MaxConn {
			logger.Warn("[ws.Start] connection size limit")
			return
		}
		// it need cerification information if websocket
		uid := 0
		if s.Options().OnConnAuth != nil {
			var ok bool
			if uid, ok = s.Options().OnConnAuth(r, s.opts.ID,cid); !ok {
				w.WriteHeader(401)
				return
			}
		}
		
		if len(r.Header.Get("Sec-Websocket-Protocol")) > 0 {
			s.upgrader.Subprotocols = websocket.Subprotocols(r)
		}
		// websocket connection
		c, err := s.upgrader.Upgrade(w, r, nil)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		conn := NewConnect(s, c, cid, uid)
		// add to manager
		s.GetManager(cid).Add(conn)
		conn.Start()
		cid++
	})
	log.Printf("[Websocket] server is listening on: %s", lis.Addr().String())
	if err = http.Serve(lis, nil); !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil

}

func (s *wsServer) GetManager(cid uint64) *Manager {
	str := strconv.FormatUint(cid, 10)
	idx := int(cityhash.CityHash32([]byte(str), 0) % uint32(len(s.managers)))
	return s.managers[idx]
}

func (s *wsServer) Range(f ConnHandlerFunc) {
	for _, manager := range s.managers {
		_ = manager.Range(f)
	}
}

func (s *wsServer) Endpoint() (*url.URL, error) {
	addr, err := util.Extract(s.opts.Addr, s.lis)
	if err != nil {
		return nil, err
	}
	s.endpoint = &url.URL{Scheme: "http", Host: addr}
	return s.endpoint, nil
}

// Total of connection
func (s *wsServer) Total() int {
	var c int
	for _, manager := range s.managers {
		c += manager.Len()
	}
	return c
}