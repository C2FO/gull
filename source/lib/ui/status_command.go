package ui

import (
	"os"

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
			Value:  "default",
		},
		cli.StringFlag{
			Name:   "etcdhost, c",
			Usage:  "url to the system running etcd in the format 'http://localhost:4002/v2/keys'",
			EnvVar: "GULL_ETCD_HOST",
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
	if sc.Environment == "default" {
		sc.Logger.Info("No target environment was provided, using 'default'")
	}

	sc.EtcdHostUrl = context.String("etcdhost")
	if sc.EtcdHostUrl == "" {
		sc.Logger.Info("No etcdhost was not provided, but it is required")
		os.Exit(1)
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
