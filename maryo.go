/*

maryo/maryo.go

written by superwhiskers, licensed under gnu gplv3.
if you want a copy, go to http://www.gnu.org/licenses/

*/

package main

import (
  "fmt"
  "runtime"
  "flag"
  "time"
  "os"
)

func setup(fileMap map[string]string) {

  // test data
  test := make([]string, 1)
  test[0] = "account"

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
  fmt.Printf(" proxy server for accessing the server.   display proxy info \n")
  fmt.Printf(" -> press enter                                              \n")
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
    fmt.Printf(" 2. custom                                display proxy info \n")
    fmt.Printf(" 3. template                                                 \n")
    fmt.Printf("                                                             \n")
    fmt.Printf(" -> (1|2|3)                                                  \n")
    fmt.Printf("=============================================================\n")
    method = input(": ")

    if ( method == "1" ) || ( method == "2" ) {
      break
    } else {
      fmt.Printf("-> please enter 1 or 2\n")
      time.Sleep(1500 * time.Millisecond)
    }
  }

  // show log when
  clear()
  fmt.Printf("== maryo -> setup ===========================================\n")
  fmt.Printf("                                                             \n")
  fmt.Printf(" configuring proxy..                                         \n")
  fmt.Printf(" current config status: %s\n", fileStat[fileMap["config"]])
  if method == "1" {
    fmt.Printf("-- beginning tests")
    fmt.Printf(" method: automatic..\n")
    fmt.Printf(" 1. attempting to detect endpoints running on this machine\n")

    // test for endpoints on this machine
    result := make([]bool, len(test))
    for x := 0; x < len(test); x++ {

      // test the endpoint
      fmt.Printf("  %s %s -> %s", utilIcons("uncertain"), endpointsFor("ninty", test[x]), endpointsFor("local", test[x]))
      res, err := get(endpointsFor("local", test[x]))
      if (res == "it works!") && (err == nil) {
        fmt.Printf("%s\n", padStrToMatchStr(fmt.Sprintf("\r  %s %s -> %s", utilIcons("success"), endpointsFor("ninty", test[x]), endpointsFor("local", test[x])), fmt.Sprintf("testing %s -> %s ", endpointsFor("ninty", test[x]), endpointsFor("local", test[x])), " "))
        result[x] = true
      } else {
        fmt.Printf("%s\n", padStrToMatchStr(fmt.Sprintf("\r  %s %s -> %s", utilIcons("failiure"), endpointsFor("ninty", test[x]), endpointsFor("local", test[x])), fmt.Sprintf("testing %s -> %s ", endpointsFor("ninty", test[x]), endpointsFor("local", test[x])), " "))
        result[x] = false
      }

    }

    // print out the results
    fmt.Printf("-- printing results of tests\n")
    for x := 0; x < len(test); x++ {

      // print the results
      if result[x] == true {
        fmt.Printf(" %s: success\n", test[x])
      } else {
        fmt.Printf(" %s: failiure\n", test[x])
      }

    }

  } else if method == "2" {
    fmt.Printf(" method: custom..\n")
  } else if method == "3" {
    fmt.Printf(" method: template..\n")
  }

}

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
