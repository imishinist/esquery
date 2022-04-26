package esquery

type NestedQuery struct {
	path  string
	query Mappable
	name  string
}

func Nested(path string, query Mappable) *NestedQuery {
	return &NestedQuery{
		path:  path,
		query: query,
	}
}

// Name sets the query name
func (n *NestedQuery) Name(s string) *NestedQuery {
	n.name = s
	return n
}

func (req *NestedQuery) Map() map[string]interface{} {
	m := make(map[string]interface{})

	m["path"] = req.path
	if req.query != nil {
		m["query"] = req.query.Map()
	}
	if req.name != "" {
		m["_name"] = req.name
	}

	return map[string]interface{}{
		"nested": m,
	}
}
