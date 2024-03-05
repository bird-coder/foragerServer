package api

import (
	"fmt"
	"foragerServer/service"
	zlog "foragerServer/service/logger"

	"github.com/spf13/cobra"
)

var (
	configYml string
	StartCmd  = &cobra.Command{
		Use:          "server",
		Short:        "Start API Server",
		Example:      "forager server -c config/settings.yml",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/settings.yml", "Start server with provided configuration file")
}

func setup() {
	fmt.Println("starting api server...")
}

func run() error {
	app := &service.App{}
	if err := app.LoadAppConfig(); err != nil {
		panic(err)
	}
	zlog.NewLogger(app.Config.Env)
	defer zlog.Sync()
	zlog.Info("forager server start")
	err := app.Init()
	return err
}
