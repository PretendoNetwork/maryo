package main

import (
	"fmt"
	"os"
	
	"github.com/shiena/ansicolor"
)

func printTitle() {
	colorizer := ansicolor.NewAnsiColorWriter(os.Stdout)

	fmt.Printf("                                     ▄▄██████▐██████▄,\n")
	fmt.Printf("                                   ▄██▀▀░░░▄▄░▄▄░░░▀▀██,\n")
	fmt.Printf("                                  █▀░▄██████▀░███████▄░▀▄\n")
	fmt.Printf("                                 ▐░▄███████░░░░▀███████▄░\n")
	fmt.Printf("                                 ░██████▀░▒░░░░░░▀██████▌▒\n")
	fmt.Printf("                                ┌▐████▀▄▄▄▓░░░░░█▄▄▄█████▒\n")
	fmt.Printf("                                ]██▀░▒░░░░▒░░░░░▒░░░▒░░▀█▌\n")
	fmt.Printf("                                ▒█▌░░░░░░░░░░░░░░░░░░░░░██,\n")
	fmt.Printf("                               ∩░▓▒MN▓███▄░░░░░░▄████▓M░▓▓░┐\n")
	fmt.Printf("                               ▒▒▓▌░░░▓█▓░▀▒░░░░▀▓██▓░░▒▓▒▒▒\n")
	fmt.Printf("                               ▒▒▓▌░░░░░░░░░░▒░░░░░░░░M▒▓▒▒┘\n")
	fmt.Printf("                                ╙▀╙░░░░░░░░░╢▒░░░░░░░░░▒╙▓\n")
	fmt.Printf("                                   ╙░░░░░░░░║▒░░░░░░░░░  ▓\n")
	fmt.Printf("                                    ╙░░░░▒▒░░░░░░░░░░▒` Æ`\n")
	fmt.Printf("                                      ▒░░░╬▓▓@▓▓╣▒░░▒╓M`\n")
	fmt.Printf("                                       ]▒░░░╙╜▒░▄█▓▓▀\n")
	fmt.Printf("                                        ║▒░░░░░░░▒╜\n")
	fmt.Printf("                                         ▒▒▒▒▒▒▒▒▒\n")
	fmt.Printf("                                        ,▒▒▒▒▒▒▒▒▒\n")
	fmt.Printf("                                      ╓╣▒▒▒▒▒▒▒▒▒▒╣@\n")
	fmt.Printf("                                  ,╓▄╣▒▒▒▒▒▒▒▒▒▒▒▒▒▒╢▄▄,\n")
	fmt.Printf("                      ¿░░¿,▄▄▄▄██▓▓▓▓▌▒▒╣▒▒▒▒▒▒▒▒▒▒▒▒▓▓▓▓▓█▓▄▄▄╖,r░░,\n")
	fmt.Printf("                    ,░░░░░░╙▓▓▓▓▓▓▓▓▓▓▓▒▒▒▒▒▒▒▒▒▒╢▒▒▓▓▓▓▓▓▓▓▓▓█░░░░░░░\n")
	fmt.Printf("                    ░░░░░░░░▐▓▓▓▓▓▓▓▓▓▓▓▓╢╢▒▒▒▒▒╢╫▓▓▓▓▓▓▓▓▓▓▓▓░░░░░░░░░\n")
	fmt.Printf("                   ░░░░░░░░░░▓▓▓▓▓▓▓▓▓▓▓▓▌░░╙╜╙░░▓▓▓▓▀║   ,▀▓▌░░░░░░░░░\n")
	fmt.Printf("                   ░░░░░░░░░░░▓▓▓▓▓▓▓▓▓▓▓▓░░░░░░▐▓▓▓ ╓▓▓▄▓▓▌ ▓W░░░░░░░░\n")
	fmt.Printf("                   ░░░░░░░░░░░▐▓▓▓▓▓▓▓▓▓▓▓█░░░░▐▓▓▓▓╓▓▌ ▀` ▓▓▐▌░░░░░░░░\n")
	fmt.Printf("                    ░░░░░░░░░░░▓▓▓▓▓▓▓▓▓▓▓▓▓░░╓▓▓▓▓▓▓▀     ╙║▓▌░░░░░░░░\n")
	fmt.Printf("                    ░░░░░░░░░░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓░█▓▓▓▓▓█▓▄▄w╓▄▓▓▀░░░░░░░░░\n")
	fmt.Printf("                    ░░░░░░░░░░░]▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓█▓█░░░░░░░░░░░\n")
	fmt.Printf("                               __  __                   ____  \n")
	fmt.Printf("                              |  \\/  |                 / __ \\ \n")
	fmt.Printf("                              | \\  / | __ _ _ __ _   _| |  | |\n")
	fmt.Printf("                              | |\\/| |/ _` | '__| | | | |  | |\n")
	fmt.Printf("                              | |  | | (_| | |  | |_| | |__| |\n")
	fmt.Printf("                              |_|  |_|\\__,_|_|   \\__, |\\____/ \n")
	fmt.Printf("                                                  __/ |       \n")
	fmt.Printf("                                                 |___/        \n")
	fmt.Fprintf(colorizer, "%s╔════════════════════════════════════════════════════════════════════════════════════════╗\n", "\x1b[91m")
	fmt.Fprintf(colorizer, "%s‖                                  %sPretendo %sServer Satus                                 %s‖\n", "\x1b[91m", "\x1b[35m", "\x1b[96m", "\x1b[91m")
	fmt.Fprintf(colorizer, "%s‖                                                                                        ‖\n", "\x1b[91m")
	fmt.Fprintf(colorizer, "%s‖  %saccount.pretendo.cc %s- %sOnline √                                                        %s‖\n", "\x1b[91m", "\x1b[35m", "\x1b[37m", "\x1b[32m", "\x1b[91m")
	fmt.Fprintf(colorizer, "%s‖  %sendpoint1.pretendo.cc %s- %sOffline ×                                                     %s‖\n", "\x1b[91m", "\x1b[35m", "\x1b[37m", "\x1b[91m", "\x1b[91m")
	fmt.Fprintf(colorizer, "%s‖  %sendpoint2.pretendo.cc %s- %sOffline ×                                                     %s‖\n", "\x1b[91m", "\x1b[35m", "\x1b[37m", "\x1b[91m", "\x1b[91m")
	fmt.Fprintf(colorizer, "%s╚════════════════════════════════════════════════════════════════════════════════════════╝\n", "\x1b[91m")
	
	fmt.Fprintf(colorizer, "%s\n", "\x1b[37m")
}

func main() {
	printTitle()
}