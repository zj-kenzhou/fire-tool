package ip

import (
	"errors"
	"net"
)

// 获取ip
func getIpFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	ip = ip.To4()
	if ip == nil {
		return nil // not an ipv4 address
	}

	return ip
}

// ExternalIp 获取本机ip
func ExternalIp() (net.IP, error) {
	faces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, face := range faces {
		if face.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if face.Flags&net.FlagLoopback != 0 {
			continue // logback interface
		}
		adders, err := face.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range adders {
			ip := getIpFromAddr(addr)
			if ip == nil {
				continue
			}
			return ip, nil
		}
	}
	return nil, errors.New("connected to the network")
}

// ExternalIpNoErr 获取ip,不返回错误
func ExternalIpNoErr() net.IP {
	ip, err := ExternalIp()
	if err != nil {
		panic(err)
	}
	return ip
}
