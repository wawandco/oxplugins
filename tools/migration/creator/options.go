package creator

import (
	"strings"

	"github.com/gobuffalo/flect"
)

type opts struct {
	Columns   []string
	TableName string
}

func newOptions(args []string) opts {
	return opts{
		TableName: flect.Underscore(flect.Pluralize(strings.ToLower(args[0]))),
	}
}
