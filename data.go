/*

maryo/data.go

written by superwhiskers, licensed under gnu gplv3.
if you want a copy, go to http://www.gnu.org/licenses/

*/

package main

// struct for the isitworking endpoint
type isitworkingStruct struct {
  Server string `json:"server"`
}

type 
// 
func endpointsFor(kind string, subdomain string) string {

  // endpoint map
  testEndpoints := make(map[string]map[string]string)

  // official
  testEndpoints["official"] = make(map[string]string)
  testEndpoints["official"]["account"] = "http://account.pretendo.cc"

  // local (with testing_env set)
  testEndpoints["local"] = make(map[string]string)
  testEndpoints["local"]["account"] = "http://127.0.0.1:8080"

  // custom (if you modified the HOSTS file or if it is on another service)
  testEndpoints["custom"] = make(map[string]string)
  testEndpoints["custom"]["protocol"] = "http://"

  // official nintendo ones
  testEndpoints["ninty"] = make(map[string]string)
  testEndpoints["ninty"]["account"] = "account.nintendo.net"

  // return the endpoint
  return testEndpoints[kind][subdomain]

}

func serverResFor(subdomain string) string {

  // response map
  resMap := make(map[string]string)

  // there is currently only one
  resMap["account"] = "account.nintendo.net"

  // return the endpoint
  return resMap[subdomain]

}

func utilIcons(kind string) string {

  // char map
  iconMap := make(map[string]string)

  // place the icons in
  iconMap["success"] = "√"
  iconMap["failiure"] = "×"
  iconMap["uncertain"] = "-"

  // return the string
  return iconMap[kind]

}
