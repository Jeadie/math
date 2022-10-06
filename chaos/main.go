package chaos

type HyperParams struct {
	initalIter   uint
	dy           float64
	maxSeriesLen uint
}

const (
	Stewart     string = "stewart"
	Bifurcation        = "bifurcation"
)

func Chaos(args []string) {
	if len(args) == 1 && (args[0] == "-h" || args[0] == "--help") {
		return
	}

	switch args[0] {
	case Stewart:
		StewartExample(args[1:])
	case Bifurcation:
		RunBifurcation(args[2:])
	default:
		return
	}
}
