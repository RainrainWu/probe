package jobs

import (
	"sync"
	"encoding/json"

	"github.com/RainrainWu/probe/utils"
)

type Runner struct {
	Series		[]func(*Runner)
	Rep			*utils.Report
	Result		chan string
	group		*sync.WaitGroup
}

func (r *Runner) Init() {

	r.Rep = new(utils.Report)
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

func (r *Runner) Wrap(item func(*Runner)) func(*sync.WaitGroup) {
	return func(wg *sync.WaitGroup) {
		defer wg.Done()
		item(r)
	}
}

func (r *Runner) WrapAll(series []func(*Runner)) []func(*sync.WaitGroup) {
	
	var _series []func(*sync.WaitGroup)
	for _, item := range series {
		_series = append(_series, r.Wrap(item))
	}
	return _series
}
