package chaos

import (
	"flag"
	"fmt"
	"math"
)

// This command runs an example recursive function found in Ian Stewart's `Does God Play Dice`, ISBN13 9780631168478.
// Stewart poses the recursive polynomial as an example of a simple determinisitc system that produces chaotic results.
//	x_{i+1} = kx_i**2 -1

func GetChaosStewartFlags(p *StewartParams) *flag.FlagSet {
	fs := flag.NewFlagSet(Stewart, flag.ExitOnError)
	fs.Float64Var(&p.x, "x", 0.42, "Parameter of the recursive formula kx**2 -1")
	fs.Float64Var(&p.k, "k", 1.24, "Parameter of the recursive formula kx**2 -1")
	fs.UintVar(&p.hyper.initalIter, "initial-iterations", 50, "Number of initial iterations "+
		"before considering output")
	fs.UintVar(&p.hyper.maxSeriesLen, "maxSeriesLen", 9, "Size of series to consider "+
		"periodicities and other patterns within. Patterns greater than this will be considered chaotic")
	return fs
}

type StewartParams struct {
	x     float64
	k     float64
	hyper *HyperParams
}

func StewartExample(args []string) {
	if len(args) == 1 && (args[0] == "-h" || args[0] == "--help") {
		GetChaosStewartFlags(&StewartParams{hyper: &HyperParams{}}).Usage()
		return
	}

	params := StewartParams{hyper: &HyperParams{}}
	fs := GetChaosStewartFlags(&params)
	err := fs.Parse(args)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Make specific recursive function based on parameter, k
	metaRecurse := func(k float64) func(x float64) float64 {
		return func(x float64) float64 {
			return k*math.Pow(x, 2) - 1.0
		}
	}
	r := Run(params.x, metaRecurse(params.k), params.hyper.initalIter, params.hyper.maxSeriesLen)
	fmt.Printf("kx**2-1 with parameters k=%f, x_0=%f is %s.\n", params.k, params.x, r.pattern.ToString())
}
