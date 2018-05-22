package main

import (
	"os"

	"github.com/urfave/cli"
)

var version = "master"

func main() {
	app := cli.NewApp()
	app.Usage = "CLI tool for ignite-agent"
	app.Version = version
	app.Commands = commands()
	app.Run(os.Args)
}

func commands() []cli.Command {
	return []cli.Command{
		{
			Name:    "init",
			Usage:   "init services",
			Action:  initImages,
			Aliases: []string{"i"},
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name: "verbose, v", Usage: "Show verbose log",
				},
			},
		},
		{
			Name:    "ls",
			Usage:   "list services",
			Action:  list,
			Aliases: []string{"l"},
		},
	}
}
