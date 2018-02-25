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
	"github.com/shiena/ansicolor"
)


/* terminal utils */

// function to clear the screen
func clear() { if runtime.GOOS == "windows" { cmd := exec.Command("cls"); err := cmd.Run(); if err != nil { fmt.Printf("[err] : Error while executing cls. (System32 is damaged)\n"); os.Exit(1); }; } else { fmt.Printf("\033[2J\033[;H"); }; }

// trick Terminal.app to respect ANSI color codes
func ansiTrick() { cmd := exec.Command("export", "TERM=xterm"); err := cmd.Run(); if err != nil { fmt.Printf("[err] : Error while executing export. (Isn't that a shell builtin?)\n"); os.Exit(1); }; }

// get terminal input
func input(prompt string) string { fmt.Printf(prompt); scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); return scanner.Text(); }

// -- give terminal style --

// set terminal title
func ttitle(title string) { fmt.Print(strings.Join([]string {"\033]0;",title,"\007"}, "")); }

// output a formatted escape code for 8/16 bit color
func tcolor(cid int) string { return strings.Join([]string {"\033[",string(cid),"m"}, ""); }

// terminal color codes
func printColor(index string, text string) {

	// map to store terminal codes
	var termCodes map[string]string
	var prefix string
	termCodes = make(map[string]string)

	// fix colors on Terminal.app
	if runtime.GOOS == "darwin" { ansiTrick(); }

	// fix prefix for windows
	if runtime.GOOS == "windows" { prefix = "\x1b"; } else { prefix = "\033"; }

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

	// join the prefix and the code
	if runtime.GOOS == "windows" {
		w := ansicolor.NewAnsiColorWriter(os.Stdout);
		fmt.Fprintf(w, strings.Join([]string {prefix,termCodes[index],text,prefix,termCodes["reset"]}, ""))
	} else {
		fmt.Printf(strings.Join([]string {prefix,termCodes[index],text,prefix,termCodes["reset"]}, ""))
	}

}
