package main

import "testing"

func runSample(t *testing.T, a []int, expect int) {
	res := solve(a)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	a := []int{-4, 1, -3, 5}
	expect := 5
	runSample(t, a, expect)
}

func TestSample2(t *testing.T) {
	a := []int{1, -2, 3, -4}
	expect := 4
	runSample(t, a, expect)
}

func TestSample3(t *testing.T) {
	a := []int{-1, 3, -5}
	expect := 2
	runSample(t, a, expect)
}

func TestSample4(t *testing.T) {
	a := []int{-1}
	expect := 0
	runSample(t, a, expect)
}
