package main

import (

	"github.com/RainrainWu/probe"
	"github.com/RainrainWu/probe/jobs"
	"github.com/RainrainWu/probe/example_cases"
)

func main() {

	jobs.AddJob("coffee", example_cases.CoffeeCase)
	jobs.AddJob("salad", example_cases.SaladCase)
	probe.Start()
}