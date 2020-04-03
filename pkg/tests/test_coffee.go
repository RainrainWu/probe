package tests

import (
	"github.com/RainrainWu/probe/pkg/utils"
)

var CoffeeCase []func(*utils.Runner) int = []func(*utils.Runner) int {
	CoffeeAmericano,
	CoffeeEspresso,
}

func CoffeeAmericano(r *utils.Runner) int {
	dets := r.Rep.InitDetail("Americano")
	dets.Append("Hot Water", "95 degree celcius")
	r.Rep.Pass()
	return 0
}

func CoffeeEspresso(r *utils.Runner) int {
	dets := r.Rep.InitDetail("Espresso")
	dets.Append("Hot Water", "70 degree celcius, not hot enough")
	r.Rep.Warning()
	return 1
}
