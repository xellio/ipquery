package main

import (
	"bytes"
	"net"
)

//
//IPRange - a struct that holds the start and end of an IP range
//
type IPRange struct {
	Start net.IP
	End   net.IP
}

//
// privateRanges ...
//
var privateRanges = []IPRange{
	IPRange{
		Start: net.ParseIP("10.0.0.0"),
		End:   net.ParseIP("10.255.255.255"),
	},
	IPRange{
		Start: net.ParseIP("100.64.0.0"),
		End:   net.ParseIP("100.127.255.255"),
	},
	IPRange{
		Start: net.ParseIP("172.16.0.0"),
		End:   net.ParseIP("172.31.255.255"),
	},
	IPRange{
		Start: net.ParseIP("192.0.0.0"),
		End:   net.ParseIP("192.0.0.255"),
	},
	IPRange{
		Start: net.ParseIP("192.168.0.0"),
		End:   net.ParseIP("192.168.255.255"),
	},
	IPRange{
		Start: net.ParseIP("198.18.0.0"),
		End:   net.ParseIP("198.19.255.255"),
	},
}

//
// InRange - check to see if a given IP address is within a given range
//
func InRange(r IPRange, ipAddress net.IP) bool {
	if bytes.Compare(ipAddress, r.Start) >= 0 && bytes.Compare(ipAddress, r.End) < 0 {
		return true
	}
	return false
}

//
// IsPrivateSubnet - check if the IP address is in a private subnet
//
func IsPrivateSubnet(ipAddress net.IP) bool {
	if ipCheck := ipAddress.To4(); ipCheck != nil {
		for _, r := range privateRanges {
			if InRange(r, ipAddress) {
				return true
			}
		}
	}
	return false
}
