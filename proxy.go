/*

maryo/proxy.go

the proxy that makes this program a reverse proxy
made by that magical 3ds person

written by Superwhiskers, licensed under gnu gplv3.
if you want a copy, go to http://www.gnu.org/licenses/

*/

package main

import (
	// internals

	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	// externals
	"github.com/elazarl/goproxy"
)

//"os"
// set this over here for no issues
var config map[string]interface{}

func startProxy(configName string, logging bool) {

	// set the terminal title
	ttitle("maryo -> proxy")

	// get the config data
	config = readJSONFile(configName)

	// check if we decrypt all connections
	decryptAll := config["config"].(map[string]interface{})["decryptOutgoing"].(string)

	// check if log file exists
	if doesFileExist("maryo/proxy.log") == false {

		// make it then
		createFile("maryo/proxy.log")

	}

	// write current timestamp to log
	t := time.Now().Format("20060102150405")
	writeFile("maryo/proxy.log", fmt.Sprintf("-> started log [%s]\n", t))

	// get ip
	ip := getIP()

	// start the console log
	fmt.Printf("-- proxy log --\n")
	consoleSequence(fmt.Sprintf("-> local IP address is %s%s%s\n", code("green"), ip, code("reset")))
	consoleSequence(fmt.Sprintf("-> hosting proxy on %s:9437%s\n", code("green"), code("reset")))
	writeFile("maryo/proxy.log", fmt.Sprintf("-> got local ip as %s, hosting on port :9437", ip))

	// load that proxy
	proxy := goproxy.NewProxyHttpServer()

	// set some settings

	// add the ninty cert and key to the proxy for decrypting
	setCA(nintyCert, nintyKey)

	// verbose mode can be a little... too verbose
	proxy.Verbose = logging

	// set up the proxy

	// this needs to be set for request data
	var reqData string

	// make it always MITM
	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)

	// request handler
	proxy.OnRequest().DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {

			// log the request
			consoleSequence(fmt.Sprintf("-> request to %s%s%s\n", code("green"), r.URL.Host, code("reset")))

			writeFile("maryo/proxy.log", fmt.Sprintf("-> got request to %s\n", r.URL.Host))

			// get prettified request
			reqData = formatRequest(r)

			// if it is said to be verbose with logging, print request data
			if logging == true {

				// log the request data, then
				fmt.Printf("\n-- request data\n")
				fmt.Printf("%s", reqData)
				fmt.Printf("\n\n")

			}

			// always log to file
			writeFile("maryo/proxy.log", fmt.Sprintf("-> request data to %s\n", r.URL.Host))
			writeFile("maryo/proxy.log", fmt.Sprintf("%s", reqData))
			writeFile("maryo/proxy.log", fmt.Sprintf("\n\n"))

			// attempt to proxy it to the servers listed in config

			// check if it is in it in the first place
			// also, strip the URL of the port
			if redirTo, isItIn := config["endpoints"].(map[string]interface{})[strings.Split(r.URL.Host, ":")[0]].(string); isItIn {

				// check if we decrypt all outgoing connections
				if decryptAll == "true" {

					// if protocol is HTTPS
					if r.URL.Scheme == "https" {

						// let the user know
						fmt.Printf("-> switching protocol to http\n")

						// set it to HTTP
						r.URL.Scheme = "http"

					}

				}

				// log the redirect
				consoleSequence(fmt.Sprintf("-> proxying %s%s%s to %s%s%s\n", code("green"), r.URL.Host, code("reset"), code("green"), redirTo, code("reset")))
				writeFile("maryo/proxy.log", fmt.Sprintf("-> proxying %s to %s", r.URL.Host, redirTo))

				// redirect it
				r.URL.Host = redirTo

			}

			// just return nil for response, since we aren't modifying it
			return r, nil

		})

	// start the proxy
	log.Fatal(http.ListenAndServe(":9437", proxy))
}
