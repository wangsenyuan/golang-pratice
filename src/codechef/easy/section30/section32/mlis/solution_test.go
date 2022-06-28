package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := solve(A)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 2, 1}
	expect := 3
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{1, 5, 8, 6, 9}
	expect := 5
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{1, 5, 8, 6, 9}
	expect := 5
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{1, 2, 2, 1}
	expect := 3
	runSample(t, A, expect)
}
