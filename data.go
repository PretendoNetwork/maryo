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

// TODO: somewhere in here include a list containing a list of templates to add

// pretendo stock config
var pretendoConf = map[string]map[string]string{"config": map[string]string{"decryptOutgoing": "false"}, "endpoints": map[string]string{"account.nintendo.net": "account.pretendo.cc"}}

// local stock config
var localConf = map[string]map[string]string{"config": map[string]string{"decryptOutgoing": "true"}, "endpoints": map[string]string{"account.nintendo.net": "127.0.0.1:8080"}}

// test endpoints
var testEndpoints = map[string]map[string]string{"official": map[string]string{"account": "account.pretendo.cc"}, "local": map[string]string{"account": "127.0.0.1:8080"}, "ninty": map[string]string{"account": "account.nintendo.net"}}

// supposed return value for custom servers
var resMap = map[string]string{"account": "account.nintendo.net"}

// icons used for displaying results
var utilIcons = map[string]string{"success": "√", "failiure": "×", "uncertain": "-"}
