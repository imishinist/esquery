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
			}},
	})
}
