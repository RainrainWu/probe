package utils

import (
	"sync"
	"encoding/json"
)

type Runner struct {
	Series		[]func(*Runner) int
	Rep			*Report
	Result		chan string
	group		*sync.WaitGroup
}

func (r *Runner) Init() {

	r.Rep = new(Report)
	r.Result = make(chan string)
	r.group = new(sync.WaitGroup)
}

func (r *Runner) Run() {

	for _, test := range r.WrapAll(r.Series) {
		r.group.Add(1)
		go test(r.group)
	}
	r.group.Wait()

	report, _ := json.MarshalIndent(r.Rep, "", "  ")
	r.Result <- string(report)
}

func (r *Runner) Wrap(item func(*Runner) int) func(*sync.WaitGroup) int {
	return func(wg *sync.WaitGroup) int {
		defer wg.Done()
		return item(r)
	}
}

func (r *Runner) WrapAll(series []func(*Runner) int) []func(*sync.WaitGroup) int {
	
	var _series []func(*sync.WaitGroup) int
	for _, item := range series {
		_series = append(_series, r.Wrap(item))
	}
	return _series
}
