// +build windows

package fing

import (
	"fmt"
	"log"
	"net"
	"syscall"
	"time"
	"unsafe"
)

var SendARP = syscall.MustLoadDLL("iphlpapi.dll").MustFindProc("SendARP")

func ip4ToUint32(ip net.IP) (uint32, error) {
	ip = ip.To4()
	if ip == nil {
		return 0, fmt.Errorf("ip address %v is not ip4", ip)
	}
	var ret uint32
	for i := 4; i > 0; i-- {
		ret <<= 8
		ret += uint32(ip[i-1])
	}
	return ret, nil
}

func Mac(ip string) (net.HardwareAddr, time.Duration, error) {
	dstIP := net.ParseIP(ip)
	haddr, err := sendARP(dstIP)
	return haddr, time.Duration(0), err
}

func sendARP(ip net.IP) (net.HardwareAddr, error) {
	dst, err := ip4ToUint32(ip)
	if err != nil {
		return nil, err
	}
	mac := []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
	n := uint32(len(mac))
	ret, _, _ := SendARP.Call(
		uintptr(dst),
		0,
		uintptr(unsafe.Pointer(&mac[0])),
		uintptr(unsafe.Pointer(&n)))
	if ret != 0 {
		return nil, syscall.Errno(ret)
	}
	return mac, nil
}

// ExampleARP -
func ExampleARP() {
	ip := "192.168.0.250"
	mac, _, err := Mac(ip)
	if err != nil {
		log.Fatalf("could not find MAC for %q: %v", ip, err)
	}
	fmt.Printf("MAC address for %v is %v\n", ip, mac)
}
