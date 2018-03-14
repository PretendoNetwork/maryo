/*

maryo/proxy.go

the proxy that makes this program a reverse proxy
made by that magical 3ds person

written by Superwhiskers, licensed under gnu gplv3.
if you want a copy, go to http://www.gnu.org/licenses/

*/

package main

import (
	"fmt"
)
//"github.com/cssivision/reverseproxy"
//"os"

func startProxy(configName string) {

	// set the terminal title
	ttitle("maryo -> proxy")

	// get the config data
	//config := readJSONFile(configName)

	consoleSequence(fmt.Sprintf("hey, just a %stest%s message\n", code("red"), code("reset")))
}
