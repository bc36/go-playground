package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInitSlice(t *testing.T) {
	// context: https://leetcode.com/problems/distribute-elements-into-two-arrays-i/
	// I would like init two slice: one (a) starts with nums[0], the other (b) starts with nums[1]
	// then append the rest of elements to them one at a time: if a[-1] > b[-1], then append nums[i] to 'a', otherwise append to 'b'.
	nums := []int{2, 1, 3}

	a := nums[:1]       // {nums[0]}
	b := nums[1:2]      // {nums[1]}, false, it still bounded with 'nums'
	c := []int{nums[1]} // {nums[1]}, right

	fmt.Printf("%+v, %p\n", a, a)
	fmt.Printf("%+v, %p\n", b, b) // still shares the address with 'nums'
	fmt.Printf("%+v, %p\n", c, c) // a new address

	// b = c // The element appended to 'a' will override the value in 'b', causing [2, 3, 3]
	for _, x := range nums[2:] {
		if a[len(a)-1] > b[len(b)-1] {
			a = append(a, x)
		} else {
			b = append(b, x)
		}
	}
	fmt.Printf("Answer is %v, which is %v\n", append(a, b...), reflect.DeepEqual(append(a, b...), []int{2, 3, 1}))
}
