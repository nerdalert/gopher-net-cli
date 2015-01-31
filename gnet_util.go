package main

import (
	"fmt"
	"net"
	"strconv"

	"errors"
	log "github.com/nerdalert/gopher-net-ctl/Godeps/_workspace/src/github.com/Sirupsen/logrus"
	"strings"
)

// convert AS number from string to uint32
func StrToUin32(s string) uint32 {
	n, err := strconv.ParseInt(s, 0, 32)
	if err != nil {
		log.Errorf("[ %d ] is not a valid 32-bit AS number, please use a valid integer: %s", n, err)
	}
	n32 := uint32(n)
	return n32
}

func isValidIp(url string) bool {
	return ValidV4RegEx.MatchString(url)
}

func dnsLookup(s string) (string, error) {
	ipAddr, err := net.ResolveIPAddr("ip", s)
	return ipAddr.IP.String(), err
}

func getHostRoute(h string) *net.IPNet {
	hr := fmt.Sprintf(h + "/32")
	_, hroute, err := net.ParseCIDR(hr)
	if err != nil {
		log.Errorf("cidr parse error: %s", err)
		return nil
	}
	return hroute
}

func GetCidr(ipstr string) (*net.IPNet, error) {
	// check for a cidr format
	if strings.Contains(ipstr, "/") {
		if _, ipnet, err := net.ParseCIDR(ipstr); err == nil {
			return ipnet, nil
		}
	}
	// check for a valid IP address. if it is valid build and return a host route.
	ok := isValidIp(ipstr)
	if ok {
		return getHostRoute(ipstr), nil
	}
	// if the string is not an ip address, try to resolve an ip from hostname
	ipstr = strings.TrimSpace(ipstr)
	if resolvedIP, err := dnsLookup(ipstr); err == nil {
		fmt.Println("hostname received, resolved host route:", getHostRoute(resolvedIP))
		return getHostRoute(resolvedIP), nil
	}
	return nil, errors.New("No valid IP address, CIDR or resolvable hostname found")
}
