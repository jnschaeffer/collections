package collections

import "testing"

type negativeShuffler struct {
	negativeBounded
	dummySwappable
}

func TestNegativeShuffler(t *testing.T) {
	n := negativeShuffler{}

	defer func() {
		if err := recover(); err == nil {
			t.Fatalf("uncaught negative shuffler length")
		}
	}()

	Shuffle(n)
}
