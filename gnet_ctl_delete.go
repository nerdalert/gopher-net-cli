package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net"
    "net/http"

    log "github.com/nerdalert/gopher-net-ctl/Godeps/_workspace/src/github.com/Sirupsen/logrus"
    "github.com/nerdalert/gopher-net-ctl/Godeps/_workspace/src/github.com/codegangsta/cli"
)

var GnetCtlDelete = cli.Command{
	Name:  "delete",
	Usage: "gnet-ctl delete <option>",
	Subcommands: []cli.Command{
		{
			Usage:  "use 'gnet-ctl delete help' for subcommand usage",
			Action: cli.ShowSubcommandHelp,
		},
		{
			Name:   "neighbor",
			Usage:  "'gnet-ctl delete neighbor <ip_address of neighbor>' delete a bgp neighbor",
			Action: DeleteNeighdor,
		},
	},
}

func DeleteNeighdor(c *cli.Context) {
    log.Debugln("get routes CLI called")
    client := NewClient()
    // TODO: replace w/ param
    ip := net.ParseIP("172.16.86.12")
    neighbor := Neighbor{NeighborIP: ip, PeerAs: 32, Description: "WTF"}

    j, err := json.Marshal(&neighbor)
    if err != nil {
        log.Println(err)
        return
    }
    req, err := http.NewRequest("POST", "http://127.0.0.1:8080"+NEIGHBOR+DEL, bytes.NewBuffer(j))
    if err != nil {
        fmt.Printf("No answer from the bgp daemon: %s \n", err)
        log.Error(error(err))
        return
    }
    parseFormErr := req.ParseForm()
    if parseFormErr != nil {
        log.Error(error(parseFormErr))
        return
    }
    // Get Request
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf("No answer from the bgp daemon: %s \n", err)
        return
    }
    // Read Response
    respBody, e := ioutil.ReadAll(resp.Body)
    if e != nil {
        return
    }
    if resp.Status == "404 Not Found" {
        log.Println("requested data type not supported")
    }
    // Display Results
    log.Debugln("response Status : ", resp.Status)
    log.Debugln("response Headers : ", resp.Header)
    log.Debugln("response Body : ", string(respBody))
    log.Infoln("results: ", string(respBody))
    return
}
