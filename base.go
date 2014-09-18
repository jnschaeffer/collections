// Package collections contains interfaces and functions meant for performing
// operations on collections in Go, such as zipping collections together and
// splitting them into smaller collections.
package collections

import "fmt"

// An Ord represents a result of comparing two items.
type Ord int

const (
	Less    Ord = 1 << iota // Less represents the idea of x < y.
	Equal                   // Equal represents the idea of x == y.
	Greater                 // Greater represents the idea of x > y.
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

type IntSlice []int

func (s IntSlice) Len() int { return len(s) }
func (s IntSlice) Compare(i, j int) Ord {
	x, y := s[i], s[j]
	var ord Ord
	switch {
	case x < y:
		ord = Less
	case x == y:
		ord = Equal
	case x > y:
		ord = Greater
	}

	return ord
}
func (s IntSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type FloatSlice []float64

func (s FloatSlice) Len() int { return len(s) }
func (s FloatSlice) Compare(i, j int) Ord {
	x, y := s[i], s[j]
	var ord Ord
	switch {
	case x < y:
		ord = Less
	case x == y:
		ord = Equal
	case x > y:
		ord = Greater
	}

	return ord
}
func (s FloatSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type StringSlice []string

func (s StringSlice) Len() int { return len(s) }
func (s StringSlice) Compare(i, j int) Ord {
	x, y := s[i], s[j]
	var ord Ord
	switch {
	case x < y:
		ord = Less
	case x == y:
		ord = Equal
	case x > y:
		ord = Greater
	}

	return ord
}
func (s StringSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type ByteSlice []byte
