package jobs

import (
	"github.com/RainrainWu/probe/pkg/utils"
	"github.com/RainrainWu/probe/pkg/tests"
	"github.com/RainrainWu/probe/pkg/config"
)

var (
	case_catalog map[string]([]func(*utils.Runner) int) = map[string]([]func(*utils.Runner) int){
		"salad": 	tests.SaladCase,
		"coffee": 	tests.CoffeeCase,
	}
	quota	chan int = make(chan int, config.WORKER_QUOTA)
)

func init() {
	for i := 0; i < cap(quota); i++ {
		remandFlag()
	}
}

func fetchFlag() int {
	return <- quota
}

func remandFlag() {
	quota <- 1
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
