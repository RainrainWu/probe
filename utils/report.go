package utils

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
}

type Report struct {
	Index   string
	Meta	Metadata
	Stat	Statistic
	Dets	map[string]Detail
}

type Filter struct {
	Index	string
}

func (r *Report) SetMeta(data Metadata) {
	r.Meta = data
	r.Dets = make(map[string]Detail)
	r.Index = r.Meta.Index
}

func (r *Report) InitDetail(name string) Detail {
	r.Dets[name] = Detail{
		Content: make(map[string]string),
	}
	return r.Dets[name]
}

func (d *Detail) Append(key, value string) {
	d.Content[key] = value
}

func (r *Report) Pass() {
	r.Stat.Total += 1
	r.Stat.Pass += 1
}

func (r *Report) Warning() {
	r.Stat.Total += 1
	r.Stat.Warning += 1
}

func (r *Report) Fail() {
	r.Stat.Total += 1
	r.Stat.Fail += 1
}
