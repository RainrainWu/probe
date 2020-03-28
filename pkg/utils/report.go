package main

import (
	"encoding/json"
	"fmt"
)

type Metadata struct {
	Id			string	`json:"id"`
	Timestamp	string	`json:"timestamp"`
	Environment	string	`json:"environment"`
	Tester		string	`json:"tester"`
	Description	string	`json:"description"`
}

type Statistic struct {
	Total	int
	Pass	int
	Warning	int
	Fail	int
}

type Message struct {
	Key		string
	Value	string
}

type Detail struct {
	Name	string
	Content	[]Message
}

type Report struct {
	Meta	Metadata	`json:"meta"`
	Stat	Statistic	`json:"stat"`
	Dets	[]Detail	`json:"dets"`
}

func main() {
	me := Metadata{
		Id: "001",
		Timestamp: "002",
		Environment: "dev",
		Tester: "CI",
		Description: "none",
	}
	st := Statistic{
		Total: 		3,
		Pass:		1,
		Warning:	1,
		Fail:		1,
	}
	ms1 := Message{
		Key:	"token",
		Value:	"12345",
	}
	ms2 := Message{
		Key:	"code",
		Value:	"200",
	}
	de := Detail{
		Name:		"Auth",
		Content:	[]Message{ms1, ms2},
	}
	re := &Report{
		Meta:	me,
		Stat:	st,
		Dets:	[]Detail{de},
	}
	data, _ :=json.Marshal(re)
	fmt.Println(string(data))
}