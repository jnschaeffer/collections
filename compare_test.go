package collections

import (
	"fmt"
	"testing"
	"time"
)

func checkResult(lt Ord, eq Ord, gt Ord) (err error) {
	if lt != Less || eq != Equal || gt != Greater {
		str := "bad ordering: expected (%s, %s, %s), got (%s, %s, %s)"
		err = fmt.Errorf(str, Less, Greater, Equal, lt, eq, gt)
	}

	return
}

// INPUT: Two ints.
// EXPECTED: Comparison results equivalent to <, =, and >.
func TestCompareInts(t *testing.T) {
	lt := CompareInts(0, 1)
	eq := CompareInts(1, 1)
	gt := CompareInts(1, 0)

	if err := checkResult(lt, eq, gt); err != nil {
		t.Fatalf("%s", err)
	}
}

// INPUT: Two floats.
// EXPECTED: Comparison results equivalent to <, =, and >.
func TestCompareFloats(t *testing.T) {
	lt := CompareFloats(0, 1)
	eq := CompareFloats(1, 1)
	gt := CompareFloats(1, 0)

	if err := checkResult(lt, eq, gt); err != nil {
		t.Fatalf("%s", err)
	}
}

// INPUT: Two strings.
// EXPECTED: Comparison results equivalent to <, =, and >.
func TestCompareStringss(t *testing.T) {
	lt := CompareStrings("a", "b")
	eq := CompareStrings("b", "b")
	gt := CompareStrings("b", "a")

	if err := checkResult(lt, eq, gt); err != nil {
		t.Fatalf("%s", err)
	}
}

// INPUT: Two times.
// EXPECTED: Comparison results equivalent to <, =, and >.
func TestCompareTimes(t *testing.T) {
	str := "2006-1-2"
	var (
		t1, t2 time.Time
		err    error
	)

	if t1, err = time.Parse(str, "2014-01-01"); err != nil {
		panic(err)
	}

	if t2, err = time.Parse(str, "2014-01-02"); err != nil {
		panic(err)
	}

	lt := CompareTimes(t1, t2)
	eq := CompareTimes(t2, t2)
	gt := CompareTimes(t2, t1)

	if err = checkResult(lt, eq, gt); err != nil {
		t.Fatalf("%s", err)
	}
}
