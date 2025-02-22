package ip

import (
	"net"
	"strings"
)

// 判断一个IP是否是内网IP
func HasLocalIPAddr(ip string) bool {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}

	// 判断IPv4
	if parsedIP.To4() != nil {
		// 判断是否为内网IPv4地址
		return isPrivateIPv4(parsedIP)
	}

	// 判断IPv6
	return isPrivateIPv6(parsedIP)
}

// 判断是否是内网IPv4
func isPrivateIPv4(ip net.IP) bool {
	//这里对于ipv4来说，只有后4位是真实的ip地址，前面都是0
	first, second := ip[12], ip[13]
	// 获取IP的前缀
	if first == 10 {
		// 10.0.0.0/8
		return true
	}
	if first == 172 && second >= 16 && second <= 31 {
		// 172.16.0.0/12
		return true
	}
	if first == 192 && second == 168 {
		// 192.168.0.0/16
		return true
	}
	if first == 127 {
		// 127.0.0.0/8 (Loopback 地址)
		return true
	}
	return false
}

// 判断是否是内网IPv6
func isPrivateIPv6(ip net.IP) bool {
	// 检查IPv6是否为ULA（Unique Local Address）
	return strings.HasPrefix(ip.String(), "fc00::") || strings.HasPrefix(ip.String(), "fd00::")
}
