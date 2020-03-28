package utils

import (
	"fmt"
	"sync"
)

const (
	LOG_FORMAT	= "[%s]\t%s"
)

type Runner struct {
	Series		[]func(*Runner) int
	Logger 		chan string
	LogLevel	int
	group		*sync.WaitGroup
}

func (r *Runner) Wrap(item func(*Runner) int) func(*sync.WaitGroup) int {
	return func(wg *sync.WaitGroup) int {
		defer wg.Done()
		return item(r)
	}
}

func (r *Runner) CaseFactory(series []func(*Runner) int) []func(*sync.WaitGroup) int {
	
	var _series []func(*sync.WaitGroup) int
	for _, item := range series {
		_series = append(_series, r.Wrap(item))
	}
	return _series
}

func (r *Runner) Init() {

	r.Logger = make(chan string)
	r.group = new(sync.WaitGroup)
}

func (r *Runner) Start() {

	for _, test := range r.CaseFactory(r.Series) {
		r.group.Add(1)
		go test(r.group)
	}
	r.group.Wait()
}

func (r *Runner) Debug(msg string) {
	if r.LogLevel >= 2 {
		r.Logger <-	fmt.Sprintf(LOG_FORMAT, "DEBUG", msg)
	}
}

func (r *Runner) Info(msg string) {
	if r.LogLevel >= 1 {
		r.Logger <- fmt.Sprintf(LOG_FORMAT, "INFO", msg)
	}
}

func (r *Runner) Warning(msg string) {
	if r.LogLevel >= 0 {
		r.Logger <- fmt.Sprintf(LOG_FORMAT, "WARN", msg)
	}
}

func (r *Runner) Error(msg string) {
	r.Logger <- fmt.Sprintf(LOG_FORMAT, "ERROR", msg)
}
