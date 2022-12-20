package config

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type httpConf struct {
	Host string
	Port int
}

var wg sync.WaitGroup

func (a *app) initHttp(handler *echo.Echo) {
	srv := &http.Server{
		Addr:    fmt.Sprintf("%v:%v", a.HttpConf.Host, a.HttpConf.Port),
		Handler: handler,

		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	shutdownError := make(chan error)
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		s := <-sig
		Logger.Info("process", zap.String("type", "signal"), zap.String("source", "signal"), zap.String("act", s.String()), zap.String("status", "delegated"))
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := srv.Shutdown(ctx)
		if err != nil {
			shutdownError <- err
		}
		Logger.Info("process", zap.String("type", "background"), zap.String("source", "httprouter"), zap.String("act", "closing"), zap.String("addr", srv.Addr), zap.String("status", "done"))

		wg.Wait()
		shutdownError <- nil
	}()

	Logger.Info("process", zap.String("type", "server"), zap.String("source", "httprouter"), zap.String("act", "serve"), zap.String("addr", srv.Addr), zap.String("status", "running"))
	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		Logger.Fatal(err.Error())
	}

	err = <-shutdownError
	if err != nil {
		Logger.Fatal(err.Error())
	}

	Logger.Info("process", zap.String("type", "server"), zap.String("source", "httprouter"), zap.String("act", "shutdown"), zap.String("addr", srv.Addr), zap.String("status", "done"))
}
