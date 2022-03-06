package chaos

import (
	"flag"
	"fmt"
	"math"
)

func GetFlagAndParams(p *StewartParams) *flag.FlagSet {
	fs := flag.NewFlagSet("stewart", flag.ExitOnError)
	fs.Float64Var(&p.x, "x", 0.42, "Parameter of the recursive formula kx**2 -1")
	fs.Float64Var(&p.k, "k", 1.24, "Parameter of the recursive formula kx**2 -1")
	fs.UintVar(&p.hyper.initalIter, "initial-iterations", 50, "Number of initial iterations before considering output")
	fs.UintVar(&p.hyper.maxSeriesLen, "maxSeriesLen", 9, "Size of series to consider periodicities and other patterns within. Patterns greater than this will be considered chaotic")
	return fs
}

type StewartParams struct {
	x float64
	k float64
	hyper *HyperParams
}

func GetParams(args []string) *StewartParams {
	param := StewartParams{hyper: &HyperParams{}}
	fs := GetFlagAndParams(&param)
	fs.Parse(args)
	return &param
}

func StewartExample(args []string) {
	// TODO: change to `--help` on first
	if len(args) == 1 && (args[0] == "-h" || args[0] == "--help") {
		GetFlagAndParams(&StewartParams{hyper: &HyperParams{}}).Usage()
		return
	}
	params := GetParams(args)

	// Make specific recursive function based on parameter, k
	metaRecurse := func(k float64) func(x float64) float64 {
		return func(x float64) float64 {
			return params.k*math.Pow(x, 2) - 1.0
		}
	}
	r := Run(params.x, metaRecurse(params.k), params.hyper.initalIter, params.hyper.maxSeriesLen)
	fmt.Println(r.pattern)
	for v := range r.v {
		fmt.Println(v)
	}
}
