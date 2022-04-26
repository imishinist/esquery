package esquery

import "testing"

func TestNested(t *testing.T) {
	runMapTests(t, []mapTest{
		{
			"Nested Query",
			Nested("path", Term("field", "value")),
			map[string]interface{}{
				"nested": map[string]interface{}{
					"path":  "path",
					"query": Term("field", "value").Map(),
				},
			},
		},
		{
			"Nested Query with name",
			Nested("path", Term("field", "value")).Name("name"),
			map[string]interface{}{
				"nested": map[string]interface{}{
					"path":  "path",
					"query": Term("field", "value").Map(),
					"_name": "name",
				},
			},
		},
	})
}
