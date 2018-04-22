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
)

// main function
func main() {

	// reset term colors
	consoleSequence(fmt.Sprintf("%s", code("reset")))

	// parse some flags here
	config := flag.String("config", "maryo-data/config.json", "value for config file path (default is maryo/config.json)")
	logging := flag.Bool("logging", false, "if set, the proxy will log all request data (only needed for debugging)")
	doSetup := flag.Bool("setup", false, "if set, maryo will go through setup again")
	generateCerts := flag.Bool("regencerts", false, "if set, maryo will generate self-signed certificates for private use")
	flag.Parse()

	// set window title
	ttitle("maryo")

	// generate the certs if needed
	if *generateCerts == true {

		// that's it, really
		doCertGen(*config)

		// clear the screen
		clear()

		// give the user a message
		fmt.Printf("your certificate and key pair have been generated\n")
		fmt.Printf("reload the program to use them.\n")

		// close the program
		os.Exit(0)

	}

	// if not forced to do setup
	if *doSetup == false {

		// can't have that ugliness there
		clear()

		// map for holding file status
		fileMap := make(map[string]string)

		// config.json -- if nonexistent, it follows the user's instruction to create one, or use a builtin copy

		// set it to nonexistent beforehand
		fileMap["config"] = "ne"

		// check if config exists
		if doesFileExist(*config) != false {

			// check if the file is valid JSON
			if checkJSONValidity(*config) != true {

				// set fileMap to have the correct status for the file
				fileMap["config"] = "iv"

			// if it is valid
			} else {

				// "ditto"
				fileMap["config"] = "va"

			}

		}

		// cert.pem and key.pem -- if nonexistent, just do setup
		fileMap["cert"] = "ne"

		// check the cert
		if doesFileExist("maryo-data/cert.pem") != false {

			// check the pubkey
			if doesFileExist("maryo-data/public-key.pem") != false {
				
				// check the privatekey
				if doesFileExist("maryo-data/private-key.pem") != false {

					// say it is valid if it is there
					fileMap["cert"] = "va"

				}

			}

		}

		// do the setup function if the file isn't completely correct
		if fileMap["config"] == "ne" {

			// perform setup
			setup(fileMap)

		// if it's invalid
		} else if fileMap["config"] == "iv" {

			// i'm not just going to perform autosetup because they might have some stuff in there
			fmt.Printf("your config is invalid.\n")
			fmt.Printf("you have three different options:\n")
			fmt.Printf(" 1. run this program with the --setup flag\n")
			fmt.Printf(" 2. delete the config and run this program\n")
			fmt.Printf(" 3. fix the config\n")
			os.Exit(1)

		// if the certificates don't exist
		} else if fileMap["cert"] == "ne" {

			// i'm not going to force you to set it up again
			fmt.Printf("you don't have any certs in the maryo-data folder\n")
			fmt.Printf("you have three different options:\n")
			fmt.Printf(" 1. run this program with the --regencerts flag\n")
			fmt.Printf(" 2. run this program with the --setup flag\n")
			fmt.Printf(" 3. provide your own certs\n")
			os.Exit(1)

		// otherwise, start the proxy
		} else {

			// start the proxy
			startProxy(*config, *logging)

		}

		// run setup function
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
