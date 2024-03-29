package main

import "testing"

func runSample(t *testing.T, a []int, expect bool) {
	res := solve(a, 1)

	if res != expect {
		t.Fatalf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{14, 5, 9}
	expect := false
	runSample(t, a, expect)
}
