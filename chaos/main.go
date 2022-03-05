package chaos

import (
	"flag"
	"fmt"
	"strconv"
)

type ChaosParams struct {
	x float64
	k float64
}

func GetParms(fs *flag.FlagSet) *ChaosParams {
	param := ChaosParams{}

	x, err := strconv.ParseFloat(fs.Arg(0), 64)
	if err != nil {
		param.x = x
		return &param
	}
	param.k, err = strconv.ParseFloat(fs.Arg(1), 64)
	return &param
}

func Chaos(fs *flag.FlagSet) {
	params := GetParms(fs)
	fmt.Println("CHAOSSS", params)
}
