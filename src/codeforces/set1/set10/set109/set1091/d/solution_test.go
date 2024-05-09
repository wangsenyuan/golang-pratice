package main

import "testing"

func runSample(t *testing.T, n int, expect int) {
	res := solve(n)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 9)
}

func TestSample2(t *testing.T) {
	runSample(t, 4, 56)
}

func TestSample3(t *testing.T) {
	runSample(t, 10, 30052700)
}
