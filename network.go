package utils

import "net"

func GetLocalIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		panic(err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}()

	localAddress := conn.LocalAddr().(*net.UDPAddr)
	return localAddress.IP
}
