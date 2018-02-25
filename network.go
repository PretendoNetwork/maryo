/*

maryo/network.go

utilities involving the network

written by Superwhiskers, licensed under gnu gplv3.
if you want a copy, go to http://www.gnu.org/licenses/

*/

package main

import (
	"io"
	"net/http"
	"path/filepath"
	"strings"
	"fmt"
	"os"
	"io/ioutil"
	"bytes"
)

/* net utils */

// function to download a file from a URL.
// based on https://www.github.com/thbar/golang-playground/blob/master/download-files.go
func downloadFile(args []string) {

	// declare this
	var downloadTo = ""

	// arg checking
	if len(args) != 2 {
		tmp := strings.Split(args[0], "/")
		downloadTo = tmp[len(tmp)-1]
	} else {
		downloadTo = args[1]
	}

	// detect if file already exists
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	_, err = os.Stat(strings.Join([]string { dir, "/", downloadTo }, ""))
	if err == nil {
		fmt.Printf("[err] : File %s already exists.. (Did you try running this program already?)\n", downloadTo)
		panic(err)
	}

	// create the file
	oput, err := os.Create(downloadTo)
	if err != nil {
		fmt.Printf("[err] : Error creating file %s.. (Does it already exist?)\n", downloadTo)
		panic(err)
	}
	defer oput.Close()

	// attempt to download the contents
	res, err := http.Get(args[0])
	if err != nil {
		fmt.Printf("[err] : Error downloading from %s.. (Is your internet working?)\n", args[0])
		panic(err)
	}
	defer res.Body.Close()

	// copy url contents to file
	bytes, err := io.Copy(oput, res.Body)
	if err != nil {
		fmt.Printf("[err] : Error copying data from %s to %s.. (Is %s in the working directory?)\n", args[0], downloadTo, downloadTo)
		panic(err)
	}

	fmt.Printf("Successfully copied %s bytes from %s to %s\n", bytes, args[0], downloadTo)
}


// function to get data from a URL.
// based on https://www.github.com/thbar/golang-playground/blob/master/download-files.go
func get(url string) string {

	// attempt to download the contents
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("[err] : Error downloading from %s.. (Is your internet working?)\n", url)
		panic(err)
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("[err] : Error reading from %s.. (Is your internet working?)\n", url)
		panic(err)
	}

	// convert the bytes to a string
	n := bytes.IndexByte(data, 0)
	ret := string(data[:n])

	return ret
}
