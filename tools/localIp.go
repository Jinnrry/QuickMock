package tools

import (
	"net"
)

var localIp = ""

func GetLocalIp() string {
	if localIp != "" {
		return localIp
	}

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		localIp = "localhost"
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				localIp = ipnet.IP.String()
				break
			}
		}
	}

	if localIp == "" {
		localIp = "localhost"
	}
	return localIp
}
