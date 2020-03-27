package utils

import (
	"sync"
)

type Runner struct {
	Series	[]func(*Runner) int
	Logger 	chan string
	group	*sync.WaitGroup
}

func (r *Runner) Wrap(item func(*Runner) int) func(*sync.WaitGroup) int {
	return func(wg *sync.WaitGroup) int {
		defer wg.Done()
		return item(r)
	}
}

func (r *Runner) Factory(series []func(*Runner) int) []func(*sync.WaitGroup) int {
	
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

	for _, test := range r.Factory(r.Series) {
		r.group.Add(1)
		go test(r.group)
	}
	r.group.Wait()
}

func (r *Runner) Debug(message string) {
	r.Logger <- "  " + message
}