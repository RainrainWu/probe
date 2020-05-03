package example_cases

import (
	"github.com/RainrainWu/probe/jobs"
)

var SaladCase []func(*jobs.Runner) int = []func(*jobs.Runner) int {
	SaladCobb,
	SaladCaesar,
}

func SaladCobb(r *jobs.Runner) int {
	r.Rep.Pass()
	return 0
}

func SaladCaesar(r *jobs.Runner) int {
	dets := r.Rep.InitDetail("Caesar")
	dets.Append("Topping", "topping not found")
	dets.Append("Plate", "Plate is broken")
	r.Rep.Fail()
	return 0
}
