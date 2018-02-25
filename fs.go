/*

maryo/fs.go

utilities involving files and the filesystem

written by Superwhiskers, licensed under gnu gplv3.
if you want a copy, go to http://www.gnu.org/licenses/

*/

package main

import (
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// check if file exists
func doesFileExist(file string) bool {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	_, err = os.Stat(strings.Join([]string { dir, "/", file }, ""))
	if err != nil {
		return false
	}
	return true

}

// create a file
func createFile(file string) {

	// detect if file already exists
	if doesFileExist(file) == true {
		fmt.Printf("[err] : %s already exists..", file)
		os.Exit(1)
	}

	// create the file
	oput, err := os.Create(file)
	if err != nil {
		fmt.Printf("[err] : error creating file %s.. (does it already exist?)\n", file)
		defer oput.Close()
		panic(err)
	}
	defer oput.Close()

}

// read a file, output data as string
func readFile(file string) string {

	// read file
	b, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("[err] : error reading file %s.. (does it exist?)\n", file)
		panic(err)
	}

	// convert byte array to string
	str := string(b)
	return str

}

// delete a file
func deleteFile(file string) {

	// delete the file
	err := os.Remove(file)
	if err != nil {
		fmt.Printf("[err] : error deleting file %s..", file)
		panic(err)
	}

}

// write to file
func writeFile(file string, data string) {

	// convert string to byte array
	bdata := []byte(data)

	// write to file
	err := ioutil.WriteFile(file, bdata, 0644)
	if err != nil {
		fmt.Printf("[err] : error writing to file %s.. (does it exist?)\n", file)
		panic(err)
	}

}

// check if file is valid JSON
func checkJSONValidity(file string) bool {

	// get JSON from file
	filedata := []byte(readFile(file))

	// this only exists because it's required to unmarshal the file
	var data map[string]interface{}

	// unmarshal the file
	err := json.Unmarshal(filedata, &data)

	// check for errors
	if err != nil {
		return false
	} else {
		return true
	}

}

// read a JSON file
func readJSONFile(file string) map[string]interface{} {

	// get json from file, and turn into byte array
	jsonObj := []byte(readFile(file))

	// initialize an interface
	var data map[string]interface{}

	// turn json into a valid golang item
	err := json.Unmarshal(jsonObj, &data)
	if err != nil {
		fmt.Printf("[err] : error converting raw JSON to valid golang item from %s.. (is this valid JSON?)\n", file)
		panic(err)
	}

	return data

}

// write to a json file
func writeJSONFile(file string, data map[string]int) {

	// turn go map into valid JSON
	fileData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("[err] : error while converting a golang map into JSON. (how did this even happen)\n")
		panic(err)
	}

	// convert to string
	sFileData := string(fileData)

	// write it to a file
	writeFile(file, sFileData)

}
