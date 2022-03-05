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
	maxPeriod uint
}

func GetFlagAndParams(p *ChaosParams) *flag.FlagSet {
	fs := flag.NewFlagSet("chaos", flag.ExitOnError)
	fs.Float64Var(&p.x, "x", 0.42, "Parameter of the recursive formula kx**2 -1")
	fs.Float64Var(&p.k, "k", 1.24, "Parameter of the recursive formula kx**2 -1")
	fs.UintVar(&p.initalIter, "initial-iterations", 50, "Number of initial iterations before considering output")
	fs.UintVar(&p.maxPeriod, "maxPeriod", 9, "Max size of perioidicity under consideration. Periodicity greater than this will be considered chaotic")
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
	state := ConstructIterationState(p)
	state.AddIteration(x)

	for i := uint(0); i < p.initalIter; i++ {
		x = k*math.Pow(x, 2) - 1.0
		state.AddIteration(x)
	}

	newX := x
	for !state.isFinished() {
		newX = k*math.Pow(x, 2) - 1.0
		state.AddIteration(newX)
		state.ttl -=1
		x = newX
		fmt.Println(state.ttl, state.previousN)
	}
	fmt.Println(state.status, state.previousN)
}

// StateChecks is list of functions where the i'th State check returns true iff status is IterationStatus(i)
var StateChecks = []func(*IterationState) bool{isFalse, isConvergent, isFalse, isFalse, isChaotic, isTrue}
type IterationStatus int64
const (
	Divergent IterationStatus = iota
	Convergent
	Periodic
	Intermittency
	Chaotic
	Incomplete
)

func isFalse(s *IterationState) bool { return false }
func isTrue(s *IterationState) bool { return true }

func isChaotic(s *IterationState) bool {
	return s.ttl <= 0
}

func isConvergent(s *IterationState) bool {
	if s.k == 0 {
		return math.Abs(s.previousN[s.k] - s.previousN[len(s.previousN)-1]) < s.dy
	}
	return math.Abs(s.previousN[s.k] - s.previousN[s.k-1]) < s.dy
}


func ConstructIterationState(p *ChaosParams) *IterationState{
	previousN := make([]float64, p.maxPeriod)
	return &IterationState{
		status:    Incomplete,
		dy:        p.dy,
		previousN: previousN,
		k:         0,
		n:         int(p.maxPeriod),
		ttl: p.maxPeriod,
	}
}

type IterationState struct {
	status IterationStatus
	dy float64
	ttl uint

	// Basic ring, of N most previous elements
	previousN []float64
	k int
	n int
}

func (s *IterationState) isFinished() bool {
	return s.status != Incomplete
}

// AddIteration adds the value to the previous list, and checks if the iteration status has changed.
func (s *IterationState) AddIteration(x float64) {
	// Add x into previous N
	s.previousN[s.k] = x
	s.k++
	if s.k >= s.n {
		s.k -= s.n
	}

	// Check for status changes, exit on first
	for i, fn := range StateChecks {
		if fn(s) {
			s.status = IterationStatus(i)
			break
		}
	}

}

