package collections

import "sort"

// Sorter represents a collection that can be sorted.
type Sorter interface {
	Bounded
	Comparable
	Swappable
}

// stdlibSorter is a wrapper around a Sorter which can be used directly in
// sort.Sort.
type stdlibSorter struct {
	Sorter
}

// newSorter creates a new stdlibSorter containing the provided Sorter.
func newSorter(s Sorter) *stdlibSorter {
	return &stdlibSorter{s}
}

// Len returns the length of the underlying Sorter. If the Sorter returns
// a negative length, Len will panic.
func (s *stdlibSorter) Len() int {
	n := s.Sorter.Len()
	if n < 0 {
		panic("Len: negative length")
	}

	return n
}

// Less returns the result of the underlying Sorter's compare operation. If
// the Sorter returns an invalid Ord value, Less will panic.
func (s *stdlibSorter) Less(i, j int) bool {
	o := s.Compare(i, j)
	checkOrd("Less", o)

	return o == Less
}

// sort uses the Go stdlib sort to sort the items in this collection.
func (s *stdlibSorter) sort() {
	sort.Sort(s)
}

// Sort sorts the items in the collection using the standard Go sorting function.
func Sort(s Sorter) {
	newSorter(s).sort()
}
