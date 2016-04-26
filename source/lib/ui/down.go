package ui

import (
	"fmt"

	"github.com/codegangsta/cli"
)

func GetDownCli() cli.Command {
	return cli.Command{
		Name:    "down",
		Aliases: []string{"d"},
		Usage:   "Migrate to a previous configuration",
		Flags:   subcommandFlags,
		Action: func(c *cli.Context) {
			ParseOptions(c).Down()
		},
	}
}

func (o *Options) Down() {
	fmt.Printf("TODO: Implement the Down subcommand\n")
}
