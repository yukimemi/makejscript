package command

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

const DebugHeader string = "@set @junk=1 /*\r\n@cscript //nologo //e:jscript \"%~f0\" %*\r\n@ping -n 30 localhost > nul\r\n@goto :eof\r\n*/"

func CmdDebug(c *cli.Context) {
	var err error

	err = commandExecute(c, DebugHeader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occured !")
	}
	fmt.Println("Done !")
}
