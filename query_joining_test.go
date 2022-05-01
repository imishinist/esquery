package esquery

import "testing"

func TestNested(t *testing.T) {
	runMapTests(t, []mapTest{
		{
			"InnerHits Query",
			InnerHits(),
			map[string]interface{}{},
		},
		{
			"InnerHits Query with options",
			InnerHits().From(1).Size(10).Sort("field", "desc").Name("name"),
			map[string]interface{}{
				"from": 1,
				"size": 10,
				"sort": Sort{
					{
						"field": map[string]interface{}{
							"order": "desc",
						},
					},
				},
				"name": "name",
			},
		},
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
			"Nested Query with inner_hits",
			Nested("path", Term("field", "value")).InnerHits(InnerHits()).Name("name"),
			map[string]interface{}{
				"nested": map[string]interface{}{
					"path":       "path",
					"query":      Term("field", "value").Map(),
					"inner_hits": InnerHits().Map(),
					"_name":      "name",
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
