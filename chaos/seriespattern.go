package chaos

import "math"


type SeriesPattern int64
const (
	Divergent SeriesPattern = iota
	Convergent
	Periodic
	Intermittency
	Chaotic
)
// SeriesChecks is list of functions where the i'th State check returns true iff pattern is SeriesPattern(i)
// Chaotic must be last check. Chaotic if no other pattern found.
var seriesChecks = []func([]float64) bool{isDivergent, isConvergent, isPeriodic, isFalse, isTrue}
var patternNames = []string{"Divergent", "Convergent", "Periodic", "Intermittency", "Chaotic"}

func FindPattern(series []float64) SeriesPattern {
	// Check for pattern changes, exit on first
	for i, fn := range seriesChecks {
		if fn(series) {
			return SeriesPattern(i)
		}
	}
	return Chaotic
}

func (p SeriesPattern) ToString() string {
	return patternNames[p]
}

func isFalse(x []float64) bool { return false }
func isTrue(x []float64) bool { return true }

func isPeriodic(x []float64) bool {

	// map value to occurrences
	f := make(map[float64][]int)
	for i, v := range x {
		l, exist := f[v]
		if exist {
			l = append(l, i)
		} else {
			f[v] = []int{i}
		}
	}

	// First element must recur
	occ := f[x[0]]
	if len(occ) < 2 {
		return false
	}

	// Try all periodicities, smallest first, from first element
	for _, p := range occ[1:] {
		if hasPeriod(x, p) {
			return true
		}
	}
	return false
}

func hasPeriod(x []float64, p int) bool {
	for i := range x[:len(x)-p] {
		if x[i] != x[i+p] {
			return false
		}
	}
	return true
}

func isDivergent(x []float64) bool {
	nInf := math.Inf(-1)
	pInf := math.Inf(+1)

	for _, v := range x {
		if v == nInf || v == pInf {
			return true
		}
	}
	return false
}

func isConvergent(x []float64) bool {
	y := x[0]
	for _, v := range x {
		if math.Abs(y-v) > 0.0001 {
			return false
		}
	}
	return true
}
