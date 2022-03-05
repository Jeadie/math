package main

import (
	"flag"
	"github.com/Jeadie/math/chaos"
)

const (
	Chaos string = "chaos"
)

func main() {

	switch flag.Arg(0) {
		case Chaos:
			chaos.Chaos(flag.NewFlagSet(Chaos, flag.ExitOnError))
	}

}
