package utils

import (
	"sync"
)

type Metadata struct {
	Index		string
	Env			string
	Tester		string
	Topic		[]string
	Subject		string
}

type Statistic struct {
	Total	int
	Pass	int
	Warning	int
	Fail	int
}

type Detail struct {
	Content	map[string]string
	lock	*sync.RWMutex
}

type Report struct {
	Index   string
	Meta	Metadata
	Stat	Statistic
	Dets	map[string]Detail
	lock	*sync.RWMutex
}

type Filter struct {
	Index	string
}

func (r *Report) SetMeta(data Metadata) {
	r.Meta = data
	r.Dets = make(map[string]Detail)
	r.Index = r.Meta.Index
	r.lock = &sync.RWMutex{}
}

func (r *Report) InitDetail(name string) Detail {
	r.lock.Lock()
	r.Dets[name] = Detail{
		Content: make(map[string]string),
		lock: &sync.RWMutex{},
	}
	r.lock.Unlock()
	return r.Dets[name]
}

func (d *Detail) Append(key, value string) {
	d.lock.Lock()
	d.Content[key] = value
	d.lock.Unlock()
}

func (r *Report) Pass() {
	r.lock.Lock()
	r.Stat.Total += 1
	r.Stat.Pass += 1
	r.lock.Unlock()
}

func (r *Report) Warning() {
	r.lock.Lock()
	r.Stat.Total += 1
	r.Stat.Warning += 1
	r.lock.Unlock()
}

func (r *Report) Fail() {
	r.lock.Lock()
	r.Stat.Total += 1
	r.Stat.Fail += 1
	r.lock.Unlock()
}
