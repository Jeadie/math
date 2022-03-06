package chaos

// Run attempts to find the pattern of a recursive function applied to an initial value.
func Run(x float64, recurse func(float64) float64, initialIterations, maxPeriodicity uint) *SeriesResult {
	path := make(chan float64, initialIterations + maxPeriodicity + 1)
	defer close(path)

	path <- x
	// Allow pattern to stabilise
	for i := uint(0); i < initialIterations; i++ {
		x = recurse(x)
		path <- x
	}

	state := &IterationState{
		previousN: make([]float64, maxPeriodicity),
		k:         0,
	}

	newX := x
	for i := maxPeriodicity; i > 0; i-- {
		newX = recurse(x)
		state.AddIteration(newX)
		x = newX
		path <- x
	}
	return &SeriesResult{
		v:       path,
		pattern: FindPattern(state.previousN),
	}
}

type SeriesResult struct {
	v chan float64
	pattern SeriesPattern
}

type IterationState struct {
	previousN []float64
	k int
}

// AddIteration adds the value to the previous list.
func (s *IterationState) AddIteration(x float64) {
	// Add x into previous N
	s.previousN[s.k] = x
	s.k++
}