package example_cases

import (
	"github.com/RainrainWu/probe/jobs"
)

var CoffeeCase []func(*jobs.Runner) = []func(*jobs.Runner) {
	CoffeeAmericano,
	CoffeeEspresso,
}

func CoffeeAmericano(r *jobs.Runner) {
	dets := r.Rep.InitDetail("Americano")
	dets.Append("Hot Water", "95 degree celcius")
	r.Rep.Pass()
}

func CoffeeEspresso(r *jobs.Runner) {
	dets := r.Rep.InitDetail("Espresso")
	dets.Append("Hot Water", "70 degree celcius, not hot enough")
	r.Rep.Warning()
}
