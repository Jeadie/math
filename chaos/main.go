package chaos

import "flag"

func GetFlagSet(name string) *flag.FlagSet {
	return flag.NewFlagSet(name, flag.ExitOnError)
}
func Chaos() {
	flag.Args()
}
