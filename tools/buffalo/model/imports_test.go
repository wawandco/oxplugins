package model

import (
	"reflect"
	"testing"
)

func Test_BuildImports(t *testing.T) {
	cases := []struct {
		attrs    []attr
		expected []string
		testName string
	}{
		{
			testName: "With Default Attributes",
			attrs:    []attr{{Name: "id", goType: "uuid"}, {Name: "created_at", goType: "timestamp"}, {Name: "updated_at", goType: "timestamp"}},
			expected: []string{"encoding/json", "github.com/gofrs/uuid", "log", "time"},
		},
		{
			testName: "All Possible Attributes",
			attrs:    []attr{{Name: "id", goType: "uuid"}, {Name: "created_at", goType: "timestamp"}, {Name: "updated_at", goType: "timestamp"}, {Name: "description", goType: "nulls.String"}, {Name: "prices", goType: "slices.Float"}},
			expected: []string{"encoding/json", "github.com/gobuffalo/nulls", "github.com/gobuffalo/pop/v5/slices", "github.com/gofrs/uuid", "log", "time"},
		},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			imports := buildImports(c.attrs)
			if !reflect.DeepEqual(c.expected, imports) {
				t.Errorf("unexpected result, it should be %v but got %v", c.expected, imports)
			}
		})
	}
}
