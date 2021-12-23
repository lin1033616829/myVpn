package main

import (
	"fmt"
	"log"
	"myVpn/vpnServer/initialize"
	"myVpn/vpnServer/service"
	"net"
	"strings"
)

func main() {
	initialize.InitLog()


	server, err := net.Listen("tcp", ":7080")
	if err != nil {
		fmt.Printf("Listen failed: %v\n", err)
		return
	}

	go initialize.NotifyBackend(server)

	for {
		client, err := server.Accept()
		if err != nil {
			if strings.Contains(err.Error(), "use of closed network connection") {
				break
			}
			log.Printf("Accept failed: %v", err)
			continue
		}

		go service.Process(client)
	}
}



