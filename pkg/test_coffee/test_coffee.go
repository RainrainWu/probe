package test_coffee

import (
	"github.com/RainrainWu/probe/pkg/utils"
)

var Cases []func(*utils.Runner) int = []func(*utils.Runner) int {
	Americano,
	Espresso,
}

func Americano(r *utils.Runner) int {
	dets := r.Rep.InitDetail("Americano")
	dets.Append("Hot Water", "95 degree celcius")
	r.Rep.Pass()
	return 0
}

func Espresso(r *utils.Runner) int {
	dets := r.Rep.InitDetail("Espresso")
	dets.Append("Hot Water", "70 degree celcius, not hot enough")
	r.Rep.Warning()
	return 1
}
