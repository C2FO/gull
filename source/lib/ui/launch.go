package ui

import (
	"fmt"
	"os"
	"strings"

	"github.com/codegangsta/cli"
)

type GullCommand interface {
	GetCliCommand() cli.Command
	ParseOptions(context *cli.Context)
	GetFlags() []cli.Flag
}

func Launch() {
	app := cli.NewApp()
	app.Name = "gull"
	app.Version = "0.11.0"
	app.Usage = "etcd configuration migration management system"
	app.Commands = []cli.Command{
		new(ConvertCommand).GetCliCommand(),
		new(DestroyCommand).GetCliCommand(),
		new(DownCommand).GetCliCommand(),
		new(NewCommand).GetCliCommand(),
		new(StatusCommand).GetCliCommand(),
		new(UpCommand).GetCliCommand(),
	}
	err := app.Run(os.Args)
	if err != nil {
		// If an invalid argument is passed, don't panic and muddy up the screen with a stacktrace
		if strings.Contains(err.Error(), "flag provided but not defined") {
			fmt.Println(err.Error())
		} else {
			panic(err)
		}
	}
}
