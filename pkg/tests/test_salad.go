package tests

import (
	"github.com/RainrainWu/probe/pkg/utils"
)

var SaladCase []func(*utils.Runner) int = []func(*utils.Runner) int {
	SaladCobb,
	SaladCaesar,
}

func SaladCobb(r *utils.Runner) int {
	r.Rep.Pass()
	return 0
}

func SaladCaesar(r *utils.Runner) int {
	dets := r.Rep.InitDetail("Caesar")
	dets.Append("Topping", "topping not found")
	dets.Append("Plate", "Plate is broken")
	r.Rep.Fail()
	return 0
}
