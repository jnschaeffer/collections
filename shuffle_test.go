package collections

import (
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func TestShuffle(t *testing.T) {
	const (
		n       = 10
		target  = float64(n-1) / 2
		epsilon = target * 0.02
		iters   = 100000
	)
	var sums [n]int64
	s := make(IntSlice, n)
	for i := 0; i < iters; i++ {
		for k := 0; k < n; k++ {
			s[k] = k
		}
		Shuffle(s)
		for k := 0; k < n; k++ {
			sums[k] += int64(s[k])
		}
	}
	for i := 0; i < n; i++ {
		avg := float64(sums[i]) / iters
		delta := avg - target
		if delta < 0 {
			delta = -delta
		}
		if delta > epsilon {
			t.Fatal("bad distribution")
		}
	}
}
