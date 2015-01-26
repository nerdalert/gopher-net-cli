package main

import (
	"os"
	"text/tabwriter"

	log "github.com/nerdalert/gopher-net-ctl/Godeps/_workspace/src/github.com/Sirupsen/logrus"
	"github.com/nerdalert/gopher-net-ctl/Godeps/_workspace/src/github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "gnet-ctl"
	app.Usage = "gopher net command line tool"
	app.Version = "0.1"
	app.Commands = []cli.Command{
		{
			Name:  "show",
			Usage: "show and modify bgp neighbor states and configurations",
			Subcommands: []cli.Command{
				{
					Name:   "",
					Usage:  "gnet-ctl show { neighbors-conf | routes | neighbors | global-conf | rib-out | rib-in }",
					Action: ShowNeighborsConfigs,
				},
				{
					Name:   "neighbors-conf",
					Usage:  "show neighbor configuration and current state",
					Action: ShowNeighborsConfigs,
				},
				{
					Name:   "routes",
					Usage:  "show best incoming destinations",
					Action: GetRoutes,
				},
				{
					Name:   "neighbors",
					Usage:  "show neighbor configurations",
					Action: ShowNeighbors,
				},
				{
					Name:   "global-conf",
					Usage:  "show the local BGP global configuration",
					Action: ShowGlobalConfig,
				},
				{
					Name:   "rib-out",
					Usage:  "show best path bgp routes",
					Action: ShowRibOut,
				},
				{
					Name:   "rib-in",
					Usage:  "show bgp rib-in",
					Action: ShowRibIn,
				},
			},
		},
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "remote",
			Value: "",
			Usage: "IP address of the remote neighbor",
		},
	}
	app.Action = func(c *cli.Context) {
		switch true {
		case c.Bool("debug"):
			// todo
		default:
			cli.ShowCommandHelp(c, "")
		}
	}
	app.Run(os.Args)
}

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
