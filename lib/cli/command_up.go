package cli

import "github.com/c2fo/gull/lib/gull"

var commandInfoUp = &Command{
	Name:    "up",
	Usage:   "",
	Summary: "Migrate ETCD to the most recent version available",
	Help:    ``,
	Run:     commandUp,
}

func commandUp(cmd *Command, args ...string) {
	gull.Migrate()
}
