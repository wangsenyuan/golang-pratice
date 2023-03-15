package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, s string, Q [][]int, expect []int64) {
	res := solve(s, Q)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := ")(()())()"
	Q := [][]int{
		{2, 3, 6},
		{2, 2, 7},
		{1, 3, 4},
		{2, 2, 7},
		{2, 2, 9},
		{1, 5, 6},
		{1, 2, 7},
		{2, 8, 9},
	}
	expect := []int64{3, 4, 2, 4, 1}
	runSample(t, s, Q, expect)
}
