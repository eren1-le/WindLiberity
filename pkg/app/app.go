/*
 * @Author: eren dengdengd1222@mail.com
 * @Date: 2026-09-23 14:02:36
 * @LastEditors: eren dengdengd1222@mail.com
 * @LastEditTime: 2026-09-24 10:50:01
 * @FilePath: /WindLiberity/pkg/app/app.go
 * @Description:
 *
 */
package app

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/xid"
	"golang.org/x/sync/errgroup"
)

type App struct {
	opts 	options
	ctx 	context.Context
	cancel  func()
}

// New create app globally
func New(opts ...Option) *App {
	o := options{id: xid.New().String()}
	for _, opt := range opts {
		opt(&o)
	}

	ctx, cancel := context.WithCancel(context.Background())
	return &App{
		opts: 	o,
		ctx: 	ctx,
		cancel: cancel,
	}
}

func (a *App) Run() error {
	eg, ctx := errgroup.WithContext(a.ctx)

	for _, srv := range a.opts.servers {
		srv := srv
		eg.Go(func() error {
			<-ctx.Done()
			return srv.Stop(ctx)
		})
		eg.Go(func() error {
			return srv.Start(ctx)
		})
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	eg.Go(func() error {
		defer log.Println(("singal defer"))
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case s := <-quit:
					log.Printf("Server receive a quit singal: %s", s.String())
					if err := a.Stop(); err != nil {
						log.Print("failed to stop app, err: %v", err)
						return err;
					}
			}
		}
	})

	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}


//	Stop  stops the application gracefully
func (a *App) Stop() error {
	if a.cancel != nil {
		a.cancel()
	}
	return nil
}