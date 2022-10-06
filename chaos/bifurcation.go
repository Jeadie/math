package chaos

import (
	"flag"
	"fmt"
	"math"
	"strconv"
)

// Constructs data for bifurcation diagram

type BifurcationArgs struct {
	x FloatPair
	k FloatPair
	d FloatPair
}

type BifurcationResult struct {
	s    *SeriesResult
	x, k float64
}

type BifurcationInput struct {
	h    *HyperParams
	x, k float64
}

func GetBifurcationArgs(b *BifurcationArgs, h *HyperParams) *flag.FlagSet {
	fs := flag.NewFlagSet("bifurcation", flag.ExitOnError)
	x := FloatPair{0.0, 1.0}
	k := FloatPair{0.0, 4.0}
	d := FloatPair{100.0, 100.0}

	fs.Var(&b.x, "x", "Bounds of initial condition")
	fs.Var(&b.k, "k", "Bounds of function parameter")
	fs.Var(&b.d, "d", "Delta to increment x,k during sampling")

	fs.UintVar(&h.initalIter, "initial-iterations", 50, "Number of initial iterations "+
		"before considering output")
	fs.UintVar(&h.maxSeriesLen, "maxSeriesLen", 9, "Size of series to consider "+
		"periodicities and other patterns within. Patterns greater than this will be considered chaotic")

	return fs
}

func RunBifurcation(args []string) {
	if len(args) == 1 && (args[0] == "-h" || args[0] == "--help") {
		GetBifurcationArgs(&BifurcationArgs{}, &HyperParams{}).Usage()
		return
	}

	bif := &BifurcationArgs{}
	h := &HyperParams{}
	fs := GetBifurcationArgs(bif, h)
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

	work := make(chan BifurcationInput, 10)
	// Instantiate units of work
	go func(bif *BifurcationArgs, output chan BifurcationInput, h *HyperParams) {
		for i := bif.x[0]; i <= bif.x[1]; i += bif.d[0] {
			for j := bif.k[0]; j <= bif.k[1]; j += bif.d[1] {
				output <- BifurcationInput{
					h: h,
					x: i,
					k: j,
				}
			}
		}
	}(bif, work, h)

	// Process units of work
	result := make(chan *BifurcationResult, 10)
	go func(in chan BifurcationInput, r chan *BifurcationResult) {
		for i := range in {
			r <- &BifurcationResult{
				s: Run(i.x, metaRecurse(i.k), i.h.initalIter, i.h.maxSeriesLen),
				x: i.x,
				k: i.k,
			}
		}
	}(work, result)

	bif.k[1]
	grid := make([][]SeriesResult, 1)
	for bifurcationResult := range result {

	}
	// Save results to disk and output result.
	// Disk:
	//  - Values from all series instantiations
	//  - Classification of all series
	//
	// stdout: Classification of all series
}

func (p *FloatPair) String() string {
	return fmt.Sprintf("[%f, %f]", p[0], p[1])
}

func (p *FloatPair) Set(x string) error {
	f, err := strconv.ParseFloat(x, 64)
	if err != nil {
		return err
	}

	// Set both to first value to keep track of ordering
	if p[0] == p[1] {
		p[1] = f
	} else {
		p[0] = f
		p[1] = f
	}
	return err
}

type FloatPair [2]float64
