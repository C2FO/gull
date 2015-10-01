package ui

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
)

type Options struct {
	Environment string
	EtcdHost    string
}

var defaultEtcdHost = "http://localhost:4001"

func ParseOptions(context *cli.Context) *Options {
	options := &Options{
		Environment: context.String("environment"),
		EtcdHost:    context.String("etcdHost"),
	}
	if options.Environment == "" {
		log.Printf("No environment was specified, but it is required. Run `gull -h` for more information.")
		os.Exit(1)
	}
	if options.EtcdHost == "" {
		log.Printf("No etcdHost was specified. Defaulting to [%s]", defaultEtcdHost)
		options.EtcdHost = defaultEtcdHost
	}
	return options
}
