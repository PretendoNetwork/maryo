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

	// verify config data

	fmt.Printf("-- proxy log --\n")
	consoleSequence(fmt.Sprintf("-> local IP address is %s%s%s\n", code("green"), getIP(), code("reset")))
	consoleSequence(fmt.Sprintf("-> hosting proxy on %s:9437%s\n", code("green"), code("reset")))

	// load that proxy
	proxy := goproxy.NewProxyHttpServer()

	// set some settings

	// verbose mode can be a little... too verbose
	proxy.Verbose = logging

	// set up the proxy

	// make it always MITM
	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)

	// request handler
	proxy.OnRequest().DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {

			// log the request
			consoleSequence(fmt.Sprintf("-> request to %s%s%s\n", code("green"), r.URL.Host, code("reset")))

			// if it is said to be verbose with logging, print request data
			fmt.Printf("\n-- request data\n")
			fmt.Printf("%s", formatRequest(r))
			fmt.Printf("\n\n")

			// attempt to proxy it to the servers listed in config

			// check if it is in it in the first place
			// also, strip the URL of the port
			if redirTo, isItIn := config[strings.Split(r.URL.Host, ":")[0]]; isItIn {

				// if protocol is HTTPS
				if r.URL.Scheme == "https" {

					// set it to HTTP
					r.URL.Scheme = "http"

				}

				// log the redirect
				consoleSequence(fmt.Sprintf("-> proxying %s%s%s to %s%s%s\n", code("green"), r.URL.Host, code("reset"), code("green"), redirTo.(string), code("reset")))

				// redirect it
				r.URL.Host = redirTo.(string)

			}

			// just return nil for response, since we aren't modifying it
			return r, nil

		})

	// start the proxy
	log.Fatal(http.ListenAndServe(":9437", proxy))
}
