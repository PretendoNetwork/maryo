/*

maryo/utils.go

contains some utilities for things like JSON

written by superwhiskers, licensed under gnu gplv3.
if you want a copy, go to http://www.gnu.org/licenses/

*/

package main

import (
	// internals
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	// externals
	"github.com/elazarl/goproxy"
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

// setting CA in goproxy
func setCA(caCert, caKey []byte) error {

	// get the keypair from cert and key data
	goproxyCa, err := tls.X509KeyPair(caCert, caKey)

	// handle error
	if err != nil {

		// return the error
		return err

	}

	// again, handle errors
	if goproxyCa.Leaf, err = x509.ParseCertificate(goproxyCa.Certificate[0]); err != nil {

		// return the error
		return err

	}

	// set the CA
	goproxy.GoproxyCa = goproxyCa

	// on connections, use it
	goproxy.OkConnect = &goproxy.ConnectAction{Action: goproxy.ConnectAccept, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}

	// on MITMed connections, use it
	goproxy.MitmConnect = &goproxy.ConnectAction{Action: goproxy.ConnectMitm, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}

	// on MITMed HTTP connections, use it
	goproxy.HTTPMitmConnect = &goproxy.ConnectAction{Action: goproxy.ConnectHTTPMitm, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}

	// on rejected connections, use it
	goproxy.RejectConnect = &goproxy.ConnectAction{Action: goproxy.ConnectReject, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}

	// then return nil since there were no errors
	return nil

}
