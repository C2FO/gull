package ui

import (
	"fmt"
	"os"

	"github.com/c2fo/gull/source/lib/common"
	"github.com/c2fo/gull/source/lib/gull"
	"github.com/codegangsta/cli"
)

type DestroyCommand struct {
	Application string
	Environment string
	EtcdHostUrl string
	Logger      gull.ILogger
}

func (dc *DestroyCommand) GetFlags() []cli.Flag {
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

func (dc *DestroyCommand) GetCliCommand() cli.Command {
	return cli.Command{
		Name:  "destroy",
		Usage: "Remove an environment from etcd",
		Flags: dc.GetFlags(),
		Action: func(c *cli.Context) error {
			dc.ParseOptions(c)
			dc.Destroy()
			return nil
		},
	}
}

func (dc *DestroyCommand) ParseOptions(context *cli.Context) {
	dc.Logger = gull.NewVerboseLogger()
	dc.Application = context.String("application")
	if dc.Application == "" {
		dc.Logger.Info("application was not found, but is required.")
		os.Exit(1)
	}

	dc.Environment = context.String("environment")
	if dc.Environment == "" {
		dc.Logger.Info("No target environment was provided, using 'default'")
		dc.Environment = "default"
	}

	dc.EtcdHostUrl = context.String("etcdhost")
	if dc.EtcdHostUrl == common.DefaultEtcdServerUrl {
		fmt.Printf("No etcdhost was provided, using [%v]\n", common.DefaultEtcdServerUrl)
	}
}

func (dc *DestroyCommand) Destroy() {
	logger := gull.NewVerboseLogger()
	logger.Info("Destroying environment [%s] on etcd host [%s]", dc.Environment, dc.EtcdHostUrl)
	target := gull.NewEtcdMigrationTarget(dc.EtcdHostUrl, dc.Application, dc.Environment, false, logger)
	destroy := gull.NewDestroy(target)
	err := destroy.Execute()
	if err != nil {
		dc.Logger.Info("An error occurred while deleting the environment: [%+v]\n", err)
		os.Exit(1)
	}
}
