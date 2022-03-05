package chaos

import (
	"flag"
	"fmt"
	"math"
)

type ChaosParams struct {
	x float64
	k float64
	hyper *HyperParams
}

type HyperParams struct {
	initalIter uint
	dy float64
	maxSeriesLen uint
}

func GetFlagAndParams(p *ChaosParams) *flag.FlagSet {
	fs := flag.NewFlagSet("chaos", flag.ExitOnError)
	fs.Float64Var(&p.x, "x", 0.42, "Parameter of the recursive formula kx**2 -1")
	fs.Float64Var(&p.k, "k", 1.24, "Parameter of the recursive formula kx**2 -1")
	fs.UintVar(&p.hyper.initalIter, "initial-iterations", 50, "Number of initial iterations before considering output")
	fs.UintVar(&p.hyper.maxSeriesLen, "maxSeriesLen", 9, "Size of series to consider periodicities and other patterns within. Patterns greater than this will be considered chaotic")
	fs.Float64Var(&p.hyper.dy, "delta", 0.001, "delta between iterations before halting (after initialIter), if not terminated for cyclic nature")
	return fs
}

func GetParams(args []string) *ChaosParams {
	param := ChaosParams{hyper: &HyperParams{}}
	fs := GetFlagAndParams(&param)
	fs.Parse(args)
	return &param
}

func Help() {
	GetFlagAndParams(&ChaosParams{hyper: &HyperParams{}}).Usage()
}

func Chaos(args []string) {
	// TODO: change to `--help` on first
	if len(args) == 1 && (args[0] == "-h" || args[0] == "--help") {
		Help()
		return
	}
	params := GetParams(args)

	// Make specific recursive function based on parameter, k
	metaRecurse := func(k float64) func(x float64) float64 {
		return func(x float64) float64 {
			return params.k*math.Pow(x, 2) - 1.0
		}
	}
	Run(params.x, metaRecurse(params.k), params.hyper)
}

func Run(x float64, recurse func(float64) float64, p *HyperParams) *IterationState{

	// Allow pattern to stabilise
	for i := uint(0); i < p.initalIter; i++ {
		x = recurse(x)
	}

	state := ConstructIterationState(p)
	newX := x
	for i := state.ttl; i >0; i-- {
		newX = recurse(x)
		state.AddIteration(newX)
		x = newX
	}
	state.UpdatePattern()
	fmt.Println(state.pattern)
	return state
}

func ConstructIterationState(p *HyperParams) *IterationState{
	previousN := make([]float64, p.maxSeriesLen)
	return &IterationState{
		dy:        p.dy,
		previousN: previousN,
		k:         0,
		n:         int(p.maxSeriesLen),
		ttl: p.maxSeriesLen,
	}
}

type IterationState struct {
	pattern SeriesPattern
	dy      float64
	ttl uint

	// Basic ring, of N most previous elements
	previousN []float64
	k int
	n int
}

// AddIteration adds the value to the previous list, and checks if the iteration pattern has changed.
func (s *IterationState) AddIteration(x float64) {
	// Add x into previous N
	s.previousN[s.k] = x
	s.k++
}

func (s *IterationState) UpdatePattern() {
	// Check for pattern changes, exit on first
	for i, fn := range SeriesChecks {
		if fn(s.previousN) {
			s.pattern = SeriesPattern(i)
			return
		}
	}
}

