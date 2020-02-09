package main

import "testing"

func runSample(t *testing.T, N, K int, A []int, expect int) {
	res := solve(N, K, A)

	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", N, K, A, expect, res)
	}
}

func TestSample1(t *testing.T) {
	N, K := 3, 3
	A := []int{1, 2, 3}
	expect := 1
	runSample(t, N, K, A, expect)
}
