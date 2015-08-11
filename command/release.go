package command

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

const ReleaseHeader string = `@set @junk=1 /*
@cscript //nologo //e:jscript "%~f0" %* & @ping -n 10 localhost > nul & @goto :eof
*/
`

func CmdRelease(c *cli.Context) {
	var err error

	err = commandExecute(c, ReleaseHeader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occured !")
	}
	fmt.Println("Done !")
}
