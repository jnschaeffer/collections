package collections

// Sortable represents a collection which can be used by the standard sort
// package (with the Sorter wrapper).
type Sortable interface {
	Bounded
	Comparable
	Swappable
}

// Sorter is a wrapper around a Sortable which can be used directly in
// sort.Sort.
type Sorter struct {
	Sortable
}

// NewSorter creates a new Sorter containing the provided Sortable.
func NewSorter(s Sortable) *Sorter {
	return &Sorter{s}
}

// Len returns the length of the underlying Sortable. If the Sortable returns
// a negative length, Len will panic.
func (s *Sorter) Len() int {
	n := s.Sortable.Len()
	if n < 0 {
		panic("Len: negative length")
	}

	return n
}

// Less returns the result of the underlying Sortable's compare operation. If
// the Sortable returns an invalid Ord value, Less will panic.
func (s *Sorter) Less(i, j int) bool {
	o := s.Compare(i, j)
	checkOrd("Less", o)

	return o == Less
}
