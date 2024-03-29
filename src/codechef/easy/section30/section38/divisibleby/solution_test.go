package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, A []int, expect []int64) {
	res := solve(A)

	if !reflect.DeepEqual(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2}
	expect := []int64{1, 2}
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{2, 2, 2}
	expect := []int64{1, 1, 1}
	runSample(t, A, expect)
}
