package main

import "testing"

func runSample(t *testing.T, n int, a int, b int, k int, expect int) {
	res := solve(n, a, b, k)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	n, a, b, k := 5, 2, 4, 1
	expect := 2
	runSample(t, n, a, b, k, expect)
}

func TestSample2(t *testing.T) {
	n, a, b, k := 5, 3, 4, 1
	expect := 0
	runSample(t, n, a, b, k, expect)
}
