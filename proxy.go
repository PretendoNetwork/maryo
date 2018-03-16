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
	consoleSequence(fmt.Sprintf("-> local IP addresss is %s%s%s\n", code("green"), getIP(), code("reset")))
	consoleSequence(fmt.Sprintf("-> hosting proxy on %s:9437%s\n", code("green"), code("reset")))

	// load that proxy
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	log.Fatal(http.ListenAndServe(":9437", proxy))
}
