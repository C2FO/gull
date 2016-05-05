package ui

import (
	"fmt"
	"os"

	"github.com/c2fo/gull/source/lib/gull"
	"github.com/codegangsta/cli"
)

type StatusCommand struct {
	Environment string
	EtcdHostUrl string
}

func (sc *StatusCommand) GetFlags() []cli.Flag {
	return []cli.Flag{
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
		Action: func(c *cli.Context) {
			sc.ParseOptions(c)
			sc.Status()
		},
	}
}

func (sc *StatusCommand) ParseOptions(context *cli.Context) {
	sc.Environment = context.String("environment")
	if sc.Environment == "default" {
		fmt.Printf("No target environment was provided, using 'default'\n")
	}

	sc.EtcdHostUrl = context.String("etcdhost")
	if sc.EtcdHostUrl == "" {
		fmt.Println("No etcdhost was not provided, but it is required")
		os.Exit(1)
	}
}

func (sc *StatusCommand) Status() {
	fmt.Printf("Checking migration status of environment [%s] on etcd host [%s]\n", sc.Environment, sc.EtcdHostUrl)
	target := gull.NewEtcdMigrationTarget(sc.EtcdHostUrl, sc.Environment)
	status := gull.NewStatus(target)
	err := status.Check()
	if err != nil {
		fmt.Printf("An error occurred while checking status: [%+v]\n", err)
		os.Exit(1)
	}
}