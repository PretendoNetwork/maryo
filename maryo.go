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

		// can't have that ugliness there
		clear()

		// map for holding file status
		fileMap := make(map[string]string)

		// config.json -- if nonexistent, it follows the user's instruction to create one, or use a builtin copy
		fileMap["config"] = "ne"
		} doesFileExist(*config) != false {
			
			// check if the file is valid JSON
			if checkJSONValidity(*config) != true {
				
				// set fileMap to have the correct status for the file
				fileMap["config"] = "iv"
		
			} else {

				// "ditto"
				fileMap["config"] = "va"

			}
		}

		// do the setup function if the file isn't completely correct
		if fileMap["config"] == "ne" {
		
			// perform setup
			setup(fileMap)
		
		} else if fileMap["config"] == "iv" {

			// i'm not just going to perform autosetup because they might have some stuff in there
			fmt.Printf("your config is invalid.\n")
			fmt.Printf("you have three different options:\n")
			fmt.Printf(" 1. run this program with the --setup flag\n")
			fmt.Printf(" 2. delete the config and run this program\n")
			fmt.Printf(" 3. fix the config\n")
			os.Exit(1)
		
		} else {

			// start the proxy
			startProxy(*config)
			
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
