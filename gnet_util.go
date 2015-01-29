package main

import (
	"strconv"

	log "github.com/nerdalert/gopher-net-ctl/Godeps/_workspace/src/github.com/Sirupsen/logrus"
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
