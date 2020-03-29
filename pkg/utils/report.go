package utils

type Metadata struct {
	Id			string	`json:"id"`
	Env			string	`json:"env"`
	Tester		string	`json:"tester"`
	Subject		string	`json:"subject"`
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
	Meta	Metadata			`json:"meta"`
	Stat	Statistic			`json:"stat"`
	Dets	map[string]Detail	`json:"dets"`
}

func (r *Report) SetMeta(data Metadata) {
	r.Meta = data
	r.Dets = make(map[string]Detail)
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

/*
func main() {
	me := Metadata{
		Id: "001",
		Time: "002",
		Env: "dev",
		Tester: "CI",
		Subject: "none",
	}
	st := Statistic{
		Total: 		3,
		Pass:		1,
		Warning:	1,
		Fail:		1,
	}
	de := Detail{
		Name:		"Auth",
		Content:	map[string]string{
			"token": "123456",
			"code": "200",
		},
	}
	re := &Report{
		Meta:	me,
		Stat:	st,
		Dets:	[]Detail{de},
	}
	fmt.Println(re.Marshal())
}
*/