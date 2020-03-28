package main

import (
	"fmt"

	"github.com/RainrainWu/probe/pkg/test_coffee"
	"github.com/RainrainWu/probe/pkg/test_salad"
	"github.com/RainrainWu/probe/pkg/utils"
)

func main() {

	coffee_runner := utils.Runner{
		Series: 	test_coffee.Cases,
		LogLevel: 	1,
	}
	coffee_runner.Init()
	go coffee_runner.Start()

	salad_runner := utils.Runner{
		Series: 	test_salad.Cases,
		LogLevel:	1,
	}
	salad_runner.Init()
	go salad_runner.Start()

	show(salad_runner)
}

func show(runner utils.Runner) {

	for {
		select {
		case msg := <- runner.Logger:
			fmt.Println(msg)
		default:
		}
	}
}