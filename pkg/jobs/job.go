package main

import (
	"fmt"

	"github.com/RainrainWu/probe/pkg/utils"
	"github.com/RainrainWu/probe/pkg/tests"
)

const (
	FLAG_NUMBER	int	= 1
)

var (
	case_catalog map[string]([]func(*utils.Runner) int) = map[string]([]func(*utils.Runner) int){
		"salad": 	tests.SaladCase,
		"coffee": 	tests.CoffeeCase,
	}
	flags	chan int = make(chan int, FLAG_NUMBER)
)

func init() {
	for i := 0; i < cap(flags); i++ {
		remandFlag()
	}
}

func fetchFlag() int {
	return <- flags
}

func remandFlag() {
	flags <- 1
}

func RunJob(meta utils.Metadata) string {
	fetchFlag()
	var series []func(*utils.Runner) int
	for _, topic := range meta.Topic {
		series = append(series, case_catalog[topic]...)
	}

	runner := utils.Runner{
		Series:	series,
	}
	runner.Init()
	runner.Rep.SetMeta(meta)
	go runner.Run()

	result := <- runner.Result
	utils.WriteReport(*runner.Rep)
	remandFlag()
	return result
}

func main() {
	s_meta := utils.Metadata{
		Topic: []string{"salad"},
	}
	c_meta := utils.Metadata{
		Topic: []string{"coffee"},
	}
	fmt.Println(RunJob(s_meta))
	fmt.Println(RunJob(c_meta))
}
