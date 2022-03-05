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
var SeriesChecks = []func([]float64) bool{isFalse, isConvergent, isFalse, isFalse, isChaotic, isTrue}


func isFalse(x []float64) bool { return false }
func isTrue(x []float64) bool { return true }

func isChaotic(x []float64) bool {
	return true
}

func isConvergent(x []float64) bool {
	return math.Abs(x[0] -  x[1]) < 0.0001
}
