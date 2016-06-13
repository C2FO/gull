package ui

import (
	"fmt"
	"os"

	"github.com/c2fo/gull/source/lib/common"
	"github.com/c2fo/gull/source/lib/gull"
	"github.com/codegangsta/cli"
)

type StatusCommand struct {
	Application string
	Environment string
	EtcdHostUrl string
	Logger      gull.ILogger
}

func (sc *StatusCommand) GetFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   "application, a",
			Usage:  "application to target for configuration migration",
			EnvVar: "GULL_APPLICATION",
		},
		cli.StringFlag{
			Name:   "environment, e",
			Usage:  "system to target for configuration migration",
			EnvVar: "GULL_ENVIRONMENT",
		},
		cli.StringFlag{
			Name:   "etcdhost, s",
			Usage:  "url to the system running etcd",
			EnvVar: "GULL_ETCD_HOST",
			Value:  common.DefaultEtcdServerUrl,
		},
	}
}

func (sc *StatusCommand) GetCliCommand() cli.Command {
	return cli.Command{
		Name:  "status",
		Usage: "Display what gull knows for a single environment",
		Flags: sc.GetFlags(),
		Action: func(c *cli.Context) error {
			sc.ParseOptions(c)
			sc.Status()
			return nil
		},
	}
}

func (sc *StatusCommand) ParseOptions(context *cli.Context) {
	sc.Logger = gull.NewVerboseLogger()
	sc.Application = context.String("application")
	if sc.Application == "" {
		sc.Logger.Info("application was not found, but is required.")
		os.Exit(1)
	}

	sc.Environment = context.String("environment")
	if sc.Environment == "" {
		sc.Logger.Info("No target environment was provided, using 'default'")
		sc.Environment = "default"
	}

	sc.EtcdHostUrl = context.String("etcdhost")
	if sc.EtcdHostUrl == common.DefaultEtcdServerUrl {
		fmt.Printf("No etcdhost was provided, using [%v]\n", common.DefaultEtcdServerUrl)
	}
}

func (sc *StatusCommand) Status() {
	logger := gull.NewVerboseLogger()
	logger.Info("Checking migration status of environment [%s] on etcd host [%s]", sc.Environment, sc.EtcdHostUrl)
	target := gull.NewEtcdMigrationTarget(sc.EtcdHostUrl, sc.Application, sc.Environment, false, logger)
	status := gull.NewStatus(target)
	err := status.Check()
	if err != nil {
		sc.Logger.Info("An error occurred while checking status: [%+v]\n", err)
		os.Exit(1)
	}
}
