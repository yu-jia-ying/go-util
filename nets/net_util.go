package nets

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func GetIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	// 192.168.1.20:61085
	ip := strings.Split(localAddr.String(), ":")[0]

	return ip
}
func GetMacAddr() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic("Poor soul, here is what you got: " + err.Error())
	}
	if len(interfaces) == 0 {
		return ""
	}

	maxIndexInterface := interfaces[0]
	for _, inter := range interfaces {
		if inter.HardwareAddr == nil {
			continue
		}
		if inter.Flags&net.FlagUp == 1 {
			maxIndexInterface = inter
		}
	}
	return maxIndexInterface.HardwareAddr.String()
}

func GetHostname() string {
	name, err := os.Hostname()
	if err != nil {
		return ""
	}
	return name
}
