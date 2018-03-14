/*

maryo/shell.go

a small collection of terminal utilities

written by Superwhiskers, licensed under gnu gplv3.
if you want a copy, go to http://www.gnu.org/licenses/

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)


/* terminal utils */

// function to clear the screen
func clear() {

	// detect if the host OS is Windows
	if runtime.GOOS == "windows" {

		// if it is, then use the native `cls` command to
		// clear the screen
		cmd := exec.Command("cmd", "/c", "cls")

		// because of the way things work, we have to pipe output
		// directly to the main stdout
		cmd.Stdout = os.Stdout

		// then run the command
		err := cmd.Run()

		// handle errors correctly
		if err != nil {
			fmt.Printf("[err] : error while executing cls. (report this issue)\n")
			panic(err)
		}

	} else {

		// if it is some ANSI escape code respecive OS,
		// use the clear code
		fmt.Printf("\033[2J\033[;H")

	}
}

// trick Terminal.app to respect ANSI color codes
func ansiTrick() {

	// run the export command to set the terminal type to
	// an escape code respective term
	// (i'm not even sure if this works)
	cmd := exec.Command("export", "TERM=xterm")

	// run the command
	err := cmd.Run()

	// handle errors correctly
	if err != nil {
		fmt.Printf("[err] : error while executing export. (isn't that a shell builtin?)\n")
		panic(err)
	}

}

// get terminal input
// (for the amount of times i have to get term input, i need this)
func input(prompt string) string {

	// print a prompt
	fmt.Printf(prompt)

	// create a scanner to get user input
	scanner := bufio.NewScanner(os.Stdin)

	// run the scanner
	scanner.Scan()

	// return the user-supplied input
	return scanner.Text()

}

// shorthand for len([]rune(x))
func length(x string) int { return len([]rune(x)); }

// pad string to match the length of another string
func padStrToMatchStr(pad string, match string, padWith string) string {

	// make sure the character to pad it is only 1 character
	if length(padWith) != 1 {

		// throw an error if it isn't
		fmt.Printf("[err] : '%s' is not 1 character long", padWith)
		os.Exit(1)

	}

	// pad the string to match the length of the other
	for x := 0; x < length(match); x++ {
		pad += padWith
	}

	// return the padded string
	return pad

}

// is it windows
func isWindows() bool { return (runtime.GOOS == "windows"); }

/* give terminal style */

// set terminal title
func ttitle(title string) { fmt.Print(strings.Join([]string {"\033]0;",title,"\007"}, "")); }

// output a formatted escape code for 8/16 bit color
func tcolor(cid int) string { return strings.Join([]string {"\033[",string(cid),"m"}, ""); }

// terminal color codes
func code(index string) string {

	// map to store terminal codes
	var termCodes map[string]string
	var prefix string
	termCodes = make(map[string]string)

	// fix colors on Terminal.app
	if runtime.GOOS == "darwin" { ansiTrick(); }

	// fix prefix for windows
	if isWindows() { prefix = "\x1b"; } else { prefix = "\033"; }

	// codes go here

	// style
	termCodes["bold"] = "[1m"
	termCodes["reset"] = "[0m"
	termCodes["underline"] = "[4m"
	termCodes["dim"] = "[2m"
	termCodes["invert"] = "[7m"
	termCodes["hide"] = "[8m"

	// colors
	termCodes["grey"] = "[90m"
	termCodes["red"] = "[91m"
	termCodes["green"] = "[92m"
	termCodes["yellow"] = "[93m"
	termCodes["blue"] = "[94m"
	termCodes["magenta"] = "[95m"
	termCodes["cyan"] = "[96m"
	termCodes["white"] = "[97m"

	// output the correct terminal color code
	return strings.Join([]string {prefix,termCodes[index]}, "")

}
