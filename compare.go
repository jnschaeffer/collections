package collections

import "time"

// CompareInts compares two ints and returns an ordering.
func CompareInts(x, y int) Ord {
	o := Equal
	switch {
	case x < y:
		o = Less
	case x > y:
		o = Greater
	}

	return o
}

// CompareFloats compares two floats and returns an ordering.
func CompareFloats(x, y float64) Ord {
	o := Equal
	switch {
	case x < y:
		o = Less
	case x > y:
		o = Greater
	}

	return o
}

// CompareStrings compares two strings and returns an ordering.
func CompareStrings(x, y string) Ord {
	o := Equal
	switch {
	case x < y:
		o = Less
	case x > y:
		o = Greater
	}

	return o
}

// CompareTimes compares two times and returns an ordering.
func CompareTimes(x, y time.Time) Ord {
	o := Equal
	switch {
	case x.Before(y):
		o = Less
	case x.After(y):
		o = Greater
	}

	return o
}
