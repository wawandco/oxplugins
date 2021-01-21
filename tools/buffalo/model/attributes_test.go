package model

import (
	"reflect"
	"testing"
)

func Test_BuildAttrs(t *testing.T) {
	defaults := []attr{{Name: "id", goType: "uuid"}, {Name: "created_at", goType: "timestamp"}, {Name: "updated_at", goType: "timestamp"}}

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
			expected: []attr{{Name: "id", goType: "uuid"}, {Name: "created_at", goType: "timestamp"}, {Name: "updated_at", goType: "timestamp"}, {Name: "description", goType: "text"}, {Name: "title", goType: "string"}},
		},
		{
			testName: "Replacing Defaults",
			args:     []string{"description:text", "id:int"},
			expected: []attr{{Name: "created_at", goType: "timestamp"}, {Name: "updated_at", goType: "timestamp"}, {Name: "description", goType: "text"}, {Name: "id", goType: "int"}},
		},
		{
			testName: "Replacing Defaults 2",
			args:     []string{"created_at:int", "description:text", "updated_at:int", "id:int"},
			expected: []attr{{Name: "created_at", goType: "int"}, {Name: "description", goType: "text"}, {Name: "updated_at", goType: "int"}, {Name: "id", goType: "int"}},
		},
		{
			testName: "Testing Camelize Capitalize",
			args:     []string{"created_at:int", "Description:text", "updated_at:int", "ID:int"},
			expected: []attr{{Name: "created_at", goType: "int"}, {Name: "description", goType: "text"}, {Name: "updated_at", goType: "int"}, {Name: "id", goType: "int"}},
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
