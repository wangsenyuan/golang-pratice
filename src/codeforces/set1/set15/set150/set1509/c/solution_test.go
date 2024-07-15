package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{3, 1, 2}
	expect := 3
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{5}
	expect := 0
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{1, 6, 3, 3, 6, 3}
	expect := 11
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{104, 943872923, 6589, 889921234, 1000000000, 69}
	expect := 2833800505
	runSample(t, a, expect)
}
