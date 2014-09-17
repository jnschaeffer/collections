package collections

// Splitter represents an ordered collection that can be split into smaller
// units.
type Splitter interface {
	Bounded
	// Implementation of comparable should at the very least be able to compare
	// the current element to its predecessor.
	Comparable
	// Split splits the collection from i up to (but not including) j.
	Split(i, j int)
}

// Split iterates over the entire collection in order, splitting it into groups
// where each element in the group is equal according to Compare.
func Split(s Splitter) {
	start := 0
	max := s.Len()

	for i := 0; i < max; i++ {
		if s.Compare(start, i) != Equal {
			s.Split(start, i)
			start = i
		}
	}

	if start < max {
		s.Split(start, max)
	}
}
