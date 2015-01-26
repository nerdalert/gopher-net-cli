package main

const (
	BASE_VERSION       = "/v1"
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
	NEIGHBOR           = BASE_VERSION + NEIGHBOR_PREFIX
	NEIGHBORS          = BASE_VERSION + NEIGHBORS_PREFIX
	ROUTE_TABLES       = BASE_VERSION + ROUTES
	GLOBAL_CONFIG      = BASE_VERSION + GLOBAL_CONF
	NEIGHBORS_CONFIG   = BASE_VERSION + NEIGHBORS_CONF
	RIB_IN             = ROUTE_TABLES + RIB_IN_PREFIX
	RIB_OUT            = ROUTE_TABLES + RIB_OUT_PREFIX
	REST_PORT          = 8080
)
