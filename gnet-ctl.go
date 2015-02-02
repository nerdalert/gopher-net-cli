package main

import (
	"os"

	log "github.com/nerdalert/gopher-net-ctl/Godeps/_workspace/src/github.com/Sirupsen/logrus"
	"github.com/nerdalert/gopher-net-ctl/Godeps/_workspace/src/github.com/codegangsta/cli"
)

func init() {
	app.EnableBashCompletion = true
	log.SetLevel(log.DebugLevel)
}

var (
	app *cli.App = cli.NewApp()
)

func main() {
	app.Name = "gnet-ctl"
	app.Usage = "command line utility for viewing and manipulating Gopher Net route peerings, " +
		"state and configuration. All commands are functions are also available via the daemon REST APIs."
	app.Version = "0.1"
	app.Commands = []cli.Command{
		GnetCtlShow,
		GnetCtlAdd,
		GnetCtlDelete,
	}
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, d",
			Usage: "set gnet-ctl logging to debug.",
		},
	}
	app.Before = cliInit
	app.Run(os.Args)
}

func cliInit(c *cli.Context) error {
	if c.GlobalBool("debug") {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
	app.EnableBashCompletion = true
	log.SetOutput(os.Stderr)
	return nil
}
