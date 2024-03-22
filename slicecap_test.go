package main

import (
	"fmt"
	"testing"
)

func TestSliceCap(t *testing.T) {
	a := []int{1}
	for i := 0; i < 20; i++ {
		preLen, preCap := len(a), cap(a)
		a = append(a, a...)
		fmt.Printf("len: %d, cap: %d, curLen/preLen = %.2f, curCap/preCap = %.2f\n", len(a), cap(a), float64(len(a))/float64(preLen), float64(cap(a))/float64(preCap))
	}

	fmt.Println()

	// The cap growing 1.25x for the large slice. Ref - https://github.com/golang/go/blob/4c08c125936b4ae3daff04cecf5309dd5dd1e2c5/src/runtime/slice.go#L280-L281
	b := []int{1}
	for i := 0; i < 20; i++ {
		preLen, preCap := len(b), cap(b)
		s := make([]int, 1000*(i+1))
		b = append(b, s...)
		fmt.Printf("len: %d, cap: %d, curLen/preLen = %.2f, curCap/preCap = %.2f\n", len(b), cap(b), float64(len(b))/float64(preLen), float64(cap(b))/float64(preCap))
	}

	fmt.Println()

	// The cap growing 2x for the small slice (< 256). Ref - https://github.com/golang/go/blob/4c08c125936b4ae3daff04cecf5309dd5dd1e2c5/src/runtime/slice.go#L279
	c := []int{1}
	for i := 0; i < 20; i++ {
		preLen, preCap := len(c), cap(c)
		s := make([]int, 15)
		c = append(c, s...)
		fmt.Printf("len: %d, cap: %d, curLen/preLen = %.2f, curCap/preCap = %.2f\n", len(c), cap(c), float64(len(c))/float64(preLen), float64(cap(c))/float64(preCap))
	}
}
