package chaos

type HyperParams struct {
	initalIter uint
	dy float64
	maxSeriesLen uint
}

func Chaos(args []string) {
	// TODO: change to `--help` on first
	if len(args) == 1 && (args[0] == "-h" || args[0] == "--help") {
		return
	}

	switch args[0] {
		case "stewart":
			StewartExample(args[1:])
		default:
			return
	}
}
