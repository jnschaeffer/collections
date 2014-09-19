package collections

import (
	"fmt"
	"testing"
)

type dummySplitter struct{}

func (s dummySplitter) Split(i, j int) {}

type negativeSplitter struct {
	negativeBounded
	dummyComparable
	dummySplitter
}

type nonsenseSplitter struct {
	badComparable
	dummyBounded
	dummySplitter
}

// INPUT: A slice of repeating ints.
// EXPECTED: n slices of ints, each containing only the same value.
func TestSplit(t *testing.T) {
	ints := []int{1, 1, 1, 2, 2, 2, 3}
	splitter := NewIntSplitter(ints)

	Split(splitter)

	res := splitter.Result
	failStr := fmt.Sprintf("bad split: expected [[1 1 1] [2 2 2] [3]], got %v",
		res)

	switch {
	case len(res) != 3:
		t.Fatalf(failStr)
	case len(res[0]) != 3:
		t.Fatalf(failStr)
	case len(res[1]) != 3:
		t.Fatalf(failStr)
	case len(res[2]) != 1:
		t.Fatalf(failStr)
	}

	for _, x := range res[0] {
		if x != 1 {
			t.Fatalf(failStr)
		}
	}

	for _, x := range res[1] {
		if x != 2 {
			t.Fatalf(failStr)
		}
	}

	for _, x := range res[2] {
		if x != 3 {
			t.Fatalf(failStr)
		}
	}
}

// INPUT: A bad splitter implementation with negative length.
// EXPECTED: A panic.
func TestNegativeSplitter(t *testing.T) {
	s := negativeSplitter{}
	defer func() {
		if err := recover(); err == nil {
			t.Fatalf("uncaught negative splitter length")
		}
	}()

	Split(s)
}

// INPUT: A bad splitter implementation with nonsense ordering.
// Expected: A panic.
func TestNonsenseSplitterCompare(t *testing.T) {
	s := nonsenseSplitter{}

	defer func() {
		if err := recover(); err == nil {
			t.Fatalf("uncaught nonsense compare result")
		}
	}()
	Split(s)
}
