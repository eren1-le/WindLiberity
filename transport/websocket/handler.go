/*
 * @Author: eren dengdengd1222@mail.com
 * @Date: 2026-10-02 18:08:53
 * @LastEditors: eren dengdengd1222@mail.com
 * @LastEditTime: 2026-10-04 13:43:28
 * @FilePath: /WindLiberity/transport/websocket/handler.go
 * @Description:
 *
 */
package websocket

// Handler deal with Message
type Handler struct {
	size 	int				//size of work pool
	engine 	*Engine			// Route
	queue	[]chan *Request //Worker get task from message queue
}

// NewHandler
func NewHandler(size int, e *Engine) *Handler {
	return &Handler{
		engine: e,
		size:	size,
		queue:  make([]chan *Request, size),
	}
}

// init work pool
func (h *Handler) Init(len int) {
	for i := 0; i < h.size; i++ {
		// init
		h.queue[i] = make(chan *Request,len)
		go h.startWork(h.queue[i])
	}
}
// Execute 
func (h *Handler) Execute(r *Request) {
	h.engine.Start(r)
}
// AsyncExecute
func (h *Handler) AsyncExecute(r *Request) {
	h.queue[r.Conn().GetID()%uint64(h.size)] <- r
}
// startWork
func (h *Handler) startWork(queue chan *Request) {
	for {
		select {
		case r := <-queue:
			h.Execute(r)
		}
	}
}