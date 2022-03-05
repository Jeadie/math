package chaos

import (
	"flag"
	"fmt"
)

type ChaosParams struct {
	x float64
	k float64

	initalIter uint // Number of initial iterations before considering output
	dy float64 // delta between iterations before halting (after initialIter), if not terminated for cyclic nature
}

func GetFlagAndParams(p *ChaosParams) *flag.FlagSet {
	fs := flag.NewFlagSet("chaos", flag.ExitOnError)
	fs.Float64Var(&p.x, "x", 0.42, "Parameter of the recursive formula kx**2 -1")
	fs.Float64Var(&p.k, "k", 0.24, "Parameter of the recursive formula kx**2 -1")
	fs.UintVar(&p.initalIter, "initial-iterations", 20, "Number of initial iterations before considering output")
	fs.Float64Var(&p.dy, "delta", 0.001, "delta between iterations before halting (after initialIter), if not terminated for cyclic nature")
	return fs
}

func GetParams(args []string) *ChaosParams {
	param := ChaosParams{}
	fs := GetFlagAndParams(&param)
	fs.Parse(args)
	return &param
}

func Help() {
	GetFlagAndParams(&ChaosParams{}).Usage()
}

func Chaos(args []string) {
	if len(args) == 0 {
		Help()
		return
	}
	params := GetParams(args)
	fmt.Println("CHAOSSS", params)
}
