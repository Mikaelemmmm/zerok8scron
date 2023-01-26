package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/conf"
	"os"
	"zerok8scron/internal/config"
	"zerok8scron/internal/logic"
	"zerok8scron/internal/svc"
)

const (
	codeFailure = 1
)

var (
	confPath string

	rootCmd = &cobra.Command{
		Use:   "cron",
		Short: "exec cron job",
		Long:  "exec cron job",
	}

	// all job ...
	helloJob = &cobra.Command{
		Use:   "hello",
		Short: "print 'hello SvcName' once per minute",
		RunE:  logic.Hello,
	}

	// add more job
)

// Execute executes the given command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(codeFailure)
	}
}

func init() {

	// init config
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&confPath, "config", "etc/cron.yaml", "config file (default is $HOME/.cobra.yaml)")

	// add subcommand
	rootCmd.AddCommand(helloJob)
}

func initConfig() {
	var c config.Config
	conf.MustLoad(confPath, &c)
	svc.InitSvcCtx(c)
}
