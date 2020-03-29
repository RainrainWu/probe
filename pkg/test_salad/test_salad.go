package test_salad

import (
	"github.com/RainrainWu/probe/pkg/utils"
)

var Cases []func(*utils.Runner) int = []func(*utils.Runner) int {
	Cobb,
	Caesar,
}

func Cobb(r *utils.Runner) int {
	r.Rep.Pass()
	return 0
}

func Caesar(r *utils.Runner) int {
	dets := r.Rep.InitDetail("Caesar")
	dets.Append("Topping", "topping not found")
	dets.Append("Plate", "Plate is broken")
	r.Rep.Fail()
	return 0
}
