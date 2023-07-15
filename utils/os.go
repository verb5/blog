package utils


import (
	"net"
	"fmt"
	"os"
)
func returnString(s string) string {
	m:=[]byte(s)
	v := make([]byte,len(m))
	copy(v,m)
	return string(v)
}


func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}
	return hostname
}

func getIPAddress() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
	}
	var ip string
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			// fmt.Println(ipnet.IP.String())
			ip = ipnet.IP.String()
		}
	}
	return ip
}