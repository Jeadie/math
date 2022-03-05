package chaos

import (
	"flag"
	"fmt"
	"math"
)

type ChaosParams struct {
	x float64
	k float64

	initalIter uint
	dy float64
	n uint
	maxSeriesLen uint
}

func GetFlagAndParams(p *ChaosParams) *flag.FlagSet {
	fs := flag.NewFlagSet("chaos", flag.ExitOnError)
	fs.Float64Var(&p.x, "x", 0.42, "Parameter of the recursive formula kx**2 -1")
	fs.Float64Var(&p.k, "k", 1.24, "Parameter of the recursive formula kx**2 -1")
	fs.UintVar(&p.initalIter, "initial-iterations", 50, "Number of initial iterations before considering output")
	fs.UintVar(&p.maxSeriesLen, "maxSeriesLen", 9, "Size of series to consider periodicities and other patterns within. Patterns greater than this will be considered chaotic")
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
	Run(params)
}

func Run(p *ChaosParams) {
	x := p.x
	k := p.k
	// Allow pattern to stabilise
	for i := uint(0); i < p.initalIter; i++ {
		x = k*math.Pow(x, 2) - 1.0
	}

	state := ConstructIterationState(p)
	newX := x
	for i := state.ttl; i >0; i-- {
		newX = k*math.Pow(x, 2) - 1.0
		state.AddIteration(newX)
		x = newX
	}
	state.UpdatePattern()
	fmt.Println(state.pattern)
}

func ConstructIterationState(p *ChaosParams) *IterationState{
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

