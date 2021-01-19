package model

import (
	"html/template"

	"github.com/gobuffalo/flect"
)

var modelTemplate string = `package models

var (
	{{- range $i := .Imports }}
	"{{$i}}"
	{{- end }}
)

// {{ properize .Name }} model struct
type {{ properize .Name }} struct {
}

// String converts the struct into a string value
func ({{ .Char }} {{ properize .Name }}) String() string {
	js, _ := json.Marshal({{ .Char }})

	return string(js)
}

// {{ pluralize .Name }} array model struct of {{ properize .Name }}
type {{ pluralize .Name }} []{{ properize .Name }}
`

var modelTestTemplate string = `package models

func (ms *ModelSuite) Test_{{ properize .Name }}() {
	ms.Fail("This test needs to be implemented!")
}`

var templateFuncs = template.FuncMap{
	"capitalize": func(field string) string {
		return flect.Capitalize(field)
	},
	"singularize": func(field string) string {
		return flect.Singularize(field)
	},
	"pluralize": func(field string) string {
		return flect.Pluralize(flect.Capitalize(field))
	},
	"properize": func(field string) string {
		return flect.Capitalize(flect.Singularize(field))
	},
}
