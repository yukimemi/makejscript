package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/yukimemi/makejscript/command"
)

func main() {

	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = "yukimemi"
	app.Email = "yukimemi@gmail.com"
	app.Usage = ""

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	app.Action = command.CmdRelease

	app.Run(os.Args)
}
