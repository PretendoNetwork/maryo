/*

maryo/maryo.go

written by superwhiskers, licensed under gnu gplv3.
if you want a copy, go to http://www.gnu.org/licenses/

*/

package main

import (
  //"github.com/cssivision/reverseproxy"
  //"strings"
  "fmt"
  "runtime"
)

func main() {
  clear()
  fmt.Printf("-- log\n")

  // i might include some os-specific code
  fmt.Printf("loading 0%%: detecting os")
  fmt.Printf("%s\n", padStrToMatchStr(fmt.Sprintf("\ros: %s", runtime.GOOS), "loading 0%%: detecting os", " "))

  // currently it checks for nothing, but checking for the following is planned:
  // config.json -- if nonexistent, it follows the user's instruction to create one, or use a builtin copy
  fmt.Printf("loading 10%%: checking files.")
  if doesFileExist("config.json") == false {
    fmt.Printf("%s\n", padStrToMatchStr("\rconfig: not found", "loading 10%%: checking files.", " "))
  } else {
    fmt.Printf("%s\n", padStrToMatchStr("\rconfig: found", "loading 10%%: checking files.", " "))
  }

  fmt.Printf("loaded..")
}
