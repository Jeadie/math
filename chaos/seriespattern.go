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
var seriesChecks = []func([]float64) bool{isFalse, isConvergent, isFalse, isFalse, isChaotic, isTrue}
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

func isChaotic(x []float64) bool {
	return true
}

func isConvergent(x []float64) bool {
	return math.Abs(x[0] -  x[1]) < 0.0001
}
