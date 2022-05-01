package esquery

type InnerHitsQuery struct {
	from *uint64
	size *uint64
	sort Sort
	name string
}

func InnerHits() *InnerHitsQuery {
	return &InnerHitsQuery{}
}

func (i *InnerHitsQuery) From(from uint64) *InnerHitsQuery {
	i.from = &from
	return i
}

func (i *InnerHitsQuery) Size(size uint64) *InnerHitsQuery {
	i.size = &size
	return i
}

func (i *InnerHitsQuery) Sort(name string, order Order) *InnerHitsQuery {
	i.sort = append(i.sort, map[string]interface{}{
		name: map[string]interface{}{
			"order": order,
		},
	})

	return i
}

func (i *InnerHitsQuery) Name(name string) *InnerHitsQuery {
	i.name = name
	return i
}

func (i *InnerHitsQuery) Map() map[string]interface{} {
	m := make(map[string]interface{})

	if i.from != nil {
		m["from"] = *i.from
	}
	if i.size != nil {
		m["size"] = *i.size
	}
	if i.sort != nil {
		m["sort"] = i.sort
	}
	if i.name != "" {
		m["name"] = i.name
	}

	return m
}

type NestedQuery struct {
	path      string
	query     Mappable
	innerHits Mappable
	name      string
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

func (n *NestedQuery) InnerHits(i Mappable) *NestedQuery {
	n.innerHits = i
	return n
}

func (req *NestedQuery) Map() map[string]interface{} {
	m := make(map[string]interface{})

	m["path"] = req.path
	if req.query != nil {
		m["query"] = req.query.Map()
	}
	if req.innerHits != nil {
		m["inner_hits"] = req.innerHits.Map()
	}
	if req.name != "" {
		m["_name"] = req.name
	}

	return map[string]interface{}{
		"nested": m,
	}
}
