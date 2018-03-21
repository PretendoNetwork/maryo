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
	"reflect"
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

// function for zeroing something
// takes a pointer (&varible)
func erase(v interface{}) {

	// get the pointer
	p := reflect.ValueOf(v).Elem()

	// zero it
	p.Set(reflect.Zero(p.Type()))

}

// clone a request into a requestable request object
func cloneReq(request *http.Request) *http.Request {

	// clone the request
	newReq := &http.Request{

		Method:     request.Method,
		URL:        request.URL,
		Proto:      request.Proto,
		ProtoMajor: request.ProtoMajor,
		ProtoMinor: request.ProtoMinor,
		Header:     request.Header,
		Body:       request.Body,
		Host:       request.Host,
	}

	// return the cloned request
	return newReq

}
