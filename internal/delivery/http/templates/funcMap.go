package templates

import (
	"html/template"
)

func GetMainTemplatesFuncMap() template.FuncMap {
	return template.FuncMap{
		"exampleAdd": func(a int, b int) int { return a + b },
	}
}
