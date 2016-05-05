package ui

import (
	"fmt"
	"os"

	"github.com/c2fo/gull/source/lib/gull"
	"github.com/codegangsta/cli"
)

type DownCommand struct {
	Environment     string
	EtcdHostUrl     string
	SourceDirectory string
	Verbose         bool
	DryRun          bool
}

func (dc *DownCommand) GetFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   "environment, e",
			Usage:  "system to target for configuration migration",
			EnvVar: "GULL_ENVIRONMENT",
			Value:  "default",
		},
		cli.StringFlag{
			Name:   "etcdhost, c",
			Usage:  "url to the system running etcd in the format 'http://localhost:4002/v2/keys'",
			EnvVar: "GULL_ETCD_HOST",
		},
		cli.BoolFlag{
			Name:   "verbose, v",
			Usage:  "display extra logging when running the command",
			EnvVar: "GULL_VERBOSE",
		},
		cli.BoolFlag{
			Name:   "dryrun, d",
			Usage:  "show the expected configuration to be deployed",
			EnvVar: "GULL_DRY_RUN",
		},
	}
}

func (dc *DownCommand) GetCliCommand() cli.Command {
	return cli.Command{
		Name:  "down",
		Usage: "Migrate to a previous configuration",
		Flags: dc.GetFlags(),
		Action: func(c *cli.Context) {
			dc.ParseOptions(c)
			dc.Down()
		},
	}
}

func (dc *DownCommand) ParseOptions(context *cli.Context) {
	dc.Verbose = context.Bool("verbose")
	dc.DryRun = context.Bool("dryrun")

	dc.Environment = context.String("environment")
	if dc.Environment == "default" {
		if dc.Verbose {
			fmt.Printf("environment was not found, migrating 'default'\n")
		}
	}

	dc.EtcdHostUrl = context.String("etcdhost")
	if dc.EtcdHostUrl == "" {
		fmt.Println("No etcdhost was provided, but it is required")
		os.Exit(1)
	}
}

func (dc *DownCommand) Down() {
	// A 'Down' will remove all the current configuration for an environment.
	// Then all migrations that were stored in etcd for that environment are applied, ignoring the latest migration.
	var target gull.MigrationTarget
	if dc.DryRun {
		target = gull.NewMockMigrationTarget(dc.Environment)
	} else {
		target = gull.NewEtcdMigrationTarget(dc.EtcdHostUrl, dc.Environment)
	}
	down := gull.NewDown(target)
	err := down.Migrate()
	if err != nil {
		fmt.Printf("An error occurred while performing the migration: [%+v]\n", err)
		os.Exit(1)
	}
	if dc.DryRun {
		target.Debug()
	}
}