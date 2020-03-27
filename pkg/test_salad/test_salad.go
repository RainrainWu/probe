package test_salad

import (
	"github.com/RainrainWu/probe/pkg/utils"
)

var Cases []func(*utils.Runner) int = []func(*utils.Runner) int {
	Cobb,
	Caesar,
	Caprese,
	ColeSlaw,
}

func Cobb(r *utils.Runner) int {
	r.Debug("Make Cobb")
	return 0
}

func Caesar(r *utils.Runner) int {
	r.Debug("Make Caesar")
	return 0
}

func Caprese(r *utils.Runner) int {
	r.Debug("Make Caprese")
	return 0
}

func ColeSlaw(r *utils.Runner) int {
	r.Debug("Make Coleslaw")
	return 0
}
