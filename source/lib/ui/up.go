package ui

import (
	"fmt"

	"github.com/codegangsta/cli"
)

func GetUpCli() cli.Command {
	return cli.Command{
		Name:    "up",
		Aliases: []string{"u"},
		Usage:   "Migrate to the latest configuration",
		Flags:   subcommandFlags,
		Action: func(c *cli.Context) {
			ParseOptions(c).Up()
		},
	}
}

func (o *Options) Up() {
	fmt.Printf("TODO: Implement Up subcommand\n")
	// An 'Up' will walk through all migrations and apply 'default' to the target environment.
	// All migrations will then be walked again, applying any migrations containing the target environment.
}
