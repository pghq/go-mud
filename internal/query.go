package internal

const (
	// DefaultLimit is the default max search results
	DefaultLimit int = 50
)

// Query for internal services
type Query struct {
	Key   []byte
	Point []float64
	Tag   string
	Limit int
}

// LimitQuery creates a new limit query
func LimitQuery(limit []int) Query {
	q := Query{
		Limit: DefaultLimit,
	}

	if len(limit) > 0 {
		q.Limit = limit[0]
	}

	return q
}

// TagQuery creates a new tag query
func TagQuery(tag string, limit []int) Query {
	q := LimitQuery(limit)
	q.Tag = tag
	return q
}

// PointQuery creates a new point query
func PointQuery(point []float64, limit []int) Query {
	q := LimitQuery(limit)
	q.Point = point
	return q
}
