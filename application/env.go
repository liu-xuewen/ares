package application

import (
	"fmt"
	"net"
	"os"
)

// BaseLogDir gets ARES_BASE_LOG_DIR.
func BaseLogDir() string {
	return os.Getenv("ARES_BASE_LOG_DIR")
}

// EnvAccessLogDir gets ARES_ACCESS_LOG_DIR.
func EnvAccessLogDir() string {
	return os.Getenv("ARES_ACCESS_LOG_DIR")
}

// EnvServerHost gets ARES_SERVER_HOST.
func EnvServerHost() string {
	return os.Getenv("ARES_SERVER_HOST")
}

// EnvAppMode gets ARES_APP_MODE.
func EnvAppMode() string {
	return os.Getenv("ARES_APP_MODE")
}

// getMacAddrs 检测当前服务器
func getMacAddrs() (macAddrs []string) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("fail to get net interfaces: %v", err)
		return macAddrs
	}

	for _, netInterface := range netInterfaces {
		macAddr := netInterface.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}

		macAddrs = append(macAddrs, macAddr)
	}
	return macAddrs
}

func getIPs() (ips []string) {
	interfaceAddr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf("fail to get net interface addrs: %v", err)
		return ips
	}

	for _, address := range interfaceAddr {
		ipNet, isValidIPNet := address.(*net.IPNet)
		if isValidIPNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}
	return ips
}
