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

func GetParams(args []string) *ChaosParams {
	fs := flag.NewFlagSet("chaos", flag.ExitOnError)
	fs.Parse(args)
	param := ChaosParams{}
	x, err := strconv.ParseFloat(fs.Arg(0), 64)
	if err != nil {
		fmt.Println(err)
		return &param
	}
	param.x = x
	param.k, err = strconv.ParseFloat(fs.Arg(1), 64)
	if err != nil {fmt.Println(err)}
	return &param
}

func Chaos(args []string) {
	params := GetParams(args)
	fmt.Println("CHAOSSS", params)
}
