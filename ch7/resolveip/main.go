// Resolving local IP addresses
package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// All network interfaces
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}

	for i, inter := range interfaces {
		// Resolve addresses for each interface
		addrs, err := inter.Addrs()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d %v\n", i, inter.Name)
		for i, a := range addrs {
			if ip, ok := a.(*net.IPNet); ok {
				fmt.Printf("%d - IP Address: %s\n", i, ip)
			}
		}
	}
}
