/*

maryo/maryo.go

written by superwhiskers, licensed under gnu gplv3.
if you want a copy, go to http://www.gnu.org/licenses/

*/

package main

import (
	// internals
	"flag"
	"fmt"
	"os"
	"runtime"
)

// main function
func main() {

	// parse some flags here
	config := flag.String("config", "config.json", "value for config file path")
	//logging := flag.Bool("logging", false, "if set, the proxy will log all requests and data")
	doSetup := flag.Bool("setup", false, "if set, maryo will go through setup again")
	flag.Parse()

	// set window title
	ttitle("maryo")

	if *doSetup == false {

		clear()
		fmt.Printf("-- log\n")

		// startup
		fmt.Printf("= startup")

		// i might include some os-specific code
		fmt.Printf("loading 0%%: detecting os")
		fmt.Printf("%s\n", padStrToMatchStr(fmt.Sprintf("\ros: %s", runtime.GOOS), "loading 0%%: detecting os", " "))

		// file checking
		fmt.Printf("loading 50%%: checking files.")

		// map for holding file status
		fileMap := make(map[string]string)

		// config.json -- if nonexistent, it follows the user's instruction to create one, or use a builtin copy
		fileMap["config"] = "ne"
		if doesFileExist(*config) == false {
			fmt.Printf("%s\n", padStrToMatchStr("\rconfig: nonexistent", "loading 50%%: checking files.", " "))
		} else {
			// check if the file is valid JSON
			if checkJSONValidity(*config) != true {
				fmt.Printf("%s\n", padStrToMatchStr("\rconfig: invalid", "loading 50%%: checking files.", " "))
				fileMap["config"] = "iv"
			} else {
				fmt.Printf("%s\n", padStrToMatchStr("\rconfig: valid", "loading 50%%: checking files.", " "))
				fileMap["config"] = "va"
			}
		}

		// print final info
		fmt.Printf("loaded..\n")

		// do the setup function if the file isn't completely correct
		if fileMap["config"] == "ne" {
			setup(fileMap)
		} else if fileMap["config"] == "iv" {
			fmt.Printf("your config is invalid.\n")
			fmt.Printf("you have three different options:\n")
			fmt.Printf(" 1. run this program with the --setup flag\n")
			fmt.Printf(" 2. delete the config and run this program\n")
			fmt.Printf(" 3. fix the config\n")
			os.Exit(1)
		} else {
			startProxy()
		}

		// run setup function and load proxy
	} else {

		// fileMap
		var fileMap map[string]string
		fileMap = make(map[string]string)

		// place config value in there
		fileMap["config"] = "uk"

		// just do the setup function
		setup(fileMap)

	}
}
