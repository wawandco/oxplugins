package model

type opts struct {
	Name    string
	Attrs   []attr
	Imports []string
}

func (o opts) Char() string {
	return o.Name[:1]
}
