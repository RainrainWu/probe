package example_cases

import (
	"github.com/RainrainWu/probe/jobs"
)

var SaladCase []func(*jobs.Runner) = []func(*jobs.Runner) {
	SaladCobb,
	SaladCaesar,
}

func SaladCobb(r *jobs.Runner) {
	r.Rep.Pass()
}

func SaladCaesar(r *jobs.Runner) {
	dets := r.Rep.InitDetail("Caesar")
	dets.Append("Topping", "topping not found")
	dets.Append("Plate", "Plate is broken")
	r.Rep.Fail()
}
