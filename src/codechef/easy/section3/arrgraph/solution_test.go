package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 4, []int{15, 14, 21, 10}, 1)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, []int{2, 4, 8, 16}, 1)
}

func TestSample3(t *testing.T) {
	runSample(t, 4, []int{2, 4, 8, 3}, 0)
}
