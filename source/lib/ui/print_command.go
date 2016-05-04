package ui

import (
	"fmt"

	"github.com/codegangsta/cli"
)

func GetPrintCli() cli.Command {
	return cli.Command{
		Name:    "print",
		Aliases: []string{"p"},
		Usage:   "Display the entire configuration contents of one environment",
		Flags:   subcommandFlags,
		Action: func(c *cli.Context) {
			ParseOptions(c).Print()
		},
	}
}

func (o *Options) Print() {
	fmt.Printf("TODO: Implement the Print subcommand\n")
}
