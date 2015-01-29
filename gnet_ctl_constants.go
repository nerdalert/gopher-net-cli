package main

import (
	"github.com/nerdalert/gopher-net-ctl/Godeps/_workspace/src/github.com/codegangsta/cli"
)

const (
	VERSION            = "/v1"
	ROUTES             = "/bgp/routes"
	ADJ_RIB_LOCAL      = "/adj-rib-local"
	RIB_LOCAL          = "/local-rib"
	NEIGHBOR_ADDR      = "remotePeerAddr"
	REMOTE_AS_ARG      = "remoteAS"
	REMOTE_NEIGHBOR_AS = "/neighbor-as"
	GLOBAL_CONF        = "/bgp/conf/global"
	NEIGHBORS_CONF     = "/bgp/conf/neighbors"
	ADD                = "/add"
	DEL                = "/delete"
	RIB_OUT_PREFIX     = "/routes-out"
	RIB_IN_PREFIX      = "/routes-in"
	NEIGHBOR_PREFIX    = "/bgp/neighbor"
	NEIGHBORS_PREFIX   = "/bgp/neighbors"
	NEIGHBOR           = VERSION + NEIGHBOR_PREFIX
	NEIGHBORS          = VERSION + NEIGHBORS_PREFIX
	ROUTE_TABLES       = VERSION + ROUTES
	GLOBAL_CONFIG      = VERSION + GLOBAL_CONF
	NEIGHBORS_CONFIG   = VERSION + NEIGHBORS_CONF
	RIB_IN             = ROUTE_TABLES + RIB_IN_PREFIX
	RIB_OUT            = ROUTE_TABLES + RIB_OUT_PREFIX
	REST_PORT          = 8080
)

var (
	NeighborIpFlag      = cli.StringFlag{Name: "neighbor-ip", Value: "", Usage: "ip address of the bgp neighbor", EnvVar: ""}
	NeighborAsFlag      = cli.StringFlag{Name: "neighbor-as", Value: "", Usage: "ip address of the bgp neighbor", EnvVar: ""}
	NeighborDescripFlag = cli.StringFlag{Name: "description", Value: "", Usage: "bgp neighbor description", EnvVar: ""}
)
