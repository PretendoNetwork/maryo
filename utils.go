/*

maryo/utils.go

contains some utilities for things like JSON

written by superwhiskers, licensed under gnu gplv3.
if you want a copy, go to http://www.gnu.org/licenses/

*/

package main

import (
	// internals
	"fmt"
	"log"
	"net"
)

// get the ip address of the machine
func getIP() string {

	// dial a connection to another ip
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Printf("[err]: error while connecting to another computer\n")
		log.Fatal(err)
	}
	defer conn.Close()

	// get the local address
	localAddr := conn.LocalAddr().(*net.UDPAddr)

	// return it
	return localAddr.IP.String()
}
