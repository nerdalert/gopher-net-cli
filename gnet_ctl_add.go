package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"

	log "github.com/nerdalert/gopher-net-ctl/Godeps/_workspace/src/github.com/Sirupsen/logrus"
	"github.com/nerdalert/gopher-net-ctl/Godeps/_workspace/src/github.com/codegangsta/cli"
)

var GnetCtlAdd = cli.Command{
	Name:  "add",
	Usage: "gnet-ctl add <option>",
	Subcommands: []cli.Command{
		{
			Usage:  "use 'gnet-ctl add help' for subcommand usage",
			Action: cli.ShowSubcommandHelp,
		},
		{
			Name: "neighbor",
			Usage: "gnet-ctl add neighbor  --neighbor-ip=<ip_address of neighbor> --neighbor-as=<AS_number> \n" +
				"\t(Example): 'gnet-ctl add neighbor --neighbor-ip=172.16.100.1 --neighbor-as=65001 --description=zone2-r1'",
			Flags: []cli.Flag{
				NeighborIpFlag,
				NeighborAsFlag,
				NeighborDescripFlag,
			},
			Action: AddNeighdor,
		},
	},
}

// TODO: temp struct
type Neighbor struct {
	NeighborIP  net.IP `json:"neighbor_ip"`
	NeighborAS  uint32 `json:"neighbor_as"`
	Description string `json:"description"`
}

func AddNeighdor(c *cli.Context) {
	client := NewClient()
	bgpNeighbor := new(Neighbor)
	rawIp := c.String("neighbor-ip")
	var neighborIp net.IP
	if rawIp != "" {
		neighborIp = net.ParseIP(rawIp)
	} else {
		log.Error("a peer IP address is required to add a bgp neighbor")
		return
	}
	var neighborAs uint32
	neighborAsStr := c.String("neighbor-as")
	if neighborAsStr != "" {
		neighborAs = StrToUin32(neighborAsStr)
	} else {
		log.Error("an AS number is required to add a bgp neighbor")
		return
	}
	bgpNeighbor.NeighborIP = neighborIp
	bgpNeighbor.NeighborAS = neighborAs
	description := c.String("description")
	if description != "" {
		bgpNeighbor.Description = description
	}
	j, err := json.Marshal(&bgpNeighbor)
	if err != nil {
		log.Println(err)
		return
	}
	req, err := http.NewRequest("POST", "http://127.0.0.1:8080"+NEIGHBOR+ADD, bytes.NewBuffer(j))
	if err != nil {
		log.Errorf("No answer from the bgp daemon: %s \n", err)
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
		log.Errorf("No answer from the bgp daemon: %s \n", err)
		return
	}
	// Read Response
	respBody, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return
	}
	if resp.Status == "404 Not Found" {
		log.Error("requested data type not supported")
	}
	// Display Results
	log.Debugln("response Status : ", resp.Status)
	log.Debugln("response Headers : ", resp.Header)
	log.Debugln("response Body : ", string(respBody))
	log.Infoln("results: ", string(respBody))
	return
}

type RestRoute struct {
	IpPrefix    string `json:"ip_prefix"`
	PrefixMask  uint8  `json:"ip_mask"`
	NextHop     string `json:"ip_nexthop"`
	LocalPref   uint32 `json:"local_pref"`
	RouteFam    string `json:"route_family"`
	ExCommunity string `json:"opaque"`
}

func AddRoute(c *cli.Context) {
	client := NewClient()
	bgpNeighbor := new(Neighbor)
	strIp := c.String("local-pref")
	var neighborIp net.IP
	if strIp != "" {
		neighborIp = net.ParseIP(strIp)
	} else {
		log.Error("a peer IP address is required to add a bgp neighbor")
		return
	}
	var neighborAs uint32
	neighborAsStr := c.String("neighbor-as")
	if neighborAsStr != "" {
		neighborAs = StrToUin32(neighborAsStr)
	} else {
		log.Error("an AS number is required to add a bgp neighbor")
		return
	}
	bgpNeighbor.NeighborIP = neighborIp
	bgpNeighbor.NeighborAS = neighborAs
	description := c.String("description")
	if description != "" {
		bgpNeighbor.Description = description
	}
	j, err := json.Marshal(&bgpNeighbor)
	if err != nil {
		log.Println(err)
		return
	}
	req, err := http.NewRequest("POST", "http://127.0.0.1:8080"+ROUTE_TABLES+ADD, bytes.NewBuffer(j))
	if err != nil {
		log.Errorf("No answer from the bgp daemon: %s \n", err)
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
		log.Errorf("No answer from the bgp daemon: %s \n", err)
		return
	}
	// Read Response
	respBody, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return
	}
	if resp.Status == "404 Not Found" {
		log.Error("requested data type not supported")
	}
	// Display Results
	log.Debugln("response Status : ", resp.Status)
	log.Debugln("response Headers : ", resp.Header)
	log.Debugln("response Body : ", string(respBody))
	log.Infoln("results: ", string(respBody))
	return
}
