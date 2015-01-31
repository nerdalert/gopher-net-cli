package main

import (
	"reflect"
	"testing"

	log "github.com/nerdalert/gopher-net-ctl/Godeps/_workspace/src/github.com/Sirupsen/logrus"
)

func TestAsStringTo32(t *testing.T) {
	var n uint32
	var as string = "650001"
	ustr := StrToUin32(as)
	us := reflect.ValueOf(ustr).Kind()
	u := reflect.ValueOf(n).Kind()
	assertEqual(us, u)
}

func TestPrefixParsing(t *testing.T) {
	ip := "172.16.203.1"
	ipCidr := "172.16.203.1/32"
	hostname := "google-public-dns-a.google.com."
	gipCidr := "8.8.8.8/32"
	// test cidr input parse
	addr, _ := GetCidr(ip)
	assertEqual(addr.String(), ipCidr)
	// test host ip only input to host route parse
	addr, _ = GetCidr(ipCidr)
	assertEqual(addr.String(), ipCidr)
	// test hostname input to host route parse
	addr, _ = GetCidr(hostname)
	assertEqual(addr.String(), gipCidr)
}

func assertEqual(v interface{}, v1 interface{}) bool {
	if v != v1 {
		log.Errorf("the values %v and %v are not equal.", v1, v)
		return false
	}
	return true
}
