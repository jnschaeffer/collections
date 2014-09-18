package collections

import "testing"

type IntZipper struct {
	Left   []int
	Right  []int
	Result [][2]int
}

func (z *IntZipper) Compare(i, j int) Ord {
	x, y := z.Left[i], z.Right[j]

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

func (z *IntZipper) LenLeft() int {
	return len(z.Left)
}

func (z *IntZipper) LenRight() int {
	return len(z.Right)
}

func (z *IntZipper) AddLeft(i int) {
	z.Result = append(z.Result, [2]int{z.Left[i], 0})
}

func (z *IntZipper) AddRight(j int) {
	z.Result = append(z.Result, [2]int{0, z.Right[j]})
}

func (z *IntZipper) AddBoth(i, j int) {
	z.Result = append(z.Result, [2]int{z.Left[i], z.Right[j]})
}

type zipperStub struct{}

func (z zipperStub) AddLeft(i int) {}

func (z zipperStub) AddRight(j int) {}

func (z zipperStub) AddBoth(i, j int) {}

type dummyLeft struct{}

func (d dummyLeft) LenLeft() int {
	return 2
}

type dummyRight struct{}

func (d dummyRight) LenRight() int {
	return 2
}

type negativeLeft struct{}

func (n negativeLeft) LenLeft() int {
	return -1
}

type negativeRight struct{}

func (n negativeRight) LenRight() int {
	return -1
}

type nonsenseZipper struct {
	zipperStub
	dummyLeft
	dummyRight
	badComparable
}

type negativeLeftZipper struct {
	zipperStub
	dummyComparable
	negativeLeft
	dummyRight
}

type negativeRightZipper struct {
	zipperStub
	dummyComparable
	dummyRight
	negativeLeft
}

// INPUT: An IntZipper where all of the left elements are odd and the right
// elements are even.
// EXPECTED: [[1 0] [0 2] [3 0] [0 4]]
func TestZipWithGaps(t *testing.T) {
	z := &IntZipper{
		[]int{1, 3},
		[]int{2, 4},
		[][2]int{},
	}

	ZipWithGaps(z)

	if len(z.Result) != 4 {
		t.Fatalf("wrong number of zipped elements: got %d", len(z.Result))
	}

	for i, r := range z.Result {
		idx := i % 2
		if r[idx] != i+1 {
			t.Fatalf("bad zip: expected [[1 0] [0 2] [3 0] [0 4], got %v",
				z.Result)
		}
	}
}

// INPUT: An IntZipper where all of the left elements are odd and the right
// elements are even.
// EXPECTED: [[1 2] [3 4]]
func TestZip(t *testing.T) {
	z := &IntZipper{
		[]int{1, 3},
		[]int{2, 4},
		[][2]int{},
	}

	Zip(z)

	if len(z.Result) != 2 {
		t.Fatalf("wrong number of zipped elements: got %d", len(z.Result))
	}

	for i, r := range z.Result {
		if r[0] != i*2+1 || r[1] != i*2+2 {
			t.Fatalf("bad zip: expected [[1 2] [3 4]], got %v", z.Result)
		}
	}
}

// INPUT: A zipper with nonsense compare values.
// EXPECTED: A panic.
func TestNonsenseZipperCompare(t *testing.T) {
	z := nonsenseZipper{}

	defer func() {
		if err := recover(); err == nil {
			t.Fatalf("uncaught nonsense compare value")
		}
	}()

	ZipWithGaps(z)
}

// INPUT: A zipper with negative left length.
// EXPECTED: A panic.
func TestNegativeLeftZipper(t *testing.T) {
	z := negativeLeftZipper{}

	defer func() {
		if err := recover(); err == nil {
			t.Fatalf("uncaught negative left length")
		}
	}()

	ZipWithGaps(z)
}

// INPUT: A zipper with negative right length.
// EXPECTED: A panic.
func TestNegativeRightZipper(t *testing.T) {
	z := negativeRightZipper{}

	defer func() {
		if err := recover(); err == nil {
			t.Fatalf("uncaught negative right length")
		}
	}()

	ZipWithGaps(z)
}
