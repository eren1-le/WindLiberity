package ws

import (
	"context"
	"encoding/json"
	"time"

	"WindLiberity/pkg/app"
	"WindLiberity/transport/websocket"

	"github.com/binbinly/pkg/logger"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

// UserConnInfo
type UserConnInfo struct {
	UserID   int    `json:"user_id"`
	ConnID   uint64 `json:"conn_id"`
	ServerID string `json:"server_id"`
}

// Server websocket server
type Server struct {
	ws 	websocket.Server
	rdb *redis.Client

}

// New websocket server
func New(ws websocket.Server, rdb *redis.Client) *Server {
	return &Server{
		ws: ws,
		rdb: rdb,
	}
}

// Send
func (w *Server) Send(ctx context.Context, c *UserConnInfo, event string, data any) (err error) {
	return w.send(ctx, c, Pack(event, data))
}

// send
func (w *Server) send(ctx context.Context, c *UserConnInfo,  msg []byte) error {
	if c.ConnID == 0 {
		w.SaveHistory(ctx, c, msg)
		return nil
	}
	conn, err := w.ws.GetManager(c.ConnID).Get(c.ConnID)
	if errors.Is(err, websocket.ErrConnNotFound) {
		w.SaveHistory(ctx, c, msg)
		return nil
	} else if err != nil {
		return errors.Wrapf(err, "[ws.Close] get conn by cid: %v", c.ConnID)
	}
	if err = conn.AsyncSend(ctx,0, msg); err != nil {
		return errors.Wrapf(err, "[ws.send] AsyncSend msg:%v", msg)
	}

	return nil
}

// BatchSendConn
func (w *Server) BatchSendConn(ctx context.Context, cs []*UserConnInfo, event string, data any) (err error) {
	msg := Pack(event, data)
	for _, c := range cs {	
		if err = w.send(ctx, c, msg); err != nil {
			logger.Warnf("[ws.batchSend] err: %v", c.UserID, err)
		}
	}
	return err
}

// BatchSendMessage
func (w *Server) BatchSendMessage(ctx context.Context, c *UserConnInfo, list []string) error {
	if len(list) == 0 {
		return nil
	}
	conn, err := w.ws.GetManager(c.ConnID).Get(c.ConnID)
	if errors.Is(err, websocket.ErrConnNotFound) {
		return nil
	} else if err != nil {
		return errors.Wrapf(err, "[ws.Close] get conn by uid: %v", c.UserID)
	}
	for _, msg := range list {
		if err = conn.AsyncSend(ctx, 0, []byte(msg)); err != nil {
			logger.Warnf("[ws.BatchSendMessage] AsyncSend uid:%d err: %v", c.UserID, err)
		}
	}
	return err
}

// Close 
func (w *Server) Close(ctx context.Context, c *UserConnInfo, data any) error {
	conn, err := w.ws.GetManager(c.ConnID).Get(c.ConnID)
	if errors.Is(err, websocket.ErrConnNotFound) {
		return nil
	} else if err != nil {
		return errors.Wrapf(err, "[ws.Close] get conn by uid: %v", c.UserID)
	}

	if data != nil {
		_ = conn.Send(ctx, 0, Pack(EventClose, data))
	}
	conn.Stop()
	return nil
}
// Broadcast
func (w *Server) Broadcast(ctx context.Context, omit int, event string, data any) (err error) {
	msg := Pack(event, data)
	w.ws.Range(func(cid uint64, conn websocket.Connection) error {
		if conn.GetUID() == omit {
			return nil
		}
		return w.send(ctx, &UserConnInfo{ConnID: cid}, msg)
	})
	return
}
// SaveHistory
func (w *Server) SaveHistory(ctx context.Context, c *UserConnInfo, msg []byte) {
	pipe := w.rdb.Pipeline()
	key := app.BuildHistoryKey(c.UserID)
	pipe.RPush(ctx,key,msg)
	pipe.Expire(ctx,key,14*time.Hour)
	if _, err := pipe.Exec(ctx); err != nil {
		logger.Warnf("[ws.saveHistory] exec err: %v", err)
	}
}

// Pack
func Pack(event string, data any) []byte {
	msg, _ := json.Marshal(struct {
		Event string `json:"event"`
		Data  any 	 `json:"data"`
	}{
		Event:  event,
		Data:	data,
	})
	return msg
}