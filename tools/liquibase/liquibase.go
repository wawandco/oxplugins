package liquibase

import (
	"fmt"

	pop4 "github.com/gobuffalo/pop"
	pop5 "github.com/gobuffalo/pop/v5"
)

func NewPlugin(conns interface{}) *Command {
	result := map[string]URLProvider{}

	switch v := conns.(type) {
	case map[string]*pop4.Connection:
		for k, conn := range v {
			result[k] = conn
		}
	case map[string]*pop5.Connection:
		for k, conn := range v {
			result[k] = conn
		}
	default:
		fmt.Println("[warning] Liquibase plugin ONLY receives pop v4 and v5 connections")
	}

	return &Command{
		connections: result,
	}
}
