package main

import (

	"io/ioutil"
	"text/template"
	"bytes"
	"fmt"

	"github.com/RainrainWu/probe/utils"
)

var data utils.Report = utils.Report{
	Index: "Test template",
	Meta: utils.Metadata{
		Index: "Test template",
		Env: "local",
		Tester: "Rain",
		Topic: []string{ "coffee", "salad" },
		Subject: "test template",
	},
	Stat: utils.Statistic{
		Total: 4,
		Pass: 2,
		Warning: 1,
		Fail: 1,
	},
	Dets: map[string]utils.Detail{
		"Americano": utils.Detail{
			Content: map[string]string{
				"Hot Water": "95 degree celcius",
			},
		},
		"Caesar": utils.Detail{
			Content: map[string]string{
			  "Plate": "Plate is broken",
			  "Topping": "topping not found",
			},
		},
	},
}

func Render(report utils.Report) (string, error) {

	message, err := ioutil.ReadFile("templates/report.gohtml")
	if err != nil {
		return "", err
	}

	temp, err := template.New("UsersPage").Parse(string(message));
	if err != nil {
		return "", err
	}

	var result bytes.Buffer
	if err = temp.Execute(&result, report); err != nil {
		return "", err
	}
	return result.String(), nil
}

func main() {

	res, _ := Render(data)
	fmt.Println(res)
}