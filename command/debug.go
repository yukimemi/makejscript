package command

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

const DebugHeader string = `@set @junk=1 /*
@cscript //nologo //e:jscript "%~f0" %* & @pause & @goto :eof
*/
`

func CmdDebug(c *cli.Context) {
	var err error

	err = commandExecute(c, DebugHeader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occured !")
	}
	fmt.Println("Done !")
}
