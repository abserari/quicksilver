// +build !windows

package fing

import (
	"net"
	"time"

	"github.com/j-keck/arping"
)

func Mac(ip string) (net.HardwareAddr, time.Duration, error) {
	dstIP := net.ParseIP(ip)
	return arping.Ping(dstIP)
}
