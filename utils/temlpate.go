package utils

import (

	"io/ioutil"
	"text/template"
	"bytes"
)

func Render(report Report) (string, error) {

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
