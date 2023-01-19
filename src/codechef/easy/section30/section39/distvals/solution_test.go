package main

import "testing"

func runSample(t *testing.T, A []int, expect int) {
	res := int(solve(A))

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	A := []int{1, 1}
	expect := 1
	runSample(t, A, expect)
}

func TestSample2(t *testing.T) {
	A := []int{4, 2, 1}
	expect := 2
	runSample(t, A, expect)
}

func TestSample3(t *testing.T) {
	A := []int{8, 1, 7, 2}
	expect := 4
	runSample(t, A, expect)
}

func TestSample4(t *testing.T) {
	A := []int{6, 9, 4, 2, 1}
	expect := 4
	runSample(t, A, expect)
}
