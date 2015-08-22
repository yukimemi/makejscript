package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/yukimemi/makejscript/command"
)

// GlobalFlags is global flags.
var GlobalFlags = []cli.Flag{
	cli.StringFlag{
		// EnvVar: "ENV_I",
		Name:  "output, o",
		Value: "cmd",
		Usage: "Output cmd file path",
	},
}

// Commands is subcommand list.
var Commands = []cli.Command{
	{
		Name:   "release",
		Usage:  "Release build",
		Action: command.CmdRelease,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "debug",
		Usage:  "Debug build",
		Action: command.CmdDebug,
		Flags:  []cli.Flag{},
	},
}

// CommandNotFound print error when using undefined subcommand.
func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
