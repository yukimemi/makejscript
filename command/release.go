package command

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

const ReleaseHeader string = "@set @junk=1 /*\r\n@cscript //nologo //e:jscript \"%~f0\" %*\r\n@exit /b %errorlevel%\r\n*/"

func CmdRelease(c *cli.Context) {
	var err error

	err = commandExecute(c, ReleaseHeader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occured !")
	}
	fmt.Println("Done !")
}
