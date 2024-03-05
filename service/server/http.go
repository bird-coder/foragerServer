package server

import (
	"context"
	"foragerServer/constants"
	"foragerServer/options"
	"foragerServer/routes"
	zlog "foragerServer/service/logger"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	*http.Server
	ctx context.Context
}

func NewHttp(ctx context.Context, c *options.AppConfig) *HttpServer {
	SetGinEnv(c.Env)
	r := gin.Default()
	routes.HandleRoutes(r)
	s := &http.Server{
		Addr:           c.Server.Http.Addr,
		Handler:        r,
		ReadTimeout:    time.Duration(c.Server.Http.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(c.Server.Http.WriteTimeout) * time.Second,
		MaxHeaderBytes: c.Server.Http.MaxHeaderBytes,
	}
	server := &HttpServer{
		s,
		ctx,
	}
	return server
}

func (server *HttpServer) Run() error {
	if err := server.ListenAndServe(); err != nil {
		zlog.Error("http server exit: [%+v]", err)
		if err == http.ErrServerClosed {
			zlog.Info("waiting for shutdown finish...")
			<-server.ctx.Done()
			zlog.Info("http server shutdown finished")
			return err
		}
		return err
	}
	return nil
}

func (server *HttpServer) Close() error {
	zlog.Info("http server shutdown...")
	if err := server.Shutdown(server.ctx); err != nil {
		zlog.Error("http server shutdown error(%v)", err)
		return err
	}
	zlog.Info("http server shutdown processed successfully")
	return nil
}

func SetGinEnv(env constants.Env) {
	if env == constants.PRO {
		gin.SetMode(gin.ReleaseMode)
	}
}
