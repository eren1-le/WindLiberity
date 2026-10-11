package websocket

import (
	"encoding/json"
)

type Request struct {
	conn  Connection
	event string
	data  []byte
}

// Message
type Message struct {
	Event string 			`json:"event"`
	Data  json.RawMessage	`json:"data"`
}

// Newrequest
func NewRequest(conn Connection, msg []byte) (*Request, error) {
	m := &Message{}
	if err := json.Unmarshal(msg, m); err != nil {
		return nil, err
	}
	return &Request{
		conn: conn,
		event: m.Event,
		data: m.Data,
	}, nil
}

func (r Request) Conn() Connection {
	return r.conn
}

func (r Request) Event() string {
	return r.event
}

func (r Request) Data() []byte {
	return r.data
}