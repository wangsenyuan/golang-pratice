package main

import "testing"

func runSample(t *testing.T, n, k int, A []int, expect int) {
	res := solve(n, k, A)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", n, k, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, k := 5, 3
	A := []int{2, 4, 2, 4, 2}
	expect := 5
	runSample(t, n, k, A, expect)
}

func TestSample2(t *testing.T) {
	n, k := 8, 5
	A := []int{9, 3, 5, 7, 8, 11, 17, 2}
	expect := 3
	runSample(t, n, k, A, expect)
}
