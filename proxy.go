/*

maryo/proxy.go

the proxy that makes this program a reverse proxy
made by that magical 3ds person

written by superwhiskers, licensed under gnu gplv3.
if you want a copy, go to http://www.gnu.org/licenses/

*/

package main

import (
	// internals
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"net/http/httputil"
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
	if doesFileExist("maryo-data/proxy.log") == false {

		// make it then
		createFile("maryo-data/proxy.log")

	}

	// write current timestamp to log
	t := time.Now().Format("20060102150405")
	writeFile("maryo-data/proxy.log", fmt.Sprintf("-> started log [%s]\n", t))

	// get ip
	ip := getIP()

	// start the console log
	fmt.Printf("-- proxy log --\n")
	consoleSequence(fmt.Sprintf("-> local IP address is %s%s%s\n", code("green"), ip, code("reset")))
	consoleSequence(fmt.Sprintf("-> hosting proxy on %s:9437%s\n", code("green"), code("reset")))
	writeFile("maryo-data/proxy.log", fmt.Sprintf("-> got local ip as %s, hosting on port :9437", ip))

	// load that proxy
	proxy := goproxy.NewProxyHttpServer()

	// set some settings

	// http client for use when performing POST requests
	httpClient := &http.Client{}

	cert, err := ioutil.ReadFile("maryo-data/cert.pem")
	if err != nil {
		panic(err)
	}

	key, err := ioutil.ReadFile("maryo-data/cert.key")
	if err != nil {
		panic(err)
	}

	// add the ninty cert and key to the proxy for decrypting
	setCA(cert, key)

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
			writeFile("maryo-data/proxy.log", fmt.Sprintf("-> got request to %s\n", r.URL.Host))

			// get prettified request

			// define these variables so no errors
			var reqData []byte
			var err error

			// no errors for post requests
			if r.Method == "POST" {

				// if it is, then tell the dumper it is
				reqData, err = httputil.DumpRequest(r, true)

			} else {

				// otherwise, don't
				reqData, err = httputil.DumpRequest(r, false)

			}

			if err != nil {

				// output error
				fmt.Printf("[err]: error occurred while dumping http request\n")
				fmt.Printf("%s\n", err.Error())

			}

			// if it is said to be verbose with logging, print request data
			if logging == true {

				// log the request data, then
				fmt.Printf("\n-- request data\n")
				fmt.Printf("%s", string(reqData[:]))
				fmt.Printf("\n\n")

			}

			// always log to file
			writeFile("maryo-data/proxy.log", fmt.Sprintf("-> request data to %s\n", r.URL.Host))
			writeFile("maryo-data/proxy.log", fmt.Sprintf("%s", string(reqData[:])))
			writeFile("maryo-data/proxy.log", fmt.Sprintf("\n\n"))

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
				writeFile("maryo-data/proxy.log", fmt.Sprintf("-> proxying %s to %s", r.URL.Host, redirTo))

				// redirect it
				r.URL.Host = redirTo

			}

			// here is a fancy workaround for POST requests
			if r.Method == "POST" {

				// show the user that we performed a POST request
				fmt.Printf("-> performing %s request to %s%s://%s%s%s\n", r.Method, code("green"), r.URL.Scheme, r.URL.Host, r.URL.Path, code("reset"))

				// clone the request
				newReq := cloneReq(r)

				// perform the request
				resp, err := httpClient.Do(newReq)

				// error handling
				if err != nil {

					// return a response
					return r, goproxy.NewResponse(newReq, goproxy.ContentTypeText, http.StatusBadGateway, strings.Join([]string{"no worries, this is an error in maryo\n", err.Error()}, ""))

				}

				// dump response

				// make these variables earlier on so no errors
				var fmtResp []byte

				// check if the request is a post request
				if r.Method == "POST" {

					// if so, tell the dumper that it is one
					fmtResp, err = httputil.DumpResponse(resp, true)

				} else {

					// otherwise, don't
					fmtResp, err = httputil.DumpResponse(resp, false)

				}

				// error handling
				if err != nil {

					// log the error
					fmt.Printf("[err]: error while dumping response")
					fmt.Printf("%s\n", err.Error())

				}

				// make sure the user wants to log response data
				if logging == true {

					// log it if they do
					fmt.Printf("\n-- response data\n")
					fmt.Printf("%s\n", string(fmtResp[:]))
					fmt.Printf("\n\n")

				}

				// return the processed response
				return r, resp

			}

			// just return nil for response, since we aren't modifying it
			return r, nil

		})

	// start the proxy
	log.Fatal(http.ListenAndServe(":9437", proxy))

}
