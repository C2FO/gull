package ui

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

type Options struct {
	Environment string
	EtcdHost    string
}

func ParseOptions(context *cli.Context) *Options {
	options := &Options{
		Environment: context.String("environment"),
		EtcdHost:    context.String("etcdHost"),
	}
	if options.Environment == "" {
		fmt.Printf("No environment was specified, but it is required. Run `gull -h` for more information.\n")
		os.Exit(1)
	}
	if options.EtcdHost == "" {
		fmt.Printf("No etcdHost was specified. Defaulting to [%s]\n", defaultEtcdHost)
		options.EtcdHost = defaultEtcdHost
	}
	return options
}
