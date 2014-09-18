package collections

type dummyBounded struct{}

func (b dummyBounded) Len() int { return 2 }

type dummyComparable struct{}

func (c dummyComparable) Compare(i, j int) Ord { return Equal }

type dummySwappable struct{}

func (c dummySwappable) Swap(i, j int) {}

type negativeBounded struct{}

func (b negativeBounded) Len() int { return -1 }

type badComparable struct{}

func (b badComparable) Compare(i, j int) Ord { return -1 }
