package main

import (
	"fmt"
	"github.com/Jeadie/math/chaos"
	"os"
)

const (
	Chaos string = "chaos"
)

func ListCommands() {
	fmt.Println("Available commands: ")
	fmt.Println(fmt.Sprintf("    %s: run iteration experiment of kx**2 -1", Chaos))
}

func main() {
	if len(os.Args) < 2 || os.Args[1] == "--help" || os.Args[1] == "-h" {
		ListCommands()
		return
	}

	switch os.Args[1] {
		case Chaos:
			chaos.Chaos(os.Args[2:])

		default:
			ListCommands()
			return
	}

}
