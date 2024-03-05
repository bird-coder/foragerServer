package service

import (
	"context"
	"foragerServer/model"
	"foragerServer/options"
	"foragerServer/service/dao"
	zlog "foragerServer/service/logger"
	"foragerServer/service/server"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bird-coder/manyo/rungroup"
)

type App struct {
	Dao    *dao.Dao
	Config *options.AppConfig
}

func (app *App) Init() (err error) {
	var (
		ctxHttp, cancelHttp = context.WithCancel(context.Background())
		httpServer          server.Server
	)
	httpServer = server.NewHttp(ctxHttp, app.Config)

	var g rungroup.Group
	{
		term := make(chan os.Signal, 1)
		signal.Notify(term, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
		closeChan := make(chan struct{})
		g.Add(
			func() error {
				select {
				case s := <-term:
					zlog.Info("get a signal %s", s.String())
					zlog.Info("forager server exit")
					time.Sleep(time.Second)
				case <-closeChan:
				}
				return nil
			},
			func(err error) {
				close(closeChan)
			},
		)
	}
	{
		// init dao
		cancel := make(chan struct{})
		g.Add(
			func() error {
				d := dao.NewDao(app.Config.Dao)
				model.Init(d)

				<-cancel

				return nil
			},
			func(err error) {
				close(cancel)
			},
		)
	}
	{
		// init server
		g.Add(
			func() error {
				err := httpServer.Run()
				zlog.Info("http server stopped...")
				return err
			},
			func(err error) {
				httpServer.Close()
				cancelHttp()
			},
		)
	}
	if err = g.Run(); err != nil {
		zlog.Error("app init error: %v", err)
	}
	return
}

func (app *App) LoadAppConfig() (err error) {
	app.Config = &options.AppConfig{}
	app.Config.Server = &options.ServerConfig{}
	app.Config.Dao = &options.DaoConfig{}
	if err = app.Config.LoadConfig(); err != nil {
		return
	}
	if err = app.Config.Server.LoadConfig(); err != nil {
		return
	}
	if err = app.Config.Dao.LoadConfig(); err != nil {
		return
	}
	return nil
}
