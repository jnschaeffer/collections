// collections contains interfaces and functions meant for performing
// operations on collections in Go, such as zipping collections together and
// splitting them into smaller collections.

package collections

// An Ord represents a result of comparing two items.
type Ord int

const (
	Less    Ord = 1 << iota // Less represents the idea of x < y.
	Equal                   // Equal represents the idea of x == y.
	Greater                 // Greater represents the idea of x > y.
)

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
