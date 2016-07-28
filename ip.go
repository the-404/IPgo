// Copyright 2016, 404@the404.me. All rights reserved.

package IPgo

import (
	"net"
)

var list []string

func incIP(ip net.IP) {

	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func Iplist(cidr string) []string {
	ip, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		ip = net.ParseIP(cidr)
		list = append(list, ip.String())
		return list
	}
	for ip := ip.Mask(ipNet.Mask); ipNet.Contains(ip); incIP(ip) {

		list = append(list, ip.String())
	}
	return list
}
