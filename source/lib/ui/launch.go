package ui

import (
	"os"

	"github.com/codegangsta/cli"
)

func Launch() {
	app := cli.NewApp()
	app.Name = "gull"
	app.Version = "0.0.1"
	app.Usage = "etcd configuration migration management system"
	subcommandFlags := []cli.Flag{
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
	app.Commands = []cli.Command{
		{
			Name:    "up",
			Aliases: []string{"u"},
			Usage:   "Migrate to the latest configuration",
			Flags:   subcommandFlags,
			Action: func(c *cli.Context) {
				ParseOptions(c).Up()
			},
		},
		{
			Name:    "down",
			Aliases: []string{"d"},
			Usage:   "Migrate to a previous configuration",
			Flags:   subcommandFlags,
			Action: func(c *cli.Context) {
				ParseOptions(c).Down()
			},
		},
		{
			Name:    "print",
			Aliases: []string{"p"},
			Usage:   "Display the entire configuration contents of one environment",
			Flags:   subcommandFlags,
			Action: func(c *cli.Context) {
				ParseOptions(c).Print()
			},
		},
		{
			Name:    "status",
			Aliases: []string{"s"},
			Usage:   "Display what gull knows about for one environment",
			Flags:   subcommandFlags,
			Action: func(c *cli.Context) {
				ParseOptions(c).Status()
			},
		},
	}

	app.Run(os.Args)
}
