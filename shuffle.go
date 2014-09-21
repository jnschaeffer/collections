package collections

import "math/rand"

// A Shuffler represents any collection where elements can be shuffled in
// random order.
type Shuffler interface {
	Bounded
	Swappable
}

// Shuffle rearranges the elements of the provided shuffler using the
// Fisher-Yates shuffle.
func Shuffle(s Shuffler) {
	for i := s.Len() - 1; i > 0; i-- {
		s.Swap(i, rand.Intn(i+1))
	}
}
