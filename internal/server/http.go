
package server

import (
	"WindLiberity/pkg/app"
	"WindLiberity/transport/http"
)
// NewHttpServer http server
func NewHttpServer(c *app.ServerConfig) *http.Server {
	srv := http.NewServer(
		http.WithAddress(c.Addr),
		http.WithReadTimeout(c.ReadTimeout),
		http.WithWrtieTimeOut(c.WriteTimeout),
	)
	return srv
}