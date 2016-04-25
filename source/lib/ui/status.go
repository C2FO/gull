package ui

import (
	"fmt"

	"github.com/codegangsta/cli"
)

func GetStatusCli() cli.Command {
	return cli.Command{
		Name:    "status",
		Aliases: []string{"s"},
		Usage:   "Display what gull knows about for one environment",
		Flags:   subcommandFlags,
		Action: func(c *cli.Context) {
			ParseOptions(c).Status()
		},
	}
}

func (o *Options) Status() {
	fmt.Printf("The environment name is [%s] and the etcdHost is [%s]\n", o.Environment, o.EtcdHost)
	fmt.Printf("TODO: Implement the Status subcommand\n")
}
