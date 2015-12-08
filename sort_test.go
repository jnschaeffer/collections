package collections

import (
	"sort"
	"testing"
)

type negativeSorter struct {
	negativeBounded
	dummyComparable
	dummySwappable
}

type nonsenseSorter struct {
	dummyBounded
	badComparable
	dummySwappable
}

// INPUT: A slice of ints as an IntSlice wrapped by a Sorter.
// EXPECTED: A sorted slice of ints.
func TestSorter(t *testing.T) {
	ints := []int{5, 1, 3, 4, 2}

	s := newSorter(IntSlice(ints))

	s.sort()

	if !sort.IsSorted(s) {
		t.Errorf("bad sort: %#v is not sorted", ints)
	}
}

// INPUT: A sorter with negative length.
// EXPECTED: A panic.
func TestNegativeSorter(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Fatalf("uncaught negative sorter length")
		}
	}()

	n := negativeSorter{}

	s := newSorter(n)

	s.sort()
}

// INPUT: A sorter with nonsense comparison results.
// EXPECTED: A panic.
func TestNonsenseSorterCompare(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Fatalf("uncaught nonsense compare result")
		}
	}()

	n := nonsenseSorter{}

	s := newSorter(n)

	s.sort()
}
