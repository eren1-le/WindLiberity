package http

import (
	"time"

)

type Option func(*options)

type options struct {
	network 		string
	address 		string
	readTimeout 	time.Duration
	writeTimeout 	time.Duration
}


// WithAddress
func WithAddress(addr string) Option {
	return func(s *options) {
		s.address = addr
	}
}

// WithReadTimeout
func WithReadTimeout(timeout time.Duration) Option {
	return func(o *options) {
		o.readTimeout = timeout
	}
}

// WithWrtieTimeOut
func WithWrtieTimeOut(timeout time.Duration) Option {
	return func(o *options) {
		o.writeTimeout = timeout
	}
}
// NewOptions
func NewOptions(opt ...Option) options {
	opts := options{
		network:	"tcp",
		address:	":9050",
		readTimeout:  5 * time.Second,
		writeTimeout: 5 * time.Second,
	}

	for _, o := range opt {
		o(&opts)
	}

	return opts
}
