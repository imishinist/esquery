package esquery

import (
	"testing"
)

func TestTermLevel(t *testing.T) {
	runMapTests(t, []mapTest{
		{
			"exists",
			Exists("title"),
			map[string]interface{}{
				"exists": map[string]interface{}{
					"field": "title",
				},
			},
		},
		{
			"exists with name",
			Exists("title").Name("name"),
			map[string]interface{}{
				"exists": map[string]interface{}{
					"field": "title",
					"_name": "name",
				},
			},
		},
		{
			"ids",
			IDs("1", "4", "100"),
			map[string]interface{}{
				"ids": map[string]interface{}{
					"values": []string{"1", "4", "100"},
				},
			},
		},
		{
			"ids with name",
			IDs("1", "4", "100").Name("name"),
			map[string]interface{}{
				"ids": map[string]interface{}{
					"values": []string{"1", "4", "100"},
					"_name":  "name",
				},
			},
		},
		{
			"simple prefix",
			Prefix("user", "ki"),
			map[string]interface{}{
				"prefix": map[string]interface{}{
					"user": map[string]interface{}{
						"value": "ki",
					},
				},
			},
		},
		{
			"complex prefix",
			Prefix("user", "ki").Rewrite("ji").Name("name"),
			map[string]interface{}{
				"prefix": map[string]interface{}{
					"user": map[string]interface{}{
						"value":   "ki",
						"rewrite": "ji",
						"_name":   "name",
					},
				},
			},
		},
		{
			"int range",
			Range("age").Gte(10).Lte(20).Boost(2.0),
			map[string]interface{}{
				"range": map[string]interface{}{
					"age": map[string]interface{}{
						"gte":   10,
						"lte":   20,
						"boost": 2.0,
					},
				},
			},
		},
		{
			"string range",
			Range("timestamp").Gte("now-1d/d").Lt("now/d").Relation(RangeIntersects).Name("name"),
			map[string]interface{}{
				"range": map[string]interface{}{
					"timestamp": map[string]interface{}{
						"gte":      "now-1d/d",
						"lt":       "now/d",
						"relation": "INTERSECTS",
						"_name":    "name",
					},
				},
			},
		},
		{
			"regexp",
			Regexp("user", "k.*y"),
			map[string]interface{}{
				"regexp": map[string]interface{}{
					"user": map[string]interface{}{
						"value": "k.*y",
					},
				},
			},
		},
		{
			"complex regexp",
			Regexp("user", "k.*y").Flags("ALL").MaxDeterminizedStates(10000).Rewrite("constant_score").Name("name"),
			map[string]interface{}{
				"regexp": map[string]interface{}{
					"user": map[string]interface{}{
						"value":                   "k.*y",
						"flags":                   "ALL",
						"max_determinized_states": 10000,
						"rewrite":                 "constant_score",
						"_name":                   "name",
					},
				},
			},
		},
		{
			"wildcard",
			Wildcard("user", "ki*y").Rewrite("constant_score"),
			map[string]interface{}{
				"wildcard": map[string]interface{}{
					"user": map[string]interface{}{
						"value":   "ki*y",
						"rewrite": "constant_score",
					},
				},
			},
		},
		{
			"wildcard with name",
			Wildcard("user", "ki*y").Rewrite("constant_score").Name("name"),
			map[string]interface{}{
				"wildcard": map[string]interface{}{
					"user": map[string]interface{}{
						"value":   "ki*y",
						"rewrite": "constant_score",
						"_name":   "name",
					},
				},
			},
		},
		{
			"fuzzy",
			Fuzzy("user", "ki"),
			map[string]interface{}{
				"fuzzy": map[string]interface{}{
					"user": map[string]interface{}{
						"value": "ki",
					},
				},
			},
		},
		{
			"complex fuzzy",
			Fuzzy("user", "ki").Fuzziness("AUTO").MaxExpansions(50).Transpositions(true).Name("name"),
			map[string]interface{}{
				"fuzzy": map[string]interface{}{
					"user": map[string]interface{}{
						"value":          "ki",
						"fuzziness":      "AUTO",
						"max_expansions": 50,
						"transpositions": true,
						"_name":          "name",
					},
				},
			},
		},
		{
			"term",
			Term("user", "Kimchy"),
			map[string]interface{}{
				"term": map[string]interface{}{
					"user": map[string]interface{}{
						"value": "Kimchy",
					},
				},
			},
		},
		{
			"complex term",
			Term("user", "Kimchy").Boost(1.3).Name("name"),
			map[string]interface{}{
				"term": map[string]interface{}{
					"user": map[string]interface{}{
						"value": "Kimchy",
						"boost": 1.3,
						"_name": "name",
					},
				},
			},
		},
		{
			"terms",
			Terms("user").Values("bla", "pl"),
			map[string]interface{}{
				"terms": map[string]interface{}{
					"user": []string{"bla", "pl"},
				},
			},
		},
		{
			"complex terms",
			Terms("user").Values("bla", "pl").Boost(1.3).Name("name"),
			map[string]interface{}{
				"terms": map[string]interface{}{
					"user":  []string{"bla", "pl"},
					"boost": 1.3,
					"_name": "name",
				},
			},
		},
		{
			"terms_set",
			TermsSet("programming_languages", "go", "rust", "COBOL").MinimumShouldMatchField("required_matches"),
			map[string]interface{}{
				"terms_set": map[string]interface{}{
					"programming_languages": map[string]interface{}{
						"terms":                      []string{"go", "rust", "COBOL"},
						"minimum_should_match_field": "required_matches",
					},
				},
			},
		},
	})
}
