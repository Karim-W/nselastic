package index

import "github.com/karim-w/nselastic"

const (
	_DEFAULT_SIZE = 1000
	_DEFAULT_FROM = 0
	_DEFAULT_SORT = "asc"
)

func build_elastic_query(
	body nselastic.Queryable,
) map[string]interface{} {
	m := map[string]interface{}{
		"from":             body.From.GetOrElse(_DEFAULT_FROM),
		"size":             body.Size.GetOrElse(_DEFAULT_SIZE),
		"track_total_hits": true,
		"query": map[string]interface{}{
			"bool": map[string]interface{}{},
		},
	}

	if len(body.Search.Searches) > 0 {
		v, ok := m["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"]
		if !ok {
			v = []map[string]interface{}{}
		}

		partials := make([]map[string]interface{}, 0, len(body.Search.Searches))

		for _, p := range body.Search.Searches {
			if p == "" {
				continue
			}

			partials = append(partials, map[string]interface{}{
				"multi_match": map[string]interface{}{
					"query":    p,
					"fields":   []string{"*"},
					"operator": "and",
					"type":     "best_fields",
				},
			})
		}

		for _, p := range partials {
			v = append(v.([]map[string]interface{}), p)
		}

		m["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = v
	}

	if len(body.SortingOptions) > 0 {
		sorting := make([]map[string]interface{}, 0, len(body.SortingOptions))

		for _, sort := range body.SortingOptions {
			if sort.SortField.IsNone() {
				continue
			}
			sorting = append(
				sorting,
				map[string]interface{}{
					sort.SortField.Unwrap(): map[string]interface{}{
						"order": sort.SortOrder.GetOrElse(_DEFAULT_SORT),
					},
				},
			)
		}

		m["sort"] = sorting
	}

	if len(body.Search.Ranges) > 0 {

		v, ok := m["query"].(map[string]interface{})["bool"].(map[string]interface{})["filter"]
		if !ok {
			v = []map[string]interface{}{}
		}

		ranges := make([]map[string]interface{}, 0, len(body.Search.Ranges))

		for _, r := range body.Search.Ranges {
			if r.Gte.IsNone() && r.Lte.IsNone() {
				continue
			}

			rng := map[string]interface{}{}

			if r.Gte.IsSome() {
				rng["gte"] = r.Gte.Unwrap()
			}

			if r.Lte.IsSome() {
				rng["lte"] = r.Lte.Unwrap()
			}

			ranges = append(ranges, map[string]interface{}{
				r.Key: rng,
			})

		}

		for _, r := range ranges {
			v = append(v.([]map[string]interface{}), map[string]interface{}{
				"range": r,
			})
		}

		m["query"].(map[string]interface{})["bool"].(map[string]interface{})["filter"] = v
	}

	if len(body.Search.Equal) > 0 {
		v, ok := m["query"].(map[string]interface{})["bool"].(map[string]interface{})["filter"]
		if !ok {
			v = []map[string]interface{}{}
		}

		filters := make(
			[]map[string]interface{},
			0,
			len(body.Search.Equal)+len(body.Search.NotEqual),
		)

		for _, f := range body.Search.Equal {
			f := map[string]interface{}{
				"term": map[string]interface{}{
					f.Key: f.Value,
				},
			}
			filters = append(filters, f)
		}

		for _, f := range filters {
			v = append(v.([]map[string]interface{}), f)
		}

		m["query"].(map[string]interface{})["bool"].(map[string]interface{})["filter"] = v
	}

	if len(body.Search.NotEqual) > 0 {
		v, ok := m["query"].(map[string]interface{})["bool"].(map[string]interface{})["must_not"]
		if !ok {
			v = []map[string]interface{}{}
		}

		filters := make(
			[]map[string]interface{},
			0,
			len(body.Search.Equal)+len(body.Search.NotEqual),
		)

		for _, f := range body.Search.NotEqual {
			f := map[string]interface{}{
				"term": map[string]interface{}{
					f.Key: f.Value,
				},
			}
			filters = append(filters, f)
		}

		for _, f := range filters {
			v = append(v.([]map[string]interface{}), f)
		}

		m["query"].(map[string]interface{})["bool"].(map[string]interface{})["must_not"] = v
	}

	return m
}
