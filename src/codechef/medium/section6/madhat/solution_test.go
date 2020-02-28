package main

import "testing"

func runSample(t *testing.T, n int, A []int, expect int) {
	res := solve(n, A)

	if res != expect {
		t.Errorf("Sample %d %v, expect %d, but got %d", n, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n := 4
	A := []int{0, 2, 1, 0}
	expect := 3
	runSample(t, n, A, expect)
}

func TestSample2(t *testing.T) {
	n := 4
	A := []int{0, 2, 1, 1}
	expect := 0
	runSample(t, n, A, expect)
}
