package websocket

import (
	"math"
	"sync"
	"time"

)

const abortIndex int8 = math.MaxInt8 / 2

type HandlerFunc func(c *Context)

// HandlerChain
type HandlerChain []HandlerFunc

// RouteGroup
type RouteGroup struct {
	Handlers 	HandlerChain
	engine		*Engine
}

// Engine
type Engine struct {
	RouteGroup

	tree map[string]HandlerChain
	pool sync.Pool
}

// Context
type Context struct {
	Req		 *Request
	handlers HandlerChain
	index	 int8
}

// NewEngine get a new engine
func NewEngine() *Engine {
	engine := &Engine{
		RouteGroup: RouteGroup{
			Handlers: nil,
		},
		tree: make(map[string]HandlerChain),
	}
	engine.RouteGroup.engine = engine
	engine.pool.New = func() any {
		return engine.allocateContext()
	}

	return engine
}

// Next
func (c *Context) Next() {
	c.index++
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}
}

// Abort
func (c *Context) Abort() {
	c.index = abortIndex
}

// Reset
func (c *Context) Reset() {
	c.handlers = nil
	c.index = -1
}
// Done always returns nil (chan which will wait forever),
// if you want to abort your work when the connection was closed
// you should use Request.Context().Done() instead.
func (c *Context) Done() <-chan struct{} {
	return nil
}

// Err always returns nil, maybe you want to use Request.Context().Err() instead.
func (c *Context) Err() error {
	return nil
}

// Deadline always returns that there is no deadline (ok==false),
// maybe you want to use Request.Context().Deadline() instead.
func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return
}

// Value returns the value associated with this context for key, or nil
// if no value is associated with key. Successive calls to Value with
// the same key returns the same result.
func (c *Context) Value(any) any {
	return c.Req
}

// allocateContext
func (e *Engine) allocateContext() *Context {
	return &Context{}
}

// Start
func (e *Engine) Start(req *Request) {
	c := e.pool.Get().(*Context)
	c.Req = req
	c.Reset()
	handlers := e.getHandlers(c.Req)
	if handlers != nil {
		c.handlers = handlers
		c.Next()
	}
}
// Add Route
func (e *Engine) AddRoute(event string, handlers HandlerChain) {
	e.tree[event] = handlers
}


func (e *Engine) getHandlers(req *Request) HandlerChain {
	handlers, ok := e.tree[req.Event()]
	if !ok {
		return nil
	}
	return handlers
}

// Use set common middleware
func (e *Engine) Use(middleware ...HandlerFunc) {
	e.RouteGroup.Use(middleware...)
}
// load middleware
func (g *RouteGroup) Use(middleware ...HandlerFunc) {
	g.Handlers = append(g.Handlers, middleware...)
}

// AddRoute specific middleware
func (g *RouteGroup) AddRoute(event string, handlers ...HandlerFunc) {
	handlers = g.mergeHandlers(handlers)
	g.engine.AddRoute(event, handlers)
}

// merge specific and common middleware
func (g *RouteGroup) mergeHandlers(handlers HandlerChain) HandlerChain {
	finalSize := len(g.Handlers) + len(handlers)
	mergedHandlers := make(HandlerChain, finalSize)
	copy(mergedHandlers, g.Handlers)
	copy(mergedHandlers[len(g.Handlers):], handlers)
	return mergedHandlers
}
