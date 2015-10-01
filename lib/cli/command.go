package cli

import (
	"flag"
)

// shamelessly snagged from liamstask/goose
// each command gets its own set of args,
// defines its own entry point, and provides its own help
type Command struct {
	Run  func(cmd *Command, args ...string)
	Flag flag.FlagSet

	Name  string
	Usage string

	Summary string
	Help    string
}

func (c *Command) Exec(args []string) {
	c.Flag.Usage = func() {}
	c.Flag.Parse(args)
	c.Run(c, c.Flag.Args()...)
}
