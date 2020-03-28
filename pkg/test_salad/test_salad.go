package test_salad

import (
	"github.com/RainrainWu/probe/pkg/utils"
)

var Cases []func(*utils.Runner) int = []func(*utils.Runner) int {
	Cobb,
	Caesar,
}

func Cobb(r *utils.Runner) int {
	r.Info("Start making Cobb")
	r.Info("Finish making Cobb")
	return 0
}

func Caesar(r *utils.Runner) int {
	r.Info("Start making Caesar")
	r.Debug("Add topping")
	r.Error("Topping not found")
	r.Info("Finish making Caesar")
	return 0
}
