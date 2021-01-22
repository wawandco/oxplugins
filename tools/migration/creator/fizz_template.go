package creator

import (
	"html/template"

	"github.com/gobuffalo/flect"
)

var fizzUPTemplate string = `create_table("{{ .TableName }}") {
}
`

var fizzDownTemplate string = `drop_table("{{ .TableName }}")`

var templateFuncs = template.FuncMap{
	"tablelize": func(field string) string {
		return flect.Underscore(flect.Pluralize(field))
	},
	// "pascalize": func(field string) string {
	// 	return flect.Pascalize(field)
	// },
	// "pluralize": func(field string) string {
	// 	return flect.Pluralize(flect.Capitalize(field))
	// },
	// "properize": func(field string) string {
	// 	return flect.Capitalize(flect.Singularize(field))
	// },
	// "singularize": func(field string) string {
	// 	return flect.Singularize(field)
	// },
	// "underscore": func(field string) string {
	// 	return flect.Underscore(field)
	// },
}
