package ui

import (
	"os"

	"github.com/codegangsta/cli"
)

var defaultEtcdHost = "http://localhost:4002"

type GullCommand interface {
	GetCliCommand() cli.Command
	ParseOptions(context *cli.Context)
	GetFlags() []cli.Flag
}

var subcommandFlags = []cli.Flag{
	cli.StringFlag{
		Name:   "environment, env, e",
		Value:  "*",
		Usage:  "location in etcd where the configuration will be written",
		EnvVar: "GULL_ENVIRONMENT,GULL_ENV",
	},
	cli.StringFlag{
		Name:   "etcdHost, etcd, d",
		Value:  "http://localhost:4001",
		Usage:  "URL to an etcd server",
		EnvVar: "GULL_ETCD_HOST,GULL_ETCD",
	},
}

func Launch() {
	app := cli.NewApp()
	app.Name = "gull"
	app.Version = "0.0.1"
	app.Usage = "etcd configuration migration management system"
	app.Commands = []cli.Command{
		new(ConvertCommand).GetCliCommand(),
	}

	app.Run(os.Args)
}
