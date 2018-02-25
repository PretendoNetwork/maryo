/*

maryo/data.go

written by superwhiskers, licensed under gnu gplv3.
if you want a copy, go to http://www.gnu.org/licenses/

*/

package main

func endpointsFor(kind string, subdomain string) string {

  // endpoint map
  testEndpoints := make(map[string]map[string]string)

  // official
  testEndpoints["official"] = make(map[string]string)
  testEndpoints["official"]["account"] = "http://account.pretendo.cc/v1/api/isthisworking"

  // local (with testing_env set)
  testEndpoints["local"] = make(map[string]string)
  testEndpoints["local"]["account"] = "http://127.0.0.1/v1/api/isthisworking"

  // custom (if you modified the HOSTS file or if it is on another service)
  testEndpoints["custom"] = make(map[string]string)
  testEndpoints["custom"]["protocol"] = "http://"
  testEndpoints["custom"]["account"] = "/v1/api/isthisworking"

  // official nintendo ones
  testEndpoints["ninty"] = make(map[string]string)
  testEndpoints["ninty"]["account"] = "https://account.nintendo.net"

  // return the endpoint
  return testEndpoints[kind][subdomain]

}

func utilIcons(kind string) string {

  // char map
  iconMap := make(map[string]string)

  // place the icons in
  iconMap["success"] = "✔"
  iconMap["failiure"] = "✖"
  iconMap["uncertain"] = "-"

  // return the string
  return iconMap[kind]

}
