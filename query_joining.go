package esquery

type NestedQuery struct {
	path  string
	query Mappable
}

func Nested(path string, query Mappable) *NestedQuery {
	return &NestedQuery{
		path:  path,
		query: query,
	}
}

func (req *NestedQuery) Map() map[string]interface{} {
	m := make(map[string]interface{})

	m["path"] = req.path
	if req.query != nil {
		m["query"] = req.query.Map()
	}

	return map[string]interface{}{
		"nested": m,
	}
}
