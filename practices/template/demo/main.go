package main

import (
	"html/template"
	"os"
	"strings"
)

func main() {
	const text = `
Output 0: {{title .Name1}}
Output 1: {{title .Name2}}
Output 2: {{.Name3 | title}}
`
	funcMap := template.FuncMap{
		"title": strings.Title,
	}
	tpl, _ := template.New("tour").Funcs(funcMap).Parse(text)
	data := map[string]string{
		"Name1": "go",
		"Name2": "php",
		"Name3": "python",
	}
	_ = tpl.Execute(os.Stdout, data)
}
