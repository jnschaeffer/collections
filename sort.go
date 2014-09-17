package collections

// Sortable represents a collection which can be used by the standard sort
// package (with the Sorter wrapper).
type Sortable interface {
	Bounded
	Comparable
}

// Sorter is a wrapper around a Sortable which can be used directly in
// sort.Sort.
type Sorter struct {
	Sortable
}

func (s *Sorter) Less(i, j int) bool {
	return s.Compare(i, j) == Less
}
