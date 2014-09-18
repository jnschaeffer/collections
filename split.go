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

// IntSplitter represents a splitter over a list of ints.
// Note: By adding IntSlice anonymously, we gain access to the underlying
// Bounded and Comparable implementations of IntSlice. This same pattern is
// true for FloatSlice, StringSlice, etc.
type IntSplitter struct {
	IntSlice
	Result []IntSlice
}

func NewIntSplitter(xs IntSlice) *IntSplitter {
	return &IntSplitter{xs, []IntSlice{}}
}

func (s *IntSplitter) Split(i, j int) {
	s.Result = append(s.Result, s.IntSlice[i:j])
}

// FloatSplitter represents a splitter over a list of ints.
type FloatSplitter struct {
	Floats FloatSlice
	Result []FloatSlice
}

func NewFloatSplitter(xs FloatSlice) *FloatSplitter {
	return &FloatSplitter{xs, []FloatSlice{}}
}

func (s *FloatSplitter) Split(i, j int) {
	s.Result = append(s.Result, s.Floats[i:j])
}

// StringSplitter represents a splitter over a list of ints.
type StringSplitter struct {
	Strings StringSlice
	Result  []StringSlice
}

func NewStringSplitter(xs StringSlice) *StringSplitter {
	return &StringSplitter{xs, []StringSlice{}}
}

func (s *StringSplitter) Split(i, j int) {
	s.Result = append(s.Result, s.Strings[i:j])
}

// ByteSplitter represents a splitter over a list of ints.
type ByteSplitter struct {
	Bytes  ByteSlice
	Result []ByteSlice
}

func NewByteSplitter(xs ByteSlice) *ByteSplitter {
	return &ByteSplitter{xs, []ByteSlice{}}
}

func (s *ByteSplitter) Split(i, j int) {
	s.Result = append(s.Result, s.Bytes[i:j])
}

// Split iterates over the entire collection in order, splitting it into groups
// where each element in the group is equal according to Compare.
func Split(s Splitter) {
	start := 0
	max := s.Len()

	// early termination/panic conditions
	switch {
	case max < 0:
		panic("Split: negative splitter length")
	case max == 0:
		return
	}

	for i := 0; i < max; i++ {
		o := s.Compare(start, i)
		checkOrd("Split", o)
		if o != Equal {
			s.Split(start, i)
			start = i
		}
	}

	if start < max {
		s.Split(start, max)
	}
}
