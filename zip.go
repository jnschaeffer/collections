package collections

import "fmt"

// Zipper represents a pair of ordered collections that can be zipped together.
// Elements of each collection are assumed to be sorted in ascending order.
type Zipper interface {
	// Comparable must compare the left and right collection elements at i and
	// j respectively, returning the comparison with respect to the left value.
	// If the left element is less than the right, Compare should return Less.
	// If the left and right elements are equal, return Equal. If the left
	// element is greater, return Greater. Any other value will cause a panic
	// during Zip.
	Comparable
	// LenLeft returns the length of the left collection.
	LenLeft() int
	// LenRight returns the length of the right collection.
	LenRight() int
	// AddLeft adds only the element from the left collection at i to the
	// zipped collection.
	AddLeft(i int)
	// AddRight adds only the element from the right collection at j to the
	// zipped collection.
	AddRight(j int)
	// AddBoth adds both the left and right elements at i and j to the zipped
	// collection.
	AddBoth(i, j int)
}

// Zip iterates through each element in the collection, comparing each leading
// element in each collection exactly once. Any two elements that are equal
// will be zipped together by AddBoth, otherwise the lesser element will be
// added on its own. Assumes the left and right collections are in ascending
// order when ordered by z.Compare.
func Zip(z Zipper) {
	i, j := 0, 0
	maxLeft, maxRight := z.LenLeft(), z.LenRight()
	for i < maxLeft || j < maxRight {
		switch {
		case i >= maxLeft:
			z.AddRight(j)
			j += 1
		case j >= maxRight:
			z.AddLeft(i)
			i += 1
		default:
			switch c := z.Compare(i, j); {
			case c == Less:
				z.AddLeft(i)
				i += 1
			case c == Greater:
				z.AddRight(j)
				j += 1
			case c == Equal:
				z.AddBoth(i, j)
				i += 1
				j += 1
			default:
				msg := fmt.Sprintf("Zip: compare returned %d: expected %s, "+
					"%s, or %s", c, Less, Equal, Greater)
				panic(msg)
			}
		}
	}
}

type alwaysEqualZipper struct {
	Zipper
}

func (z *alwaysEqualZipper) Compare(i, j int) Ord {
	return Equal
}

// ZipWithoutGaps zips both collections in the zipper together but assumes
// elements are always equal.
func ZipWithoutGaps(z Zipper) {
	Zip(&alwaysEqualZipper{z})
}
