// Package collections contains interfaces and functions meant for performing
// operations on collections in Go, such as zipping collections together and
// splitting them into smaller collections.
package collections

import (
	"fmt"
	"time"
)

// An Ord represents a result of comparing two items.
type Ord int

const (
	// Less represents the idea of x < y.
	Less Ord = 1 << iota
	// Equal represents the idea of x == y.
	Equal
	// Greater represents the idea of x > y.
	Greater
)

// checkOrd returns an error if the ord value is not Less, Equal, or Greater.
func checkOrd(prefix string, o Ord) {
	if v := o &^ (Less | Equal | Greater); v != 0 {
		panic(fmt.Sprintf("%s: bad compare value %d", prefix, o))
	}
}

// String converts an Ord to a string for pretty printing.
func (o Ord) String() string {
	switch o {
	case Less:
		return "Less"
	case Equal:
		return "Equal"
	case Greater:
		return "Greater"
	default:
		return "?"
	}
}

// Bounded represents any collection with a known length, such as a list.
type Bounded interface {
	Len() int
}

// Comparable represents any collection where two elements can be accessed and
// compared.
type Comparable interface {
	Compare(i, j int) Ord
}

// Swappable represents any collction where two elements can be swapped.
type Swappable interface {
	Swap(i, j int)
}

// IntSlice is an alias for a slice of ints.
type IntSlice []int

// Len calculates the length of the underlying int slice.
func (s IntSlice) Len() int { return len(s) }

// Compare compares two ints in a slice, returning their ordering.
func (s IntSlice) Compare(i, j int) Ord {
	x, y := s[i], s[j]
	return CompareInts(x, y)
}

// Swap swaps two ints in a slice.
func (s IntSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// FloatSlice is an alias for a slice of floats.
type FloatSlice []float64

// Len calculates the length of the underlying float slice.
func (s FloatSlice) Len() int { return len(s) }

// Compare compares two floats in a slice, returning their ordering.
func (s FloatSlice) Compare(i, j int) Ord {
	x, y := s[i], s[j]
	return CompareFloats(x, y)
}

// Swap swaps two floats in a slice.
func (s FloatSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// StringSlice is an alias for a slice of strings.
type StringSlice []string

// Len calculates the length of the underlying string slice.
func (s StringSlice) Len() int { return len(s) }

// Compare compares two strings in a slice, returning their ordering.
func (s StringSlice) Compare(i, j int) Ord {
	x, y := s[i], s[j]
	return CompareStrings(x, y)
}

// Swap swaps two strings in a slice.
func (s StringSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// TimeSlice is an alias for a slice of times.
type TimeSlice []time.Time

// Len calculates the length of the underlying time slice.
func (s TimeSlice) Len() int { return len(s) }

// Compare compares two times in a slice, returning their ordering.
func (s TimeSlice) Compare(i, j int) Ord {
	x, y := s[i], s[j]
	return CompareTimes(x, y)
}

// Swap swaps two times in a slice.
func (s TimeSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
