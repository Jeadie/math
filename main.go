package main

import (
	"github.com/Jeadie/math/chaos"
	"os"
)

const (
	Chaos string = "chaos"
)

func main() {
	switch os.Args[1] {
		case Chaos:
			chaos.Chaos(os.Args[2:])
	}

}
