package test_coffee

import (
	"github.com/RainrainWu/probe/pkg/utils"
)

var Cases []func(*utils.Runner) int = []func(*utils.Runner) int {
	Americano,
	Espresso,
	IceDripCoffee,
	Latte,
}

func Americano(r *utils.Runner) int {
	r.Debug("Make americano")
	return 0
}

func Espresso(r *utils.Runner) int {
	r.Debug("Make espresso")
	return 0
}

func IceDripCoffee(r *utils.Runner) int {
	r.Debug("Make ice drip coffee")
	return 0
}

func Latte(r *utils.Runner) int {
	r.Debug("Make latte")
	return 0
}