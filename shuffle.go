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
	n := s.Len()
	if n < 0 {
		panic("Shuffle: negative length")
	}
	for i := n - 1; i >= 1; i-- {
		j := rand.Intn(i + 1)
		for j == i {
			j = rand.Intn(i + 1)
		}
		s.Swap(i, j)
	}
}
