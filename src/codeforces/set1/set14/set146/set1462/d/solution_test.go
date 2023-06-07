package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{3, 1, 6, 6, 2}
	expect := 4
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 2, 2, 1}
	expect := 2
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{6, 3, 2, 1}
	expect := 2
	runSample(t, A, expect)
}
