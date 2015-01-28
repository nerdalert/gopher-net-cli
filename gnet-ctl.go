package main

import (
	"os"
	"text/tabwriter"

	log "github.com/nerdalert/gopher-net-ctl/Godeps/_workspace/src/github.com/Sirupsen/logrus"
	"github.com/nerdalert/gopher-net-ctl/Godeps/_workspace/src/github.com/codegangsta/cli"
)

func init() {
    out = new(tabwriter.Writer)
    out.Init(os.Stdout, 0, 8, 1, '\t', 0)
    app.EnableBashCompletion = true
    log.SetLevel(log.DebugLevel)
}

var (
    out *tabwriter.Writer
    app *cli.App = cli.NewApp()
)

func main() {
	app := cli.NewApp()
	app.Name = "gnet-ctl"
	app.Usage = "command line utility for viewing and manipulating Gopher Net route peerings, " +
		"state and configuration. All commands are functions are available via the REST APIs also."
	app.Version = "0.1"
	app.Commands = []cli.Command{
		GnetCtlShow,
		GnetCtlAdd,
		GnetCtlDelete,
	}
	app.Run(os.Args)
}
