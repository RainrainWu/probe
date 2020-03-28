package test_coffee

import (
	"github.com/RainrainWu/probe/pkg/utils"
)

var Cases []func(*utils.Runner) int = []func(*utils.Runner) int {
	Americano,
	Espresso,
}

func Americano(r *utils.Runner) int {
	r.Info("Start making americano")
	r.Debug("Prepare hot water")
	r.Info("Finish making americano")
	return 0
}

func Espresso(r *utils.Runner) int {
	r.Info("Start making espresso")
	r.Debug("Prepare hot water")
	r.Warning("Hot water is not hot enough")
	r.Info("Finish making espresso")
	return 0
}
