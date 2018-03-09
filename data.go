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

// pretendo stock config
var pretendoConf = map[string]string{"account.nintendo.net": "http://account.pretendo.cc"}

// local stock config
var localConf = map[string]string{"account.nintendo.net": "http://127.0.0.1:8080"}

// test endpoints
var testEndpoints = map[string]map[string]string{"official": map[string]string{"account": "http://account.pretendo.cc"}, "local": map[string]string{"account": "http://127.0.0.1:8080"}, "ninty": map[string]string{"account": "https://account.nintendo.net"}}

// supposed return value for custom servers
var resMap = map[string]string{"account": "account.nintendo.net"}

// icons used for displaying results
var utilIcons = map[string]string{"success": "√", "failiure": "×", "uncertain": "-"}
