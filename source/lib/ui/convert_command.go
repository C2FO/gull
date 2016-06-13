package ui

import (
	"fmt"
	"os"

	"github.com/c2fo/gull/source/lib/common"
	"github.com/c2fo/gull/source/lib/gull"
	"github.com/codegangsta/cli"
)

type ConvertCommand struct {
	Source                string
	Destination           string
	Verbose               bool
	FileNameIsEnvironment bool
	JsonEncode            bool
	Logger                gull.ILogger
}

func (cc *ConvertCommand) GetFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   "source, i",
			Usage:  "configuration file or directory to convert",
			EnvVar: "GULL_SOURCE",
		},
		cli.StringFlag{
			Name:   "destination, o",
			Value:  common.DefaultGullDirectory,
			Usage:  "directory that will contain the converted configuration migration(s)",
			EnvVar: "GULL_DESTINATION",
		},
		cli.BoolFlag{
			Name:   "verbose, v",
			Usage:  "display extra logging when running the command",
			EnvVar: "GULL_VERBOSE",
		},
		cli.BoolFlag{
			Name:   "filenameisenvironment, f",
			Usage:  "when config files don't contain an environment at the top level, use the file name as the environment",
			EnvVar: "GULL_FILE_NAME_IS_ENVIRONMENT",
		},
		cli.BoolFlag{
			Name:   "jsonencode, j",
			Usage:  "if enabled, store converted config values as JSON encoded strings instead of golang formatted string values",
			EnvVar: "GULL_JSON_ENCODE",
		},
	}
}

func (cc *ConvertCommand) GetCliCommand() cli.Command {
	return cli.Command{
		Name:  "convert",
		Usage: "Convert application configuration into gull migrations",
		Flags: cc.GetFlags(),
		Action: func(c *cli.Context) error {
			cc.ParseOptions(c)
			cc.Convert()
			return nil
		},
	}
}

func (cc *ConvertCommand) ParseOptions(context *cli.Context) {
	cc.Verbose = context.Bool("verbose")
	cc.Logger = gull.NewLogger(cc.Verbose)
	cc.FileNameIsEnvironment = context.Bool("filenameisenvironment")
	cc.JsonEncode = context.Bool("jsonencode")

	cc.Source = context.String("source")
	if cc.Source == "" {
		fmt.Printf("No source was specified, but it is required. Run `gull convert -h` for more information.\n")
		os.Exit(1)
	}

	cc.Destination = context.String("destination")
	if cc.Destination == "" || cc.Destination == common.DefaultGullDirectory {
		if cc.Verbose {
			fmt.Printf("No destination was specified. Defaulting to [%s]\n", cc.Destination)
		}
	}
}

func (cc *ConvertCommand) Convert() {
	convert, err := gull.NewConvert(cc.Destination, cc.FileNameIsEnvironment, cc.JsonEncode, cc.Logger)
	if err != nil {
		if cc.Verbose {
			fmt.Printf("An error occurred while instantiating a converter: [%+v]\n", err)
		}
		os.Exit(1)
	}
	info, err := os.Stat(cc.Source)
	if err != nil {
		if cc.Verbose {
			fmt.Printf("An error occurred while attempting to open [%v]: [%+v]\n", cc.Destination, err)
		}
		os.Exit(1)
	}
	if info.IsDir() {
		err := convert.ConvertDirectory(cc.Source)
		if err != nil {
			fmt.Printf("An error occurred while converting the directory [%v]: [%+v]\n", cc.Destination, err)
			panic(err)
		}

	} else {
		err := convert.ConvertFile(cc.Source)
		if err != nil {
			fmt.Printf("An error occurred while converting the file [%v]: [%+v]\n", cc.Destination, err)
		}
	}
	os.Exit(1)
}
