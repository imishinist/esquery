package esquery

import (
	"testing"
	"time"
)

func TestSearchMaps(t *testing.T) {
	runMapTests(t, []mapTest{
		{
			"a simple match_all query with a size and no aggs",
			Search().Query(MatchAll()).Size(20),
			map[string]interface{}{
				"query": map[string]interface{}{
					"match_all": map[string]interface{}{},
				},
				"size": 20,
			},
		},
		{
			"a complex query with an aggregation and various other options",
			Search().
				Query(
					Bool().
						Must(
							Range("date").
								Gt("some time in the past").
								Lte("now").
								Relation(RangeContains).
								TimeZone("Asia/Jerusalem").
								Boost(2.3),

							Match("author").
								Query("some guy").
								Analyzer("analyzer?").
								Fuzziness("fuzz"),
						).
						Boost(3.1),
				).
				Aggs(
					Sum("total_score", "score"),
					StringStats("tag_stats", "tags").
						ShowDistribution(true),
				).
				Size(30).
				From(5).
				Explain(true).
				Timeout(time.Duration(20000000000)),
			map[string]interface{}{
				"query": map[string]interface{}{
					"bool": map[string]interface{}{
						"must": []map[string]interface{}{
							{
								"range": map[string]interface{}{
									"date": map[string]interface{}{
										"gt":        "some time in the past",
										"lte":       "now",
										"relation":  "CONTAINS",
										"time_zone": "Asia/Jerusalem",
										"boost":     2.3,
									},
								},
							},
							{
								"match": map[string]interface{}{
									"author": map[string]interface{}{
										"query":     "some guy",
										"analyzer":  "analyzer?",
										"fuzziness": "fuzz",
									},
								},
							},
						},
						"boost": 3.1,
					},
				},
				"aggs": map[string]interface{}{
					"total_score": map[string]interface{}{
						"sum": map[string]interface{}{
							"field": "score",
						},
					},
					"tag_stats": map[string]interface{}{
						"string_stats": map[string]interface{}{
							"field":             "tags",
							"show_distribution": true,
						},
					},
				},
				"size":    30,
				"from":    5,
				"explain": true,
				"timeout": "20s",
			},
		},
	})
}
