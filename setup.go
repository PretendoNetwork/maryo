/*

maryo/setup.go

the setup function is here for easier to read code

written by superwhiskers, licensed under gnu gplv3.
if you want a copy, go to http://www.gnu.org/licenses/

*/

package main

import (
	// internals
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/kabukky/httpscerts"
	// externals
)

// cert generation function here so i don't need to rewrite it in maryo.go
func doCertGen() {

	// clear the screen because why not?
	clear()

	// get the ip address for the cert
	ip := getIP()

	// show a neat info snippet
	fmt.Printf("- generating certificate and key for %s\n", ip)

	// create necissary directories for this

	// maryo folder
	if doesDirExist("maryo") == false {

		// make it
		makeDirectory("maryo")
	
	}

	// clean the cert and key if they exist

	// cert.pem
	if doesFileExist("maryo/cert.pem") == true {

		// delete the cert
		deleteFile("maryo/cert.pem")

	}

	// key.pem
	if doesFileExist("maryo/key.pem") == true {

		// delete the key
		deleteFile("maryo/key.pem")

	}

	// generate the needed cert and key
	err := httpscerts.Generate("maryo/cert.pem", "maryo/key.pem", fmt.Sprintf("%s:9437", ip))

	// handle the error (if there is one)
	if err != nil {
		fmt.Printf("[err]: error while generating a certificate and key for %s\n", ip)
		panic(err)
	}

	// then say they were made
	fmt.Printf("  finished\n")
	fmt.Printf("\npress enter to continue...\n")
	_ = input("")

}

// setup function goes here now
func setup(fileMap map[string]string) {

	// test data
	test := make([]string, 1)
	test[0] = "account"

	// test the 'official' pretendo servers
	testOfficial := make([]string, 1)
	testOfficial[0] = "account"

	// file status map
	fileStat := make(map[string]string)
	fileStat["ne"] = "nonexistent"
	fileStat["iv"] = "invalid"
	fileStat["va"] = "valid"
	fileStat["uk"] = "unknown"

	// setup environment
	clear()
	ttitle("maryo -> setup")

	// show setup screen
	fmt.Printf("== maryo -> setup ===========================================\n")
	fmt.Printf("                                        Steps:               \n")
	fmt.Printf(" welcome to the maryo setup wizard.     > intro              \n")
	fmt.Printf(" this program will walk you through       config creation    \n")
	fmt.Printf(" setting up your very own Pretendo        confirm prefs      \n")
	fmt.Printf(" proxy server for accessing the server.   generate cert      \n")
	fmt.Printf(" -> press enter                           profit???          \n")
	fmt.Printf("                                                             \n")
	fmt.Printf("                                                             \n")
	fmt.Printf("                                                             \n")
	fmt.Printf("=============================================================\n")
	input("")

	// show config creation screen
	var method string
	for true {
		clear()
		fmt.Printf("== maryo -> setup ===========================================\n")
		fmt.Printf("                                        Steps:               \n")
		fmt.Printf(" how would you like to configure the      intro              \n")
		fmt.Printf(" proxy?                                 > config creation    \n")
		fmt.Printf(" 1. automatic                             confirm prefs      \n")
		fmt.Printf(" 2. custom                                generate cart      \n")
		fmt.Printf(" 3. template                              profit???          \n")
		fmt.Printf(" 4. skip this                                                \n")
		fmt.Printf("                                                             \n")
		fmt.Printf(" -> (1|2|3|4)                                                \n")
		fmt.Printf("=============================================================\n")
		method = input(": ")

		if (method == "1") || (method == "2") || (method == "3") || (method == "4") {
			break
		} else {
			fmt.Printf("-> please enter 1, 2, 3, or 4\n")
			time.Sleep(1500 * time.Millisecond)
		}
	}

	// create config var
	var config map[string]map[string]string

	// show log when
	clear()
	fmt.Printf("== maryo -> setup ===========================================\n")
	fmt.Printf("                                                             \n")
	fmt.Printf(" configuring proxy..                                         \n")
	fmt.Printf(" current config status: %s\n", fileStat[fileMap["config"]])
	// automatic config making
	if method == "1" {
		fmt.Printf(" method: automatic..\n")
		fmt.Printf("-- beginning tests\n")
		fmt.Printf(" 1. attempting to detect endpoints running on this machine\n")

		// test for endpoints on this machine
		result := make([]bool, len(test))
		for x := 0; x < len(test); x++ {

			// test the endpoint
			fmt.Printf("  %s %s -> %s", utilIcons["uncertain"], testEndpoints["ninty"][test[x]], testEndpoints["local"][test[x]])

			// get the json
			res, err := get(strings.Join([]string{"http://", testEndpoints["local"][test[x]], "/isthisworking"}, ""))

			// prepare a struct for the json
			var parsedRes isitworkingStruct

			// parse it
			err2 := json.Unmarshal([]byte(res), &parsedRes)
			if res != "" {
				if err2 != nil {
					if err == nil {
						fmt.Printf("\n[err] : error when parsing JSON during validating %s server\n", test[x])
						panic(err2)
					}
				}
			}

			// handle the results
			if (parsedRes.Server == resMap[test[x]]) && (err == nil) && (res != "") {

				// MS, step your game up and support ansi escape codes
				consoleSequence(fmt.Sprintf("%s\n", padStrToMatchStr(fmt.Sprintf("\r  %s%s%s%s %s -> %s", code("green"), code("bold"), utilIcons["success"], code("reset"), testEndpoints["ninty"][test[x]], testEndpoints["local"][test[x]]), fmt.Sprintf("  %s %s -> %s", utilIcons["uncertain"], testEndpoints["ninty"][test[x]], testEndpoints["local"][test[x]]), " ")))
				result[x] = true

			} else {

				// make windows great again (as if it ever was)
				consoleSequence(fmt.Sprintf("%s\n", padStrToMatchStr(fmt.Sprintf("\r  %s%s%s%s %s -> %s", code("red"), code("bold"), utilIcons["failiure"], code("reset"), testEndpoints["ninty"][test[x]], testEndpoints["local"][test[x]]), fmt.Sprintf("  %s %s -> %s", utilIcons["uncertain"], testEndpoints["ninty"][test[x]], testEndpoints["local"][test[x]]), " ")))
				result[x] = false

			}

		}

		// show the user that we are detecting the official server
		fmt.Printf(" 2. attempting to test endpoints on the official server\n")

		// test for endpoints on the official pretendo servers
		resultOfficial := make([]bool, len(testOfficial))
		for x := 0; x < len(testOfficial); x++ {

			// test the endpoint
			fmt.Printf("  %s %s -> %s", utilIcons["uncertain"], testEndpoints["ninty"][testOfficial[x]], testEndpoints["official"][testOfficial[x]])

			// get the json
			res2, err3 := get(strings.Join([]string{"http://", testEndpoints["official"][testOfficial[x]], "/isthisworking"}, ""))
			// prepare a struct for the json
			var parsedRes2 isitworkingStruct

			// parse it
			err4 := json.Unmarshal([]byte(res2), &parsedRes2)
			if res2 != "" {
				if err4 != nil {
					if err3 == nil {
						fmt.Printf("\n[err] : error when parsing JSON during validating %s server\n", testOfficial[x])
						panic(err4)
					}
				}
			}

			// handle the results
			if (parsedRes2.Server == resMap[testOfficial[x]]) && (err3 == nil) && (res2 != "") {

				// why is windows like this
				consoleSequence(fmt.Sprintf("%s\n", padStrToMatchStr(fmt.Sprintf("\r  %s%s%s%s %s -> %s", code("green"), code("bold"), utilIcons["success"], code("reset"), testEndpoints["ninty"][testOfficial[x]], testEndpoints["official"][testOfficial[x]]), fmt.Sprintf("  %s %s -> %s", utilIcons["uncertain"], testEndpoints["ninty"][testOfficial[x]], testEndpoints["official"][testOfficial[x]]), " ")))
				resultOfficial[x] = true

			} else {

				// thank goodness for my shorthand function
				consoleSequence(fmt.Sprintf("%s\n", padStrToMatchStr(fmt.Sprintf("\r  %s%s%s%s %s -> %s", code("red"), code("bold"), utilIcons["failiure"], code("reset"), testEndpoints["ninty"][testOfficial[x]], testEndpoints["official"][testOfficial[x]]), fmt.Sprintf("  %s %s -> %s", utilIcons["uncertain"], testEndpoints["ninty"][testOfficial[x]], testEndpoints["official"][testOfficial[x]]), " ")))
				resultOfficial[x] = false

			}

		}

		// print out the results
		fmt.Printf("-- printing results of tests\n")
		for x := 0; x < len(result); x++ {

			// add a header saying that these are local results
			fmt.Printf("- local\n")

			// print the results
			if result[x] == true {
				fmt.Printf(" %s: success\n", test[x])
			} else {
				fmt.Printf(" %s: failiure\n", test[x])
			}

		}
		for x := 0; x < len(resultOfficial); x++ {

			// add a header saying that these are official results
			fmt.Printf("- pretendo\n")

			// print the results
			if resultOfficial[x] == true {
				fmt.Printf(" %s: success\n", testOfficial[x])
			} else {
				fmt.Printf(" %s: failiure\n", testOfficial[x])
			}

		}

		// begin generating the config
		fmt.Printf("-- generating config\n")

		// create cfgTest, cfgResult, using, and useLocal variables
		// also make the really long name variable
		var using string
		var cfgTest []string
		var cfgResult []bool
		var useLocal bool
		var doesOfficialHaveAnyWorkingEndpoints bool

		// scan the local server result list to see if any are true, since they have priority
		useLocal = false
		for x := 0; x < len(result); x++ {
			if result[x] == true {
				useLocal = true
			}
		}

		// local servers have priority
		if useLocal == true {
			using = "local"
			cfgTest = test
			cfgResult = result
		} else {

			// check this first to see if official even works
			for x := 0; x < len(resultOfficial); x++ {
				doesOfficialHaveAnyWorkingEndpoints = false
				if resultOfficial[x] == true {
					doesOfficialHaveAnyWorkingEndpoints = true
				}
			}

			if doesOfficialHaveAnyWorkingEndpoints == true {

				// set these if it does
				using = "official"
				cfgTest = testOfficial
				cfgResult = resultOfficial

			} else {

				// exit the program
				clear()
				fmt.Printf("no servers are running currently, please try again later.")
				os.Exit(0)

			}
		}

		// make a map for the config
		config = make(map[string]map[string]string)

		// make the endpoints and config a map[string]string
		config["endpoints"] = make(map[string]string)
		config["config"] = make(map[string]string)

		// apply a nice helping of all of the working endpoints to the config
		for x := 0; x < len(cfgTest); x++ {
			if cfgResult[x] == true {
				config["endpoints"][testEndpoints["ninty"][cfgTest[x]]] = testEndpoints[using][cfgTest[x]]
			}
		}

		// set some config vars
		config["config"]["decryptOutgoing"] = "true"

		// wait for them to press enter
		fmt.Printf("\npress enter to continue...\n")
		_ = input("")

		// creating a custom config
	} else if method == "2" {
		fmt.Printf(" method: custom..\n")

		// number of values in the config
		numVals := 0

		// config
		config = make(map[string]map[string]string)

		// make the endpoints and config a map[string]string
		config["endpoints"] = make(map[string]string)
		config["config"] = make(map[string]string)

		// temp vars
		var inputtedFrom string
		var inputtedTo string

		for true {
			clear()

			// reset temporary vars
			inputtedFrom = ""
			inputtedTo = ""

			// display ui
			fmt.Printf(" you have %s redirection(s) already in\n", strconv.Itoa(numVals))
			fmt.Printf(" press <Enter> on an empty line to stop\n")

			// ask for conf vals
			inputtedFrom = input("from: ")
			if inputtedFrom == "" {
				break
			}
			inputtedTo = input("to: ")
			if inputtedTo == "" {
				break
			}

			// place them in the config var
			config["endpoints"][inputtedFrom] = inputtedTo

			// update info
			numVals++

		}

		// set default config vars
		config["config"]["decryptOutgoing"] = "true"

		// loading a template
	} else if method == "3" {

		// template variable since i have to reserve it
		var tmpl string

		// ask for choice
		fmt.Printf(" method: template..\n")
		for true {
			clear()

			// show ui
			fmt.Printf("-- please select a template\n")
			fmt.Printf(" 1. local server\n")
			fmt.Printf(" 2. pretendo servers\n")
			// TODO: make official servers work
			// not adding this one until i can figure out how
			// to make ClientCertificate.cer work
			// fmt.Printf(" 3. official servers\n")

			// ask for input
			tmpl = input(": ")

			// break if it's a valid option
			if (tmpl == "1") || (tmpl == "2") {
				break
			} else {
				fmt.Printf("-> please enter 1 or 2\n")
				time.Sleep(1500 * time.Millisecond)
			}
		}

		// load the selected template into the config var
		if tmpl == "1" {
			config = localConf
		} else if tmpl == "2" {
			config = pretendoConf
		}

	}

	// not everyone wants to generate a new config
	if method != "4" {

		// prettify the JSON
		pretty, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			fmt.Printf("[err] : error while prettifying JSON\n")
			panic(err)
		}
		prettifiedJSON := string(pretty[:])

		// confirm the preferences
		var areSettingsOkay string
		for true {
			clear()

			// display the UI
			fmt.Printf("== maryo -> setup ===========================================\n")
			fmt.Printf("                                        Steps:               \n")
			fmt.Printf(" are you okay with the settings below?    intro              \n")
			fmt.Printf("                                          config creation    \n")
			fmt.Printf("                                        > confirm prefs      \n")
			fmt.Printf("                                          generate cert      \n")
			fmt.Printf("                                          profit???          \n")
			fmt.Printf("                                                             \n")
			fmt.Printf(prettifiedJSON)
			fmt.Printf("\n                                                             \n")
			fmt.Printf("-> (y|n)                                                     \n")
			fmt.Printf("=============================================================\n")
			areSettingsOkay = input(": ")

			if (areSettingsOkay == "y") || (areSettingsOkay == "n") {
				break
			} else {
				fmt.Printf("-> please enter y or n")
				time.Sleep(1500 * time.Millisecond)
			}
		}

		// check if they answered no
		if areSettingsOkay == "n" {
			os.Exit(0)
		}

		// idk what to do after this
		stringifiedConfig, err := json.Marshal(config)
		if err != nil {
			fmt.Printf("[err] : error when stringifying json")
		}
		
		// make sure the maryo folder exists
		if doesDirExist("maryo") == false {

			// make it if it doesn't
			makeDirectory("maryo")
		
		}
		
		// place it into the file
		if fileMap["config"] == "iv" {
			deleteFile("maryo/config.json")
			createFile("maryo/config.json")
			writeByteToFile("maryo/config.json", stringifiedConfig)
		} else if fileMap["config"] == "ne" {
			createFile("maryo/config.json")
			writeByteToFile("maryo/config.json", stringifiedConfig)
		} else if fileMap["config"] == "uk" {
			// detect status of config and do the
			// things to write to it.
			if doesFileExist("maryo/config.json") == true {
				deleteFile("maryo/config.json")
				createFile("maryo/config.json")
				writeByteToFile("maryo/config.json", stringifiedConfig)
			} else {
				createFile("maryo/config.json")
				writeByteToFile("maryo/config.json", stringifiedConfig)
			}
		}
	}

	// generate a https cert
	clear()

	fmt.Printf("== maryo -> setup ===========================================\n")
	fmt.Printf("                                        Steps:               \n")
	fmt.Printf(" now, it is time to generate a https      intro              \n")
	fmt.Printf(" cert to encrypt your data                config creation    \n")
	fmt.Printf(" -> press enter                           confirm prefs      \n")
	fmt.Printf("                                        > generate cert      \n")
	fmt.Printf("                                          profit???          \n")
	fmt.Printf("                                                             \n")
	fmt.Printf("                                                             \n")
	fmt.Printf("                                                             \n")
	fmt.Printf("=============================================================\n")
	_ = input("")

	// generate the certificates
	doCertGen()

	// show them the finished screen
	clear()

	// display the UI
	fmt.Printf("== maryo -> setup ===========================================\n")
	fmt.Printf("                                        Steps:               \n")
	fmt.Printf(" congratulations, you are finished        intro              \n")
	fmt.Printf(" setting up maryo!                        config creation    \n")
	fmt.Printf(" -> press enter                           confirm prefs      \n")
	fmt.Printf("                                          generate cert      \n")
	fmt.Printf("                                        > profit???          \n")
	fmt.Printf("                                                             \n")
	fmt.Printf("                                                             \n")
	fmt.Printf("                                                             \n")
	fmt.Printf("=============================================================\n")
	_ = input("")

	// display a message saying that they need to restart the program to use the new config
	clear()
	fmt.Printf("run this program again to use the new configuration\n")

}
