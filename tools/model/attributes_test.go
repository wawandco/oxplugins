package model

import (
	"reflect"
	"testing"
)

func Test_BuildAttrs(t *testing.T) {
	defaults := []attr{{name: "id", goType: "uuid"}, {name: "createdat", goType: "timestamp"}, {name: "updatedat", goType: "timestamp"}}

	cases := []struct {
		args     []string
		expected []attr
		testName string
	}{
		{
			testName: "Empty Args",
			args:     []string{},
			expected: defaults,
		},
		{
			testName: "Some Args Without Type",
			args:     []string{"description:text", "title"},
			expected: []attr{{name: "id", goType: "uuid"}, {name: "createdat", goType: "timestamp"}, {name: "updatedat", goType: "timestamp"}, {name: "description", goType: "text"}, {name: "title", goType: "string"}},
		},
		{
			testName: "Replacing Defaults",
			args:     []string{"description:text", "id:int"},
			expected: []attr{{name: "createdat", goType: "timestamp"}, {name: "updatedat", goType: "timestamp"}, {name: "description", goType: "text"}, {name: "id", goType: "int"}},
		},
		{
			testName: "Replacing Defaults 2",
			args:     []string{"createdat:int", "description:text", "updatedat:int", "id:int"},
			expected: []attr{{name: "createdat", goType: "int"}, {name: "description", goType: "text"}, {name: "updatedat", goType: "int"}, {name: "id", goType: "int"}},
		},
		{
			testName: "Testing Camelize Capitalize",
			args:     []string{"CreatedAt:int", "Description:text", "UpdatedAt:int", "ID:int"},
			expected: []attr{{name: "createdat", goType: "int"}, {name: "description", goType: "text"}, {name: "updatedat", goType: "int"}, {name: "id", goType: "int"}},
		},
	}

	for _, c := range cases {
		t.Run(c.testName, func(t *testing.T) {
			attrs := buildAttrs(c.args)
			if !reflect.DeepEqual(c.expected, attrs) {
				t.Errorf("unexpected result, it should be %v but got %v", c.expected, attrs)
			}
		})
	}
}
