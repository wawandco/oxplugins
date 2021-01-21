package model

import "strings"

type opts struct {
	Name    string
	Attrs   []attr
	Imports []string
}

func (o opts) Char() string {
	return strings.ToLower(o.Name[:1])
}
