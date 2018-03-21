/*

maryo/network.go

utilities involving the network

written by Superwhiskers, licensed under gnu gplv3.
if you want a copy, go to http://www.gnu.org/licenses/

*/

package main

import (
	// internals
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

/* net utils */

// function to download a file from a URL.
// based on https://www.github.com/thbar/golang-playground/blob/master/download-files.go
func downloadFile(args []string) {

	// declare this
	var downloadTo = ""

	// arg checking
	if len(args) != 2 {
		
		// split the url by /
		tmp := strings.Split(args[0], "/")
		
		// figure out where to download the file to
		downloadTo = tmp[len(tmp)-1]
	
	// set the download path to the 2nd arg	
	} else {
		
		// set it
		downloadTo = args[1]
		
	}

	// detect if file already exists
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	_, err = os.Stat(strings.Join([]string{dir, "/", downloadTo}, ""))
	
	// error handling
	if err == nil {
		
		// show error message
		fmt.Printf("[err] : file %s already exists.. (did you try running this program already?)\n", downloadTo)
		
		// show traceback
		panic(err)
		
	}

	// create the file
	oput, err := os.Create(downloadTo)
	
	// error handling
	if err != nil {
		
		// show error message
		fmt.Printf("[err] : error creating file %s.. (does it already exist?)\n", downloadTo)
		
		// show traceback
		panic(err)
		
	}
	
	// close the file stream
	defer oput.Close()

	// attempt to download the contents
	res, err := http.Get(args[0])
	
	// error handling
	if err != nil {
		
		// show error message
		fmt.Printf("[err] : error downloading from %s.. (is your internet working?)\n", args[0])
		
		// show traceback
		panic(err)
		
	}
	
	// close response body stream
	defer res.Body.Close()

	// copy url contents to file
	_, err = io.Copy(oput, res.Body)
	
	// error handlong
	if err != nil {
		
		// show error message
		fmt.Printf("[err] : error copying data from %s to %s.. (is %s in the working directory?)\n", args[0], downloadTo, downloadTo)
	
		// show traceback
		panic(err)
		
	}

}

// function to get data from a URL.
// based on https://www.github.com/thbar/golang-playground/blob/master/download-files.go
func get(url string) (string, error) {

	// attempt to download the contents
	res, err := http.Get(url)
	
	// error handling
	if err != nil {
		
		// return an empty string, and the error
		return "", err
		
	}
	
	// close request body stream once finished
	defer res.Body.Close()
	
	// read all data from body
	data, err := ioutil.ReadAll(res.Body)
	
	// error handling
	if err != nil {
		
		// return an empty string, and the error
		return "", err
		
	}

	// convert the bytes to a string
	ret := string(data[:])
	
	// return the request response
	return ret, nil
	
}
