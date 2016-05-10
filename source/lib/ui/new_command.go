package ui

import (
	"fmt"
	"os"

	"github.com/c2fo/gull/source/lib/common"
	"github.com/c2fo/gull/source/lib/gull"
	"github.com/codegangsta/cli"
)

type NewCommand struct {
	Destination string
	Name        string
}

func (nc *NewCommand) GetFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   "destination, d",
			Value:  common.DefaultGullDirectory,
			Usage:  "directory that will contain the converted configuration migration(s)",
			EnvVar: "GULL_DESTINATION",
		},
		cli.StringFlag{
			Name:   "name, n",
			Usage:  "What to call the migration. This will be prepended with a migration ID. Spaces are replaced with dashes.",
			EnvVar: "GULL_MIGRATION_NAME",
		},
	}
}

func (nc *NewCommand) GetCliCommand() cli.Command {
	return cli.Command{
		Name:  "new",
		Usage: "Create an empty migration file",
		Flags: nc.GetFlags(),
		Action: func(c *cli.Context) error {
			nc.ParseOptions(c)
			nc.New()
			return nil
		},
	}
}

func (nc *NewCommand) ParseOptions(context *cli.Context) {
	nc.Destination = context.String("destination")
	if nc.Destination == "" || nc.Destination == common.DefaultGullDirectory {
		fmt.Printf("No destination was specified. Defaulting to [%s]\n", nc.Destination)
	}
	nc.Name = context.String("name")
	if nc.Name == "" {
		fmt.Println("A name for the migration is required")
		os.Exit(1)
	}
}

func (nc *NewCommand) New() {
	migrationPath, err := gull.CreateNewMigrationFile(nc.Name, nc.Destination)
	if err != nil {
		fmt.Printf("An error occurred while creating the migration file: [%+v]\n", err)
		os.Exit(1)
	}
	fmt.Printf("The migration file was written to [%v]\n", migrationPath)
}
