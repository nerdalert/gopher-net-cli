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
			Usage: "Add a BGP Neighbor to the specified node to form a peering " +
				"(note: must be added to both sides of the peers): " +
				"'gnet-ctl add neighbor  --neighbor-ip=<ip_address of neighbor> --neighbor-as=<AS_number>' \n" +
				"\t(Example): 'gnet-ctl add neighbor --neighbor-ip=172.16.100.1 --neighbor-as=65001 --description=zone2-r1'",
			Flags: []cli.Flag{
				NeighborIpFlag,
				NeighborAsFlag,
				NeighborDescripFlag,
			},
			Action: AddNeighdor,
		},
		{
			Name: "route",
			Usage: "Advertise a new IP Route from the specified node:" +
				"  'gnet-ctl add route --prefix=<ip_network/prefix> --nexthop=<ip_of_nexthop>' \n" +
				"\t(Example): 'gnet-ctl add route --prefix=172.16.100.100/32 --nexthop=172.16.100.1'",
			Flags: []cli.Flag{
				RouteIpPrefix,
				RouteNextHop,
			},
			Action: AddRoute,
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

type IpRoute struct {
	IpPrefix    string `json:"ip_prefix"`
	PrefixMask  uint8  `json:"ip_mask"`
	NextHop     string `json:"ip_nexthop"`
	LocalPref   uint32 `json:"local_pref"`
	RouteFam    string `json:"route_family"`
	ExCommunity string `json:"opaque"`
}

func AddRoute(c *cli.Context) {
	var (
		cidr         *net.IPNet
		err          error
		preflen      uint8
		mlen         int
		typedNextHop net.IP
	)
	client := NewClient()
	ipRoute := new(IpRoute)
	// Parse cli flag input
	strIp := c.String("prefix")
	strNextHop := c.String("nexthop")

	if strIp != "" {
		cidr, err = GetCidr(strIp)
		if err != nil {
			log.Error("Error parsing a valid ip prefix.")
			return
		}
	} else {
		log.Error("a valid ip prefix or resolvable hostname is required")
		return
	}
	// validate/type the provided ip via net.IP then cast back to string
	if strNextHop != "" {
		typedNextHop = net.ParseIP(strNextHop)
	} else {
		log.Error("a valid ip nexthop is required")
		return
	}
	if cidr.Mask != nil {
		mlen, _ = cidr.Mask.Size()
	}
	preflen = uint8(mlen)
	ipRoute.NextHop = typedNextHop.String()
	ipRoute.IpPrefix = cidr.IP.String()
	ipRoute.PrefixMask = preflen
	log.Debugf("sending route addition request for Prefix: %s/%d Nexthop: %s",
		ipRoute.IpPrefix, ipRoute.PrefixMask, ipRoute.NextHop)
	j, err := json.Marshal(&ipRoute)
	if err != nil {
		log.Println(err)
		return
	}
	log.Debugf("marshalled route rest request: %s ", j)
	// http post rout request
	req, err := http.NewRequest("POST", "http://127.0.0.1:8080"+ROUTE_TABLES+ADD, bytes.NewBuffer(j))
	if err != nil {
		log.Errorf("no answer from the bgp daemon: %s \n", err)
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
		log.Errorf("no answer from the bgp daemon: %s \n", err)
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
	// Display Results for debugging
	log.Debugln("response Status : ", resp.Status)
	log.Debugln("response Headers : ", resp.Header)
	log.Debugln("response Body : ", string(respBody))
	log.Infoln("results: ", string(respBody))
	return
}
