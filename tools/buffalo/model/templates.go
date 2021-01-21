package model

import (
	"html/template"

	"github.com/gobuffalo/flect"
)

var modelTemplate string = `package models

import (
	{{- range $i := .Imports }}
	"{{$i}}"
	{{- end }}
)

// {{ properize .Name }} model struct
type {{ properize .Name }} struct {
	{{- range $attr := .Attrs }}
	{{ pascalize $attr.Name }}	{{$attr.GoType }} ` + "`" + `json:"{{ underscore $attr.Name }}" db:"{{ underscore $attr.Name }}"` + "`" + `
	{{- end }}
}

// {{ pluralize .Name }} array model struct of {{ properize .Name }}
type {{ pluralize .Name }} []{{ properize .Name }}

// String converts the struct into a string value
func ({{ .Char }} {{ properize .Name }}) String() string {
	return fmt.Sprintf("%+v\n", {{ .Char }})
}
`

var modelTestTemplate string = `package models

func (ms *ModelSuite) Test_{{ properize .Name }}() {
	ms.Fail("This test needs to be implemented!")
}`

var templateFuncs = template.FuncMap{
	"capitalize": func(field string) string {
		return flect.Capitalize(field)
	},
	"pascalize": func(field string) string {
		return flect.Pascalize(field)
	},
	"pluralize": func(field string) string {
		return flect.Pluralize(flect.Capitalize(field))
	},
	"properize": func(field string) string {
		return flect.Capitalize(flect.Singularize(field))
	},
	"singularize": func(field string) string {
		return flect.Singularize(field)
	},
	"underscore": func(field string) string {
		return flect.Underscore(field)
	},
}
