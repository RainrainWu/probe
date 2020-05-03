package jobs

import (

	"errors"

	"go.uber.org/zap"

	"github.com/RainrainWu/probe/utils"
	"github.com/RainrainWu/probe/config"
)

var (
	catalog map[string]([]func(*Runner) int) = make(map[string]([]func(*Runner) int))
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

func AddJob(subject string, content []func(*Runner) int) error {
	if _, ok := catalog[subject]; ok {
		return errors.New("Subject already exist in catalog.")
	}
	if len(content) <= 0 {
		return errors.New("Empty content.")
	}
	catalog[subject] = content
	return nil
}

func RunJob(meta utils.Metadata) string {
	fetchFlag()
	var series []func(*Runner) int
	for _, topic := range meta.Topic {
		series = append(series, catalog[topic]...)
	}

	runner := Runner{
		Series:	series,
	}
	runner.Init()
	runner.Rep.SetMeta(meta)
	utils.Logger.Info("Start running jobs",
		zap.String("Index", meta.Index),
		zap.String("Env", meta.Env),
		zap.String("Tester", meta.Tester),
		zap.Any("Topic", meta.Topic),
		zap.String("Subject", meta.Subject),
	)
	go runner.Run()

	result := <- runner.Result
	utils.WriteReport(*runner.Rep)
	remandFlag()
	return result
}
