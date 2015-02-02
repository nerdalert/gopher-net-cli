[![Build Status](https://travis-ci.org/nerdalert/gopher-net-ctl.svg?branch=master)](https://travis-ci.org/nerdalert/gopher-net-ctl)


# gopher-net-cli
CLI for Gopher Net Router

### Example Advertise a new route

-Example adding a prefix to advertise into the tables (only a couple of fields are plumbed in until it uses the daemon data structures):

	gopher-net-ctl -d add route --prefix=172.15.14.0/24  --nexthop=172.16.86.1

-Example Output in Quagga

	ub134(config-router)# do sho ip route
	Codes: K - kernel route, C - connected, S - static, R - RIP,
	       O - OSPF, I - IS-IS, B - BGP, A - Babel,
	       > - selected route, * - FIB route

	K>* 0.0.0.0/0 via 172.16.86.2, eth0
	C>* 127.0.0.0/8 is directly connected, lo
	B>* 172.15.14.0/24 [200/0] via 172.16.86.1, eth0, 00:00:03
	C>* 172.16.86.0/24 is directly connected, eth0

### Example remove the advertised route

-This will stop advertising the requested prefix. There is some cidr from the specified node.from the specified note a new route

	gopher-net-ctl -d delete route --prefix=172.15.14.0/24

Example Output in Quagga

	ub134(config-router)# do sho ip route
	Codes: K - kernel route, C - connected, S - static, R - RIP,
	       O - OSPF, I - IS-IS, B - BGP, A - Babel,
	       > - selected route, * - FIB route

	K>* 0.0.0.0/0 via 172.16.86.2, eth0
	C>* 127.0.0.0/8 is directly connected, lo
	C>* 172.16.86.0/24 is directly connected, eth0

### Host Route Support

This and the API will also support hostnames that are resolvable by the target node/bgp speaker.

-Add a hosts DNS name (or VIP) to be advertised.

	gopher-net-ctl -d add route --prefix=google-public-dns-a.google.com  --nexthop=172.16.86.1

-The the resolved address will be advertised into the IGP/EGP as in the Quagga interface output below that is peered to an instance of the gnet daemon.

	ub134(config-router)# do sho ip route
	Codes: K - kernel route, C - connected, S - static, R - RIP,
	       O - OSPF, I - IS-IS, B - BGP, A - Babel,
	       > - selected route, * - FIB route

	K>* 0.0.0.0/0 via 172.16.86.2, eth0
	B>* 8.8.8.8/32 [200/0] via 172.16.86.1, eth0, 00:00:08
	C>* 127.0.0.0/8 is directly connected, lo
	B>* 172.15.14.0/24 [200/0] via 172.16.86.1, eth0, 00:07:46
	C>* 172.16.86.0/24 is directly connected, eth0


