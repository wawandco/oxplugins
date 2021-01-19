package model

type opts struct {
	Name    string
	Attrs   map[string]string
	Imports []string
}

func (o opts) Char() string {
	return o.Name[:1]
}
