package ui

import (
	"fmt"
	"os"

	"github.com/c2fo/gull/source/lib/common"
	"github.com/c2fo/gull/source/lib/gull"
	"github.com/codegangsta/cli"
)

type UpCommand struct {
	Application     string
	Environment     string
	EtcdHostUrl     string
	SourceDirectory string
	Verbose         bool
	DryRun          bool
	Full            bool
}

func (uc *UpCommand) GetFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   "application, a",
			Usage:  "application to target for configuration migration",
			EnvVar: "GULL_APPLICATION",
		},
		cli.StringFlag{
			Name:   "environment, e",
			Usage:  "environment to target for configuration migration",
			EnvVar: "GULL_ENVIRONMENT",
			Value:  "default",
		},
		cli.StringFlag{
			Name:   "etcdhost, c",
			Usage:  "url to the system running etcd in the format 'http://localhost:4002/v2/keys'",
			EnvVar: "GULL_ETCD_HOST",
		},
		cli.StringFlag{
			Name:   "source, s",
			Usage:  "directory containing gull migrations",
			EnvVar: "GULL_SOURCE",
			Value:  common.DefaultGullDirectory,
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
		cli.BoolFlag{
			Name:   "full",
			Usage:  "wipe an environment's configuration and reapply all migrations",
			EnvVar: "GULL_FULL",
		},
	}
}

func (uc *UpCommand) GetCliCommand() cli.Command {
	return cli.Command{
		Name:  "up",
		Usage: "Migrate to the latest configuration",
		Flags: uc.GetFlags(),
		Action: func(c *cli.Context) {
			uc.ParseOptions(c)
			uc.Up()
		},
	}
}

func (uc *UpCommand) ParseOptions(context *cli.Context) {
	uc.Verbose = context.Bool("verbose")
	uc.DryRun = context.Bool("dryrun")

	uc.SourceDirectory = context.String("source")

	uc.Application = context.String("application")
	if uc.Application == "" {
		fmt.Println("application was not found, but is required.")
		os.Exit(1)
	}

	uc.Environment = context.String("environment")
	if uc.Environment == "default" {
		if uc.Verbose {
			fmt.Printf("environment was not found, migrating 'default'\n")
		}
	}

	uc.EtcdHostUrl = context.String("etcdhost")
	if uc.EtcdHostUrl == "" {
		fmt.Println("No etcdhost was provided, but it is required")
		os.Exit(1)
	}

	uc.Full = context.Bool("full")
}

func (uc *UpCommand) Up() {
	// An 'Up' will walk through all migrations and apply 'default' to the target environment.
	// All migrations will then be walked again, applying any migrations containing the target environment.
	var target gull.MigrationTarget
	if uc.DryRun {
		target = gull.NewMockMigrationTarget(uc.Application, uc.Environment)
	} else {
		target = gull.NewEtcdMigrationTarget(uc.EtcdHostUrl, uc.Application, uc.Environment, uc.Full)
	}
	up := gull.NewUp(uc.SourceDirectory, target)
	err := up.Migrate()
	if err != nil {
		fmt.Printf("An error occurred while performing the migration: [%+v]\n", err)
		os.Exit(1)
	}
	if uc.DryRun {
		target.Debug()
	}
}
