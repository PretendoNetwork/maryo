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
	"net/http"
	"strings"
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

// pretty-print HTTP requests
func formatRequest(r *http.Request) string {

	// create return string
	var request []string

	// add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)

	// add the host
	request = append(request, fmt.Sprintf("host: %v", r.Host))

	// loop through headers
	for name, headers := range r.Header {

		// lowercase any headers
		name = strings.ToLower(name)

		// loop through the header data
		for _, h := range headers {

			// add it to the output
			request = append(request, fmt.Sprintf("%v: %v", name, h))

		}
	}

	// if this is a POST request, add post data
	if r.Method == "POST" {

		// parse the form
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())

	}

	// return the request as a string
	return strings.Join(request, "\n")
}
