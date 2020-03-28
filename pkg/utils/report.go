package utils

import (
	"time"
)

type Metadata struct {
	Id			string	`json:"id"`
	Time		string	`json:"time"`
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
	Name	string
	Content	map[string]string
}

type Report struct {
	Meta	Metadata	`json:"meta"`
	Stat	Statistic	`json:"stat"`
	Dets	[]Detail	`json:"dets"`
}

func (r *Report) SetMeta(id, env, tester, subject string) {
	data := Metadata{
		Id: 		id,
		Time: 		string(time.Now().Unix()),
		Env: 		env,
		Tester:		tester,
		Subject:	subject,
	}
	r.Meta = data
}

func (r *Report) Pass(detail Detail) {
	r.Stat.Total += 1
	r.Stat.Pass += 1
	r.Dets = append(r.Dets, detail)
}

func (r *Report) Warning(detail Detail) {
	r.Stat.Total += 1
	r.Stat.Warning += 1
	r.Dets = append(r.Dets, detail)
}

func (r *Report) Fail(detail Detail) {
	r.Stat.Total += 1
	r.Stat.Fail += 1
	r.Dets = append(r.Dets, detail)
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